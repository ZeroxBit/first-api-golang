package models

// Message para el cliente del api
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
