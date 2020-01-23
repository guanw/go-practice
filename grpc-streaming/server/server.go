package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	stream "github.com/guanw/go-practice/grpc-streaming/proto"
	"google.golang.org/grpc"
)

type dnsServer struct {
}

func NewServer() stream.DnsServiceServer {
	return &dnsServer{}
}

func (d *dnsServer) Echo(req *stream.Req, st stream.DnsService_EchoServer) error {
	if req.Id == 1 {
		for _, item := range []struct {
			Service string
			Host    string
		}{
			{
				Service: "s1",
				Host:    "192.0.0.1",
			},
			{
				Service: "s2",
				Host:    "192.0.0,2",
			},
		} {
			st.Send(&stream.Resp{
				Service: item.Service,
				Host:    item.Host,
			})
		}
	} else {
		return errors.New("id not found")
	}
	return nil
}

func main() {
	server := NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	stream.RegisterDnsServiceServer(grpcServer, server)
	grpcServer.Serve(lis)
}
