package main

import (
	"context"
	"fmt"
	logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_http_server "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/server"
	users_transport_http "github.com/yunishuseynov99/go_todolist/internal/features/users/transport/http"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer cancel()

	newLogger, err := logger.NewLogger(logger.NewConfigMust())

	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
		os.Exit(1)
	}
	defer newLogger.Close()

	newLogger.Debug("Starting ToDo application !")

	usersTransportHttp := users_transport_http.NewUsersHTTPHandler(nil)

	usersRoutes := usersTransportHttp.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)

	apiVersionRouter.RegisterRoutes(usersRoutes...)
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		newLogger)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		newLogger.Error("HTTP server run error", zap.Error(err))
	}

}
