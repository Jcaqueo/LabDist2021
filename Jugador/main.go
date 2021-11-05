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
var veredicto int32 = 1
var eliminado = false

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
	log.Printf("%v",r.Answer)
	if r.Answer == 1{
		esperandoorden = 0
	} else{
		esperandoorden = 1
	}

	return 

}

func PedirEstadoLider() (*pb.Status, error){
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
	r, err := cLider.EstadoLider(ctx, &pb.Name{Name: "Jugador 1",})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r,nil

}

func EnviarRespuesta(numero int32) (*pb.Status , error){

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
	r, err := cLider.MandarALider(ctx, &pb.Playermove{Move: numero,})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r,nil

}

func EnviarRespuesta2(numero int32) (){

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
	r, err := cLider.MandarALider2(ctx, &pb.Playermove{Move: numero,})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Message)
	return 

}

func RetornarEstado() (*pb.Status){

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


	r, err := cLider.RetornarEstado(ctx, &pb.Msg{Message: "Jugador 1 solicita nuevo estado"},)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r

}





func main() {

	var decision int = 0
	var respuesta int32
	//var verposo int32

	fmt.Println("Jugador 1, Bienvenido al Juego del Calamar")
	fmt.Println("[1] Para ingresar al juego")
	fmt.Scanln(&decision)

	Conect()


	for !eliminado {
		if decision == 1 {
			for esperandoorden >= 1{
				//estamos esperando orden
				//llamamos a la funcion
				PedirinicioDejuego()
				
				fmt.Println("Esperando que inicie la ronda 1")
				time.Sleep(1 * time.Second)
			}
			//Juego 1
			fmt.Println("Comienza el juego número 1") // dentro de decision == 2
			for ronda := 1; ronda <= 4; ronda++ {
				fmt.Print("Jugador 1, Ingrese su número: ")
				fmt.Scanln(&respuesta)
				//esperamos que el Lider responda 
				for {
					fmt.Println("Esperando a que el líder de su respuesta")
					//time.Sleep(10 * time.Second)
					r,_ := PedirEstadoLider()
					if r.Status{
				    	fmt.Println("Respondio")
						break
						}
				}
				//Enviar respuesta jugador
				s,_ := EnviarRespuesta(respuesta)
				//retornar true or false
				status := s.Status//func que recibe la data
				//almacenamos el estado
				//fmt.Println("status: ", status)
				//si ganaste el juego
				// if s.ganador == true {
				// 	fmt.Println("Felicidades ganaste el juego!")
				// }
				//status -> false moriste
				if !status{
					eliminado = true
					fmt.Println("eliminado: ", eliminado)
					decision = 3
					break
				} 
			}
			decision +=1 
		} else if decision == 2 {
			//esperando veredicto
			for {
				//estamos esperando orden
				//llamamos a la funcion
				s := RetornarEstado()
				fmt.Println("Esperando veredicto del juego")
				//miramos si ganaste el juego
				fmt.Println(s)
				if s.Time == 1 {
					if s.Ganador == true{
						fmt.Println("Felicidades Ganaste el Juego")
						return
					} else if s.Status == false{
						fmt.Println("Fusite eliminado del juego")
						return
					}
					break
				}
				
				fmt.Println("Esperando veredicto del juego")
				time.Sleep(1 * time.Second)
			}

			fmt.Println("Sobreviviste al juego")

			fmt.Println("Juego 2")
			//esperar a que el lider inicie la ronda del juego
			esperandoorden = 1
			for esperandoorden >= 1{
				//estamos esperando orden
				//llamamos a la funcion
				PedirinicioDejuego()
				
				fmt.Println("Esperando que inicie la ronda 2")
				time.Sleep(1 * time.Second)
			}

			//Mandamos al respuesta del juego 2
			fmt.Print("Jugador 1, Ingrese su número: ")
			fmt.Scanln(&respuesta)
			EnviarRespuesta2(respuesta)
			//esperamos que el lider nos mande nuestro nuevo estado
			for {
				s:= RetornarEstado()
				if s.Time == 1{
					//miramos si el jugador esta muerto 
					//status = false -> muerte
					if s.Status == false{
						fmt.Println("Fuiste eliminado del juego")
						//terminamos el juego
						return
					}
					break
				}
			} 
			//si no retorna es porque estamos vivos
			decision += 1
			
		} else if decision == 3 {
			fmt.Println("Sobreviviste al juego")
			fmt.Println("Juego 3")
			decision +=1 
		}
		if decision == 4 {
			fmt.Println("Fuiste Eliminado")
			break
		}
		
		// fmt.Println("¿Quieres ver el pozo acumulado?")
		// fmt.Println("[1] Si")
		// fmt.Println("[2] No")
		// fmt.Print("Jugador 1, Decida que acción realizar: ")
		// fmt.Scanln(&verposo)
		// // Cambiarverposo
	}
	return
}
