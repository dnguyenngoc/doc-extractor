package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dnguyenngoc/doc-extractor/pkg/database"
	"github.com/dnguyenngoc/doc-extractor/pkg/handler"
	"github.com/dnguyenngoc/doc-extractor/pkg/middleware"
	"github.com/dnguyenngoc/doc-extractor/pkg/setting"
	"github.com/gorilla/mux"
)

type BaseRecord struct {
	RequestID   string                 `json:"request_id"`
	RequestType string                 `json:"request_type"`
	Pipeline    map[string]interface{} `json:"pipeline"`
	Data        map[string]interface{} `json:"data"`
}

func routes() *mux.Router {

	r := mux.NewRouter()

	// Hello page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello Mấy Bố!")) })

	// Login get token
	r.HandleFunc("/api/v1/account/login/access-token", func(w http.ResponseWriter, r *http.Request) {
		db := database.GetDBConnection()
		handler.LoginAccessToken(w, r, db)
	}).Methods(http.MethodPost)

	// Api doc processing
	r_doc_process := r.PathPrefix("/api/v1/document-processing").Subrouter()
	r_doc_process.HandleFunc("/vniddocs/async/", middleware.ValidateJWT(handler.VnIddocsAsync)).Methods(http.MethodPost)

	return r
}

func main() {

	setting.InitConfig()
	fmt.Println("Init config successed!")

	// Set up the database connection pool
	_ = database.CreateDBPool()

	fmt.Println("Load DB successed!")

	// Handle router
	r := routes()

	// start sever
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
