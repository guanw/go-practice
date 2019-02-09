package main

import (
	"fmt"
	"time"

	tchannel "github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/json"
)

var log = tchannel.SimpleLogger

type Ping struct {
	Message string `json:"message"`
}

// Ping is the response type.
type Pong Ping

func main() {
	// Create a new TChannel for sending requests.
	client, _ := tchannel.NewChannel("ping-client", nil)
	ctx, _ := json.NewContext(time.Second * 100)
	peer := client.Peers().Add("127.0.0.1:10500")

	var pong Pong
	if err := json.CallPeer(ctx, peer, "PingService", "ping", &Ping{"Hello World"}, &pong); err != nil {
		log.WithFields(tchannel.ErrField(err)).Fatal("json.Call failed.")
	}

	fmt.Println(pong.Message)
}
