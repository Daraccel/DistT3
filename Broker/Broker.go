package main 


import {
	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
}

type server struct{
	pb.UnimplementedFuncionesServiceServer
}

func (s *server) inf_bro (context.Context, *pb.InformanteBroker) (*pb.BrokerInformante, err){
	return &pb.BrokerInformante{
		Direccion:"50061"
	}, nil
}

func main() {
	listner, err := net.Listen(network: "tcp", address: "50051")
	
	if err != nil {
		panic ("No se pudo crear coneccion tcp de servidor broker " + err.Error())
	}
	
	serv := grpc.NerServer() 
	pb.RegisterFuncionesServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("No se pudo inicializar servidor broker " + err)
	}
}

