package main

import (
	pb "Lider/proto"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var siguientepaso int = 0
var esperandoporjugadores int = 16
var Respuestajuego1 int32 = 0

var siguientejuego int32 = 0
var respuestaLider int32 = 0

var round int32 = 1

// Server
type server struct {
	pb.UnimplementedStartServerServer
}

type jugador struct {
	numero int
	state  bool
	suma   int32
}

var jugadores []jugador

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

func (s *server) AgregarJugador(ctx context.Context, jugador *pb.Name) (*pb.Status, error) {
	fmt.Println("El jugador", jugador.Name, "fue anadido al juego")

	//cuando cambie el estado devolvemos la variable
	siguientepaso = 0
	//retornamos el estado del jugador
	retorno := &pb.Status{
		Status: true,
	}
	return retorno, nil
}

func (s *server) Siguientejuego(ctx context.Context, jugador *pb.Name) (*pb.Nextgame, error) {
	fmt.Println("Un jugador esta esperando empezar a jugar")
	//retornamos la orden del juego
	retorno := &pb.Nextgame{
		Answer: siguientejuego,
	}
	return retorno, nil
}

func (s *server) SeSolicitoPozo(ctx context.Context, name *pb.Name) (*pb.Amount, error) {
	fmt.Println("El jugador %v solicitio el valor del pozo", name.Name)

	//Conneccion al pozo
	connPozo, errLider := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if errLider != nil {
		log.Fatalf("did not connect: %v", errLider)
	}
	defer connPozo.Close()
	//retornamos la instancia con el lider
	cPozo := pb.NewStartServerClient(connPozo)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := cPozo.PedirPozo(ctx, &pb.Msg{Message: "El Lider esta pidiendo el valor del pozo"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r, nil

}

func (s *server) EstadoLider(ctx context.Context, name *pb.Name) (*pb.Status, error) {
	fmt.Println("Estado solicitado por", name.Name)

	if respuestaLider == 0 {
		retorno := &pb.Status{
			Status: false,
		}
		return retorno, nil
	} else {
		respuestaLider = 0
		retorno := &pb.Status{
			Status: true,
		}
		return retorno, nil

	}

}

func (s *server) MandarALider(ctx context.Context, movidajugador *pb.Playermove) (*pb.Status, error) {

	if jugadores[0].numero != 1 {
		movidajugador.Move = 11
	}
	fmt.Println("La movida del jugador fue:", movidajugador.Move)
	jugadores[0].suma += movidajugador.Move
	//ya se escogio el numero random por lo tanto haremos el check
	if movidajugador.Move >= Respuestajuego1 {

		if movidajugador.Move != 11 {
			jugadores[0].state = false
		}
		//retornamos el estado del jugador
		retorno := &pb.Status{
			Status: jugadores[0].state,
		}

		for i := 0; i < len(jugadores); i++ {
			if jugadores[i].suma < 21 {
				numeroBot := rand.Int31n(8) + 3
				jugadores[i].suma = +numeroBot
				fmt.Println("El numero escogido por ", jugadores[i].numero, "es", numeroBot)
				if numeroBot >= Respuestajuego1 {
					fmt.Println("Se actualizo jugador: ", jugadores[i].numero)
					jugadores[i].state = false
				}
			}
		}

		Respuestajuego1 = 0
		round += 1

		return retorno, nil
	} else {

		if round == 4 {
			if jugadores[0].suma < 21 {
				jugadores[0].state = false
			}
		}

		fmt.Println("ronda: ", round)
		//retornamos el estado del jugador
		retorno := &pb.Status{
			Status: jugadores[0].state,
		}

		for i := 1; i < len(jugadores); i++ {
			if jugadores[i].suma < 21 {
				numeroBot := rand.Int31n(7) + 3
				jugadores[i].suma = +numeroBot
				fmt.Println("El numero escogido por ", jugadores[i].numero, "es", numeroBot)
				if numeroBot >= Respuestajuego1 {
					fmt.Println("Se actualizo jugador: ", jugadores[i].numero)
					jugadores[i].state = false
				}
			}
		}
		if round == 4 {
			fmt.Println("Se revisaran si son mayor o igual a 21")
			for i := 1; i < len(jugadores); i++ {
				if jugadores[i].suma < 21 {
					fmt.Println("Se actualizo jugador: ", jugadores[i].numero)
					jugadores[i].state = false
				}
			}
		}

		Respuestajuego1 = 0
		round += 1

		return retorno, nil
	}

}

func removeJugador(jugadores []jugador, index int) []jugador {
	return append(jugadores[:index], jugadores[index+1:]...)
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
	for estadodeljuego < 5 {
		var decision int
		if estadodeljuego == 0 {
			//realizamos la primera orden
			fmt.Println("Desea iniciar el juego?")
			fmt.Println("[1] Si")
			fmt.Println("[2] No")
			fmt.Scan(&decision)
			if decision == 1 {
				estadodeljuego += 1
				siguientejuego = 1
				//Creamos nuestra lista de jugadores, con (id_jugador, estado = vivo)
				for i := 1; i <= 16; i++ {
					jugadores = append(jugadores, jugador{i, true, 0})
				}
			} else {
				fmt.Println("Comando no reconocido")

			}
		} else if estadodeljuego == 1 {
			//actualizar el valor de la respuesta
			if Respuestajuego1 == 0 {
				//actualizamos el valor random
				Respuestajuego1 = rand.Int31n(5) + 6

				//el Lider respondio
				respuestaLider = 1
				//esperamos un tiempito jiji
				//time.Sleep(5 * time.Second)
				for i := 0; i < len(jugadores); i++ {
					if !(jugadores[i].state) {
						fmt.Println("Jugador ", jugadores[i].numero, "Ha muerto")
						jugadores = removeJugador(jugadores, i)
						i--
					}
				}
				for i := 0; i < len(jugadores); i++ {
					fmt.Println("Jugador vivo: ", jugadores[i].numero, jugadores[i].state)
				}
				if round > 5 {
					estadodeljuego += 1
				}

			} else {
				fmt.Println("stop, numero lider: ", Respuestajuego1)
				time.Sleep(5 * time.Second)
			}

		}

	}

}
