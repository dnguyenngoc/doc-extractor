package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dnguyenngoc/doc-extractor/pkg/dto"
	"github.com/dnguyenngoc/doc-extractor/pkg/httputil"
	"github.com/dnguyenngoc/doc-extractor/pkg/security"
	"github.com/dnguyenngoc/doc-extractor/pkg/setting"
)

func LoginAccessToken(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	clientID := request.FormValue("client_id")
	grantType := request.FormValue("grant_type")
	clientSecret := request.FormValue("client_secret")

	query := fmt.Sprintf("SELECT hashed_password FROM xtractcore.clients WHERE client_name='%s' LIMIT 1;", clientID)
	client, err := db.Query(query)

	if err != nil {
		resp := dto.XtractSyncMultiImgsResponseError{StatusCode: "Invalid/Bad request", StatusMessage: "Not Found"}
		httputil.ResponseWithJson(writer, http.StatusBadRequest, resp)
		return
	}

	var hashedPassword string
	if client.Next() {
		err := client.Scan(&hashedPassword)
		if err != nil {
			resp := dto.XtractSyncMultiImgsResponseError{StatusCode: "Invalid/Bad request", StatusMessage: "Not Found"}
			httputil.ResponseWithJson(writer, http.StatusBadRequest, resp)
			return
		}
		if !security.VerifyHash(clientSecret, hashedPassword) {
			resp := dto.XtractSyncMultiImgsResponseError{StatusCode: "Invalid/Bad request", StatusMessage: "Password not Match"}
			httputil.ResponseWithJson(writer, http.StatusBadRequest, resp)
			return
		}
	} else {
		resp := dto.XtractSyncMultiImgsResponseError{StatusCode: "Invalid/Bad request", StatusMessage: "Client not found"}
		httputil.ResponseWithJson(writer, http.StatusBadRequest, resp)
		return
	}

	input := dto.InfoAccount{
		ClientID:     clientID,
		GrantType:    grantType,
		ClientSecret: clientSecret,
	}

	// create bearer token
	token, err := security.CreateBearerToken(input)
	if err != nil {
		resp := dto.XtractSyncMultiImgsResponseError{StatusCode: "InternalServerError", StatusMessage: "Failed to create token"}
		httputil.ResponseWithJson(writer, http.StatusInternalServerError, resp)
		return
	}

	// return token in response
	res := dto.ResponseAccessToken{
		TokenType:   setting.Config.GetString("TokenType"),
		AccessToken: token,
		ExpiresIn:   setting.Config.GetInt("tokenExpire"),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(res)
}
