package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	setting "doc-extractor/internal/settings"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello Mấy Bố!")) })
	return r
}

func main() {

	setting.InitConfig()
	fmt.Println("Init config successed!")

	// Set up the database connection pool
	// _ = database.GetMongoClient()

	fmt.Println("Load DB successed!")

	// Handle router
	r := routes()

	// start sever
	log.Println("Starting server on port 3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
