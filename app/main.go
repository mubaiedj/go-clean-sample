package main

import (
	"github.com/mubaiedj/go-clean-sample/app/application/process_order"
	"github.com/mubaiedj/go-clean-sample/app/infrastructure/memory_database"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/config"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Info("Preparing app configurations")
	config.LoadSettings("test", "go-clean-sample", "config.yaml")

	//Repository
	orderRepository := memory_database.NewOrdersRepository() // On Memory Repository

	// SQL Repository - TODO: PLease uncomment following lines if you want to inject a different database
	//connectionCockroach := cockroach.CreateCockroachDbConnection()
	//cockroach.AutoMigrateEntities(connectionCockroach)
	//orderRepository := cockroach.NewOrdersRepository(connectionCockroach)

	//UseCase
	orderUseCase := process_order.NewOrderUseCase(orderRepository)

	//Signs Catcher
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//Start WebServer
	web.NewWebServer()
	web.InitRoutes(orderUseCase)
	go web.Start(config.GetString("web.port"))

	//Graceful Shutdown process
	sig := <-quit
	gracefulShutdown(sig)
}

func gracefulShutdown(sig os.Signal) {
	web.Shutdown()
	log.Info("Shutdown process completed for signal: %v", sig)
}
