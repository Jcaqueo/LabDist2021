package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"fmt"
	pb "Pozo/proto"
)

const (
	port = ":50051"
)

// Server
type server struct {
	pb.UnimplementedGetAmountServer
}

// Total amount from proto
func (s *server) PedirPozo(ctx context.Context, msg *pb.Msg) (*pb.Amount, error) {
	fmt.Println(Msg.message)
	Toreturn := &pb.Amount{
		Amount: Total_Amount,
	  }
	return Toreturn, nil
}

var Total_Amount int32 = 0;

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetAmountServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	






}



