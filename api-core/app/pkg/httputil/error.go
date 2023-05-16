package httputil

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func NewError(writer http.ResponseWriter, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(er)
}
