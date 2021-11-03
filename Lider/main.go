package main

import (
	"context"
	"net"
	"fmt"
	"log"
	"math/rand"
	"google.golang.org/grpc"
	pb "Lider/proto"
	"time"
)

const (
	port = ":50051"


)

var siguientepaso int = 0
var esperandoporjugadores int = 16
var juego1 int32

// Server
type server struct {
	pb.UnimplementedStartServerServer
}


func InitServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStartServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return 
}

func (s *server) AgregarJugador(ctx context.Context, jugador *pb.Name ) (*pb.Status, error){
	fmt.Println("El jugador %v fue anadido al juego",jugador)
	//tiempo muerto esperando que el lider inicie el juego
	for siguientepaso < 1 {
		
	}

	//cuando cambie el estado devolvemos la variable
	siguientepaso = 0
	//retornamos el estado del jugador
	retorno := &pb.Status{
		Status: true,
	  }
	return retorno, nil
}

func (s *server) Juego1(ctx context.Context, movidajugador *pb.Playermove ) (*pb.Status , error){
	//ya se escogio el numero random por lo tanto haremos el check
	if movidajugador.Move >= juego1 {
		//cuando cambie el estado devolvemos la variable
		siguientepaso = 0
		//retornamos el estado del jugador
		retorno := &pb.Status{
			Status: false,
		  }
		return retorno, nil
	} else{
		//cuando cambie el estado devolvemos la variable
		siguientepaso = 0
		//retornamos el estado del jugador
		retorno := &pb.Status{
			Status: true,
		  }
		return retorno, nil
	}

}


func main() {
	//Iniciamos el servidor del del lider
	go InitServer(port)

	//Inicio de la interface
	fmt.Println("Inicio el juego")

	var estadodeljuego int = 0

	//iniciamos la semilla
	rand.Seed(time.Now().UnixNano())

	// matenemos la interfaz andando 
	for estadodeljuego < 5{
		var decision int
		
		if estadodeljuego == 0{
			//realizamos la primera orden
			fmt.Println("Desea iniciar el juego?")
			fmt.Println("[1] Si")
			fmt.Println("[2] No")
			fmt.Scan(&decision)
			if decision == 1{
				estadodeljuego += 1
				siguientepaso = 1
			} else {
				fmt.Println("Comando no reconocido")
			}
		} else if estadodeljuego == 1{
			//Esperamos a todos los jugadores

			//Se ecoge el numero aleatorio
			juego1 := rand.Intn(10 - 6 + 1) + 6
			fmt.Println("numero de la ronda: %v",juego1)

		}

	}

	
}



