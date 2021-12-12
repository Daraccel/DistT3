package main

import (
	"context"
	"fmt"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

var conn_brok = "localhost:50051"

func mensaje() (string, string) {
	var n_planeta string
	var n_ciudad string

	fmt.Println("Ingrese numero de planeta. \n ")
	fmt.Scan(&n_planeta)
	fmt.Println("Ingrese numero de planeta. \n ")
	fmt.Scan(&n_ciudad)

	return n_planeta, n_ciudad
}

func main() {
					//nombreCIudad : rebeldes, reloj vectorial
	visits := make(map[string][]string)

	conn1, err := grpc.Dial(conn_brok, grpc.WithInsecure()) //Conecta con el server
	if err != nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient := pb.NewFuncionesServiceClient(conn1)
	//serviceClient := pb.NewFuncionesServiceClient(":50052")
	p_name, p_siti := mensaje()

	res, err: =serviceClient.InfBrolei(context.Background(), &pb.LeiaBroker{
		Nombre:      p_name,
		Ciudad :     p_siti,
	}) 
	
	(*BrokerLeia, error)
	//res, err := serviceClient.InfBro(context.Background(), &pb.InformanteBroker{ //Envía al broker todas las acciones y recibe una respuesta en res donde le llega la direccion	})
	if err != nil {
		panic("Error en mensaje de broker " + err.Error())
	}

	fmt.Println(res.Direccion)

	conn2, err := grpc.Dial("localhost"+res.Direccion, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient2 := pb.NewFuncionesServiceClient(conn2)

	res2, err2 := serviceClient2.InfServ(context.Background(), &pb.InformanteServidor{
		Accion:          a1,
		PlanetaAfectado: a2,
		CiudadAfectada:  a3,
		NuevoValor:      a4,
	})

	if err2 != nil {
		panic("Error en mensaje de broker " + err2.Error())
	}
	fmt.Println(res2)

}
