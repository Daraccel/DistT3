package main

import (
	"fmt"
	"context"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

var conn_brok = "localhost:50051"

func menu() (string, string, string, string) {
	var accion_int int
	var accion_str1 string
	var accion_str2 string
	var accion_str3 string
	var accion_str4 string

	fmt.Println("Ingrese numero de accion a realizar. \n 1: AddCity \n 2: UpdateName \n 3: UpdateNumber \n 4: DeleteCity")
	fmt.Scan(&accion_int)
	switch {
	case accion_int == 1:
		accion_str1 = "AddCity"
	case accion_int == 2:
		accion_str1 = "UpdateName"
	case accion_int == 3:
		accion_str1 = "UpdateNumber"
	case accion_int == 4:
		accion_str1 = "DeleteCity"
	}
	
	fmt.Println("Ingrese nombre de planeta:")
	fmt.Scan(&accion_str2)
	
	fmt.Println("Ingrese nombre de ciudad:")
	fmt.Scan(&accion_str3)

	if accion_int == 4{
		accion_str4 = "0"
	} else {
		fmt.Println("Ingrese nuevo valor:")
		fmt.Scan(&accion_str4)
	}

	return accion_str1, accion_str2, accion_str3, accion_str4
}

func main() {

	conn1, err := grpc.Dial(conn_brok, grpc.WithInsecure())
	if err!= nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}
	
	serviceClient := pb.NewFuncionesServiceClient(conn1)
	a1,a2,a3,a4 := menu()

	res,err := serviceClient.InfBro(context.Background(), &pb.InformanteBroker{
		Accion:				a1,
		PlanetaAfectado:	a2,
		CiudadAfectada:		a3,
		NuevoValor:			a4,
	})

	if err!= nil {
		panic("Error en mensaje de broker " + err.Error())
	}

	fmt.Println(res.Direccion)




	conn2, err := grpc.Dial("localhost" + res.Direccion, grpc.WithInsecure())
	if err!= nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient2 := pb.NewFuncionesServiceClient(conn2)

	res2,err2 := serviceClient2.InfServ(context.Background(), &pb.InformanteServidor{
		Accion:				a1,
		PlanetaAfectado:	a2,
		CiudadAfectada:		a3,
		NuevoValor:			a4,
	})

	if err2!= nil {
		panic("Error en mensaje de broker " + err2.Error())
	}
	fmt.Println(res2)

}