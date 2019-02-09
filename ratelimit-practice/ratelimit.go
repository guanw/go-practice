package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

//type ratelimiter struct {
//	rl Limiter
//}
//
//func newRateLimiter() {
//	rl := ratelimit.New(200)
//	r := ratelimiter{
//		rl: rl,
//	}
//}
func main() {
	rl := ratelimit.New(10) // per second

	prev := time.Now()
	for i := 0; i < 100; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
