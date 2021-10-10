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
const _ = grpc.SupportPackageIsVersion7

// MovieHandlerClient is the client API for MovieHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieHandlerClient interface {
	FetchMovie(ctx context.Context, in *FetchMovieRequest, opts ...grpc.CallOption) (*MovieData, error)
}

type movieHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieHandlerClient(cc grpc.ClientConnInterface) MovieHandlerClient {
	return &movieHandlerClient{cc}
}

func (c *movieHandlerClient) FetchMovie(ctx context.Context, in *FetchMovieRequest, opts ...grpc.CallOption) (*MovieData, error) {
	out := new(MovieData)
	err := c.cc.Invoke(ctx, "/grpc.MovieHandler/FetchMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieHandlerServer is the server API for MovieHandler service.
// All implementations should embed UnimplementedMovieHandlerServer
// for forward compatibility
type MovieHandlerServer interface {
	FetchMovie(context.Context, *FetchMovieRequest) (*MovieData, error)
}

// UnimplementedMovieHandlerServer should be embedded to have forward compatible implementations.
type UnimplementedMovieHandlerServer struct {
}

func (UnimplementedMovieHandlerServer) FetchMovie(context.Context, *FetchMovieRequest) (*MovieData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMovie not implemented")
}

// UnsafeMovieHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieHandlerServer will
// result in compilation errors.
type UnsafeMovieHandlerServer interface {
	mustEmbedUnimplementedMovieHandlerServer()
}

func RegisterMovieHandlerServer(s *grpc.Server, srv MovieHandlerServer) {
	s.RegisterService(&_MovieHandler_serviceDesc, srv)
}

func _MovieHandler_FetchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieHandlerServer).FetchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MovieHandler/FetchMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieHandlerServer).FetchMovie(ctx, req.(*FetchMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MovieHandler",
	HandlerType: (*MovieHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMovie",
			Handler:    _MovieHandler_FetchMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
