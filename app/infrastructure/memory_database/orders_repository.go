package memory_database

import (
	"github.com/mubaiedj/go-clean-sample/app/domain/entity"
)

type ordersRepository struct {
	orders []entity.Order
}

func NewOrdersRepository() *ordersRepository {
	return &ordersRepository{}
}

func (c *ordersRepository) Create(order *entity.Order) (*entity.Order, error) {
	c.orders = append(c.orders, *order)
	return order, nil
}

func (c *ordersRepository) Find(id string) (*entity.Order, error) {
	for i := range c.orders {
		if c.orders[i].Id == id {
			return &c.orders[i], nil
		}
	}
	return nil, nil
}
