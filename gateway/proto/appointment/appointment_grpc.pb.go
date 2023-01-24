// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appointment

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

// AppointmentClient is the client API for Appointment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppointmentClient interface {
	CreateSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
}

type appointmentClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentClient(cc grpc.ClientConnInterface) AppointmentClient {
	return &appointmentClient{cc}
}

func (c *appointmentClient) CreateSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/CreateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServer is the server API for Appointment service.
// All implementations must embed UnimplementedAppointmentServer
// for forward compatibility
type AppointmentServer interface {
	CreateSchedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	mustEmbedUnimplementedAppointmentServer()
}

// UnimplementedAppointmentServer must be embedded to have forward compatible implementations.
type UnimplementedAppointmentServer struct {
}

func (UnimplementedAppointmentServer) CreateSchedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedAppointmentServer) mustEmbedUnimplementedAppointmentServer() {}

// UnsafeAppointmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServer will
// result in compilation errors.
type UnsafeAppointmentServer interface {
	mustEmbedUnimplementedAppointmentServer()
}

func RegisterAppointmentServer(s grpc.ServiceRegistrar, srv AppointmentServer) {
	s.RegisterService(&Appointment_ServiceDesc, srv)
}

func _Appointment_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/CreateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).CreateSchedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Appointment_ServiceDesc is the grpc.ServiceDesc for Appointment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Appointment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appointment.Appointment",
	HandlerType: (*AppointmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchedule",
			Handler:    _Appointment_CreateSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/appointment/appointment.proto",
}
