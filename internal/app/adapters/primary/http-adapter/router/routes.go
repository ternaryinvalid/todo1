package router

import (
	api_controller "github.com/ternaryinvalid/todo1/internal/app/adapters/primary/http-adapter/api-controller"
	auth_service "github.com/ternaryinvalid/todo1/internal/app/application/auth-service"
	"github.com/ternaryinvalid/todo1/internal/app/domain/config"
	"net/http"
)

func (r *Router) AppendRoutes(config config.Router, ctr *api_controller.Controller) {
	r.config = config

	apiV1Subrouter := r.router.PathPrefix(apiV1Prefix).Subrouter()

	authMiddleware := auth_service.New(r.config.AuthorizationConfig)

	routes := []Route{
		{
			Name:    "/reports/get",
			Path:    "/reports/get",
			Method:  http.MethodPost,
			Handler: authMiddleware.AuthMiddleware(),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}
