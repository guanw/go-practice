package main

import (
	"context"
	"fmt"
	"io"
	"log"

	stream "github.com/guanw/go-practice/grpc-streaming/proto"
	"google.golang.org/grpc"
)

type dnsClient struct {
}

func NewClient() stream.DnsServiceClient {
	return &dnsClient{}
}

func (d *dnsClient) Echo(ctx context.Context, in *stream.Req, opts ...grpc.CallOption) (stream.DnsService_EchoClient, error) {
	return nil, nil
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := stream.NewDnsServiceClient(conn)
	stream, err := client.Echo(context.Background(), &stream.Req{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("echo failed %v", err)
	}
	for {
		h, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("service and host", h.Service, h.Host)
	}

}
