package core_http_request

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var RequestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}

	if err := RequestValidator.Struct(dest); err != nil {
		return fmt.Errorf("request validation: %w", err)
	}

	return nil
}
