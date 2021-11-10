package kafka_models

type MessageStatusKafka struct {
	OrderId   string `json:"order_id"`
	Channel   string `json:"channel"`
	MessageId string `json:"message_id"`
	Status    string `json:"status"`
	Event     string `json:"event"`
	AppClient string `json:"app_client"`
}
