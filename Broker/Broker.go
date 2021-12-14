package main

import (
	"context"
	// "fmt"
	// "log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

var onlyOnce sync.Once
var conn = ":50051"
var connS1 = ":50052"
var connS2 = ":50053"
var connS3 = ":50054"

type dats struct {
	direc string
}

type infoplanet struct {
	rebels              int32
	cordx, cordy, cordz int32
	server              string
}

var map_planet map[int]infoplanet
var map_random map[int]dats

type server struct {
	pb.UnimplementedFuncionesServiceServer
}

// func server_random() string {
// 	onlyOnce.Do(func(){
// 		rand.Seed(time.Now().UnixNano())
// 	})

// 	// var num int
// 	var str string
// 	// num = rand.Intn(3) + 1
// 	// log.Println(num)
// 	// switch {
// 	// case num == 1:
// 	// 	str = ":50052"
// 	// case num == 2:
// 	// 	str = ":50053"
// 	// case num == 3:
// 	// 	str = ":50054"
// 	// }
// 	str = connS1
// 	return str
// }

func con_serv(planet string, dir string) (int32, int32, int32) {
	conn1, err := grpc.Dial(dir, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient1 := pb.NewFuncionesServiceClient(conn1)

	res, err := serviceClient1.Inf_BroServCoord(context.Background(), &pb.BrokerServidorCoord{
		Planeta: planet,
	})
	if err != nil {
		panic("Error en mensaje de servidor " + err.Error())
	}
	return res.X, res.Y, res.Z
}

//Conexion fulcrum broker desde leia      //retorna rebeldes, reloj vectorial, serverName
func con_servB(planet string, ciudad string, dir string) (int32, int32, int32, int32, string) {
	conn1, err := grpc.Dial(dir, grpc.WithInsecure())
	if err != nil {
		panic("Fallo en la conexion con el server, esto debe ser Obra del Imperio" + err.Error())
	}
	serviceClient2 := pb.NewFuncionesServiceClient(conn1)

	res, err := serviceClient2.InfBroful(context.Background(), &pb.BrokerFulcrum{
		Nombreplaneta: planet,
		Nombreciudad:  ciudad,
	})

	if err != nil {
		panic("Error en mensaje de servidor, el imperio debe estar detras de esto " + err.Error())
	}

	return res.Rebeldes, res.X, res.Y, res.Z, res.Nombreserver
}

func (s *server) InfBrolei(ctx context.Context, req *pb.LeiaBroker) (*pb.BrokerLeia, error) {
	map_planet = make(map[int]infoplanet)
	cont := 0

	rbls, x1, y1, z1, sername := con_servB(req.Planeta, req.Ciudad, connS1)
	if (x1 != 0) || (y1 != 0) || (z1 != 0) {
		map_planet[cont] = infoplanet{rbls, x1, y1, z1, sername}
		cont++
	}

	rbls, x2, y2, z2, sername := con_servB(req.Planeta, req.Ciudad, connS2)
	if (x2 != 0) || (y2 != 0) || (z2 != 0) {
		map_planet[cont] = infoplanet{rbls, x2, y2, z2, sername}
		cont++
	}
	rbls, x3, y3, z3, sername := con_servB(req.Planeta, req.Ciudad, connS3)
	if (x3 != 0) || (y3 != 0) || (z3 != 0) {
		map_planet[cont] = infoplanet{rbls, x3, y3, z3, sername}
		cont++
	}
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	eleccion := rand.Intn(len(map_planet))

	return &pb.BrokerLeia{

		Rebeldes: map_planet[eleccion].rebels,
		X:        map_planet[eleccion].cordx,
		Y:        map_planet[eleccion].cordy,
		Z:        map_planet[eleccion].cordz,
		Sfulcrum: map_planet[eleccion].server,
	}, nil
}

func (s *server) InfBro(ctx context.Context, req *pb.InformanteBroker) (*pb.BrokerInformante, error) {
	map_random = make(map[int]dats) //mapa temporal de esta parte
	cont := 0                       //Contador

	x1, y1, z1 := con_serv(req.PlanetaAfectado, connS1)
	if (x1 >= req.X) && (y1 >= req.Y) && (z1 >= req.Z) {
		map_random[cont] = dats{connS1}
		cont++
	}
	x2, y2, z2 := con_serv(req.PlanetaAfectado, connS2)
	if (x2 >= req.X) && (y2 >= req.Y) && (z2 >= req.Z) {
		map_random[cont] = dats{connS2}
		cont++
	}
	x3, y3, z3 := con_serv(req.PlanetaAfectado, connS3)
	if (x3 >= req.X) && (y3 >= req.Y) && (z3 >= req.Z) {
		map_random[cont] = dats{connS3}
		cont++
	}

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return &pb.BrokerInformante{
		Direccion: map_random[rand.Intn(len(map_random))].direc,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", conn)

	if err != nil {
		panic("No se pudo crear coneccion tcp de servidor broker " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterFuncionesServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("No se pudo inicializar servidor broker " + err.Error())
	}
}
