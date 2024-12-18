// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pkg/grpc/protobuf/digimon.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PrimaryNodeService_ReceiveEncryptedMessage_FullMethodName = "/digimon.PrimaryNodeService/ReceiveEncryptedMessage"
	PrimaryNodeService_GetAttackData_FullMethodName           = "/digimon.PrimaryNodeService/GetAttackData"
	PrimaryNodeService_SendTerminationSignal_FullMethodName   = "/digimon.PrimaryNodeService/SendTerminationSignal"
)

// PrimaryNodeServiceClient is the client API for PrimaryNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio del Primary Node
type PrimaryNodeServiceClient interface {
	ReceiveEncryptedMessage(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (*Empty, error)
	GetAttackData(ctx context.Context, in *TaiRequest, opts ...grpc.CallOption) (*AttackDataResponse, error)
	SendTerminationSignal(ctx context.Context, in *TerminateProcess, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type primaryNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimaryNodeServiceClient(cc grpc.ClientConnInterface) PrimaryNodeServiceClient {
	return &primaryNodeServiceClient{cc}
}

func (c *primaryNodeServiceClient) ReceiveEncryptedMessage(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, PrimaryNodeService_ReceiveEncryptedMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *primaryNodeServiceClient) GetAttackData(ctx context.Context, in *TaiRequest, opts ...grpc.CallOption) (*AttackDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttackDataResponse)
	err := c.cc.Invoke(ctx, PrimaryNodeService_GetAttackData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *primaryNodeServiceClient) SendTerminationSignal(ctx context.Context, in *TerminateProcess, opts ...grpc.CallOption) (*TerminateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, PrimaryNodeService_SendTerminationSignal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrimaryNodeServiceServer is the server API for PrimaryNodeService service.
// All implementations must embed UnimplementedPrimaryNodeServiceServer
// for forward compatibility.
//
// Servicio del Primary Node
type PrimaryNodeServiceServer interface {
	ReceiveEncryptedMessage(context.Context, *EncryptedMessage) (*Empty, error)
	GetAttackData(context.Context, *TaiRequest) (*AttackDataResponse, error)
	SendTerminationSignal(context.Context, *TerminateProcess) (*TerminateResponse, error)
	mustEmbedUnimplementedPrimaryNodeServiceServer()
}

// UnimplementedPrimaryNodeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPrimaryNodeServiceServer struct{}

func (UnimplementedPrimaryNodeServiceServer) ReceiveEncryptedMessage(context.Context, *EncryptedMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveEncryptedMessage not implemented")
}
func (UnimplementedPrimaryNodeServiceServer) GetAttackData(context.Context, *TaiRequest) (*AttackDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttackData not implemented")
}
func (UnimplementedPrimaryNodeServiceServer) SendTerminationSignal(context.Context, *TerminateProcess) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTerminationSignal not implemented")
}
func (UnimplementedPrimaryNodeServiceServer) mustEmbedUnimplementedPrimaryNodeServiceServer() {}
func (UnimplementedPrimaryNodeServiceServer) testEmbeddedByValue()                            {}

// UnsafePrimaryNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimaryNodeServiceServer will
// result in compilation errors.
type UnsafePrimaryNodeServiceServer interface {
	mustEmbedUnimplementedPrimaryNodeServiceServer()
}

func RegisterPrimaryNodeServiceServer(s grpc.ServiceRegistrar, srv PrimaryNodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedPrimaryNodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PrimaryNodeService_ServiceDesc, srv)
}

func _PrimaryNodeService_ReceiveEncryptedMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrimaryNodeServiceServer).ReceiveEncryptedMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrimaryNodeService_ReceiveEncryptedMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrimaryNodeServiceServer).ReceiveEncryptedMessage(ctx, req.(*EncryptedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrimaryNodeService_GetAttackData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrimaryNodeServiceServer).GetAttackData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrimaryNodeService_GetAttackData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrimaryNodeServiceServer).GetAttackData(ctx, req.(*TaiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrimaryNodeService_SendTerminationSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateProcess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrimaryNodeServiceServer).SendTerminationSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrimaryNodeService_SendTerminationSignal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrimaryNodeServiceServer).SendTerminationSignal(ctx, req.(*TerminateProcess))
	}
	return interceptor(ctx, in, info, handler)
}

// PrimaryNodeService_ServiceDesc is the grpc.ServiceDesc for PrimaryNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimaryNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.PrimaryNodeService",
	HandlerType: (*PrimaryNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveEncryptedMessage",
			Handler:    _PrimaryNodeService_ReceiveEncryptedMessage_Handler,
		},
		{
			MethodName: "GetAttackData",
			Handler:    _PrimaryNodeService_GetAttackData_Handler,
		},
		{
			MethodName: "SendTerminationSignal",
			Handler:    _PrimaryNodeService_SendTerminationSignal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/protobuf/digimon.proto",
}

const (
	DataNodeService_StoreDigimon_FullMethodName        = "/digimon.DataNodeService/StoreDigimon"
	DataNodeService_GetDigimonAttribute_FullMethodName = "/digimon.DataNodeService/GetDigimonAttribute"
	DataNodeService_Terminate_FullMethodName           = "/digimon.DataNodeService/Terminate"
)

// DataNodeServiceClient is the client API for DataNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio del Data Node
type DataNodeServiceClient interface {
	StoreDigimon(ctx context.Context, in *DigimonInfo, opts ...grpc.CallOption) (*StoreDigimonResponse, error)
	GetDigimonAttribute(ctx context.Context, in *DigimonRequest, opts ...grpc.CallOption) (*DigimonResponse, error)
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type dataNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeServiceClient(cc grpc.ClientConnInterface) DataNodeServiceClient {
	return &dataNodeServiceClient{cc}
}

func (c *dataNodeServiceClient) StoreDigimon(ctx context.Context, in *DigimonInfo, opts ...grpc.CallOption) (*StoreDigimonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreDigimonResponse)
	err := c.cc.Invoke(ctx, DataNodeService_StoreDigimon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) GetDigimonAttribute(ctx context.Context, in *DigimonRequest, opts ...grpc.CallOption) (*DigimonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DigimonResponse)
	err := c.cc.Invoke(ctx, DataNodeService_GetDigimonAttribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, DataNodeService_Terminate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServiceServer is the server API for DataNodeService service.
// All implementations must embed UnimplementedDataNodeServiceServer
// for forward compatibility.
//
// Servicio del Data Node
type DataNodeServiceServer interface {
	StoreDigimon(context.Context, *DigimonInfo) (*StoreDigimonResponse, error)
	GetDigimonAttribute(context.Context, *DigimonRequest) (*DigimonResponse, error)
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
	mustEmbedUnimplementedDataNodeServiceServer()
}

// UnimplementedDataNodeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataNodeServiceServer struct{}

func (UnimplementedDataNodeServiceServer) StoreDigimon(context.Context, *DigimonInfo) (*StoreDigimonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDigimon not implemented")
}
func (UnimplementedDataNodeServiceServer) GetDigimonAttribute(context.Context, *DigimonRequest) (*DigimonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDigimonAttribute not implemented")
}
func (UnimplementedDataNodeServiceServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}
func (UnimplementedDataNodeServiceServer) mustEmbedUnimplementedDataNodeServiceServer() {}
func (UnimplementedDataNodeServiceServer) testEmbeddedByValue()                         {}

// UnsafeDataNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServiceServer will
// result in compilation errors.
type UnsafeDataNodeServiceServer interface {
	mustEmbedUnimplementedDataNodeServiceServer()
}

func RegisterDataNodeServiceServer(s grpc.ServiceRegistrar, srv DataNodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataNodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataNodeService_ServiceDesc, srv)
}

func _DataNodeService_StoreDigimon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigimonInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).StoreDigimon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNodeService_StoreDigimon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).StoreDigimon(ctx, req.(*DigimonInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_GetDigimonAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigimonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).GetDigimonAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNodeService_GetDigimonAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).GetDigimonAttribute(ctx, req.(*DigimonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNodeService_Terminate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNodeService_ServiceDesc is the grpc.ServiceDesc for DataNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.DataNodeService",
	HandlerType: (*DataNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreDigimon",
			Handler:    _DataNodeService_StoreDigimon_Handler,
		},
		{
			MethodName: "GetDigimonAttribute",
			Handler:    _DataNodeService_GetDigimonAttribute_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _DataNodeService_Terminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/protobuf/digimon.proto",
}

const (
	RegionalServerService_TerminateRegional_FullMethodName = "/digimon.RegionalServerService/TerminateRegional"
)

// RegionalServerServiceClient is the client API for RegionalServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio del Servidor Regional
type RegionalServerServiceClient interface {
	TerminateRegional(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type regionalServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionalServerServiceClient(cc grpc.ClientConnInterface) RegionalServerServiceClient {
	return &regionalServerServiceClient{cc}
}

func (c *regionalServerServiceClient) TerminateRegional(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, RegionalServerService_TerminateRegional_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionalServerServiceServer is the server API for RegionalServerService service.
// All implementations must embed UnimplementedRegionalServerServiceServer
// for forward compatibility.
//
// Servicio del Servidor Regional
type RegionalServerServiceServer interface {
	TerminateRegional(context.Context, *TerminateRequest) (*TerminateResponse, error)
	mustEmbedUnimplementedRegionalServerServiceServer()
}

// UnimplementedRegionalServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRegionalServerServiceServer struct{}

func (UnimplementedRegionalServerServiceServer) TerminateRegional(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateRegional not implemented")
}
func (UnimplementedRegionalServerServiceServer) mustEmbedUnimplementedRegionalServerServiceServer() {}
func (UnimplementedRegionalServerServiceServer) testEmbeddedByValue()                               {}

// UnsafeRegionalServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionalServerServiceServer will
// result in compilation errors.
type UnsafeRegionalServerServiceServer interface {
	mustEmbedUnimplementedRegionalServerServiceServer()
}

func RegisterRegionalServerServiceServer(s grpc.ServiceRegistrar, srv RegionalServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedRegionalServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RegionalServerService_ServiceDesc, srv)
}

func _RegionalServerService_TerminateRegional_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionalServerServiceServer).TerminateRegional(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionalServerService_TerminateRegional_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionalServerServiceServer).TerminateRegional(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionalServerService_ServiceDesc is the grpc.ServiceDesc for RegionalServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionalServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.RegionalServerService",
	HandlerType: (*RegionalServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TerminateRegional",
			Handler:    _RegionalServerService_TerminateRegional_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/protobuf/digimon.proto",
}

const (
	DiaboromonService_StartDiaboromon_FullMethodName  = "/digimon.DiaboromonService/StartDiaboromon"
	DiaboromonService_AttackDiaboromon_FullMethodName = "/digimon.DiaboromonService/AttackDiaboromon"
)

// DiaboromonServiceClient is the client API for DiaboromonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio de Diaboromon
type DiaboromonServiceClient interface {
	StartDiaboromon(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StartResponse, error)
	AttackDiaboromon(ctx context.Context, in *AttackRequest, opts ...grpc.CallOption) (*AttackResponse, error)
}

type diaboromonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiaboromonServiceClient(cc grpc.ClientConnInterface) DiaboromonServiceClient {
	return &diaboromonServiceClient{cc}
}

func (c *diaboromonServiceClient) StartDiaboromon(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, DiaboromonService_StartDiaboromon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaboromonServiceClient) AttackDiaboromon(ctx context.Context, in *AttackRequest, opts ...grpc.CallOption) (*AttackResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttackResponse)
	err := c.cc.Invoke(ctx, DiaboromonService_AttackDiaboromon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiaboromonServiceServer is the server API for DiaboromonService service.
// All implementations must embed UnimplementedDiaboromonServiceServer
// for forward compatibility.
//
// Servicio de Diaboromon
type DiaboromonServiceServer interface {
	StartDiaboromon(context.Context, *Empty) (*StartResponse, error)
	AttackDiaboromon(context.Context, *AttackRequest) (*AttackResponse, error)
	mustEmbedUnimplementedDiaboromonServiceServer()
}

// UnimplementedDiaboromonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDiaboromonServiceServer struct{}

func (UnimplementedDiaboromonServiceServer) StartDiaboromon(context.Context, *Empty) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDiaboromon not implemented")
}
func (UnimplementedDiaboromonServiceServer) AttackDiaboromon(context.Context, *AttackRequest) (*AttackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttackDiaboromon not implemented")
}
func (UnimplementedDiaboromonServiceServer) mustEmbedUnimplementedDiaboromonServiceServer() {}
func (UnimplementedDiaboromonServiceServer) testEmbeddedByValue()                           {}

// UnsafeDiaboromonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiaboromonServiceServer will
// result in compilation errors.
type UnsafeDiaboromonServiceServer interface {
	mustEmbedUnimplementedDiaboromonServiceServer()
}

func RegisterDiaboromonServiceServer(s grpc.ServiceRegistrar, srv DiaboromonServiceServer) {
	// If the following call pancis, it indicates UnimplementedDiaboromonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DiaboromonService_ServiceDesc, srv)
}

func _DiaboromonService_StartDiaboromon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaboromonServiceServer).StartDiaboromon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiaboromonService_StartDiaboromon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaboromonServiceServer).StartDiaboromon(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiaboromonService_AttackDiaboromon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaboromonServiceServer).AttackDiaboromon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiaboromonService_AttackDiaboromon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaboromonServiceServer).AttackDiaboromon(ctx, req.(*AttackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiaboromonService_ServiceDesc is the grpc.ServiceDesc for DiaboromonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiaboromonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.DiaboromonService",
	HandlerType: (*DiaboromonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartDiaboromon",
			Handler:    _DiaboromonService_StartDiaboromon_Handler,
		},
		{
			MethodName: "AttackDiaboromon",
			Handler:    _DiaboromonService_AttackDiaboromon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/protobuf/digimon.proto",
}

const (
	TaiService_DiaboromonAttack_FullMethodName = "/digimon.TaiService/DiaboromonAttack"
)

// TaiServiceClient is the client API for TaiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio de Tai
type TaiServiceClient interface {
	DiaboromonAttack(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AttackResponse, error)
}

type taiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaiServiceClient(cc grpc.ClientConnInterface) TaiServiceClient {
	return &taiServiceClient{cc}
}

func (c *taiServiceClient) DiaboromonAttack(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AttackResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttackResponse)
	err := c.cc.Invoke(ctx, TaiService_DiaboromonAttack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaiServiceServer is the server API for TaiService service.
// All implementations must embed UnimplementedTaiServiceServer
// for forward compatibility.
//
// Servicio de Tai
type TaiServiceServer interface {
	DiaboromonAttack(context.Context, *Empty) (*AttackResponse, error)
	mustEmbedUnimplementedTaiServiceServer()
}

// UnimplementedTaiServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaiServiceServer struct{}

func (UnimplementedTaiServiceServer) DiaboromonAttack(context.Context, *Empty) (*AttackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiaboromonAttack not implemented")
}
func (UnimplementedTaiServiceServer) mustEmbedUnimplementedTaiServiceServer() {}
func (UnimplementedTaiServiceServer) testEmbeddedByValue()                    {}

// UnsafeTaiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaiServiceServer will
// result in compilation errors.
type UnsafeTaiServiceServer interface {
	mustEmbedUnimplementedTaiServiceServer()
}

func RegisterTaiServiceServer(s grpc.ServiceRegistrar, srv TaiServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaiServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaiService_ServiceDesc, srv)
}

func _TaiService_DiaboromonAttack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiServiceServer).DiaboromonAttack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaiService_DiaboromonAttack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiServiceServer).DiaboromonAttack(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TaiService_ServiceDesc is the grpc.ServiceDesc for TaiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.TaiService",
	HandlerType: (*TaiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiaboromonAttack",
			Handler:    _TaiService_DiaboromonAttack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/protobuf/digimon.proto",
}
