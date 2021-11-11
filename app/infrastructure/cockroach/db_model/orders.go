package db_model

type Order struct {
	OrderId string `gorm:"column:order_id;type:varchar(100);index:order_id"`
	Channel string `gorm:"column:channel;type:varchar(100)"`
	Status  string `gorm:"column:status;type:varchar(100)"`
}

func (Order) TableName() string {
	return "orders"
}
