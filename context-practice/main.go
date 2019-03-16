package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)

		s.Scan()
		fmt.Println("scanned: ", s.Text())
		cancel()
	} ()

	sleepAndTalk(ctx, 5 * time.Second, "hello")
}

func sleepAndTalk(c context.Context, t time.Duration, message string) {
	//time.Sleep(t)
	select {
		case <- time.After(t):
			fmt.Println(message)
		case <- c.Done():
			log.Print(c.Err())

	}

}
