package db_model

import (
	"github.com/google/uuid"
	"time"
)

type MessageStatus struct {
	MessageId uuid.UUID `gorm:"column:message_id;type:uuid;unique_index"`
	OrderId   string    `gorm:"column:order_id;type:varchar(100);index:order_event"`
	Recipient string    `gorm:"column:recipient;type:varchar(100)"`
	Event     string    `gorm:"column:event;type:varchar(100);index:order_event"`
	AppClient string    `gorm:"column:app_client;type:varchar(100)"`
	Channel   string    `gorm:"column:channel;type:varchar(100)"`
	Type      string    `gorm:"column:type;type:varchar(100)"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP; default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP; default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Status    string    `gorm:"column:status;type:varchar(100)"`
	Payload   string    `gorm:"column:payload;type:JSONB"`
}

func (MessageStatus) TableName() string {
	return "message_status"
}
