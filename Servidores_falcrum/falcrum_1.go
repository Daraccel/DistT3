package main 


import (
	"context"
	// "fmt"
	// "log"
	// "math/rand"
	"net"
	// "sync"
	// "time"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

var conn = ":50052"

type server struct{
	pb.UnimplementedFuncionesServiceServer
}


// Por modificar esto, x,y,z es por archivo y aumenta con cada uso
func (s *server) InfServ(ctx context.Context, req *pb.InformanteServidor) (*pb.ServidorInformante, error){
	return &pb.ServidorInformante{
		X: 1,
		Y: 1,
		Z: 1,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp",conn)
	
	if err != nil {
		panic ("No se pudo crear coneccion tcp de servidor serv 1 " + err.Error())
	}
	
	serv := grpc.NewServer() 
	pb.RegisterFuncionesServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("No se pudo inicializar servidor serv 1 " + err.Error())
	}

}