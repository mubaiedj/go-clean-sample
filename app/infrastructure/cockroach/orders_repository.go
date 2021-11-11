package cockroach

import (
	"errors"
	"github.com/mubaiedj/go-clean-sample/app/domain/entity"
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/cockroach/cockroach_connection"
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/cockroach/db_model"
)

type ordersRepository struct {
	connection cockroach_connection.CockroachConnection
}

func NewOrdersRepository(connection cockroach_connection.CockroachConnection) *ordersRepository {
	return &ordersRepository{
		connection: connection,
	}
}

func (r *ordersRepository) Create(order *entity.Order) (*entity.Order, error) {
	db, err := r.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	orderCockroach := db_model.Order{
		OrderId: order.Id,
		Channel: order.Channel,
		Status:  order.State,
	}

	db.Where("order_id = ?", order.Id).FirstOrCreate(&orderCockroach)
	if db.Error != nil {
		return &entity.Order{}, errors.New("error trying to Find or Create message status")
	}

	orderFound := &entity.Order{
		Id:      orderCockroach.OrderId,
		Channel: orderCockroach.Channel,
		State:   orderCockroach.Status,
	}

	return orderFound, nil
}

func (r *ordersRepository) Find(id string) (*entity.Order, error) {
	return nil, nil
}
