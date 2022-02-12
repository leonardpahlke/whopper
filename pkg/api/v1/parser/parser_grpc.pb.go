// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/v1/parser/parser.proto

package parser

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

// ArticleParserClient is the client API for ArticleParser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleParserClient interface {
	Parse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error)
	List(ctx context.Context, in *ListParsedArticlesRequest, opts ...grpc.CallOption) (*ListParsedArticlesResponse, error)
	Get(ctx context.Context, in *GetParsedArticleRequest, opts ...grpc.CallOption) (*GetParsedArticleResponse, error)
	// Get a list of parsers that can be used to parse a newspaper website
	// Its possible to have multiple parsers per website
	GetParsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetParsersResponse, error)
	// Get a list of supported newspapers
	GetNewspapers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNewspapersResponse, error)
	Update(ctx context.Context, in *ParsedArticle, opts ...grpc.CallOption) (*ParsedArticle, error)
	Delete(ctx context.Context, in *DeleteParsedArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type articleParserClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleParserClient(cc grpc.ClientConnInterface) ArticleParserClient {
	return &articleParserClient{cc}
}

func (c *articleParserClient) Parse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error) {
	out := new(ParserResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleParserClient) List(ctx context.Context, in *ListParsedArticlesRequest, opts ...grpc.CallOption) (*ListParsedArticlesResponse, error) {
	out := new(ListParsedArticlesResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleParserClient) Get(ctx context.Context, in *GetParsedArticleRequest, opts ...grpc.CallOption) (*GetParsedArticleResponse, error) {
	out := new(GetParsedArticleResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleParserClient) GetParsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetParsersResponse, error) {
	out := new(GetParsersResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/GetParsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleParserClient) GetNewspapers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNewspapersResponse, error) {
	out := new(GetNewspapersResponse)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/GetNewspapers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleParserClient) Update(ctx context.Context, in *ParsedArticle, opts ...grpc.CallOption) (*ParsedArticle, error) {
	out := new(ParsedArticle)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleParserClient) Delete(ctx context.Context, in *DeleteParsedArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/whopper.api.v1.ArticleParser/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleParserServer is the server API for ArticleParser service.
// All implementations must embed UnimplementedArticleParserServer
// for forward compatibility
type ArticleParserServer interface {
	Parse(context.Context, *ParserRequest) (*ParserResponse, error)
	List(context.Context, *ListParsedArticlesRequest) (*ListParsedArticlesResponse, error)
	Get(context.Context, *GetParsedArticleRequest) (*GetParsedArticleResponse, error)
	// Get a list of parsers that can be used to parse a newspaper website
	// Its possible to have multiple parsers per website
	GetParsers(context.Context, *emptypb.Empty) (*GetParsersResponse, error)
	// Get a list of supported newspapers
	GetNewspapers(context.Context, *emptypb.Empty) (*GetNewspapersResponse, error)
	Update(context.Context, *ParsedArticle) (*ParsedArticle, error)
	Delete(context.Context, *DeleteParsedArticleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedArticleParserServer()
}

// UnimplementedArticleParserServer must be embedded to have forward compatible implementations.
type UnimplementedArticleParserServer struct {
}

func (UnimplementedArticleParserServer) Parse(context.Context, *ParserRequest) (*ParserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedArticleParserServer) List(context.Context, *ListParsedArticlesRequest) (*ListParsedArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedArticleParserServer) Get(context.Context, *GetParsedArticleRequest) (*GetParsedArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedArticleParserServer) GetParsers(context.Context, *emptypb.Empty) (*GetParsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParsers not implemented")
}
func (UnimplementedArticleParserServer) GetNewspapers(context.Context, *emptypb.Empty) (*GetNewspapersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewspapers not implemented")
}
func (UnimplementedArticleParserServer) Update(context.Context, *ParsedArticle) (*ParsedArticle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArticleParserServer) Delete(context.Context, *DeleteParsedArticleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArticleParserServer) mustEmbedUnimplementedArticleParserServer() {}

// UnsafeArticleParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleParserServer will
// result in compilation errors.
type UnsafeArticleParserServer interface {
	mustEmbedUnimplementedArticleParserServer()
}

func RegisterArticleParserServer(s grpc.ServiceRegistrar, srv ArticleParserServer) {
	s.RegisterService(&ArticleParser_ServiceDesc, srv)
}

func _ArticleParser_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).Parse(ctx, req.(*ParserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleParser_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListParsedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).List(ctx, req.(*ListParsedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleParser_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParsedArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).Get(ctx, req.(*GetParsedArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleParser_GetParsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).GetParsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/GetParsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).GetParsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleParser_GetNewspapers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).GetNewspapers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/GetNewspapers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).GetNewspapers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleParser_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParsedArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).Update(ctx, req.(*ParsedArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleParser_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParsedArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleParserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whopper.api.v1.ArticleParser/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleParserServer).Delete(ctx, req.(*DeleteParsedArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleParser_ServiceDesc is the grpc.ServiceDesc for ArticleParser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleParser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whopper.api.v1.ArticleParser",
	HandlerType: (*ArticleParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _ArticleParser_Parse_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ArticleParser_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ArticleParser_Get_Handler,
		},
		{
			MethodName: "GetParsers",
			Handler:    _ArticleParser_GetParsers_Handler,
		},
		{
			MethodName: "GetNewspapers",
			Handler:    _ArticleParser_GetNewspapers_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ArticleParser_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ArticleParser_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/parser/parser.proto",
}
