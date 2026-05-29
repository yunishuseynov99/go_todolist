package users_transport_http

import (
	"fmt"
	core_logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_http_response "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/response"
	core_http_utils "github.com/yunishuseynov99/go_todolist/internal/features/users/transport/utils"
	"net/http"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit' and 'offset' query params",
		)
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)
		return
	}

	response := GetUsersResponse(usersDTOFromDomain(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")

	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query parameter: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query parameter: %w", err)
	}

	return limit, offset, nil
}
