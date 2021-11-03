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
func (s *server) AskAmount(ctx context.Context, in *pb.Message) (*pb.Amount, error) {
	log.Printf("%v is asking for the amount", in.GetMsg() )
	Toreturn := &pb.Amount{
		Port: port,
		Amount: Total_Amount,
	  }
	return Toreturn, nil
}

var Total_Amount int32 = 86;

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetAmountServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	var word string
	sum := 1

	for sum < 2{
		fmt.Println("ingrese una palabra, si desea terminar el proceso aprete enter")
		fmt.Scanln(&word)
		fmt.Println("Su palabra fue ",word)
		if	(word == "exit"){
			sum += 1
			fmt.Println(sum)
		}
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	






}



