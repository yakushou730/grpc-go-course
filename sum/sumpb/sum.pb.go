// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum/sumpb/sum.proto

package sumpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SumRequest struct {
	FirstNum             int32    `protobuf:"varint,1,opt,name=first_num,json=firstNum,proto3" json:"first_num,omitempty"`
	SecondNum            int32    `protobuf:"varint,2,opt,name=second_num,json=secondNum,proto3" json:"second_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e82cd6e6acf87a3, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstNum() int32 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

func (m *SumRequest) GetSecondNum() int32 {
	if m != nil {
		return m.SecondNum
	}
	return 0
}

type SumResponse struct {
	Sum                  int32    `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e82cd6e6acf87a3, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSum() int32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sum.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sum.SumResponse")
}

func init() { proto.RegisterFile("sum/sumpb/sum.proto", fileDescriptor_1e82cd6e6acf87a3) }

var fileDescriptor_1e82cd6e6acf87a3 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0xcd, 0xd5,
	0x2f, 0x2e, 0xcd, 0x2d, 0x48, 0x02, 0x91, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xcc, 0xc5,
	0xa5, 0xb9, 0x4a, 0x1e, 0x5c, 0x5c, 0xc1, 0xa5, 0xb9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0xd2, 0x5c, 0x9c, 0x69, 0x99, 0x45, 0xc5, 0x25, 0xf1, 0x79, 0xa5, 0xb9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xac, 0x41, 0x1c, 0x60, 0x01, 0xbf, 0xd2, 0x5c, 0x21, 0x59, 0x2e, 0xae, 0xe2, 0xd4,
	0xe4, 0xfc, 0xbc, 0x14, 0xb0, 0x2c, 0x13, 0x58, 0x96, 0x13, 0x22, 0xe2, 0x57, 0x9a, 0xab, 0x24,
	0xcf, 0xc5, 0x0d, 0x36, 0xa9, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x0b, 0x64, 0x3e,
	0xd4, 0x10, 0x10, 0xd3, 0xc8, 0x02, 0x6c, 0x55, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90,
	0x16, 0x17, 0x73, 0x70, 0x69, 0xae, 0x10, 0xbf, 0x1e, 0xc8, 0x41, 0x08, 0x27, 0x48, 0x09, 0x20,
	0x04, 0x20, 0x26, 0x29, 0x31, 0x38, 0xb1, 0x47, 0xb1, 0x82, 0x1d, 0x9f, 0xc4, 0x06, 0x76, 0xb9,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x89, 0x1e, 0x57, 0xd0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sum.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum/sumpb/sum.proto",
}
