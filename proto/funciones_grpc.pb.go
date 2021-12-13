// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package DistT3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FuncionesServiceClient is the client API for FuncionesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FuncionesServiceClient interface {
	InfBro(ctx context.Context, in *InformanteBroker, opts ...grpc.CallOption) (*BrokerInformante, error)
	InfServ(ctx context.Context, in *InformanteServidor, opts ...grpc.CallOption) (*ServidorInformante, error)
	Inf_BroServCoord(ctx context.Context, in *BrokerServidorCoord, opts ...grpc.CallOption) (*ServidorBrokerCoord, error)
	InfBrolei(ctx context.Context, in *LeiaBroker, opts ...grpc.CallOption) (*BrokerLeia, error)
	InfBroful(ctx context.Context, in *BrokerFulcrum, opts ...grpc.CallOption) (*FulcrumBroker, error)
}

type funcionesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFuncionesServiceClient(cc grpc.ClientConnInterface) FuncionesServiceClient {
	return &funcionesServiceClient{cc}
}

func (c *funcionesServiceClient) InfBro(ctx context.Context, in *InformanteBroker, opts ...grpc.CallOption) (*BrokerInformante, error) {
	out := new(BrokerInformante)
	err := c.cc.Invoke(ctx, "/grpc.FuncionesService/inf_bro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *funcionesServiceClient) InfServ(ctx context.Context, in *InformanteServidor, opts ...grpc.CallOption) (*ServidorInformante, error) {
	out := new(ServidorInformante)
	err := c.cc.Invoke(ctx, "/grpc.FuncionesService/inf_serv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *funcionesServiceClient) Inf_BroServCoord(ctx context.Context, in *BrokerServidorCoord, opts ...grpc.CallOption) (*ServidorBrokerCoord, error) {
	out := new(ServidorBrokerCoord)
	err := c.cc.Invoke(ctx, "/grpc.FuncionesService/inf_BroServCoord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *funcionesServiceClient) InfBrolei(ctx context.Context, in *LeiaBroker, opts ...grpc.CallOption) (*BrokerLeia, error) {
	out := new(BrokerLeia)
	err := c.cc.Invoke(ctx, "/grpc.FuncionesService/inf_brolei", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *funcionesServiceClient) InfBroful(ctx context.Context, in *BrokerFulcrum, opts ...grpc.CallOption) (*FulcrumBroker, error) {
	out := new(FulcrumBroker)
	err := c.cc.Invoke(ctx, "/grpc.FuncionesService/inf_broful", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FuncionesServiceServer is the server API for FuncionesService service.
// All implementations must embed UnimplementedFuncionesServiceServer
// for forward compatibility
type FuncionesServiceServer interface {
	InfBro(context.Context, *InformanteBroker) (*BrokerInformante, error)
	InfServ(context.Context, *InformanteServidor) (*ServidorInformante, error)
	Inf_BroServCoord(context.Context, *BrokerServidorCoord) (*ServidorBrokerCoord, error)
	InfBrolei(context.Context, *LeiaBroker) (*BrokerLeia, error)
	InfBroful(context.Context, *BrokerFulcrum) (*FulcrumBroker, error)
	mustEmbedUnimplementedFuncionesServiceServer()
}

// UnimplementedFuncionesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFuncionesServiceServer struct {
}

func (UnimplementedFuncionesServiceServer) InfBro(context.Context, *InformanteBroker) (*BrokerInformante, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfBro not implemented")
}
func (UnimplementedFuncionesServiceServer) InfServ(context.Context, *InformanteServidor) (*ServidorInformante, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfServ not implemented")
}
func (UnimplementedFuncionesServiceServer) Inf_BroServCoord(context.Context, *BrokerServidorCoord) (*ServidorBrokerCoord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inf_BroServCoord not implemented")
}
func (UnimplementedFuncionesServiceServer) InfBrolei(context.Context, *LeiaBroker) (*BrokerLeia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfBrolei not implemented")
}
func (UnimplementedFuncionesServiceServer) InfBroful(context.Context, *BrokerFulcrum) (*FulcrumBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfBroful not implemented")
}
func (UnimplementedFuncionesServiceServer) mustEmbedUnimplementedFuncionesServiceServer() {}

// UnsafeFuncionesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FuncionesServiceServer will
// result in compilation errors.
type UnsafeFuncionesServiceServer interface {
	mustEmbedUnimplementedFuncionesServiceServer()
}

func RegisterFuncionesServiceServer(s grpc.ServiceRegistrar, srv FuncionesServiceServer) {
	s.RegisterService(&FuncionesService_ServiceDesc, srv)
}

func _FuncionesService_InfBro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InformanteBroker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuncionesServiceServer).InfBro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FuncionesService/inf_bro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuncionesServiceServer).InfBro(ctx, req.(*InformanteBroker))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuncionesService_InfServ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InformanteServidor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuncionesServiceServer).InfServ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FuncionesService/inf_serv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuncionesServiceServer).InfServ(ctx, req.(*InformanteServidor))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuncionesService_Inf_BroServCoord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerServidorCoord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuncionesServiceServer).Inf_BroServCoord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FuncionesService/inf_BroServCoord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuncionesServiceServer).Inf_BroServCoord(ctx, req.(*BrokerServidorCoord))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuncionesService_InfBrolei_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeiaBroker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuncionesServiceServer).InfBrolei(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FuncionesService/inf_brolei",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuncionesServiceServer).InfBrolei(ctx, req.(*LeiaBroker))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuncionesService_InfBroful_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerFulcrum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuncionesServiceServer).InfBroful(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.FuncionesService/inf_broful",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuncionesServiceServer).InfBroful(ctx, req.(*BrokerFulcrum))
	}
	return interceptor(ctx, in, info, handler)
}

// FuncionesService_ServiceDesc is the grpc.ServiceDesc for FuncionesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FuncionesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.FuncionesService",
	HandlerType: (*FuncionesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "inf_bro",
			Handler:    _FuncionesService_InfBro_Handler,
		},
		{
			MethodName: "inf_serv",
			Handler:    _FuncionesService_InfServ_Handler,
		},
		{
			MethodName: "inf_BroServCoord",
			Handler:    _FuncionesService_Inf_BroServCoord_Handler,
		},
		{
			MethodName: "inf_brolei",
			Handler:    _FuncionesService_InfBrolei_Handler,
		},
		{
			MethodName: "inf_broful",
			Handler:    _FuncionesService_InfBroful_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "funciones.proto",
}
