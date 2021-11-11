package main

import (
	"github.com/mubaiedj/go-clean-sample/app/application/process_order"
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/cockroach"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/config"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
)

func main() {
	log.Info("Preparing app configurations")
	config.LoadSettings("test", "go-clean-sample", "config.yaml")

	// Database connection
	connectionCockroach := cockroach.CreateCockroachDbConnection()
	cockroach.AutoMigrateEntities(connectionCockroach)

	//Repository
	//orderRepository := memory_database.NewOrdersRepository() // On Memory Repository
	orderRepository := cockroach.NewOrdersRepository(connectionCockroach) // SQL Repository

	//UseCase
	orderUseCase := process_order.NewOrderUseCase(orderRepository)

	//WebServer
	web.NewWebServer()
	web.InitRoutes(orderUseCase)
	web.Start(config.GetString("web.port"))
}
