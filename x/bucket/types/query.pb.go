// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bucket/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_193313c7681ac565, []int{0}
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

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_193313c7681ac565, []int{1}
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

type QueryGetWhereIsRequest struct {
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
}

func (m *QueryGetWhereIsRequest) Reset()         { *m = QueryGetWhereIsRequest{} }
func (m *QueryGetWhereIsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetWhereIsRequest) ProtoMessage()    {}
func (*QueryGetWhereIsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_193313c7681ac565, []int{2}
}
func (m *QueryGetWhereIsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetWhereIsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetWhereIsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetWhereIsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetWhereIsRequest.Merge(m, src)
}
func (m *QueryGetWhereIsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetWhereIsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetWhereIsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetWhereIsRequest proto.InternalMessageInfo

func (m *QueryGetWhereIsRequest) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

type QueryGetWhereIsResponse struct {
	WhereIs WhereIs `protobuf:"bytes,1,opt,name=WhereIs,proto3" json:"WhereIs"`
}

func (m *QueryGetWhereIsResponse) Reset()         { *m = QueryGetWhereIsResponse{} }
func (m *QueryGetWhereIsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetWhereIsResponse) ProtoMessage()    {}
func (*QueryGetWhereIsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_193313c7681ac565, []int{3}
}
func (m *QueryGetWhereIsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetWhereIsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetWhereIsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetWhereIsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetWhereIsResponse.Merge(m, src)
}
func (m *QueryGetWhereIsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetWhereIsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetWhereIsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetWhereIsResponse proto.InternalMessageInfo

func (m *QueryGetWhereIsResponse) GetWhereIs() WhereIs {
	if m != nil {
		return m.WhereIs
	}
	return WhereIs{}
}

type QueryAllWhereIsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllWhereIsRequest) Reset()         { *m = QueryAllWhereIsRequest{} }
func (m *QueryAllWhereIsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllWhereIsRequest) ProtoMessage()    {}
func (*QueryAllWhereIsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_193313c7681ac565, []int{4}
}
func (m *QueryAllWhereIsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllWhereIsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWhereIsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllWhereIsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWhereIsRequest.Merge(m, src)
}
func (m *QueryAllWhereIsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllWhereIsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllWhereIsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllWhereIsRequest proto.InternalMessageInfo

func (m *QueryAllWhereIsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllWhereIsResponse struct {
	WhereIs    []WhereIs           `protobuf:"bytes,1,rep,name=WhereIs,proto3" json:"WhereIs"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllWhereIsResponse) Reset()         { *m = QueryAllWhereIsResponse{} }
func (m *QueryAllWhereIsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllWhereIsResponse) ProtoMessage()    {}
func (*QueryAllWhereIsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_193313c7681ac565, []int{5}
}
func (m *QueryAllWhereIsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllWhereIsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWhereIsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllWhereIsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWhereIsResponse.Merge(m, src)
}
func (m *QueryAllWhereIsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllWhereIsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllWhereIsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllWhereIsResponse proto.InternalMessageInfo

func (m *QueryAllWhereIsResponse) GetWhereIs() []WhereIs {
	if m != nil {
		return m.WhereIs
	}
	return nil
}

func (m *QueryAllWhereIsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "sonrio.sonr.bucket.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "sonrio.sonr.bucket.QueryParamsResponse")
	proto.RegisterType((*QueryGetWhereIsRequest)(nil), "sonrio.sonr.bucket.QueryGetWhereIsRequest")
	proto.RegisterType((*QueryGetWhereIsResponse)(nil), "sonrio.sonr.bucket.QueryGetWhereIsResponse")
	proto.RegisterType((*QueryAllWhereIsRequest)(nil), "sonrio.sonr.bucket.QueryAllWhereIsRequest")
	proto.RegisterType((*QueryAllWhereIsResponse)(nil), "sonrio.sonr.bucket.QueryAllWhereIsResponse")
}

func init() { proto.RegisterFile("bucket/query.proto", fileDescriptor_193313c7681ac565) }

var fileDescriptor_193313c7681ac565 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe3, 0x06, 0x82, 0x30, 0x17, 0xe4, 0x96, 0x3f, 0xda, 0xd2, 0x2d, 0x72, 0xd5, 0x14,
	0xa5, 0xc2, 0x56, 0xcb, 0x05, 0x89, 0x0b, 0xed, 0x81, 0x8a, 0x13, 0x25, 0x07, 0x90, 0xb8, 0x80,
	0x37, 0xb1, 0xb6, 0x16, 0x9b, 0xf5, 0x76, 0xed, 0x00, 0x15, 0x82, 0x03, 0x67, 0x24, 0x90, 0x78,
	0x00, 0xae, 0x3c, 0x4a, 0x8f, 0x95, 0xb8, 0x70, 0x42, 0x28, 0xe1, 0x41, 0xd0, 0xda, 0xb3, 0x90,
	0xb0, 0x4b, 0x42, 0x4f, 0x6b, 0xcd, 0x7c, 0xf3, 0xcd, 0x6f, 0xc7, 0x63, 0x4c, 0xa2, 0x61, 0xef,
	0xb9, 0xb4, 0xfc, 0x70, 0x28, 0xf3, 0x23, 0x96, 0xe5, 0xda, 0x6a, 0x42, 0x8c, 0x4e, 0x73, 0xa5,
	0x59, 0xf1, 0x61, 0x3e, 0x1f, 0x2c, 0xc5, 0x3a, 0xd6, 0x2e, 0xcd, 0x8b, 0x93, 0x57, 0x06, 0xd7,
	0x62, 0xad, 0xe3, 0x44, 0x72, 0x91, 0x29, 0x2e, 0xd2, 0x54, 0x5b, 0x61, 0x95, 0x4e, 0x0d, 0x64,
	0x3b, 0x3d, 0x6d, 0x06, 0xda, 0xf0, 0x48, 0x18, 0xe9, 0x1b, 0xf0, 0x17, 0x5b, 0x91, 0xb4, 0x62,
	0x8b, 0x67, 0x22, 0x56, 0xa9, 0x13, 0x83, 0x76, 0x11, 0x38, 0x32, 0x91, 0x8b, 0x41, 0x69, 0x70,
	0x09, 0x82, 0x2f, 0x0f, 0x64, 0x2e, 0x9f, 0x2a, 0x08, 0xd3, 0x25, 0x4c, 0x1e, 0x16, 0x6e, 0xfb,
	0x4e, 0xdb, 0x95, 0x87, 0x43, 0x69, 0x2c, 0x7d, 0x80, 0x17, 0xa7, 0xa2, 0x26, 0xd3, 0xa9, 0x91,
	0xe4, 0x36, 0x6e, 0x79, 0xcf, 0xab, 0xe8, 0x3a, 0xba, 0x71, 0x61, 0x3b, 0x60, 0xd5, 0xbf, 0x63,
	0xbe, 0x66, 0xf7, 0xcc, 0xf1, 0xf7, 0xd5, 0x46, 0x17, 0xf4, 0xb4, 0x83, 0x2f, 0x3b, 0xc3, 0x3d,
	0x69, 0x1f, 0x17, 0x00, 0xf7, 0xcb, 0x56, 0xe4, 0x22, 0x6e, 0xf6, 0x55, 0xdf, 0x19, 0x9e, 0xef,
	0x16, 0x47, 0xfa, 0x08, 0x5f, 0xa9, 0x68, 0x01, 0xe0, 0x0e, 0x3e, 0x07, 0x21, 0x20, 0x58, 0xae,
	0x23, 0x00, 0x09, 0x20, 0x94, 0x15, 0xf4, 0x19, 0x30, 0xec, 0x24, 0xc9, 0x5f, 0x0c, 0xf7, 0x30,
	0xfe, 0x33, 0x44, 0x70, 0x6e, 0x33, 0x3f, 0x71, 0x56, 0x4c, 0x9c, 0xf9, 0x2b, 0x85, 0x89, 0xb3,
	0x7d, 0x11, 0x4b, 0xa8, 0xed, 0x4e, 0x54, 0xd2, 0xcf, 0x08, 0xd0, 0x27, 0x5b, 0xd4, 0xa1, 0x37,
	0x4f, 0x87, 0x4e, 0xf6, 0xa6, 0x00, 0x17, 0x1c, 0xe0, 0xc6, 0x5c, 0x40, 0xdf, 0x79, 0x92, 0x70,
	0xfb, 0x4b, 0x13, 0x9f, 0x75, 0x84, 0xe4, 0x2d, 0x6e, 0xf9, 0x9b, 0x22, 0xed, 0x3a, 0x90, 0xea,
	0x52, 0x04, 0x1b, 0x73, 0x75, 0xbe, 0x21, 0x5d, 0x7b, 0xf7, 0xf5, 0xe7, 0xa7, 0x85, 0x15, 0xb2,
	0xcc, 0x0b, 0xe5, 0x4d, 0xa5, 0xdd, 0x97, 0x4f, 0x6d, 0x25, 0xf9, 0x80, 0x7e, 0x0f, 0x84, 0x74,
	0xfe, 0xe9, 0x5c, 0xd9, 0x97, 0x60, 0xf3, 0xbf, 0xb4, 0x40, 0xb2, 0xe9, 0x48, 0xd6, 0xc9, 0x5a,
	0x2d, 0x49, 0xf9, 0x14, 0xf8, 0xeb, 0xbe, 0xea, 0xbf, 0x21, 0xef, 0x11, 0xc6, 0x60, 0xb0, 0x93,
	0x24, 0x33, 0xa0, 0x2a, 0x0b, 0x34, 0x03, 0xaa, 0xba, 0x09, 0x74, 0xdd, 0x41, 0xad, 0x92, 0x95,
	0x99, 0x50, 0xbb, 0x77, 0x8f, 0x47, 0x21, 0x3a, 0x19, 0x85, 0xe8, 0xc7, 0x28, 0x44, 0x1f, 0xc7,
	0x61, 0xe3, 0x64, 0x1c, 0x36, 0xbe, 0x8d, 0xc3, 0xc6, 0x93, 0x76, 0xac, 0xec, 0xc1, 0x30, 0x62,
	0x3d, 0x3d, 0x98, 0xb6, 0x78, 0x55, 0x9a, 0xd8, 0xa3, 0x4c, 0x9a, 0xa8, 0xe5, 0x9e, 0xf8, 0xad,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x8c, 0x2b, 0x01, 0x98, 0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a WhereIs by id.
	WhereIs(ctx context.Context, in *QueryGetWhereIsRequest, opts ...grpc.CallOption) (*QueryGetWhereIsResponse, error)
	// Queries a list of WhereIs items.
	WhereIsAll(ctx context.Context, in *QueryAllWhereIsRequest, opts ...grpc.CallOption) (*QueryAllWhereIsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.bucket.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WhereIs(ctx context.Context, in *QueryGetWhereIsRequest, opts ...grpc.CallOption) (*QueryGetWhereIsResponse, error) {
	out := new(QueryGetWhereIsResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.bucket.Query/WhereIs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WhereIsAll(ctx context.Context, in *QueryAllWhereIsRequest, opts ...grpc.CallOption) (*QueryAllWhereIsResponse, error) {
	out := new(QueryAllWhereIsResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.bucket.Query/WhereIsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a WhereIs by id.
	WhereIs(context.Context, *QueryGetWhereIsRequest) (*QueryGetWhereIsResponse, error)
	// Queries a list of WhereIs items.
	WhereIsAll(context.Context, *QueryAllWhereIsRequest) (*QueryAllWhereIsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) WhereIs(ctx context.Context, req *QueryGetWhereIsRequest) (*QueryGetWhereIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhereIs not implemented")
}
func (*UnimplementedQueryServer) WhereIsAll(ctx context.Context, req *QueryAllWhereIsRequest) (*QueryAllWhereIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhereIsAll not implemented")
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
		FullMethod: "/sonrio.sonr.bucket.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WhereIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetWhereIsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WhereIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.bucket.Query/WhereIs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WhereIs(ctx, req.(*QueryGetWhereIsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WhereIsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllWhereIsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WhereIsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.bucket.Query/WhereIsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WhereIsAll(ctx, req.(*QueryAllWhereIsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonrio.sonr.bucket.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "WhereIs",
			Handler:    _Query_WhereIs_Handler,
		},
		{
			MethodName: "WhereIsAll",
			Handler:    _Query_WhereIsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bucket/query.proto",
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

func (m *QueryGetWhereIsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetWhereIsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetWhereIsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetWhereIsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetWhereIsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetWhereIsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.WhereIs.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllWhereIsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWhereIsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWhereIsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllWhereIsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWhereIsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWhereIsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.WhereIs) > 0 {
		for iNdEx := len(m.WhereIs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhereIs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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

func (m *QueryGetWhereIsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetWhereIsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.WhereIs.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllWhereIsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllWhereIsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhereIs) > 0 {
		for _, e := range m.WhereIs {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
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
func (m *QueryGetWhereIsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetWhereIsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetWhereIsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
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
			m.Did = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetWhereIsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetWhereIsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetWhereIsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhereIs", wireType)
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
			if err := m.WhereIs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllWhereIsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllWhereIsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWhereIsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllWhereIsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllWhereIsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWhereIsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhereIs", wireType)
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
			m.WhereIs = append(m.WhereIs, WhereIs{})
			if err := m.WhereIs[len(m.WhereIs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
