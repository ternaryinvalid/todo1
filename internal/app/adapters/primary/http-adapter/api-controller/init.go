package api_controller

import (
	api_service "github.com/ternaryinvalid/todo1/internal/app/application/api-service"
	auth_service "github.com/ternaryinvalid/todo1/internal/app/application/auth-service"
)

type Controller struct {
	service *api_service.ApiService
	auth    *auth_service.AuthService
}

func New(service *api_service.ApiService, authService *auth_service.AuthService) *Controller {
	return &Controller{
		service: service,
		auth:    authService,
	}
}
