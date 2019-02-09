package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gorilla/mux"

	lru "github.com/hashicorp/golang-lru"
)

type traceID struct {
	Low  uint64
	High uint64
}

func enablePprof() {
	// Create a new router
	router := mux.NewRouter()

	// Register pprof handlers
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("pprof is enabled now!")

}

func main() {
	cache, err := lru.New(150000)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		enablePprof()
	}()

	count := 0
	t := traceID{
		Low:  uint64(0),
		High: uint64(1),
	}
	for {
		fmt.Println("Inserting...", count)
		for i := 0; i < 150000; i++ {
			t.Low = uint64(i)
			t.High = uint64(i + 1)
			cache.Add(t, count)
		}
		count++
		fmt.Println("Sleep...", count)
		time.Sleep(1 * time.Second)
	}
}
