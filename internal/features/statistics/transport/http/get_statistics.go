package statistics_transport_http

import ("net/http"
core_logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
core_http_response "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/response")

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
}
