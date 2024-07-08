package http_adapter

import (
	api_controller "github.com/ternaryinvalid/todo1/internal/app/adapters/primary/http-adapter/api-controller"
	"github.com/ternaryinvalid/todo1/internal/app/adapters/primary/http-adapter/router"
	api_service "github.com/ternaryinvalid/todo1/internal/app/application/api-service"
	"github.com/ternaryinvalid/todo1/internal/app/domain/config"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout       = 30 * time.Second
	_defaultWriteTimeout      = 500 * time.Second
	_defaultReadHeaderTimeout = 30 * time.Second
	_defaultShutdownTimeout   = 3 * time.Second
)

type HttpAdapter struct {
	config          config.HttpAdapter
	router          http.Handler
	server          *http.Server
	shutdownTimeout time.Duration
	notify          chan error
}

func New(config config.HttpAdapter, svc *api_service.ApiService) HttpAdapter {
	r := router.New()

	ctr := api_controller.New(svc)

	r.AppendRoutes(config.Router, ctr)

	router := r.Router()

	httpServer := &http.Server{
		Handler:           router,
		ReadTimeout:       _defaultReadTimeout,
		WriteTimeout:      _defaultWriteTimeout,
		ReadHeaderTimeout: _defaultReadHeaderTimeout,
		Addr:              config.Server.Port,
	}

	return HttpAdapter{
		config:          config,
		router:          router,
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
		notify:          make(chan error, 1),
	}
}
