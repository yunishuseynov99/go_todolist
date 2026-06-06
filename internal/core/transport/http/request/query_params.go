package core_http_request

import (
	"fmt"
	core_errors "github.com/yunishuseynov99/go_todolist/internal/core/errors"
	"net/http"
	"strconv"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)

	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' is not a valid integer: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	return &val, nil
}
