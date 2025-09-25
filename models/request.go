package models

type Request struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Method string `json:"method"`
}
