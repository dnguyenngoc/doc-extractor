package main

import (
	"log"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/dnguyenngoc/doc-extractor/pkg/handler"

)

type BaseRecord struct {
	RequestID   string                 `json:"request_id"`
	RequestType string                 `json:"request_type"`
	Pipeline    map[string]interface{} `json:"pipeline"`
	Data        map[string]interface{} `json:"data"`
}


func routes(db *sql.DB) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello Mấy Bố!")) })

	r.HandleFunc("/api/v1/account/login/access-token", func(w http.ResponseWriter, r *http.Request) {
		handler.LoginAccessToken(w, r, db)
	}).Methods(http.MethodPost)

	return r
}


func main() {

	setting.InitConfig()
	fmt.Println("Init config successed!")

	// Set up the database connection
	db, err := database.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Load DB successed!")

	// Handle router
	r := routes(db)

	// start sever
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// func publishMessage(w http.ResponseWriter, r *http.Request) {
// 	// Parse the request body into a BaseRecord struct
// 	var record BaseRecord
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Create a Kafka writer to publish messages to the 'section_gpu' topic
// 	writer := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers: []string{"kafka:9092"},
// 		Topic:   "section_gpu",
// 	})

// 	// Convert the record to JSON bytes
// 	valueBytes, err := json.Marshal(record)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Publish the record to the 'section_gpu' topic
// 	err = writer.WriteMessages(r.Context(), kafka.Message{
// 		Value: valueBytes,
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Close the Kafka writer
// 	writer.Close()

// 	// Return a success message
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Message published successfully"))
