package statistics_transport_http

import (
	"context"
	"github.com/yunishuseynov99/go_todolist/internal/core/domain"
	core_http_server "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/server"
	"net/http"
	"time"
)

type StatisticsHTTPHandler struct {
	StatisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userId *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		StatisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
