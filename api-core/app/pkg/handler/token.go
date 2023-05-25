package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dnguyenngoc/doc-extractor/pkg/dto"
	"github.com/dnguyenngoc/doc-extractor/pkg/httputil"
	"github.com/dnguyenngoc/doc-extractor/pkg/security"
	"github.com/dnguyenngoc/doc-extractor/pkg/setting"
)

func LoginAccessToken(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	clientID := request.FormValue("client_id")
	grantType := request.FormValue("grant_type")
	// clientSecret := request.FormValue("client_secret")

	input := dto.InfoAccount{
		ClientID:  clientID,
		GrantType: grantType,
	}

	// create bearer token
	token, err := security.CreateBearerToken(input)
	if err != nil {
		resp := dto.VnIddocsResponseError{StatusCode: "InternalServerError", StatusMessage: "Failed to create token"}
		httputil.ResponseWithJson(writer, http.StatusInternalServerError, resp)
		return
	}

	// return token in response
	res := dto.ResponseAccessToken{
		TokenType:   setting.Config.GetString("ApiTokenType"),
		AccessToken: token,
		ExpiresIn:   setting.Config.GetInt("ApiTokenExpire"),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(res)
}
