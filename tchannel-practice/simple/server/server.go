package main

import (
	"fmt"

	tchannel "github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/json"
	"golang.org/x/net/context"
)

var log = tchannel.SimpleLogger

type Ping struct {
	Message string `json:"message"`
}

// Ping is the response type.
type Pong Ping

func pingHandler(ctx json.Context, ping *Ping) (*Pong, error) {
	return &Pong{
		Message: fmt.Sprintf("ping %v", ping),
	}, nil
}
func onError(ctx context.Context, err error) {
	log.WithFields(tchannel.ErrField(err)).Fatal("onError handler triggered.")
}

func listenAndHandle(s *tchannel.Channel, hostPort string) {
	log.Infof("Service %s", hostPort)

	// If no error is returned, the listen was successful. Serving happens in the background.
	if err := s.ListenAndServe(hostPort); err != nil {
		log.WithFields(
			tchannel.LogField{Key: "hostPort", Value: hostPort},
			tchannel.ErrField(err),
		).Fatal("Couldn't listen.")
	}
}

func main() {
	ch, err := tchannel.NewChannel("PingService", &tchannel.ChannelOptions{Logger: tchannel.SimpleLogger})
	if err != nil {
		log.WithFields(tchannel.ErrField(err)).Fatal("Couldn't create new channel.")
	}

	// Register a handler for the ping message on the PingService
	json.Register(ch, json.Handlers{
		"ping": pingHandler,
	}, onError)

	// Listen for incoming requests
	listenAndHandle(ch, "127.0.0.1:10500")

	select {}
}
