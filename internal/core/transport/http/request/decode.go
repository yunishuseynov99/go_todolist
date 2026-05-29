package core_http_request

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	core_errors "github.com/yunishuseynov99/go_todolist/internal/core/errors"
	"net/http"
)

var RequestValidator = validator.New()

type Validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf(
			"decode json: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	var err error

	v, ok := dest.(Validatable)
	if ok {
		err = v.Validate()
	} else {
		err = RequestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf(
			"request validation: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
