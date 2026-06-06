package users_transport_http

import (
	core_logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	"github.com/yunishuseynov99/go_todolist/internal/core/transport/http/request"
	core_http_response "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/response"
	"net/http"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, err := core_http_request.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId path value")
		return
	}

	user, err := h.usersService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
