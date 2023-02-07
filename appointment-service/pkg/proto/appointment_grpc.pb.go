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
	CreateDoctor(ctx context.Context, in *DoctorSingle, opts ...grpc.CallOption) (*Empty, error)
	GetDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*DoctorSingle, error)
	UpdateDoctor(ctx context.Context, in *DoctorSingle, opts ...grpc.CallOption) (*DoctorSingle, error)
	DeleteDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAllDoctors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DoctorMultiple, error)
	CreateSchedule(ctx context.Context, in *ScheduleSingle, opts ...grpc.CallOption) (*Empty, error)
	GetOpenEventsByDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error)
	GetReservedEventsByDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error)
	GetReservedEventsByClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error)
	GetAllEventsByClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error)
	RegisterToEvent(ctx context.Context, in *EventSingle, opts ...grpc.CallOption) (*EventSingle, error)
	UnregisterEvent(ctx context.Context, in *EventSingle, opts ...grpc.CallOption) (*Empty, error)
}

type appointmentClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentClient(cc grpc.ClientConnInterface) AppointmentClient {
	return &appointmentClient{cc}
}

func (c *appointmentClient) CreateDoctor(ctx context.Context, in *DoctorSingle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/CreateDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) GetDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*DoctorSingle, error) {
	out := new(DoctorSingle)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/GetDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) UpdateDoctor(ctx context.Context, in *DoctorSingle, opts ...grpc.CallOption) (*DoctorSingle, error) {
	out := new(DoctorSingle)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/UpdateDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) DeleteDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/DeleteDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) GetAllDoctors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DoctorMultiple, error) {
	out := new(DoctorMultiple)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/GetAllDoctors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) CreateSchedule(ctx context.Context, in *ScheduleSingle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/CreateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) GetOpenEventsByDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error) {
	out := new(EventMultiple)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/GetOpenEventsByDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) GetReservedEventsByDoctor(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error) {
	out := new(EventMultiple)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/GetReservedEventsByDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) GetReservedEventsByClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error) {
	out := new(EventMultiple)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/GetReservedEventsByClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) GetAllEventsByClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EventMultiple, error) {
	out := new(EventMultiple)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/GetAllEventsByClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) RegisterToEvent(ctx context.Context, in *EventSingle, opts ...grpc.CallOption) (*EventSingle, error) {
	out := new(EventSingle)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/RegisterToEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentClient) UnregisterEvent(ctx context.Context, in *EventSingle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/appointment.Appointment/UnregisterEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServer is the server API for Appointment service.
// All implementations must embed UnimplementedAppointmentServer
// for forward compatibility
type AppointmentServer interface {
	CreateDoctor(context.Context, *DoctorSingle) (*Empty, error)
	GetDoctor(context.Context, *IdRequest) (*DoctorSingle, error)
	UpdateDoctor(context.Context, *DoctorSingle) (*DoctorSingle, error)
	DeleteDoctor(context.Context, *IdRequest) (*Empty, error)
	GetAllDoctors(context.Context, *Empty) (*DoctorMultiple, error)
	CreateSchedule(context.Context, *ScheduleSingle) (*Empty, error)
	GetOpenEventsByDoctor(context.Context, *IdRequest) (*EventMultiple, error)
	GetReservedEventsByDoctor(context.Context, *IdRequest) (*EventMultiple, error)
	GetReservedEventsByClient(context.Context, *IdRequest) (*EventMultiple, error)
	GetAllEventsByClient(context.Context, *IdRequest) (*EventMultiple, error)
	RegisterToEvent(context.Context, *EventSingle) (*EventSingle, error)
	UnregisterEvent(context.Context, *EventSingle) (*Empty, error)
	mustEmbedUnimplementedAppointmentServer()
}

// UnimplementedAppointmentServer must be embedded to have forward compatible implementations.
type UnimplementedAppointmentServer struct {
}

func (UnimplementedAppointmentServer) CreateDoctor(context.Context, *DoctorSingle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDoctor not implemented")
}
func (UnimplementedAppointmentServer) GetDoctor(context.Context, *IdRequest) (*DoctorSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctor not implemented")
}
func (UnimplementedAppointmentServer) UpdateDoctor(context.Context, *DoctorSingle) (*DoctorSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctor not implemented")
}
func (UnimplementedAppointmentServer) DeleteDoctor(context.Context, *IdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoctor not implemented")
}
func (UnimplementedAppointmentServer) GetAllDoctors(context.Context, *Empty) (*DoctorMultiple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDoctors not implemented")
}
func (UnimplementedAppointmentServer) CreateSchedule(context.Context, *ScheduleSingle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedAppointmentServer) GetOpenEventsByDoctor(context.Context, *IdRequest) (*EventMultiple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenEventsByDoctor not implemented")
}
func (UnimplementedAppointmentServer) GetReservedEventsByDoctor(context.Context, *IdRequest) (*EventMultiple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservedEventsByDoctor not implemented")
}
func (UnimplementedAppointmentServer) GetReservedEventsByClient(context.Context, *IdRequest) (*EventMultiple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservedEventsByClient not implemented")
}
func (UnimplementedAppointmentServer) GetAllEventsByClient(context.Context, *IdRequest) (*EventMultiple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEventsByClient not implemented")
}
func (UnimplementedAppointmentServer) RegisterToEvent(context.Context, *EventSingle) (*EventSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterToEvent not implemented")
}
func (UnimplementedAppointmentServer) UnregisterEvent(context.Context, *EventSingle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterEvent not implemented")
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

func _Appointment_CreateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).CreateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/CreateDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).CreateDoctor(ctx, req.(*DoctorSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_GetDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).GetDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/GetDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).GetDoctor(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_UpdateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).UpdateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/UpdateDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).UpdateDoctor(ctx, req.(*DoctorSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_DeleteDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).DeleteDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/DeleteDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).DeleteDoctor(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_GetAllDoctors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).GetAllDoctors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/GetAllDoctors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).GetAllDoctors(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleSingle)
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
		return srv.(AppointmentServer).CreateSchedule(ctx, req.(*ScheduleSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_GetOpenEventsByDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).GetOpenEventsByDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/GetOpenEventsByDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).GetOpenEventsByDoctor(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_GetReservedEventsByDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).GetReservedEventsByDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/GetReservedEventsByDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).GetReservedEventsByDoctor(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_GetReservedEventsByClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).GetReservedEventsByClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/GetReservedEventsByClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).GetReservedEventsByClient(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_GetAllEventsByClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).GetAllEventsByClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/GetAllEventsByClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).GetAllEventsByClient(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_RegisterToEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).RegisterToEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/RegisterToEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).RegisterToEvent(ctx, req.(*EventSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appointment_UnregisterEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServer).UnregisterEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.Appointment/UnregisterEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServer).UnregisterEvent(ctx, req.(*EventSingle))
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
			MethodName: "CreateDoctor",
			Handler:    _Appointment_CreateDoctor_Handler,
		},
		{
			MethodName: "GetDoctor",
			Handler:    _Appointment_GetDoctor_Handler,
		},
		{
			MethodName: "UpdateDoctor",
			Handler:    _Appointment_UpdateDoctor_Handler,
		},
		{
			MethodName: "DeleteDoctor",
			Handler:    _Appointment_DeleteDoctor_Handler,
		},
		{
			MethodName: "GetAllDoctors",
			Handler:    _Appointment_GetAllDoctors_Handler,
		},
		{
			MethodName: "CreateSchedule",
			Handler:    _Appointment_CreateSchedule_Handler,
		},
		{
			MethodName: "GetOpenEventsByDoctor",
			Handler:    _Appointment_GetOpenEventsByDoctor_Handler,
		},
		{
			MethodName: "GetReservedEventsByDoctor",
			Handler:    _Appointment_GetReservedEventsByDoctor_Handler,
		},
		{
			MethodName: "GetReservedEventsByClient",
			Handler:    _Appointment_GetReservedEventsByClient_Handler,
		},
		{
			MethodName: "GetAllEventsByClient",
			Handler:    _Appointment_GetAllEventsByClient_Handler,
		},
		{
			MethodName: "RegisterToEvent",
			Handler:    _Appointment_RegisterToEvent_Handler,
		},
		{
			MethodName: "UnregisterEvent",
			Handler:    _Appointment_UnregisterEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/appointment.proto",
}
