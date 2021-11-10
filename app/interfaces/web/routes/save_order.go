package routes

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mubaiedj/go-clean-sample/app/application/process_order"
	"github.com/mubaiedj/go-clean-sample/app/domain/entity"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web/middleware/authentication"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web/models"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/config"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/custom_errors"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type saveOrderHandler struct {
	processOrderUseCase process_order.OrderUseCase
}

func NewSaveOrderHandler(e *echo.Echo, processOrderUseCase process_order.OrderUseCase) *saveOrderHandler {
	saveOrderHandler := &saveOrderHandler{
		processOrderUseCase: processOrderUseCase,
	}
	e.POST("/order/save", saveOrderHandler.Save, authentication.GetMiddlewareConfig())
	return saveOrderHandler
}

func (s *saveOrderHandler) Save(c echo.Context) error {
	if !config.GetBool("feature.flags.save") {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Description: "save order disabled"})
	}

	saveOrderRequest := new(models.SaveOrderRequest)

	if err := c.Bind(saveOrderRequest); err != nil {
		err := errors.New("error getting order payload")
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Description: err.Error()})
	}

	if err := c.Validate(saveOrderRequest); err != nil {
		var msgError string
		var split string
		for _, e := range err.(validator.ValidationErrors) {
			msgError = fmt.Sprintf("%s%s%s", msgError, split, e)
			split = ", "
		}
		err := errors.New(fmt.Sprintf("error validating data structure: %s", msgError))
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Description: err.Error()})
	}

	order := &entity.Order{Id: saveOrderRequest.OrderID, Channel: saveOrderRequest.Channel}

	orderValue, err := s.processOrderUseCase.Process(order)
	if err != nil {
		customErr, ok := err.(*custom_errors.RequestError)
		if ok {
			switch customErr.Kind() {
			case custom_errors.DataBaseError:
				return c.JSON(http.StatusFailedDependency, models.ErrorResponse{
					Kind:        customErr.Kind(),
					Description: err.Error(),
				})
			default:
				return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Kind: custom_errors.Unknown, Description: err.Error()})
			}
		}
	}

	return c.JSON(http.StatusCreated, models.SaveResponse{
		OrderId:     orderValue.Id,
		Description: "message save successfully",
	})
}
