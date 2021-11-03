package main

import (
	"context"
	"fmt"
	"log"

	//"math/rand"
	"time"

	pb "Jugador/proto"

	"google.golang.org/grpc"
)

const (
	//Coneccion al Lider
	addressLider = "localhost:50051"
)

var esperandoorden int32 = 1

//(pb.StartServerClient, context.Context)
func Conect() (){
		//Conneccion al Lider
		connLider, errLider := grpc.Dial(addressLider, grpc.WithInsecure(), grpc.WithBlock())
		if errLider != nil {
			log.Fatalf("did not connect: %v", errLider)
		}
		defer connLider.Close()
		//retornamos la instancia con el lider
		cLider := pb.NewStartServerClient(connLider)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := cLider.AgregarJugador(ctx, &pb.Name{Name: "Benja",})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Fuiste añadio al juego!")

		return 

}

func PedirinicioDejuego() (){
	//Conneccion al Lider
	connLider, errLider := grpc.Dial(addressLider, grpc.WithInsecure(), grpc.WithBlock())
	if errLider != nil {
		log.Fatalf("did not connect: %v", errLider)
	}
	defer connLider.Close()
	//retornamos la instancia con el lider
	cLider := pb.NewStartServerClient(connLider)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := cLider.Siguientejuego(ctx, &pb.Name{Name: "Benja",})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//log.Printf("%v",r.Answer)
	if r.Answer == 1{
		esperandoorden = 0
	} else{
		esperandoorden = 1
	}

	return 

}

// func (s *server) AgregarJugador(ctx context.Context, jugador *pb.Name ) (*pb.Status, error){

// 	retorno := &pb.Status{
// 		Status: true,
// 	  }
// 	return retorno, nil
// }

func main() {

	var decision int = 0
	var eliminado = false
	var juego int = 1
	var respuesta int32

	fmt.Println("Jugador 1, Bienvenido al Juego del Calamar")

	for !eliminado {
		if decision == 0{
			fmt.Println("[1] Para ingresar a la ronda ", juego, " del Juego del Calamar")
			fmt.Println("[2] Para ver el monto acumulado del pozo")
			fmt.Print("Jugador 1, Decida que acción realizar: ")
			fmt.Scanln(&decision)
		}

		if decision == 1 {
			Conect()
			//esperamos que el lider empiece el juego
			for esperandoorden >= 1{
				//estamos esperando orden
				//llamamos a la funcion
				PedirinicioDejuego()
				
				time.Sleep(2 * time.Second)
				fmt.Println("Esperando que inicie el juego")
			}
			var err *int = nil
			status := true
			//status, err := cLider.AgregarJugador(ctx, &pb.AgregarJugador{Name: "Pedrito",})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			fmt.Println("Comienza el juego número ", juego) // Deberia ser un mensaje entregado por el lider
			for juego < 4 {
				if juego == 1{
					for ronda := 1; ronda <= 4; ronda++ {
						fmt.Print("Jugador 1, Ingrese su número: ")
						fmt.Scanln(&respuesta)
						if err != nil {
							log.Fatalf("could not greet: %v", err)
						}
						//almacenamos el estado
						eliminado = status 
						if eliminado == true {
							juego = 0
							break
						}
					}
				fmt.Println("Sobreviviste al primer juego xd")
				juego += 1
				decision = 0
				break

				}

			}
		}

		if decision == 2 {
			fmt.Println("El monto acumulado actualmente es ...")
		}
	}
	fmt.Println("Fuiste Eliminado")
}
