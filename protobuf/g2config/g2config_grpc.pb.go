// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: g2config.proto

package g2configprotobuf

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

// G2DiagnosticClient is the client API for G2Diagnostic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type G2DiagnosticClient interface {
	AddDataSource(ctx context.Context, in *AddDataSourceRequest, opts ...grpc.CallOption) (*AddDataSourceResponse, error)
	ClearLastException(ctx context.Context, in *ClearLastExceptionRequest, opts ...grpc.CallOption) (*ClearLastExceptionResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	DeleteDataSource(ctx context.Context, in *DeleteDataSourceRequest, opts ...grpc.CallOption) (*DeleteDataSourceResponse, error)
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyResponse, error)
	GetLastException(ctx context.Context, in *GetLastExceptionRequest, opts ...grpc.CallOption) (*GetLastExceptionResponse, error)
	GetLastExceptionCode(ctx context.Context, in *GetLastExceptionCodeRequest, opts ...grpc.CallOption) (*GetLastExceptionCodeResponse, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	ListDataSources(ctx context.Context, in *ListDataSourcesRequest, opts ...grpc.CallOption) (*ListDataSourcesResponse, error)
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
}

type g2DiagnosticClient struct {
	cc grpc.ClientConnInterface
}

func NewG2DiagnosticClient(cc grpc.ClientConnInterface) G2DiagnosticClient {
	return &g2DiagnosticClient{cc}
}

func (c *g2DiagnosticClient) AddDataSource(ctx context.Context, in *AddDataSourceRequest, opts ...grpc.CallOption) (*AddDataSourceResponse, error) {
	out := new(AddDataSourceResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/AddDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) ClearLastException(ctx context.Context, in *ClearLastExceptionRequest, opts ...grpc.CallOption) (*ClearLastExceptionResponse, error) {
	out := new(ClearLastExceptionResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/ClearLastException", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) DeleteDataSource(ctx context.Context, in *DeleteDataSourceRequest, opts ...grpc.CallOption) (*DeleteDataSourceResponse, error) {
	out := new(DeleteDataSourceResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/DeleteDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyResponse, error) {
	out := new(DestroyResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) GetLastException(ctx context.Context, in *GetLastExceptionRequest, opts ...grpc.CallOption) (*GetLastExceptionResponse, error) {
	out := new(GetLastExceptionResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/GetLastException", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) GetLastExceptionCode(ctx context.Context, in *GetLastExceptionCodeRequest, opts ...grpc.CallOption) (*GetLastExceptionCodeResponse, error) {
	out := new(GetLastExceptionCodeResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/GetLastExceptionCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) ListDataSources(ctx context.Context, in *ListDataSourcesRequest, opts ...grpc.CallOption) (*ListDataSourcesResponse, error) {
	out := new(ListDataSourcesResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/ListDataSources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *g2DiagnosticClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/g2configprotobuf.G2Diagnostic/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// G2DiagnosticServer is the server API for G2Diagnostic service.
// All implementations must embed UnimplementedG2DiagnosticServer
// for forward compatibility
type G2DiagnosticServer interface {
	AddDataSource(context.Context, *AddDataSourceRequest) (*AddDataSourceResponse, error)
	ClearLastException(context.Context, *ClearLastExceptionRequest) (*ClearLastExceptionResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	DeleteDataSource(context.Context, *DeleteDataSourceRequest) (*DeleteDataSourceResponse, error)
	Destroy(context.Context, *DestroyRequest) (*DestroyResponse, error)
	GetLastException(context.Context, *GetLastExceptionRequest) (*GetLastExceptionResponse, error)
	GetLastExceptionCode(context.Context, *GetLastExceptionCodeRequest) (*GetLastExceptionCodeResponse, error)
	Init(context.Context, *InitRequest) (*InitResponse, error)
	ListDataSources(context.Context, *ListDataSourcesRequest) (*ListDataSourcesResponse, error)
	Load(context.Context, *LoadRequest) (*LoadResponse, error)
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	mustEmbedUnimplementedG2DiagnosticServer()
}

// UnimplementedG2DiagnosticServer must be embedded to have forward compatible implementations.
type UnimplementedG2DiagnosticServer struct {
}

func (UnimplementedG2DiagnosticServer) AddDataSource(context.Context, *AddDataSourceRequest) (*AddDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDataSource not implemented")
}
func (UnimplementedG2DiagnosticServer) ClearLastException(context.Context, *ClearLastExceptionRequest) (*ClearLastExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearLastException not implemented")
}
func (UnimplementedG2DiagnosticServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedG2DiagnosticServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedG2DiagnosticServer) DeleteDataSource(context.Context, *DeleteDataSourceRequest) (*DeleteDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataSource not implemented")
}
func (UnimplementedG2DiagnosticServer) Destroy(context.Context, *DestroyRequest) (*DestroyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedG2DiagnosticServer) GetLastException(context.Context, *GetLastExceptionRequest) (*GetLastExceptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastException not implemented")
}
func (UnimplementedG2DiagnosticServer) GetLastExceptionCode(context.Context, *GetLastExceptionCodeRequest) (*GetLastExceptionCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastExceptionCode not implemented")
}
func (UnimplementedG2DiagnosticServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedG2DiagnosticServer) ListDataSources(context.Context, *ListDataSourcesRequest) (*ListDataSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataSources not implemented")
}
func (UnimplementedG2DiagnosticServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedG2DiagnosticServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedG2DiagnosticServer) mustEmbedUnimplementedG2DiagnosticServer() {}

// UnsafeG2DiagnosticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to G2DiagnosticServer will
// result in compilation errors.
type UnsafeG2DiagnosticServer interface {
	mustEmbedUnimplementedG2DiagnosticServer()
}

func RegisterG2DiagnosticServer(s grpc.ServiceRegistrar, srv G2DiagnosticServer) {
	s.RegisterService(&G2Diagnostic_ServiceDesc, srv)
}

func _G2Diagnostic_AddDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).AddDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/AddDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).AddDataSource(ctx, req.(*AddDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_ClearLastException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLastExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).ClearLastException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/ClearLastException",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).ClearLastException(ctx, req.(*ClearLastExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_DeleteDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).DeleteDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/DeleteDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).DeleteDataSource(ctx, req.(*DeleteDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).Destroy(ctx, req.(*DestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_GetLastException_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastExceptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).GetLastException(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/GetLastException",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).GetLastException(ctx, req.(*GetLastExceptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_GetLastExceptionCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastExceptionCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).GetLastExceptionCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/GetLastExceptionCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).GetLastExceptionCode(ctx, req.(*GetLastExceptionCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_ListDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).ListDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/ListDataSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).ListDataSources(ctx, req.(*ListDataSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _G2Diagnostic_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(G2DiagnosticServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g2configprotobuf.G2Diagnostic/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(G2DiagnosticServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// G2Diagnostic_ServiceDesc is the grpc.ServiceDesc for G2Diagnostic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var G2Diagnostic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g2configprotobuf.G2Diagnostic",
	HandlerType: (*G2DiagnosticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDataSource",
			Handler:    _G2Diagnostic_AddDataSource_Handler,
		},
		{
			MethodName: "ClearLastException",
			Handler:    _G2Diagnostic_ClearLastException_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _G2Diagnostic_Close_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _G2Diagnostic_Create_Handler,
		},
		{
			MethodName: "DeleteDataSource",
			Handler:    _G2Diagnostic_DeleteDataSource_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _G2Diagnostic_Destroy_Handler,
		},
		{
			MethodName: "GetLastException",
			Handler:    _G2Diagnostic_GetLastException_Handler,
		},
		{
			MethodName: "GetLastExceptionCode",
			Handler:    _G2Diagnostic_GetLastExceptionCode_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _G2Diagnostic_Init_Handler,
		},
		{
			MethodName: "ListDataSources",
			Handler:    _G2Diagnostic_ListDataSources_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _G2Diagnostic_Load_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _G2Diagnostic_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "g2config.proto",
}
