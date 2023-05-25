package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dnguyenngoc/doc-extractor/pkg/dto"
	"github.com/dnguyenngoc/doc-extractor/pkg/httputil"
	"github.com/dnguyenngoc/doc-extractor/pkg/setting"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			resp := dto.VnIddocsResponseError{StatusCode: "Unauthorized", StatusMessage: "Not Autherization"}
			httputil.ResponseWithJson(writer, http.StatusUnauthorized, resp)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if t.Method == nil || t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(setting.Config.GetString("ApiTokenSecret")), nil
		})

		if err != nil || !token.Valid {
			if strings.Contains(err.Error(), "Token is expired") {
				resp := dto.VnIddocsResponseError{StatusCode: "Forbidden", StatusMessage: "Token is expired"}
				httputil.ResponseWithJson(writer, http.StatusForbidden, resp)
				return
			}

			resp := dto.VnIddocsResponseError{StatusCode: "Unauthorized", StatusMessage: "Invalid Token"}
			httputil.ResponseWithJson(writer, http.StatusUnauthorized, resp)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			resp := dto.VnIddocsResponseError{StatusCode: "Unauthorized", StatusMessage: "Error parsing claims"}
			httputil.ResponseWithJson(writer, http.StatusUnauthorized, resp)
			return
		}

		if next != nil {
			next(writer, request)
		}

	}
}
