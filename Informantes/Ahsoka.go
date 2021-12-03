func main() {
	conn, err :=grpc.Dial( target:"localhost50051", grpc.WithInsecure())
	if err!= nil {
		panic("No se pudo conectar con el servidor " + err.Error())
	}

	serviceClient := pb.NewFuncionesServiceClient(conn)

	serviceClient.Create(context.Background)(, )
}