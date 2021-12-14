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

type infoplanet struct {
	rebels              int32
	cordx, cordy, cordz int32
	server              string
}

func main() {
	//nombreCIudad : rebeldes, reloj vectorial, server de la info
	visits := make(map[string]infoplanet)

	conn1, err := grpc.Dial(conn_brok, grpc.WithInsecure()) //Conecta con el server
	if err != nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient := pb.NewFuncionesServiceClient(conn1)
	//serviceClient := pb.NewFuncionesServiceClient(":50052")
	p_name, p_siti := mensaje()

	res, err := serviceClient.InfBrolei(context.Background(), &pb.LeiaBroker{
		Planeta: p_name,
		Ciudad:  p_siti,
	})
	//res, err := serviceClient.InfBro(context.Background(), &pb.InformanteBroker{ //Envía al broker todas las acciones y recibe una respuesta en res donde le llega la direccion	})
	if err != nil {
		panic("Error en mensaje de broker " + err.Error())
	}

	fmt.Println("Rebeldes: ", res.Rebeldes)
	fmt.Println("Cordenadas: ", res.X, res.Y, res.Z)
	fmt.Println("Rebeldes: ", res.Sfulcrum)

}
