package main

import (
	//"context"
	//"log"
	//"net"
	"fmt"
	//"google.golang.org/grpc"
	//"container/list"
	//pb "Lider/proto"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	//pb.Uni
}

type player struct {
	name string
	status bool
  }



func main() {
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterServiceMessageServer(s, &server{})
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
	
	//lista con los jugadores
	//players := list.New()

	var GameState int32 = 0

	for GameState < 4 {
		if GameState == 0 {
			fmt.Println("El juego esta en estado 0")
		} else if GameState == 1 {
			fmt.Println("El juego esta en estado 1")
		} else if GameState == 2 {
			fmt.Println("El juego esta en estado 2")
		} else if GameState == 3 {
			fmt.Println("El juego esta en estado 3")
		} else {
			fmt.Println("Accion fuera de rango")
		}

		GameState += 1
	}



}



