// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: code-analyzer/code-analyzer.proto

package protoc

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

// CodeAnalyzerServiceClient is the client API for CodeAnalyzerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeAnalyzerServiceClient interface {
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	AnalyzeUploader(ctx context.Context, opts ...grpc.CallOption) (CodeAnalyzerService_AnalyzeUploaderClient, error)
	FileUpload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error)
}

type codeAnalyzerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeAnalyzerServiceClient(cc grpc.ClientConnInterface) CodeAnalyzerServiceClient {
	return &codeAnalyzerServiceClient{cc}
}

func (c *codeAnalyzerServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/codeanalyzer.CodeAnalyzerService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeAnalyzerServiceClient) AnalyzeUploader(ctx context.Context, opts ...grpc.CallOption) (CodeAnalyzerService_AnalyzeUploaderClient, error) {
	stream, err := c.cc.NewStream(ctx, &CodeAnalyzerService_ServiceDesc.Streams[0], "/codeanalyzer.CodeAnalyzerService/AnalyzeUploader", opts...)
	if err != nil {
		return nil, err
	}
	x := &codeAnalyzerServiceAnalyzeUploaderClient{stream}
	return x, nil
}

type CodeAnalyzerService_AnalyzeUploaderClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type codeAnalyzerServiceAnalyzeUploaderClient struct {
	grpc.ClientStream
}

func (x *codeAnalyzerServiceAnalyzeUploaderClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *codeAnalyzerServiceAnalyzeUploaderClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *codeAnalyzerServiceClient) FileUpload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error) {
	out := new(FileUploadResponse)
	err := c.cc.Invoke(ctx, "/codeanalyzer.CodeAnalyzerService/FileUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeAnalyzerServiceServer is the server API for CodeAnalyzerService service.
// All implementations must embed UnimplementedCodeAnalyzerServiceServer
// for forward compatibility
type CodeAnalyzerServiceServer interface {
	HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error)
	AnalyzeUploader(CodeAnalyzerService_AnalyzeUploaderServer) error
	FileUpload(context.Context, *FileUploadRequest) (*FileUploadResponse, error)
	mustEmbedUnimplementedCodeAnalyzerServiceServer()
}

// UnimplementedCodeAnalyzerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCodeAnalyzerServiceServer struct {
}

func (UnimplementedCodeAnalyzerServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedCodeAnalyzerServiceServer) AnalyzeUploader(CodeAnalyzerService_AnalyzeUploaderServer) error {
	return status.Errorf(codes.Unimplemented, "method AnalyzeUploader not implemented")
}
func (UnimplementedCodeAnalyzerServiceServer) FileUpload(context.Context, *FileUploadRequest) (*FileUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUpload not implemented")
}
func (UnimplementedCodeAnalyzerServiceServer) mustEmbedUnimplementedCodeAnalyzerServiceServer() {}

// UnsafeCodeAnalyzerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeAnalyzerServiceServer will
// result in compilation errors.
type UnsafeCodeAnalyzerServiceServer interface {
	mustEmbedUnimplementedCodeAnalyzerServiceServer()
}

func RegisterCodeAnalyzerServiceServer(s grpc.ServiceRegistrar, srv CodeAnalyzerServiceServer) {
	s.RegisterService(&CodeAnalyzerService_ServiceDesc, srv)
}

func _CodeAnalyzerService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeAnalyzerServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codeanalyzer.CodeAnalyzerService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeAnalyzerServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeAnalyzerService_AnalyzeUploader_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CodeAnalyzerServiceServer).AnalyzeUploader(&codeAnalyzerServiceAnalyzeUploaderServer{stream})
}

type CodeAnalyzerService_AnalyzeUploaderServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type codeAnalyzerServiceAnalyzeUploaderServer struct {
	grpc.ServerStream
}

func (x *codeAnalyzerServiceAnalyzeUploaderServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *codeAnalyzerServiceAnalyzeUploaderServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CodeAnalyzerService_FileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeAnalyzerServiceServer).FileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codeanalyzer.CodeAnalyzerService/FileUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeAnalyzerServiceServer).FileUpload(ctx, req.(*FileUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeAnalyzerService_ServiceDesc is the grpc.ServiceDesc for CodeAnalyzerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeAnalyzerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codeanalyzer.CodeAnalyzerService",
	HandlerType: (*CodeAnalyzerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _CodeAnalyzerService_HealthCheck_Handler,
		},
		{
			MethodName: "FileUpload",
			Handler:    _CodeAnalyzerService_FileUpload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AnalyzeUploader",
			Handler:       _CodeAnalyzerService_AnalyzeUploader_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "code-analyzer/code-analyzer.proto",
}
