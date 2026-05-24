package main

import (
	"context"
	"fmt"
	logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_http_middleware "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/middleware"
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
		newLogger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(newLogger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		newLogger.Error("HTTP server run error", zap.Error(err))
	}

}
