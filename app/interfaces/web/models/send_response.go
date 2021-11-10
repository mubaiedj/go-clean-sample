package models

type ErrorResponse struct {
	Kind        string `json:"kind"`
	Description string `json:"description"`
}

type SentResponse struct {
	MessageId   string `json:"message_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type SaveResponse struct {
	OrderId     string `json:"order_id"`
	Description string `json:"description"`
}
