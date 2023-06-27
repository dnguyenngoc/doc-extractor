package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	setting "doc-extractor/config"
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
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
