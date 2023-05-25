package dto

type InfoAccount struct {
	ClientID     string    `json:"client_id"`
	GrantType    string    `json:"grant_type"`
	ClientSecret string    `json:"client_secret"`
}