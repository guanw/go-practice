package mux

import (
	"context"
	"go.uber.org/fx"
	"log"
	"net/http"
)

var Module = fx.Provide(NewMux)

type MuxParams struct {
	fx.In
	Logger *log.Logger
}

func NewMux(lc fx.Lifecycle, params MuxParams) *http.ServeMux {
	logger := params.Logger

	logger.Print("Executing NewMux.")

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}
