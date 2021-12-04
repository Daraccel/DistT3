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

type server struct{
	pb.UnimplementedFuncionesServiceServer
}

func server_random() string {
	onlyOnce.Do(func(){
		rand.Seed(time.Now().UnixNano())
	})

	var num int
	var str string
	num = rand.Intn(3) + 1
	// log.Println(num)
	switch {
	case num == 1:
		str = ":50052"
	case num == 2:
		str = ":50053"
	case num == 3:
		str = ":50054"
	}
	return str
}

func (s *server) InfBro(ctx context.Context, req *pb.InformanteBroker) (*pb.BrokerInformante, error){
	return &pb.BrokerInformante{
		Direccion: server_random(),
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

