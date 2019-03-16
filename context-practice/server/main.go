package main

import (
	"context"
	"fmt"
	//"log"
	"net/http"
	"time"
	"github.com/judew/go-practice/context-practice/log"
)



func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100))
	log.Println(ctx,"handler started")
	defer log.Println(ctx,"handler ended")

	select {
	case <- time.After(5 * time.Second):
		fmt.Println(w, "hello")
	case <- ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), 500)
	}
}
