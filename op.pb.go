// Code generated by protoc-gen-go.
// source: op.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	op.proto

It has these top-level messages:
	OpRequest
	OpReply
*/
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math"
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

type OpRequest struct {
	Op     string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Bucket string `protobuf:"bytes,4,opt,name=bucket" json:"bucket,omitempty"`
}

func (m *OpRequest) Reset()                    { *m = OpRequest{} }
func (m *OpRequest) String() string            { return proto.CompactTextString(m) }
func (*OpRequest) ProtoMessage()               {}
func (*OpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OpRequest) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *OpRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *OpRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *OpRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

type OpReply struct {
	Status  int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Value   []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	ErrCode int32  `protobuf:"varint,4,opt,name=errCode" json:"errCode,omitempty"`
}

func (m *OpReply) Reset()                    { *m = OpReply{} }
func (m *OpReply) String() string            { return proto.CompactTextString(m) }
func (*OpReply) ProtoMessage()               {}
func (*OpReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OpReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OpReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *OpReply) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *OpReply) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func init() {
	proto.RegisterType((*OpRequest)(nil), "pb.OpRequest")
	proto.RegisterType((*OpReply)(nil), "pb.OpReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RaftRpc service

type RaftRpcClient interface {
	OpRPC(ctx context.Context, in *OpRequest, opts ...grpc.CallOption) (*OpReply, error)
}

type raftRpcClient struct {
	cc *grpc.ClientConn
}

func NewRaftRpcClient(cc *grpc.ClientConn) RaftRpcClient {
	return &raftRpcClient{cc}
}

func (c *raftRpcClient) OpRPC(ctx context.Context, in *OpRequest, opts ...grpc.CallOption) (*OpReply, error) {
	out := new(OpReply)
	err := grpc.Invoke(ctx, "/pb.RaftRpc/OpRPC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RaftRpc service

type RaftRpcServer interface {
	OpRPC(context.Context, *OpRequest) (*OpReply, error)
}

func RegisterRaftRpcServer(s *grpc.Server, srv RaftRpcServer) {
	s.RegisterService(&_RaftRpc_serviceDesc, srv)
}

func _RaftRpc_OpRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftRpcServer).OpRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RaftRpc/OpRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftRpcServer).OpRPC(ctx, req.(*OpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RaftRpc",
	HandlerType: (*RaftRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpRPC",
			Handler:    _RaftRpc_OpRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "op.proto",
}

func init() { proto.RegisterFile("op.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8f, 0xcf, 0x8a, 0xc2, 0x30,
	0x10, 0xc6, 0xb7, 0xed, 0xa6, 0xdd, 0xce, 0xfe, 0x61, 0x09, 0x22, 0xc1, 0x93, 0x14, 0x04, 0x4f,
	0x45, 0xf4, 0x11, 0xfa, 0x00, 0x4a, 0xae, 0x9e, 0x9a, 0x1a, 0x3d, 0xb4, 0x92, 0xb1, 0x49, 0x84,
	0xbe, 0xbd, 0x49, 0x6c, 0xbd, 0x79, 0xfb, 0x7e, 0xc3, 0x30, 0xbf, 0xf9, 0xe0, 0x4b, 0x61, 0x89,
	0xbd, 0x32, 0x8a, 0xc6, 0x28, 0x8a, 0x23, 0xe4, 0x7b, 0xe4, 0xf2, 0x66, 0xa5, 0x36, 0xf4, 0x0f,
	0x62, 0x85, 0x2c, 0x5a, 0x46, 0xeb, 0x9c, 0xbb, 0x44, 0xff, 0x21, 0x69, 0xe5, 0xc0, 0xe2, 0x30,
	0xf0, 0x91, 0xce, 0x80, 0xdc, 0xeb, 0xce, 0x4a, 0x96, 0xb8, 0xd9, 0x0f, 0x7f, 0x02, 0x9d, 0x43,
	0x2a, 0x6c, 0xd3, 0x4a, 0xc3, 0x3e, 0xc3, 0xea, 0x48, 0x45, 0x03, 0x99, 0x3f, 0x8e, 0xdd, 0xe0,
	0x57, 0xb4, 0xa9, 0x8d, 0xd5, 0xe1, 0x3c, 0xe1, 0x23, 0x79, 0xc5, 0x55, 0x5f, 0x26, 0x85, 0x8b,
	0x6f, 0x14, 0x0c, 0x32, 0xd9, 0xf7, 0x95, 0x3a, 0xc9, 0xe0, 0x20, 0x7c, 0xc2, 0xed, 0x06, 0x32,
	0x5e, 0x9f, 0x0d, 0xc7, 0x86, 0xae, 0x80, 0x38, 0xdf, 0xa1, 0xa2, 0xbf, 0x25, 0x8a, 0xf2, 0xd5,
	0x6b, 0xf1, 0x3d, 0xa1, 0xfb, 0xa4, 0xf8, 0x10, 0x69, 0xa8, 0xbf, 0x7b, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0x2b, 0xcb, 0x73, 0x0a, 0x01, 0x00, 0x00,
}
