package main

import (
	"github.com/mubaiedj/go-clean-sample/app/application/process_order"
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/memory_database"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/config"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
)

func main() {
	log.Info("Preparing app configurations")
	config.LoadSettings("test", "go-clean-sample", "config.yaml")

	//Repository
	orderRepository := memory_database.NewOrdersRepository()

	//UseCase
	orderUseCase := process_order.NewOrderUseCase(orderRepository)

	//WebServer
	web.NewWebServer()
	web.InitRoutes(orderUseCase)
	web.Start(config.GetString("web.port"))
}
