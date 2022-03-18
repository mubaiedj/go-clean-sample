package web

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mubaiedj/go-clean-sample/app/application/process_order"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web/middleware/json_validator"
	"github.com/mubaiedj/go-clean-sample/app/interfaces/web/routes"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"net/http"
	"time"
)

var echoServer *echo.Echo

func NewWebServer() {
	echoServer = echo.New()
	echoServer.HideBanner = true
	echoServer.HidePort = true
	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.RequestID())
	echoServer.Validator = json_validator.NewJsonValidator()
}

func InitRoutes(saveOrderUseCase process_order.OrderUseCase) {
	routes.NewHealthHandler(echoServer)
	routes.NewMetricsHandler(echoServer)
	routes.NewPingHandler(echoServer)
	routes.NewLoginHandler(echoServer)
	routes.NewSaveOrderHandler(echoServer, saveOrderUseCase)
}

func Start(port string) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	log.Info("App listen in port %s", port)
	echoServer.Logger.Fatal(echoServer.StartServer(server))
}

func Shutdown() {
	log.Info("Shutting down web server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoServer.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down web server")
	}
}
