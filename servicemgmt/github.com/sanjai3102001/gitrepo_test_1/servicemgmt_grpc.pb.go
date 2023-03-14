// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: servicemgmt.proto

package gitrepo_test_1

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

// ServiceManagementClient is the client API for ServiceManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceManagementClient interface {
	CreateNewservice(ctx context.Context, in *Newservice, opts ...grpc.CallOption) (*Service, error)
	Getservices(ctx context.Context, in *GetservicesParams, opts ...grpc.CallOption) (*ServicesList, error)
}

type serviceManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceManagementClient(cc grpc.ClientConnInterface) ServiceManagementClient {
	return &serviceManagementClient{cc}
}

func (c *serviceManagementClient) CreateNewservice(ctx context.Context, in *Newservice, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/servicemgmt.serviceManagement/CreateNewservice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceManagementClient) Getservices(ctx context.Context, in *GetservicesParams, opts ...grpc.CallOption) (*ServicesList, error) {
	out := new(ServicesList)
	err := c.cc.Invoke(ctx, "/servicemgmt.serviceManagement/Getservices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceManagementServer is the server API for ServiceManagement service.
// All implementations must embed UnimplementedServiceManagementServer
// for forward compatibility
type ServiceManagementServer interface {
	CreateNewservice(context.Context, *Newservice) (*Service, error)
	Getservices(context.Context, *GetservicesParams) (*ServicesList, error)
	mustEmbedUnimplementedServiceManagementServer()
}

// UnimplementedServiceManagementServer must be embedded to have forward compatible implementations.
type UnimplementedServiceManagementServer struct {
}

func (UnimplementedServiceManagementServer) CreateNewservice(context.Context, *Newservice) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewservice not implemented")
}
func (UnimplementedServiceManagementServer) Getservices(context.Context, *GetservicesParams) (*ServicesList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getservices not implemented")
}
func (UnimplementedServiceManagementServer) mustEmbedUnimplementedServiceManagementServer() {}

// UnsafeServiceManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceManagementServer will
// result in compilation errors.
type UnsafeServiceManagementServer interface {
	mustEmbedUnimplementedServiceManagementServer()
}

func RegisterServiceManagementServer(s grpc.ServiceRegistrar, srv ServiceManagementServer) {
	s.RegisterService(&ServiceManagement_ServiceDesc, srv)
}

func _ServiceManagement_CreateNewservice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Newservice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceManagementServer).CreateNewservice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicemgmt.serviceManagement/CreateNewservice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceManagementServer).CreateNewservice(ctx, req.(*Newservice))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceManagement_Getservices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetservicesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceManagementServer).Getservices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicemgmt.serviceManagement/Getservices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceManagementServer).Getservices(ctx, req.(*GetservicesParams))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceManagement_ServiceDesc is the grpc.ServiceDesc for ServiceManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "servicemgmt.serviceManagement",
	HandlerType: (*ServiceManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewservice",
			Handler:    _ServiceManagement_CreateNewservice_Handler,
		},
		{
			MethodName: "Getservices",
			Handler:    _ServiceManagement_Getservices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servicemgmt.proto",
}
