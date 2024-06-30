// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: attendances.proto

package schedule_service

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

const (
	AttendanceService_Create_FullMethodName  = "/attendances.AttendanceService/Create"
	AttendanceService_GetById_FullMethodName = "/attendances.AttendanceService/GetById"
	AttendanceService_GetAll_FullMethodName  = "/attendances.AttendanceService/GetAll"
	AttendanceService_Update_FullMethodName  = "/attendances.AttendanceService/Update"
	AttendanceService_Delete_FullMethodName  = "/attendances.AttendanceService/Delete"
)

// AttendanceServiceClient is the client API for AttendanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttendanceServiceClient interface {
	Create(ctx context.Context, in *CreateAttendance, opts ...grpc.CallOption) (*Attendance, error)
	GetById(ctx context.Context, in *AttendancePrimaryKey, opts ...grpc.CallOption) (*Attendance, error)
	GetAll(ctx context.Context, in *GetListAttendanceRequest, opts ...grpc.CallOption) (*GetListAttendanceResponse, error)
	Update(ctx context.Context, in *UpdateAttendance, opts ...grpc.CallOption) (*Attendance, error)
	Delete(ctx context.Context, in *AttendancePrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type attendanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttendanceServiceClient(cc grpc.ClientConnInterface) AttendanceServiceClient {
	return &attendanceServiceClient{cc}
}

func (c *attendanceServiceClient) Create(ctx context.Context, in *CreateAttendance, opts ...grpc.CallOption) (*Attendance, error) {
	out := new(Attendance)
	err := c.cc.Invoke(ctx, AttendanceService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) GetById(ctx context.Context, in *AttendancePrimaryKey, opts ...grpc.CallOption) (*Attendance, error) {
	out := new(Attendance)
	err := c.cc.Invoke(ctx, AttendanceService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) GetAll(ctx context.Context, in *GetListAttendanceRequest, opts ...grpc.CallOption) (*GetListAttendanceResponse, error) {
	out := new(GetListAttendanceResponse)
	err := c.cc.Invoke(ctx, AttendanceService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) Update(ctx context.Context, in *UpdateAttendance, opts ...grpc.CallOption) (*Attendance, error) {
	out := new(Attendance)
	err := c.cc.Invoke(ctx, AttendanceService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) Delete(ctx context.Context, in *AttendancePrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AttendanceService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttendanceServiceServer is the server API for AttendanceService service.
// All implementations must embed UnimplementedAttendanceServiceServer
// for forward compatibility
type AttendanceServiceServer interface {
	Create(context.Context, *CreateAttendance) (*Attendance, error)
	GetById(context.Context, *AttendancePrimaryKey) (*Attendance, error)
	GetAll(context.Context, *GetListAttendanceRequest) (*GetListAttendanceResponse, error)
	Update(context.Context, *UpdateAttendance) (*Attendance, error)
	Delete(context.Context, *AttendancePrimaryKey) (*Empty, error)
	mustEmbedUnimplementedAttendanceServiceServer()
}

// UnimplementedAttendanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttendanceServiceServer struct {
}

func (UnimplementedAttendanceServiceServer) Create(context.Context, *CreateAttendance) (*Attendance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAttendanceServiceServer) GetById(context.Context, *AttendancePrimaryKey) (*Attendance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedAttendanceServiceServer) GetAll(context.Context, *GetListAttendanceRequest) (*GetListAttendanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAttendanceServiceServer) Update(context.Context, *UpdateAttendance) (*Attendance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAttendanceServiceServer) Delete(context.Context, *AttendancePrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAttendanceServiceServer) mustEmbedUnimplementedAttendanceServiceServer() {}

// UnsafeAttendanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttendanceServiceServer will
// result in compilation errors.
type UnsafeAttendanceServiceServer interface {
	mustEmbedUnimplementedAttendanceServiceServer()
}

func RegisterAttendanceServiceServer(s grpc.ServiceRegistrar, srv AttendanceServiceServer) {
	s.RegisterService(&AttendanceService_ServiceDesc, srv)
}

func _AttendanceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttendance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).Create(ctx, req.(*CreateAttendance))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendancePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).GetById(ctx, req.(*AttendancePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAttendanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).GetAll(ctx, req.(*GetListAttendanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttendance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).Update(ctx, req.(*UpdateAttendance))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendancePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).Delete(ctx, req.(*AttendancePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// AttendanceService_ServiceDesc is the grpc.ServiceDesc for AttendanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttendanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attendances.AttendanceService",
	HandlerType: (*AttendanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AttendanceService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _AttendanceService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AttendanceService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AttendanceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AttendanceService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendances.proto",
}