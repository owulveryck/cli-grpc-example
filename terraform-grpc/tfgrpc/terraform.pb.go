// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tfgrpc/terraform.proto

/*
Package tfgrpc is a generated protocol buffer package.

It is generated from these files:
	tfgrpc/terraform.proto

It has these top-level messages:
	Arg
	Output
*/
package tfgrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type Arg struct {
	WorkingDir string   `protobuf:"bytes,1,opt,name=workingDir" json:"workingDir,omitempty"`
	Args       []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *Arg) Reset()                    { *m = Arg{} }
func (m *Arg) String() string            { return proto.CompactTextString(m) }
func (*Arg) ProtoMessage()               {}
func (*Arg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Arg) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *Arg) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// The response message containing the greetings
type Output struct {
	Retcode int32  `protobuf:"varint,1,opt,name=retcode" json:"retcode,omitempty"`
	Stdout  []byte `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr  []byte `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Output) GetRetcode() int32 {
	if m != nil {
		return m.Retcode
	}
	return 0
}

func (m *Output) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Output) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func init() {
	proto.RegisterType((*Arg)(nil), "tfgrpc.Arg")
	proto.RegisterType((*Output)(nil), "tfgrpc.Output")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Terraform service

type TerraformClient interface {
	Plan(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error)
	Apply(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error)
}

type terraformClient struct {
	cc *grpc.ClientConn
}

func NewTerraformClient(cc *grpc.ClientConn) TerraformClient {
	return &terraformClient{cc}
}

func (c *terraformClient) Plan(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := grpc.Invoke(ctx, "/tfgrpc.Terraform/Plan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraformClient) Apply(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := grpc.Invoke(ctx, "/tfgrpc.Terraform/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Terraform service

type TerraformServer interface {
	Plan(context.Context, *Arg) (*Output, error)
	Apply(context.Context, *Arg) (*Output, error)
}

func RegisterTerraformServer(s *grpc.Server, srv TerraformServer) {
	s.RegisterService(&_Terraform_serviceDesc, srv)
}

func _Terraform_Plan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraformServer).Plan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfgrpc.Terraform/Plan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraformServer).Plan(ctx, req.(*Arg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terraform_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraformServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfgrpc.Terraform/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraformServer).Apply(ctx, req.(*Arg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Terraform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tfgrpc.Terraform",
	HandlerType: (*TerraformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plan",
			Handler:    _Terraform_Plan_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Terraform_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tfgrpc/terraform.proto",
}

func init() { proto.RegisterFile("tfgrpc/terraform.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x86, 0xcd, 0xd7, 0x4a, 0x46, 0xb1, 0x98, 0x22, 0x2c, 0x16, 0x12, 0x02, 0x4a, 0xaa, 0x08,
	0x5a, 0x59, 0x06, 0xec, 0x95, 0xc5, 0xca, 0x2e, 0x97, 0x6c, 0x96, 0x70, 0xb9, 0xec, 0x32, 0x99,
	0x70, 0xdc, 0xbf, 0x3f, 0xc8, 0xc7, 0x71, 0xdd, 0x75, 0xf3, 0x3c, 0xcc, 0x30, 0xef, 0x0b, 0x09,
	0xb7, 0x86, 0x5c, 0xfd, 0xce, 0x9a, 0xa8, 0x6a, 0x2d, 0x1d, 0x0a, 0x47, 0x96, 0x2d, 0x8a, 0xc5,
	0x67, 0x5f, 0x10, 0x94, 0x64, 0xf0, 0x05, 0xe0, 0x68, 0x69, 0xdf, 0x0d, 0xe6, 0xbb, 0x23, 0xe9,
	0xa5, 0x5e, 0x1e, 0xab, 0x2b, 0x83, 0x08, 0x61, 0x45, 0x66, 0x94, 0x7e, 0x1a, 0xe4, 0xb1, 0x9a,
	0xe7, 0x4c, 0x81, 0xf8, 0x99, 0xd8, 0x4d, 0x8c, 0x12, 0xee, 0x49, 0x73, 0x6d, 0x1b, 0x3d, 0x9f,
	0x46, 0x6a, 0x43, 0x4c, 0x40, 0x8c, 0xdc, 0xd8, 0x89, 0xa5, 0x9f, 0x7a, 0xf9, 0xa3, 0x5a, 0x69,
	0xf5, 0x9a, 0x48, 0x06, 0x17, 0xaf, 0x89, 0x3e, 0xfe, 0x21, 0xfe, 0xdb, 0x92, 0xe2, 0x2b, 0x84,
	0xbf, 0x7d, 0x35, 0xe0, 0x43, 0xb1, 0x84, 0x2d, 0x4a, 0x32, 0xcf, 0x4f, 0x1b, 0x2c, 0xbf, 0xb3,
	0x3b, 0x7c, 0x83, 0xa8, 0x74, 0xae, 0x3f, 0xdd, 0xd8, 0xdb, 0x89, 0xb9, 0xf9, 0xe7, 0x39, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x49, 0x82, 0x25, 0x13, 0x01, 0x00, 0x00,
}