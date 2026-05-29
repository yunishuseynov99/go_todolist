package users_transport_http

import (
	core_logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	core_http_response "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/response"
	core_http_utils "github.com/yunishuseynov99/go_todolist/internal/features/users/transport/utils"
	"net/http"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID value")
		return
	}

	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.NoContentResponse()
}
