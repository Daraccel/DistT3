package main

import "google.golang.org/grpc"

func main() {
	conn, err :=grpc.Dial( target:"localhost50051", grpc.WithInsecure())
	if err!= nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient := pb.NewFuncionesServiceClient(conn)

	res,err : =serviceClient.Create(context.Background(), &pb.InformanteBroker {
		Accion:				"AddCity",
		PlanetaAfectado:	"Tu mama",
		CiudadAfectada:		"este",
		Nuevo_Valor:		"69",
	})

	if err!= nil {
		panic("Error en mensaje de broker " + err.Error())
	}

	fmt.Println(res)
}