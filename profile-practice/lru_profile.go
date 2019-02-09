package main

import (
	lru "github.com/hashicorp/golang-lru"

	"github.com/pkg/profile"
)
import _ "net/http/pprof"

type traceID struct {
	Low  uint64 `json:"lo"`
	High uint64 `json:"hi"`
}

func main() {
	// Memory profiling
	defer profile.Start(profile.MemProfile).Stop()
	cache, _ := lru.New(2700000)
	for i := 0; i < 2700000; i++ {
		trace := traceID{
			Low:  uint64(i + 1),
			High: uint64(i + 2),
		}
		cache.Add(trace, 100000)
	}
}
