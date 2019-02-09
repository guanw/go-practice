package handler

import (
	"go.uber.org/fx"
	"log"
	"net/http"
)

var Module = fx.Provide(HandlerMap)

func HandlerMap(logger *log.Logger) (map[string]http.Handler, error) {

	logger.Print("Executing handlerMap")

	res1 := make(map[string]http.Handler)
	res1["/"] = rootHandler(logger)
	res1["/hello"] = helloHandler(logger)
	return res1, nil
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
