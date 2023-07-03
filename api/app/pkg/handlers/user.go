// doc-extractor/pkg/handlers/user.go

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"doc-extractor/pkg/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve users from the database
	users := []models.User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	// Convert users to JSON
	response, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
