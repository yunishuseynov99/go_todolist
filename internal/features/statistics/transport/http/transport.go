package statistics_transport_http

type StatisticsHTTPHandler struct {
	StatisticsService StatisticsService
}

type StatisticsService interface {
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		StatisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{}
}