package entity

type Order struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
	State   string `json:"state"`
}
