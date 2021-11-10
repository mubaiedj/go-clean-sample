package repository

import "github.com/mubaiedj/go-clean-sample/app/domain/entity"

type OrdersRepository interface {
	Create(order *entity.Order) (*entity.Order, error)
	Find(id string) (*entity.Order, error)
}
