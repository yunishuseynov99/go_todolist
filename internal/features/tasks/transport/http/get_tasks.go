package tasks_transport_http

import (
	"fmt"
	"github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_http_request "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/request"
	core_http_response "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/response"
	"net/http"
)

type GetTasksResponse []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIDULimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user_id/limit/offset query params",
		)

		return
	}

	tasksDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks",
		)
		return
	}

	response := GetTasksResponse(TaskDTOsFromDomains(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDULimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		userIDQueryParamKey = "user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query parameter: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query parameter: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query parameter: %w", err)
	}

	return userID, limit, offset, nil
}
