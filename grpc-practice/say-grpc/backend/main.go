package main

import (
	"flag"
	"fmt"
	pb "github.com/judew/go-practice/grpc-practice/say-grpc/backend/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
)

func main() {


	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	log.Printf("listening to port %d", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal("could not serve: %v", err)
	}
}

type server struct {

}

func (server) Say(ctx context.Context, text *pb.Text) (speech *pb.Speech, err error) {
	f, err := ioutil.TempFile("","")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}

	return &pb.Speech{Audio: data}, nil
}