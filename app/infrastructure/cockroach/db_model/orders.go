package db_model

import (
	"github.com/google/uuid"
)

type MessageStatus struct {
	UniqueId uuid.UUID `gorm:"column:message_id;type:uuid;unique_index"`
	OrderId  string    `gorm:"column:order_id;type:varchar(100);index:order_event"`
	Channel  string    `gorm:"column:channel;type:varchar(100)"`
	Status   string    `gorm:"column:status;type:varchar(100)"`
}

func (MessageStatus) TableName() string {
	return "orders"
}
