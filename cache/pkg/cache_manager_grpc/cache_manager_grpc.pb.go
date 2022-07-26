// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/cache_manager.proto

package cache_manager_grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CacheManagerClient is the client API for CacheManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheManagerClient interface {
	GetByKey(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Clear(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Has(ctx context.Context, in *HasRequest, opts ...grpc.CallOption) (*HasResponse, error)
}

type cacheManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheManagerClient(cc grpc.ClientConnInterface) CacheManagerClient {
	return &cacheManagerClient{cc}
}

func (c *cacheManagerClient) GetByKey(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/cache_manager.CacheManager/GetByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheManagerClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/cache_manager.CacheManager/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheManagerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cache_manager.CacheManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheManagerClient) Clear(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cache_manager.CacheManager/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheManagerClient) Has(ctx context.Context, in *HasRequest, opts ...grpc.CallOption) (*HasResponse, error) {
	out := new(HasResponse)
	err := c.cc.Invoke(ctx, "/cache_manager.CacheManager/Has", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheManagerServer is the server API for CacheManager service.
// All implementations must embed UnimplementedCacheManagerServer
// for forward compatibility
type CacheManagerServer interface {
	GetByKey(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	Clear(context.Context, *empty.Empty) (*empty.Empty, error)
	Has(context.Context, *HasRequest) (*HasResponse, error)
	mustEmbedUnimplementedCacheManagerServer()
}

// UnimplementedCacheManagerServer must be embedded to have forward compatible implementations.
type UnimplementedCacheManagerServer struct {
}

func (UnimplementedCacheManagerServer) GetByKey(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByKey not implemented")
}
func (UnimplementedCacheManagerServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheManagerServer) Delete(context.Context, *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCacheManagerServer) Clear(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedCacheManagerServer) Has(context.Context, *HasRequest) (*HasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Has not implemented")
}
func (UnimplementedCacheManagerServer) mustEmbedUnimplementedCacheManagerServer() {}

// UnsafeCacheManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheManagerServer will
// result in compilation errors.
type UnsafeCacheManagerServer interface {
	mustEmbedUnimplementedCacheManagerServer()
}

func RegisterCacheManagerServer(s grpc.ServiceRegistrar, srv CacheManagerServer) {
	s.RegisterService(&CacheManager_ServiceDesc, srv)
}

func _CacheManager_GetByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheManagerServer).GetByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache_manager.CacheManager/GetByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheManagerServer).GetByKey(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheManager_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheManagerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache_manager.CacheManager/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheManagerServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache_manager.CacheManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheManagerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheManager_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheManagerServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache_manager.CacheManager/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheManagerServer).Clear(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheManager_Has_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheManagerServer).Has(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache_manager.CacheManager/Has",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheManagerServer).Has(ctx, req.(*HasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheManager_ServiceDesc is the grpc.ServiceDesc for CacheManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cache_manager.CacheManager",
	HandlerType: (*CacheManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByKey",
			Handler:    _CacheManager_GetByKey_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CacheManager_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CacheManager_Delete_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _CacheManager_Clear_Handler,
		},
		{
			MethodName: "Has",
			Handler:    _CacheManager_Has_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cache_manager.proto",
}
