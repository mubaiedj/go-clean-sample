package process_order

import (
	"github.com/mubaiedj/go-clean-sample/app/application"
	"github.com/mubaiedj/go-clean-sample/app/domain/constant/status"
	"github.com/mubaiedj/go-clean-sample/app/domain/entity"
	"github.com/mubaiedj/go-clean-sample/app/domain/repository"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/custom_errors"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/metrics"
)

type OrderUseCase interface {
	application.UseCase
	Process(order *entity.Order) (*entity.Order, error)
}

type orderUseCase struct {
	ordersRepository repository.OrdersRepository
}

func NewOrderUseCase(repository repository.OrdersRepository) *orderUseCase {
	return &orderUseCase{
		ordersRepository: repository,
	}
}

func (s *orderUseCase) Process(order *entity.Order) (*entity.Order, error) {
	order.Channel = status.INIT
	savedOrder, err := s.ordersRepository.Create(order)
	if err != nil {
		log.WithError(err).Error("error trying to save order")
		return nil, custom_errors.NewWithError(err, custom_errors.DataBaseError)
	}
	log.WithFields(log.Field("order", savedOrder.Id)).Info("Order saved successfully")
	metrics.IncrementTestMetrics("SAVE")
	return savedOrder, nil
}
