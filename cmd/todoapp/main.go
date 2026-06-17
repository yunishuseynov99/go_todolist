package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/yunishuseynov99/go_todolist/internal/core/config"
	logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_pgx_pool "github.com/yunishuseynov99/go_todolist/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/middleware"
	core_http_server "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/server"
	tasks_postgres_repository "github.com/yunishuseynov99/go_todolist/internal/features/tasks/repository/postgres"
	tasks_service "github.com/yunishuseynov99/go_todolist/internal/features/tasks/service"
	tasks_transport_http "github.com/yunishuseynov99/go_todolist/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/yunishuseynov99/go_todolist/internal/features/users/repository/postgres"
	users_service "github.com/yunishuseynov99/go_todolist/internal/features/users/service"
	users_transport_http "github.com/yunishuseynov99/go_todolist/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer cancel()

	newLogger, err := logger.NewLogger(logger.NewConfigMust())

	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
		os.Exit(1)
	}
	defer newLogger.Close()

	newLogger.Debug("app time zone", zap.Any("zone", time.Local))

	newLogger.Debug("initializing postgres connection pool")

	pool, err := core_pgx_pool.NewConnectionPool(
		ctx,
		core_pgx_pool.NewConfigMust())
	if err != nil {
		newLogger.Fatal("Failed to initialize postgres connection pool:", zap.Error(err))
	}

	defer pool.Close()

	newLogger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHttp := users_transport_http.NewUsersHTTPHandler(usersService)

	newLogger.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHttp := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	newLogger.Debug("initializing http server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		newLogger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(newLogger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHttp.Routes()...)
	apiVersionRouterV1.RegisterRoutes(tasksTransportHttp.Routes()...)

	//apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion2,
	//	core_http_middleware.Dummy("api v2 middleware"),
	//)
	//apiVersionRouterV2.RegisterRoutes(usersTransportHttp.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1 /*apiVersionRouterV2*/)

	if err := httpServer.Run(ctx); err != nil {
		newLogger.Error("HTTP server run error", zap.Error(err))
	}

}
