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
var conn   = ":50051"
var connS1 = ":50052"
var connS2 = ":50053"
var connS3 = ":50054"

type dats struct {
	direc string
}

var map_random map[int]dats

type server struct{
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
	if err!= nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient1 := pb.NewFuncionesServiceClient(conn1)

	res,err := serviceClient1.Inf_BroServCoord(context.Background(), &pb.BrokerServidorCoord{
		Planeta: planet,
	})
	if err!= nil {
		panic("Error en mensaje de servidor " + err.Error())
	}
	return res.X, res.Y, res.Z
}

func (s *server) InfBro(ctx context.Context, req *pb.InformanteBroker) (*pb.BrokerInformante, error){
	map_random = make(map[int]dats)
	cont := 0

	x1,y1,z1 := con_serv(req.PlanetaAfectado,connS1)
	if (x1 >= req.X) && (y1 >= req.Y) && (z1 >= req.Z){
		map_random[cont] = dats{connS1}
		cont ++
	}
	// x2,y2,z2 := con_serv(req.PlanetaAfectado,connS2)
	// if and (x2 >= req.X) (y2 >= req.Y) (z2 >= req.Z){
	// 	map_random[cont] = dats{connS2}
	// 	cont ++
	// }
	// x3,y3,z3 := con_serv(req.PlanetaAfectado,connS3)
	// if and (x3 >= req.X) (y3 >= req.Y) (z3 >= req.Z){
	// 	map_random[cont] = dats{connS3}
	// 	cont ++
	// }

	onlyOnce.Do(func(){
		rand.Seed(time.Now().UnixNano())
	})

	return &pb.BrokerInformante{
		Direccion: map_random[rand.Intn(len(map_random))].direc,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp",conn)
	
	if err != nil {
		panic ("No se pudo crear coneccion tcp de servidor broker " + err.Error())
	}
	
	serv := grpc.NewServer() 
	pb.RegisterFuncionesServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("No se pudo inicializar servidor broker " + err.Error())
	}
}

