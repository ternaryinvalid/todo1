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
			Name:    "/todo/create",
			Path:    "/todo/create",
			Method:  http.MethodPost,
			Handler: authMiddleware.AuthMiddleware(http.HandlerFunc(ctr.CreateTODO)),
		},
		{
			Name:    "/todo/get",
			Path:    "/todo/get",
			Method:  http.MethodGet,
			Handler: authMiddleware.AuthMiddleware(http.HandlerFunc(ctr.GetAllTODO)),
		},
		{
			Name:    "/todo/delete",
			Path:    "/todo/delete",
			Method:  http.MethodPost,
			Handler: authMiddleware.AuthMiddleware(http.HandlerFunc(ctr.DeleteTODO)),
		},
		{
			Name:    "/todo/done",
			Path:    "/todo/done",
			Method:  http.MethodPost,
			Handler: authMiddleware.AuthMiddleware(http.HandlerFunc(ctr.Done)),
		},
		{
			Name:    "/auth",
			Path:    "/auth",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(ctr.SignUP),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}
