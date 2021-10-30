package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "proyecto/proto"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedServiceMessageServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Sayhello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Received: %v", in.GetName())
	Total_Amount += in.GetAmount()
	log.Printf("TotalAmount: %v", Total_Amount)
	return in, nil
}

var Total_Amount int32 = 0;

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceMessageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



