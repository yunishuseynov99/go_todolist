package core_http_middleware

import (
	"fmt"
	core_logger "github.com/yunishuseynov99/go_todolist/internal/core/logger"
	"net/http"
)

func Dummy(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)

			log.Debug(fmt.Sprintf("-> befor: %s", s))
			next.ServeHTTP(w, r)
			log.Debug(fmt.Sprintf("<- after: %s", s))

		})
	}
}
