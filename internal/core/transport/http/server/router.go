package core_http_server

import (
	"fmt"
	core_http_middleware "github.com/yunishuseynov99/go_todolist/internal/core/transport/http/middleware"
	"net/http"
)

type APIVersion string

var (
	ApiVersion1 = APIVersion("v1")
	ApiVersion2 = APIVersion("v2")
	ApiVersion3 = APIVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion  APIVersion
	middlewares []core_http_middleware.Middleware
}

func NewAPIVersionRouter(apiVersion APIVersion, middlewares ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:    http.NewServeMux(),
		apiVersion:  apiVersion,
		middlewares: middlewares,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(
		r,
		r.middlewares...)

}
