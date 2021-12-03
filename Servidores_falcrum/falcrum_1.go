package main 


import {
	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
}

type server struct{
	pb.UnimplementedFuncionesServiceServer
}

func (s *server) inf_serv (context.Context, *pb.informante_servidor) (*pb.servidor_informante, err){
	
}