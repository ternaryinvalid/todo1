package api_controller

import (
	api_service "github.com/ternaryinvalid/todo1/internal/app/application/api-service"
)

type Controller struct {
	service *api_service.ApiService
}

func New(service *api_service.ApiService) *Controller {
	return &Controller{
		service: service,
	}
}
