package main

import (
	"fmt"
	"sync"
)

type ParallelProcessor struct {
	messages    chan string
	numRoutines int

	closed chan struct{}
	wg     sync.WaitGroup
}

func main() {
	messages := make(chan string)
	k := ParallelProcessor{
		messages: messages,
	}
	numRoutines := 10
	for i := 0; i < numRoutines; i++ {
		k.wg.Add(1)
		go func() {
			for {
				select {
				case msg := <-k.messages:
					fmt.Println(msg)
				case <-k.closed:
					k.wg.Done()
					return
				}
			}
		}()
	}
}
