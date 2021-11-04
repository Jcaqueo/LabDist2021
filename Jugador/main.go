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
				for 2 >= 1{
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
				fmt.Println("status: ", status)
				if !status{
					eliminado = true
				} 
				fmt.Println("eliminado: ", eliminado)
				if eliminado == true {
					decision = 3
					break
				}
			}
			decision +=1 
		} else if decision == 2 {
			fmt.Println("Sobreviviste al juego")
			fmt.Println("Juego 2")
			decision +=1 
			
		} else if decision == 3 {
			fmt.Println("Sobreviviste al juego")
			fmt.Println("Juego 2")
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
