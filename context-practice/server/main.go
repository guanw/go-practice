package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler started")
	defer log.Println("handler ended")

	ctx := r.Context()

	select {
	case <- time.After(5 * time.Second):
		fmt.Println(w, "hello")
	case <- ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), 500)
	}
}
