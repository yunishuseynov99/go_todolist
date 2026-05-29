package main

import (
	"context"
	"fmt"
	logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_postgres_pool "github.com/yunishuseynov99/go_todolist/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/middleware"
	core_http_server "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/server"
	users_postgres_repository "github.com/yunishuseynov99/go_todolist/internal/features/users/repository/postgres"
	users_service "github.com/yunishuseynov99/go_todolist/internal/features/users/service"
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

	newLogger.Debug("initializing postgres connection pool")

	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewConfigMust())
	if err != nil {
		newLogger.Fatal("Failed to initialize postgres connection pool:", zap.Error(err))
	}

	defer pool.Close()

	newLogger.Debug("initializing feature", zap.String("feature", "users"))

	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHttp := users_transport_http.NewUsersHTTPHandler(usersService)

	newLogger.Debug("initializing http server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		newLogger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(newLogger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHttp.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		newLogger.Error("HTTP server run error", zap.Error(err))
	}

}
