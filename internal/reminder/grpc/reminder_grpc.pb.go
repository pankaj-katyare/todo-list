// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// TodoReminderClient is the client API for TodoReminder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoReminderClient interface {
	CreateReminder(ctx context.Context, in *CreateReminderRequest, opts ...grpc.CallOption) (*CreateReminderResponse, error)
	GetReminder(ctx context.Context, in *GetReminderRequest, opts ...grpc.CallOption) (*GetReminderResponse, error)
}

type todoReminderClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoReminderClient(cc grpc.ClientConnInterface) TodoReminderClient {
	return &todoReminderClient{cc}
}

func (c *todoReminderClient) CreateReminder(ctx context.Context, in *CreateReminderRequest, opts ...grpc.CallOption) (*CreateReminderResponse, error) {
	out := new(CreateReminderResponse)
	err := c.cc.Invoke(ctx, "/TodoReminder/CreateReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoReminderClient) GetReminder(ctx context.Context, in *GetReminderRequest, opts ...grpc.CallOption) (*GetReminderResponse, error) {
	out := new(GetReminderResponse)
	err := c.cc.Invoke(ctx, "/TodoReminder/GetReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoReminderServer is the server API for TodoReminder service.
// All implementations must embed UnimplementedTodoReminderServer
// for forward compatibility
type TodoReminderServer interface {
	CreateReminder(context.Context, *CreateReminderRequest) (*CreateReminderResponse, error)
	GetReminder(context.Context, *GetReminderRequest) (*GetReminderResponse, error)
	mustEmbedUnimplementedTodoReminderServer()
}

// UnimplementedTodoReminderServer must be embedded to have forward compatible implementations.
type UnimplementedTodoReminderServer struct {
}

func (UnimplementedTodoReminderServer) CreateReminder(context.Context, *CreateReminderRequest) (*CreateReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReminder not implemented")
}
func (UnimplementedTodoReminderServer) GetReminder(context.Context, *GetReminderRequest) (*GetReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReminder not implemented")
}
func (UnimplementedTodoReminderServer) mustEmbedUnimplementedTodoReminderServer() {}

// UnsafeTodoReminderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoReminderServer will
// result in compilation errors.
type UnsafeTodoReminderServer interface {
	mustEmbedUnimplementedTodoReminderServer()
}

func RegisterTodoReminderServer(s grpc.ServiceRegistrar, srv TodoReminderServer) {
	s.RegisterService(&TodoReminder_ServiceDesc, srv)
}

func _TodoReminder_CreateReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoReminderServer).CreateReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoReminder/CreateReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoReminderServer).CreateReminder(ctx, req.(*CreateReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoReminder_GetReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoReminderServer).GetReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoReminder/GetReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoReminderServer).GetReminder(ctx, req.(*GetReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoReminder_ServiceDesc is the grpc.ServiceDesc for TodoReminder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoReminder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TodoReminder",
	HandlerType: (*TodoReminderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReminder",
			Handler:    _TodoReminder_CreateReminder_Handler,
		},
		{
			MethodName: "GetReminder",
			Handler:    _TodoReminder_GetReminder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/reminder/grpc/reminder.proto",
}


