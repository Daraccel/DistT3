package main

import (
	"context"
	"fmt"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

var conn_brok = "localhost:50051"

type registro struct {
	coor_x,coor_y,coor_z int32
	direc string
}

var mapeo_datos map[string]registro

func menu() (string, string, string, string, int) {
	var accion_int int
	var accion_str1 string
	var accion_str2 string
	var accion_str3 string
	var accion_str4 string
	var correr int = 1

	fmt.Println("Ingrese numero de accion a realizar. \n 1: AddCity \n 2: UpdateName \n 3: UpdateNumber \n 4: DeleteCity \n 5: Cerrar programa")
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
	case accion_int == 5:
		correr = 2
	}

<<<<<<< HEAD
	if correr < 2{
		fmt.Println("Ingrese nombre de planeta:")
=======
	fmt.Println("Ingrese nombre de planeta:")
	fmt.Scan(&accion_str2)

	fmt.Println("Ingrese nombre de ciudad:")
	fmt.Scan(&accion_str3)

	if accion_int == 4 {
		accion_str4 = "0"
	} else {
		fmt.Println("Ingrese nuevo valor:")
>>>>>>> Cbranch
		fmt.Scan(&accion_str2)
		
		fmt.Println("Ingrese nombre de ciudad:")
		fmt.Scan(&accion_str3)

		if accion_int == 4{
			accion_str4 = "0"
		} else {
			fmt.Println("Ingrese nuevo valor:")
			fmt.Scan(&accion_str4)
		}
	}
	return accion_str1, accion_str2, accion_str3, accion_str4, correr
}

func main() {
<<<<<<< HEAD
	var valX,valY,valZ int32
	var a1,a2,a3,a4 string

	mapeo_datos= make(map[string]registro)

	var correr int = 1

	for correr < 2 {
		conn1, err := grpc.Dial(conn_brok, grpc.WithInsecure())
		if err!= nil {
			panic("No se pudo conectar con el servidor " + err.Error())
		}

		serviceClient := pb.NewFuncionesServiceClient(conn1)
		a1,a2,a3,a4,correr = menu()

		if correr < 2 {
			fmt.Println("correr es true")
			datos_map, ok := mapeo_datos[a2]
			if ok {
				valX = datos_map.coor_x
				valY = datos_map.coor_y
				valZ = datos_map.coor_z
			} else {
				valX = 0
				valY = 0
				valZ = 0
			}

			res,err := serviceClient.InfBro(context.Background(), &pb.InformanteBroker{
				Accion:				a1,
				PlanetaAfectado:	a2,
				CiudadAfectada:		a3,
				NuevoValor:			a4,
				X:					valX,
				Y:					valY,
				Z:					valZ,
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

			mapeo_datos[a2] = registro{res2.X,res2.Y,res2.Z,res.Direccion}
			
		}
	}		
}
=======

	conn1, err := grpc.Dial(conn_brok, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient := pb.NewFuncionesServiceClient(conn1)

	a1, a2, a3, a4 := menu()

	res, err := serviceClient.InfBro(context.Background(), &pb.InformanteBroker{
		Accion:          a1,
		PlanetaAfectado: a2,
		CiudadAfectada:  a3,
		NuevoValor:      a4,
	})

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
>>>>>>> Cbranch
