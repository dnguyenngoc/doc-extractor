package dto

type BaseRecord struct {
	RequestID   string                 `json:"request_id"`
	RequestType string                 `json:"request_type"`
	Pipeline    map[string]interface{} `json:"pipeline"`
	Data        map[string]interface{} `json:"data"`
}
