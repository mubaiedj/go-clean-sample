package models

type SaveOrderRequest struct {
	OrderID string `json:"order_id" validate:"required"`
	Channel string `json:"channel" validate:"required"`
}
