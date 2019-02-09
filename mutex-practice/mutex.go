package main

import (
	"fmt"
	"sync"
	"time"
)

type simple struct {
	sync.Mutex
	count int
}

func main() {
	s := &simple{
		count: 0,
	}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				s.Lock()
				s.count++
				s.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(s.count)
}
