// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/v1/analyzer/analyzer.proto

package analyzer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnalyzerClient is the client API for Analyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyzerClient interface {
	Analyze(ctx context.Context, in *AnalyzerRequest, opts ...grpc.CallOption) (*AnalyzerResponse, error)
	List(ctx context.Context, in *ListAnalyzedTextsRequest, opts ...grpc.CallOption) (*ListAnalyzedTextsResponse, error)
	Get(ctx context.Context, in *GetAnalyzedTextRequest, opts ...grpc.CallOption) (*GetAnalyzedTextResponse, error)
	GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	Update(ctx context.Context, in *AnalyzedText, opts ...grpc.CallOption) (*AnalyzedText, error)
	Delete(ctx context.Context, in *DeleteAnalyzedTextRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type analyzerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyzerClient(cc grpc.ClientConnInterface) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzerRequest, opts ...grpc.CallOption) (*AnalyzerResponse, error) {
	out := new(AnalyzerResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.Analyzer/Analyze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) List(ctx context.Context, in *ListAnalyzedTextsRequest, opts ...grpc.CallOption) (*ListAnalyzedTextsResponse, error) {
	out := new(ListAnalyzedTextsResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.Analyzer/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) Get(ctx context.Context, in *GetAnalyzedTextRequest, opts ...grpc.CallOption) (*GetAnalyzedTextResponse, error) {
	out := new(GetAnalyzedTextResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.Analyzer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.Analyzer/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) Update(ctx context.Context, in *AnalyzedText, opts ...grpc.CallOption) (*AnalyzedText, error) {
	out := new(AnalyzedText)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.Analyzer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) Delete(ctx context.Context, in *DeleteAnalyzedTextRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.Analyzer/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzerServer is the server API for Analyzer service.
// All implementations must embed UnimplementedAnalyzerServer
// for forward compatibility
type AnalyzerServer interface {
	Analyze(context.Context, *AnalyzerRequest) (*AnalyzerResponse, error)
	List(context.Context, *ListAnalyzedTextsRequest) (*ListAnalyzedTextsResponse, error)
	Get(context.Context, *GetAnalyzedTextRequest) (*GetAnalyzedTextResponse, error)
	GetGroups(context.Context, *emptypb.Empty) (*GetGroupsResponse, error)
	Update(context.Context, *AnalyzedText) (*AnalyzedText, error)
	Delete(context.Context, *DeleteAnalyzedTextRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAnalyzerServer()
}

// UnimplementedAnalyzerServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyzerServer struct {
}

func (UnimplementedAnalyzerServer) Analyze(context.Context, *AnalyzerRequest) (*AnalyzerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedAnalyzerServer) List(context.Context, *ListAnalyzedTextsRequest) (*ListAnalyzedTextsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAnalyzerServer) Get(context.Context, *GetAnalyzedTextRequest) (*GetAnalyzedTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAnalyzerServer) GetGroups(context.Context, *emptypb.Empty) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedAnalyzerServer) Update(context.Context, *AnalyzedText) (*AnalyzedText, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAnalyzerServer) Delete(context.Context, *DeleteAnalyzedTextRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAnalyzerServer) mustEmbedUnimplementedAnalyzerServer() {}

// UnsafeAnalyzerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyzerServer will
// result in compilation errors.
type UnsafeAnalyzerServer interface {
	mustEmbedUnimplementedAnalyzerServer()
}

func RegisterAnalyzerServer(s grpc.ServiceRegistrar, srv AnalyzerServer) {
	s.RegisterService(&Analyzer_ServiceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnalyzedTextsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.Analyzer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).List(ctx, req.(*ListAnalyzedTextsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnalyzedTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.Analyzer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Get(ctx, req.(*GetAnalyzedTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.Analyzer/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzedText)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.Analyzer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Update(ctx, req.(*AnalyzedText))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnalyzedTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.Analyzer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Delete(ctx, req.(*DeleteAnalyzedTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Analyzer_ServiceDesc is the grpc.ServiceDesc for Analyzer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analyzer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.v1.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Analyzer_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Analyzer_Get_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _Analyzer_GetGroups_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Analyzer_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Analyzer_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/analyzer/analyzer.proto",
}
