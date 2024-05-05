// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/feemarket/v1/query.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest defines the request type for querying x/evm parameters.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse defines the response type for querying x/evm parameters.
type QueryParamsResponse struct {
	// params define the evm module parameters.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryBaseFeeRequest defines the request type for querying the EIP1559 base
// fee.
type QueryBaseFeeRequest struct {
}

func (m *QueryBaseFeeRequest) Reset()         { *m = QueryBaseFeeRequest{} }
func (m *QueryBaseFeeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBaseFeeRequest) ProtoMessage()    {}
func (*QueryBaseFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{2}
}
func (m *QueryBaseFeeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBaseFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBaseFeeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBaseFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBaseFeeRequest.Merge(m, src)
}
func (m *QueryBaseFeeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBaseFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBaseFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBaseFeeRequest proto.InternalMessageInfo

// QueryBaseFeeResponse returns the EIP1559 base fee.
type QueryBaseFeeResponse struct {
	// base_fee is the EIP1559 base fee
	BaseFee *cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=base_fee,json=baseFee,proto3,customtype=cosmossdk.io/math.Int" json:"base_fee,omitempty"`
}

func (m *QueryBaseFeeResponse) Reset()         { *m = QueryBaseFeeResponse{} }
func (m *QueryBaseFeeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBaseFeeResponse) ProtoMessage()    {}
func (*QueryBaseFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{3}
}
func (m *QueryBaseFeeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBaseFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBaseFeeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBaseFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBaseFeeResponse.Merge(m, src)
}
func (m *QueryBaseFeeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBaseFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBaseFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBaseFeeResponse proto.InternalMessageInfo

// QueryBlockGasRequest defines the request type for querying the EIP1559 base
// fee.
type QueryBlockGasRequest struct {
}

func (m *QueryBlockGasRequest) Reset()         { *m = QueryBlockGasRequest{} }
func (m *QueryBlockGasRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasRequest) ProtoMessage()    {}
func (*QueryBlockGasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{4}
}
func (m *QueryBlockGasRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasRequest.Merge(m, src)
}
func (m *QueryBlockGasRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasRequest proto.InternalMessageInfo

// QueryBlockGasResponse returns block gas used for a given height.
type QueryBlockGasResponse struct {
	// gas is the returned block gas
	Gas int64 `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
}

func (m *QueryBlockGasResponse) Reset()         { *m = QueryBlockGasResponse{} }
func (m *QueryBlockGasResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasResponse) ProtoMessage()    {}
func (*QueryBlockGasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{5}
}
func (m *QueryBlockGasResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasResponse.Merge(m, src)
}
func (m *QueryBlockGasResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasResponse proto.InternalMessageInfo

func (m *QueryBlockGasResponse) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "ethermint.feemarket.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "ethermint.feemarket.v1.QueryParamsResponse")
	proto.RegisterType((*QueryBaseFeeRequest)(nil), "ethermint.feemarket.v1.QueryBaseFeeRequest")
	proto.RegisterType((*QueryBaseFeeResponse)(nil), "ethermint.feemarket.v1.QueryBaseFeeResponse")
	proto.RegisterType((*QueryBlockGasRequest)(nil), "ethermint.feemarket.v1.QueryBlockGasRequest")
	proto.RegisterType((*QueryBlockGasResponse)(nil), "ethermint.feemarket.v1.QueryBlockGasResponse")
}

func init() {
	proto.RegisterFile("ethermint/feemarket/v1/query.proto", fileDescriptor_71a07c1ffd85fde2)
}

var fileDescriptor_71a07c1ffd85fde2 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x18, 0x4c, 0xdb, 0xf1, 0xa2, 0x63, 0x52, 0x74, 0x89, 0x1b, 0x59, 0x54, 0xb4, 0xd6,
	0x19, 0x52, 0xc5, 0xbb, 0x39, 0x54, 0x04, 0x0f, 0x9a, 0x93, 0x78, 0x29, 0x93, 0xf8, 0xba, 0x59,
	0xd2, 0xd9, 0xb7, 0xdd, 0x99, 0x04, 0x7b, 0x15, 0xbc, 0x78, 0x10, 0xc1, 0x3f, 0xe1, 0x45, 0xf0,
	0x67, 0xf4, 0x58, 0xf0, 0x22, 0x1e, 0x8a, 0x24, 0x82, 0x7f, 0x43, 0x32, 0x33, 0x1b, 0x5d, 0x4d,
	0x34, 0x97, 0xe1, 0xf1, 0xed, 0xf7, 0xde, 0xf7, 0xbd, 0xef, 0xb1, 0x34, 0x02, 0x33, 0x80, 0x5c,
	0x25, 0xa9, 0x11, 0xfb, 0x00, 0x4a, 0xe6, 0x43, 0x30, 0x62, 0xdc, 0x16, 0x87, 0x23, 0xc8, 0x8f,
	0x78, 0x96, 0xa3, 0x41, 0xb6, 0x39, 0xe7, 0xf0, 0x39, 0x87, 0x8f, 0xdb, 0xc1, 0x8d, 0x25, 0xbd,
	0xbf, 0x48, 0xb6, 0x3f, 0xa8, 0xc7, 0x18, 0xa3, 0x2d, 0xc5, 0xac, 0xf2, 0x68, 0x33, 0x46, 0x8c,
	0x0f, 0x40, 0xc8, 0x2c, 0x11, 0x32, 0x4d, 0xd1, 0x48, 0x93, 0x60, 0xaa, 0xfd, 0xd7, 0x0b, 0x52,
	0x25, 0x29, 0x0a, 0xfb, 0x3a, 0x28, 0xaa, 0x53, 0xf6, 0x74, 0xe6, 0xea, 0x89, 0xcc, 0xa5, 0xd2,
	0x5d, 0x38, 0x1c, 0x81, 0x36, 0xd1, 0x33, 0x7a, 0xb1, 0x84, 0xea, 0x0c, 0x53, 0x0d, 0xec, 0x01,
	0xad, 0x65, 0x16, 0xb9, 0x44, 0xae, 0x92, 0x9b, 0xe7, 0x76, 0x42, 0xbe, 0x78, 0x09, 0xee, 0xfa,
	0x3a, 0x1b, 0xc7, 0xa7, 0xad, 0xca, 0x87, 0x1f, 0x9f, 0xb6, 0x48, 0xd7, 0x37, 0x46, 0x0d, 0x3f,
	0xb9, 0x23, 0x35, 0xec, 0x02, 0x14, 0x82, 0x8f, 0x69, 0xbd, 0x0c, 0x7b, 0xc5, 0x7b, 0x74, 0xbd,
	0x27, 0x35, 0xec, 0xed, 0x03, 0x58, 0xcd, 0x8d, 0xce, 0xe5, 0xaf, 0xa7, 0xad, 0x46, 0x1f, 0xb5,
	0x42, 0xad, 0x5f, 0x0c, 0x79, 0x82, 0x42, 0x49, 0x33, 0xe0, 0x8f, 0x52, 0xd3, 0x5d, 0xeb, 0xb9,
	0xee, 0x68, 0xb3, 0x98, 0x76, 0x80, 0xfd, 0xe1, 0x43, 0x39, 0x5f, 0xeb, 0x16, 0x6d, 0xfc, 0x81,
	0x7b, 0x99, 0xf3, 0xb4, 0x1a, 0x4b, 0xb7, 0x55, 0xb5, 0x3b, 0x2b, 0x77, 0x3e, 0x56, 0xe9, 0x59,
	0xcb, 0x65, 0xaf, 0x09, 0xad, 0xb9, 0x7d, 0xd8, 0xd6, 0xb2, 0x7d, 0xff, 0x8e, 0x30, 0xb8, 0xbd,
	0x12, 0xd7, 0xe9, 0x47, 0xd1, 0xab, 0xcf, 0xdf, 0xdf, 0x9f, 0x69, 0xb2, 0x40, 0xc0, 0x58, 0xa1,
	0x2e, 0x5f, 0xde, 0x25, 0xc7, 0xde, 0x10, 0xba, 0xe6, 0xe3, 0x61, 0xff, 0x1e, 0x5e, 0xce, 0x36,
	0xd8, 0x5e, 0x8d, 0xec, 0xad, 0x5c, 0xb3, 0x56, 0x42, 0xd6, 0x5c, 0x64, 0xa5, 0xb8, 0x05, 0x7b,
	0x4b, 0xe8, 0x7a, 0x91, 0x22, 0xfb, 0x8f, 0x40, 0xf9, 0x08, 0xc1, 0x9d, 0x15, 0xd9, 0xde, 0xcf,
	0x75, 0xeb, 0xa7, 0xc5, 0xae, 0x2c, 0xf4, 0x33, 0x63, 0xef, 0xc5, 0x52, 0x77, 0x76, 0x8f, 0x27,
	0x21, 0x39, 0x99, 0x84, 0xe4, 0xdb, 0x24, 0x24, 0xef, 0xa6, 0x61, 0xe5, 0x64, 0x1a, 0x56, 0xbe,
	0x4c, 0xc3, 0xca, 0xf3, 0xed, 0x38, 0x31, 0x83, 0x51, 0x8f, 0xf7, 0x51, 0xf9, 0x11, 0xee, 0x1d,
	0xb7, 0xef, 0x8b, 0x97, 0xbf, 0x8d, 0x33, 0x47, 0x19, 0xe8, 0x5e, 0xcd, 0xfe, 0x16, 0x77, 0x7f,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x4c, 0x1b, 0x97, 0xc3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of x/feemarket module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// BaseFee queries the base fee of the parent block of the current block.
	BaseFee(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error)
	// BlockGas queries the gas used at a given block height
	BlockGas(ctx context.Context, in *QueryBlockGasRequest, opts ...grpc.CallOption) (*QueryBlockGasResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ethermint.feemarket.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BaseFee(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error) {
	out := new(QueryBaseFeeResponse)
	err := c.cc.Invoke(ctx, "/ethermint.feemarket.v1.Query/BaseFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BlockGas(ctx context.Context, in *QueryBlockGasRequest, opts ...grpc.CallOption) (*QueryBlockGasResponse, error) {
	out := new(QueryBlockGasResponse)
	err := c.cc.Invoke(ctx, "/ethermint.feemarket.v1.Query/BlockGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of x/feemarket module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// BaseFee queries the base fee of the parent block of the current block.
	BaseFee(context.Context, *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error)
	// BlockGas queries the gas used at a given block height
	BlockGas(context.Context, *QueryBlockGasRequest) (*QueryBlockGasResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) BaseFee(ctx context.Context, req *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseFee not implemented")
}
func (*UnimplementedQueryServer) BlockGas(ctx context.Context, req *QueryBlockGasRequest) (*QueryBlockGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGas not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.feemarket.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BaseFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBaseFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.feemarket.v1.Query/BaseFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseFee(ctx, req.(*QueryBaseFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BlockGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.feemarket.v1.Query/BlockGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockGas(ctx, req.(*QueryBlockGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethermint.feemarket.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "BaseFee",
			Handler:    _Query_BaseFee_Handler,
		},
		{
			MethodName: "BlockGas",
			Handler:    _Query_BlockGas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethermint/feemarket/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryBaseFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBaseFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBaseFeeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBaseFeeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBaseFeeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBaseFeeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseFee != nil {
		{
			size := m.BaseFee.Size()
			i -= size
			if _, err := m.BaseFee.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBlockGasRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBlockGasResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Gas != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Gas))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryBaseFeeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBaseFeeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseFee != nil {
		l = m.BaseFee.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBlockGasRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBlockGasResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Gas != 0 {
		n += 1 + sovQuery(uint64(m.Gas))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBaseFeeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBaseFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBaseFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBaseFeeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBaseFeeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBaseFeeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.BaseFee = &v
			if err := m.BaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBlockGasRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBlockGasRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBlockGasResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBlockGasResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			m.Gas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gas |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
