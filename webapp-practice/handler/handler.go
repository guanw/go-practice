package handler

import (
	"go.uber.org/fx"
	"log"
	"net/http"
)

var Module = fx.Provide(HandlerMap)

//type HandlerValue struct {
//	handler http.Handler
//}

func HandlerMap(logger *log.Logger) ([]http.Handler, error) {

	logger.Print("Executing handlerMap")

	res := []http.Handler{
		rootHandler(logger),
		helloHandler(logger),
	}
	return res, nil
}

func rootHandler(logger *log.Logger) http.Handler {
	logger.Print("Executing rootHandler")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Print("Go a root request")
		response := []byte("this is the root")
		w.Write(response)
	})
}

func helloHandler(logger *log.Logger) http.Handler {
	logger.Print("Executing helloHandler")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Print("Go a hello request")
		response := []byte("hello!")
		w.Write(response)
	})
}
