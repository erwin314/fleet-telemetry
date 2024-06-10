// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: protos/data_connector_service.proto

package protos

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

// VehicleServiceClient is the client API for VehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleServiceClient interface {
	VinAllowed(ctx context.Context, in *VinAllowedRequest, opts ...grpc.CallOption) (*VinAllowedResponse, error)
}

type vehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleServiceClient(cc grpc.ClientConnInterface) VehicleServiceClient {
	return &vehicleServiceClient{cc}
}

func (c *vehicleServiceClient) VinAllowed(ctx context.Context, in *VinAllowedRequest, opts ...grpc.CallOption) (*VinAllowedResponse, error) {
	out := new(VinAllowedResponse)
	err := c.cc.Invoke(ctx, "/telemetry.connectors.VehicleService/VinAllowed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServiceServer is the server API for VehicleService service.
// All implementations must embed UnimplementedVehicleServiceServer
// for forward compatibility
type VehicleServiceServer interface {
	VinAllowed(context.Context, *VinAllowedRequest) (*VinAllowedResponse, error)
	mustEmbedUnimplementedVehicleServiceServer()
}

// UnimplementedVehicleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleServiceServer struct {
}

func (UnimplementedVehicleServiceServer) VinAllowed(context.Context, *VinAllowedRequest) (*VinAllowedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VinAllowed not implemented")
}
func (UnimplementedVehicleServiceServer) mustEmbedUnimplementedVehicleServiceServer() {}

// UnsafeVehicleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServiceServer will
// result in compilation errors.
type UnsafeVehicleServiceServer interface {
	mustEmbedUnimplementedVehicleServiceServer()
}

func RegisterVehicleServiceServer(s grpc.ServiceRegistrar, srv VehicleServiceServer) {
	s.RegisterService(&VehicleService_ServiceDesc, srv)
}

func _VehicleService_VinAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VinAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).VinAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telemetry.connectors.VehicleService/VinAllowed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).VinAllowed(ctx, req.(*VinAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleService_ServiceDesc is the grpc.ServiceDesc for VehicleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry.connectors.VehicleService",
	HandlerType: (*VehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VinAllowed",
			Handler:    _VehicleService_VinAllowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/data_connector_service.proto",
}