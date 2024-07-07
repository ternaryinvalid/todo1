package http_adapter

import (
	"context"
	"log"
)

func (a *HttpAdapter) Start() {
	startMsg := "Сервер запущен."

	log.Println(startMsg)

	a.notify <- a.server.ListenAndServe()
	close(a.notify)
}

func (a *HttpAdapter) Notify() <-chan error {
	return a.notify
}

func (a *HttpAdapter) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
