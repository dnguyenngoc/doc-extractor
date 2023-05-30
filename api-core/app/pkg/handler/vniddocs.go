package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dnguyenngoc/doc-extractor/pkg/dto"
	"github.com/segmentio/kafka-go"
)

func VnIddocsAsync(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a BaseRecord struct
	var record dto.BaseRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a Kafka writer to publish messages to the 'upload_document' topic
	writer := &kafka.Writer{
		Addr:     kafka.TCP("kafka1:9092", "kafka2:9092"),
		Topic:    "upload_document",
		Balancer: &kafka.LeastBytes{},
	}

	// Convert the record to JSON bytes
	valueBytes, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Publish the record to the 'upload_document' topic
	err = writer.WriteMessages(r.Context(), kafka.Message{
		Value: valueBytes,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Close the Kafka writer
	writer.Close()

	// Return a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message published successfully"))
}
