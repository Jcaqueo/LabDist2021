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

var JugadorPos jugadorPos

var jugadores []jugador
var jugadores2 []jugador

 
// respuesta que da el Lider a los juegos
var Respuestajuego int32 = 0
//variable que se le dara al jugador para esperar el siguiente juego
var siguientejuego int32 = 0
//Si el lider ya respondio, nos yuda a coodinar el juegador con el Lider
var respuestaLider int32 = 0
//saber si el jugador respondido
var respuesta bool = true;


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

type jugadorPos struct {
	index int
	lista int
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

func (s *server) AgregarJugador(ctx context.Context, jugador *pb.Name) (*pb.Status, error) {
	fmt.Println("El jugador", jugador.Name, "fue anadido al juego")

	//cuando cambie el estado devolvemos la variable
	siguientepaso = 0
	//retornamos el estado del jugador
	retorno := &pb.Status{
		Status: true,
		Time : 0,
	}
	return retorno, nil
}

func (s *server) Siguientejuego(ctx context.Context, jugador *pb.Name) (*pb.Nextgame, error) {
	//fmt.Println("Un jugador esta esperando empezar a jugar")
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
			Time : 0,
		}
		return retorno, nil
	} else {
		respuestaLider = 0
		retorno := &pb.Status{
			Status: true,
			Time : 0,
		}
		return retorno, nil

	}

}

//por cada juego tenemos 2 funciones
// actualizar cada bot
// actualizar al jugador

//func Respuestajugadorjuego2

// func Juegobot1 (jugadores,int) (jugadores){
	//checkeoi si hay un jugador humano
	//si -> me lo salto y hago logica del juego
	//no -> misma logica del juego
//}

func Juego1Bot(jugadores []jugador, numeroLider int32) ([]jugador){

	if jugadores[0].numero == 1{    //Jugador (cliente) sigue vivo
		for i := 1; i < len(jugadores); i++ {  
			if jugadores[i].suma < 21 {
				numeroBot := rand.Int31n(7) + 3
				jugadores[i].suma = +numeroBot
				//fmt.Println("El numero escogido por ", jugadores[i].numero, "es", numeroBot)
				if numeroBot >= numeroLider {
					//fmt.Println("Se actualizo jugador: ", jugadores[i].numero)
					jugadores[i].state = false
				}
			}
		}
		
	} else{                     //Jugador (cliente) esta muerto por lo que solo quedan bots
		for i := 0; i < len(jugadores); i++ {
			if jugadores[i].suma < 21 {
				numeroBot := rand.Int31n(8) + 3
				jugadores[i].suma = +numeroBot
				//fmt.Println("El numero escogido por ", jugadores[i].numero, "es", numeroBot)
				if numeroBot >= numeroLider {
					//fmt.Println("Se actualizo jugador: ", jugadores[i].numero)
					jugadores[i].state = false
				}
			}
		}
		// aca <- respuestajuego = 0
	fmt.Println("Ronda sin el Jugador")
	Respuestajuego = 0
	round += 1
	}
	if round == 4 { // Esto tambien -> juego bot1
		fmt.Println("Se revisaran si son mayor o igual a 21")
		for i := 1; i < len(jugadores); i++ {
			if jugadores[i].suma < 21 {
				//fmt.Println("Se actualizo jugador: ", jugadores[i].numero)
				jugadores[i].state = false
			}
		}
	}
	return jugadores

}

func EmparejarJuego2() (){
	if len(jugadores)%2 == 1{
		index_eliminar := rand.Intn(len(jugadores))
		fmt.Println("Lista impar, se elimino al jugador",index_eliminar)
		removeJugador(jugadores, index_eliminar)
	}
	
	j2 := len(jugadores)

	for i := 1; i <= j2/2; i++{
		index_cambio := rand.Intn(len(jugadores))
		jugadores2 = append(jugadores2, jugadores[index_cambio])
		removeJugador(jugadores, index_cambio)
	}

}

func Juego2(valorLider int32) {
	var suma_eq1 int32
	var suma_eq2 int32

	for i := 0; i<len(jugadores); i++{
		suma_eq1 += jugadores[i].suma
	}
	for i := 0; i<len(jugadores2); i++{
		suma_eq2 += jugadores2[i].suma
	}
	
		
	if (suma_eq1 % 2) == (valorLider % 2) && (suma_eq2 % 2) != (valorLider % 2){ //Gana equipo 1
		for i :=0; i<len(jugadores2); i++{
			jugadores2[i].state = false
		}
		fmt.Println("Gano equipo 1")
		if JugadorPos.lista == 2 {
			JugadorPos.lista = 1
			JugadorPos.index += len(jugadores)
		}

		for i :=0; i<len(jugadores2); i++{
			jugadores = append(jugadores, jugadores2[i])
		}

	} else if (suma_eq1 % 2) != (valorLider % 2) && (suma_eq2 % 2) == (valorLider % 2){ //Gana equipo 2
		for i :=0; i<len(jugadores); i++{
			jugadores[i].state = false
		}
		fmt.Println("Gano equipo 2")
		if JugadorPos.lista == 2 {
			JugadorPos.lista = 1
			JugadorPos.index += len(jugadores)
		}

		for i :=0; i<len(jugadores2); i++{
			jugadores = append(jugadores, jugadores2[i])
		}

	} else if (suma_eq1 % 2) == (valorLider % 2) && (suma_eq2 % 2) == (valorLider % 2){ //Ganan ambos
		if JugadorPos.lista == 2 {
			JugadorPos.lista = 1
			JugadorPos.index += len(jugadores)
		}
		fmt.Println("Ganaron ambos")

		for i :=0; i<len(jugadores2); i++{
			jugadores = append(jugadores, jugadores2[i])
		}
	
	} else{ //Pierden ambos
		var equipoG = rand.Intn(2) + 1
		if equipoG == 1{
			for i :=0; i<len(jugadores2); i++{
				jugadores2[i].state = false
			}
		} else{
			for i :=0; i<len(jugadores); i++{
				jugadores[i].state = false
			}
		}
		fmt.Println("Perdieron ambos, pasa el equipo: ",equipoG)
		if JugadorPos.lista == 2 {
			JugadorPos.lista = 1
			JugadorPos.index += len(jugadores)
		}

		for i :=0; i<len(jugadores2); i++{
			jugadores = append(jugadores, jugadores2[i])
		}

	}
}


// esta funcion es llamada por el cliente
//si el cliente muere -> nunca se llama
func (s *server) MandarALider(ctx context.Context, movidajugador *pb.Playermove) (*pb.Status, error) {

	fmt.Println("La movida del jugador fue:", movidajugador.Move)
	jugadores[0].suma += movidajugador.Move

	//ya se escogio el numero random por lo tanto haremos el check
	//si el jugador suma mas de 21 se salvo del juego
	if (jugadores[0].suma >= 21){
		//retornamos que el jugador esta vivo
		retorno := &pb.Status{
			Status: jugadores[0].state,
			Time : 0,
		}
		Respuestajuego = 0
		round += 1
		return retorno, nil
	//si estamos en la ronda 4 y no llegamos a 21 significa que nuestro jugador murio
	} else if round == 4 {  
		//retornamos que el jugador esta muerto
		//actualizamos su estado
		jugadores[0].state = false	
		//creamos el paquete
		retorno := &pb.Status{
			Status: jugadores[0].state,
			Time : 0,
		}
		Respuestajuego = 0
		round += 1
		return retorno, nil
	}

	if movidajugador.Move >= Respuestajuego {
		jugadores[0].state = false		
		//retornamos el estado del jugador
		retorno := &pb.Status{
			Status: jugadores[0].state,
			Time : 0,
		}
		Respuestajuego = 0
		round += 1
		return retorno, nil
	} else {
		fmt.Println("ronda: ", round)
		//retornamos el estado del jugador
		retorno := &pb.Status{
			Status: jugadores[0].state,
			Time : 0,
		}
		Respuestajuego = 0
		round += 1
		return retorno, nil
	}

}

func (s *server) MandarALider2(ctx context.Context, movidajugador *pb.Playermove) (*pb.Msg, error) {

	fmt.Println("La movida del jugador fue:", movidajugador.Move)
	//actualizamos la respuesta del jugador en la lista correspondiente
	if JugadorPos.lista == 1{
		jugadores[JugadorPos.index].suma = movidajugador.Move
	} else { 
		jugadores2[JugadorPos.index].suma = movidajugador.Move
	}

	retorno := &pb.Msg{
		Message: "Respuesta recibida",
	}
	//el jugador respondio -> actualizamos la variable
	respuesta = false
	return retorno, nil


}

func (s *server) RetornarEstado(ctx context.Context, Mensajejugador *pb.Msg) (*pb.Status, error) {

	fmt.Println(Mensajejugador)
	if respuesta {
		retorno := &pb.Status{
			Status: jugadores[JugadorPos.index].state,
			Time : 1,
	
		}
		return retorno, nil
	} else {
		retorno := &pb.Status{
			Status: false,
			Time : 0,
	
		}
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
		} else if estadodeljuego == 1{
			//actualizar el valor de la respuesta
			if Respuestajuego == 0 {
				//actualizamos el valor random
				Respuestajuego = rand.Int31n(5) + 6

				//cambiamos la respuesta del juego 1
				//Respuestajuego = 1
				//el Lider respondio
				respuestaLider = 1

				// aca jugar a los bots
				//eliminamos cada ronda eliminamos a los jugadores muertos

				

				for i := 0; i < len(jugadores); i++ {
					if !(jugadores[i].state) {
						fmt.Println("Jugador ", jugadores[i].numero, "Ha muerto")
						jugadores = removeJugador(jugadores, i)
						i--
					}
				}

				//aca miramos si aun tenemos un jugador
				if len(jugadores) == 0{
					fmt.Println("Todos los jugadores han fallecido")
					return
				}
				
				jugadores = Juego1Bot(jugadores, Respuestajuego)

				//Mostramos a los jugadores vivos?
				//no es necesarios, pero para debug es util
				for i := 0; i < len(jugadores); i++ {
					fmt.Println("Jugador vivo: ", jugadores[i].numero, jugadores[i].state)
				}
				
				//si pasamos a la ronda 5 termina el juego
				if round > 5 {
					estadodeljuego += 1
				}

			} else {
				//jugador juega aca
				fmt.Println("stop, numero lider: ", Respuestajuego)
				time.Sleep(5 * time.Second)
			}

		} else if estadodeljuego == 2{
			//lista con todos los jugadores que pasaron el juego 1
			fmt.Println("Juego 2")
			//jugador vivo para esta ronde
			var jugadorvivo bool = false;
			//generamos el equipo 2
			
			

			//cambiar
			EmparejarJuego2()

			//generar las respuestas para el equipo 1
			for i := 0; i < len(jugadores); i++ {
				// != 1 -> bot
				jogador := jugadores[i]
				if jogador.numero != 1 {
					jogador.suma = (rand.Int31n(4)+1)
				} else {
					JugadorPos.lista = 1
					JugadorPos.index = i
					jugadorvivo =true
				}
			}

			//generar las respuestas para el equipo 2
			for i := 0; i < len(jugadores2); i++ {
				// != 1 -> bot
				jogador := jugadores[i]
				if jogador.numero != 1 {
					jogador.suma = (rand.Int31n(4)+1)
				} else{
					JugadorPos.lista = 2
					JugadorPos.index = i
					jugadorvivo =true
				}
			}

			
			if jugadorvivo {
				//esperar que el jugador responda
				for respuesta == true{

				}
			}

			eleccionLider := rand.Int31n(2)+1
			Juego2(eleccionLider)

			//if innecesario
			if jugadorvivo {
				respuesta = true 
			}
			
			for i := 0; i < len(jugadores); i++ {
				if !(jugadores[i].state) {
					fmt.Println("Jugador ", jugadores[i].numero, "Ha muerto")
					jugadores = removeJugador(jugadores, i)
					i--
				}
			}
			//Mostramos los sobrevivientes
			for i := 0; i < len(jugadores); i++ {
				fmt.Println("Jugador vivo: ", jugadores[i].numero, jugadores[i].state)
			}

			return
		}

	}

}
