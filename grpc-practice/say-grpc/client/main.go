package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	pb "github.com/judew/go-practice/grpc-practice/say-grpc/backend/api"
	"os"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of say backend")
	output := flag.String("o", "output.wav", "wav output file")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("usage:\n\t%s \"text to speak\"", os.Args[0])
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}

	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{
		Text: os.Args[1],
	}
	speech, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatal("could not call say %v", err)
	}
	if ioutil.WriteFile(*output, speech.Audio, 0666); err != nil {
		log.Fatal("could not write to wav file: %v", err)
	}
}
