// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: funciones.proto

package DistT3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InformanteBroker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accion          string `protobuf:"bytes,1,opt,name=accion,proto3" json:"accion,omitempty"`
	PlanetaAfectado string `protobuf:"bytes,2,opt,name=planeta_afectado,json=planetaAfectado,proto3" json:"planeta_afectado,omitempty"`
	CiudadAfectada  string `protobuf:"bytes,3,opt,name=ciudad_afectada,json=ciudadAfectada,proto3" json:"ciudad_afectada,omitempty"`
	NuevoValor      string `protobuf:"bytes,4,opt,name=nuevo_valor,json=nuevoValor,proto3" json:"nuevo_valor,omitempty"`
	X               int32  `protobuf:"varint,5,opt,name=x,proto3" json:"x,omitempty"`
	Y               int32  `protobuf:"varint,6,opt,name=y,proto3" json:"y,omitempty"`
	Z               int32  `protobuf:"varint,7,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *InformanteBroker) Reset() {
	*x = InformanteBroker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_funciones_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformanteBroker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformanteBroker) ProtoMessage() {}

func (x *InformanteBroker) ProtoReflect() protoreflect.Message {
	mi := &file_funciones_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformanteBroker.ProtoReflect.Descriptor instead.
func (*InformanteBroker) Descriptor() ([]byte, []int) {
	return file_funciones_proto_rawDescGZIP(), []int{0}
}

func (x *InformanteBroker) GetAccion() string {
	if x != nil {
		return x.Accion
	}
	return ""
}

func (x *InformanteBroker) GetPlanetaAfectado() string {
	if x != nil {
		return x.PlanetaAfectado
	}
	return ""
}

func (x *InformanteBroker) GetCiudadAfectada() string {
	if x != nil {
		return x.CiudadAfectada
	}
	return ""
}

func (x *InformanteBroker) GetNuevoValor() string {
	if x != nil {
		return x.NuevoValor
	}
	return ""
}

func (x *InformanteBroker) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *InformanteBroker) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *InformanteBroker) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type BrokerInformante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direccion string `protobuf:"bytes,1,opt,name=direccion,proto3" json:"direccion,omitempty"`
}

func (x *BrokerInformante) Reset() {
	*x = BrokerInformante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_funciones_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerInformante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerInformante) ProtoMessage() {}

func (x *BrokerInformante) ProtoReflect() protoreflect.Message {
	mi := &file_funciones_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerInformante.ProtoReflect.Descriptor instead.
func (*BrokerInformante) Descriptor() ([]byte, []int) {
	return file_funciones_proto_rawDescGZIP(), []int{1}
}

func (x *BrokerInformante) GetDireccion() string {
	if x != nil {
		return x.Direccion
	}
	return ""
}

// message informante_servidor
// strings accion, planeta_afectado, ciudad_afectada, Nuevo_Valor
type InformanteServidor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accion          string `protobuf:"bytes,1,opt,name=accion,proto3" json:"accion,omitempty"`
	PlanetaAfectado string `protobuf:"bytes,2,opt,name=planeta_afectado,json=planetaAfectado,proto3" json:"planeta_afectado,omitempty"`
	CiudadAfectada  string `protobuf:"bytes,3,opt,name=ciudad_afectada,json=ciudadAfectada,proto3" json:"ciudad_afectada,omitempty"`
	NuevoValor      string `protobuf:"bytes,4,opt,name=nuevo_valor,json=nuevoValor,proto3" json:"nuevo_valor,omitempty"`
}

func (x *InformanteServidor) Reset() {
	*x = InformanteServidor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_funciones_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformanteServidor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformanteServidor) ProtoMessage() {}

func (x *InformanteServidor) ProtoReflect() protoreflect.Message {
	mi := &file_funciones_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformanteServidor.ProtoReflect.Descriptor instead.
func (*InformanteServidor) Descriptor() ([]byte, []int) {
	return file_funciones_proto_rawDescGZIP(), []int{2}
}

func (x *InformanteServidor) GetAccion() string {
	if x != nil {
		return x.Accion
	}
	return ""
}

func (x *InformanteServidor) GetPlanetaAfectado() string {
	if x != nil {
		return x.PlanetaAfectado
	}
	return ""
}

func (x *InformanteServidor) GetCiudadAfectada() string {
	if x != nil {
		return x.CiudadAfectada
	}
	return ""
}

func (x *InformanteServidor) GetNuevoValor() string {
	if x != nil {
		return x.NuevoValor
	}
	return ""
}

// message servidor_informante
// ints x, y, z
type ServidorInformante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *ServidorInformante) Reset() {
	*x = ServidorInformante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_funciones_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServidorInformante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServidorInformante) ProtoMessage() {}

func (x *ServidorInformante) ProtoReflect() protoreflect.Message {
	mi := &file_funciones_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServidorInformante.ProtoReflect.Descriptor instead.
func (*ServidorInformante) Descriptor() ([]byte, []int) {
	return file_funciones_proto_rawDescGZIP(), []int{3}
}

func (x *ServidorInformante) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ServidorInformante) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *ServidorInformante) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type BrokerServidorCoord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
}

func (x *BrokerServidorCoord) Reset() {
	*x = BrokerServidorCoord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_funciones_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerServidorCoord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerServidorCoord) ProtoMessage() {}

func (x *BrokerServidorCoord) ProtoReflect() protoreflect.Message {
	mi := &file_funciones_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerServidorCoord.ProtoReflect.Descriptor instead.
func (*BrokerServidorCoord) Descriptor() ([]byte, []int) {
	return file_funciones_proto_rawDescGZIP(), []int{4}
}

func (x *BrokerServidorCoord) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

type ServidorBrokerCoord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *ServidorBrokerCoord) Reset() {
	*x = ServidorBrokerCoord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_funciones_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServidorBrokerCoord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServidorBrokerCoord) ProtoMessage() {}

func (x *ServidorBrokerCoord) ProtoReflect() protoreflect.Message {
	mi := &file_funciones_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServidorBrokerCoord.ProtoReflect.Descriptor instead.
func (*ServidorBrokerCoord) Descriptor() ([]byte, []int) {
	return file_funciones_proto_rawDescGZIP(), []int{5}
}

func (x *ServidorBrokerCoord) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ServidorBrokerCoord) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *ServidorBrokerCoord) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_funciones_proto protoreflect.FileDescriptor

var file_funciones_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xca, 0x01, 0x0a, 0x11, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x5f, 0x61, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x6f,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x5f, 0x61, 0x66, 0x65, 0x63, 0x74,
	0x61, 0x64, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x69, 0x75, 0x64, 0x61,
	0x64, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x65,
	0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x7a, 0x22, 0x31, 0x0a, 0x11, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x13, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x61, 0x5f, 0x61, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61,
	0x64, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x5f, 0x61, 0x66, 0x65,
	0x63, 0x74, 0x61, 0x64, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x69, 0x75,
	0x64, 0x61, 0x64, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6e,
	0x75, 0x65, 0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x13,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12,
	0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x22, 0x31, 0x0a,
	0x15, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x22, 0x41, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x7a, 0x32, 0xdf, 0x01, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x63, 0x69, 0x6f, 0x6e, 0x65,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x69, 0x6e, 0x66, 0x5f,
	0x62, 0x72, 0x6f, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x74, 0x65, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x1a, 0x19, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x10, 0x69, 0x6e, 0x66, 0x5f, 0x42,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x72, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x2f, 0x44, 0x69, 0x73,
	0x74, 0x54, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_funciones_proto_rawDescOnce sync.Once
	file_funciones_proto_rawDescData = file_funciones_proto_rawDesc
)

func file_funciones_proto_rawDescGZIP() []byte {
	file_funciones_proto_rawDescOnce.Do(func() {
		file_funciones_proto_rawDescData = protoimpl.X.CompressGZIP(file_funciones_proto_rawDescData)
	})
	return file_funciones_proto_rawDescData
}

var file_funciones_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_funciones_proto_goTypes = []interface{}{
	(*InformanteBroker)(nil),    // 0: grpc.informante_broker
	(*BrokerInformante)(nil),    // 1: grpc.broker_informante
	(*InformanteServidor)(nil),  // 2: grpc.informante_servidor
	(*ServidorInformante)(nil),  // 3: grpc.servidor_informante
	(*BrokerServidorCoord)(nil), // 4: grpc.broker_servidor_coord
	(*ServidorBrokerCoord)(nil), // 5: grpc.servidor_broker_coord
}
var file_funciones_proto_depIdxs = []int32{
	0, // 0: grpc.FuncionesService.inf_bro:input_type -> grpc.informante_broker
	2, // 1: grpc.FuncionesService.inf_serv:input_type -> grpc.informante_servidor
	4, // 2: grpc.FuncionesService.inf_BroServCoord:input_type -> grpc.broker_servidor_coord
	1, // 3: grpc.FuncionesService.inf_bro:output_type -> grpc.broker_informante
	3, // 4: grpc.FuncionesService.inf_serv:output_type -> grpc.servidor_informante
	5, // 5: grpc.FuncionesService.inf_BroServCoord:output_type -> grpc.servidor_broker_coord
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_funciones_proto_init() }
func file_funciones_proto_init() {
	if File_funciones_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_funciones_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformanteBroker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_funciones_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerInformante); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_funciones_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformanteServidor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_funciones_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServidorInformante); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_funciones_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerServidorCoord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_funciones_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServidorBrokerCoord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_funciones_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_funciones_proto_goTypes,
		DependencyIndexes: file_funciones_proto_depIdxs,
		MessageInfos:      file_funciones_proto_msgTypes,
	}.Build()
	File_funciones_proto = out.File
	file_funciones_proto_rawDesc = nil
	file_funciones_proto_goTypes = nil
	file_funciones_proto_depIdxs = nil
}
