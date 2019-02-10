package main

import (
	"github.com/go-practice/webapp-practice/handler"
	"github.com/go-practice/webapp-practice/logger"
	"github.com/go-practice/webapp-practice/mux"
	"go.uber.org/fx"
	"net/http"
)

func Register(mux *http.ServeMux, m map[string]http.Handler) {
	for k, v := range m {
		mux.Handle(k, v)
	}
}

func main() {
	app := fx.New(
		logger.Module,
		mux.Module,
		handler.Module,
		fx.Invoke(Register),
	)
	app.Run()
}
