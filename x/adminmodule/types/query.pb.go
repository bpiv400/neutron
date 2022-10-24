// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: adminmodule/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/gov/types"
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

type QueryAdminsRequest struct {
}

func (m *QueryAdminsRequest) Reset()         { *m = QueryAdminsRequest{} }
func (m *QueryAdminsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAdminsRequest) ProtoMessage()    {}
func (*QueryAdminsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f1ff4730a29e9dd, []int{0}
}
func (m *QueryAdminsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminsRequest.Merge(m, src)
}
func (m *QueryAdminsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminsRequest proto.InternalMessageInfo

type QueryAdminsResponse struct {
	Admins []string `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (m *QueryAdminsResponse) Reset()         { *m = QueryAdminsResponse{} }
func (m *QueryAdminsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAdminsResponse) ProtoMessage()    {}
func (*QueryAdminsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f1ff4730a29e9dd, []int{1}
}
func (m *QueryAdminsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminsResponse.Merge(m, src)
}
func (m *QueryAdminsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminsResponse proto.InternalMessageInfo

func (m *QueryAdminsResponse) GetAdmins() []string {
	if m != nil {
		return m.Admins
	}
	return nil
}

type QueryArchivedProposalsRequest struct {
}

func (m *QueryArchivedProposalsRequest) Reset()         { *m = QueryArchivedProposalsRequest{} }
func (m *QueryArchivedProposalsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryArchivedProposalsRequest) ProtoMessage()    {}
func (*QueryArchivedProposalsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f1ff4730a29e9dd, []int{2}
}
func (m *QueryArchivedProposalsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryArchivedProposalsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryArchivedProposalsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryArchivedProposalsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryArchivedProposalsRequest.Merge(m, src)
}
func (m *QueryArchivedProposalsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryArchivedProposalsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryArchivedProposalsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryArchivedProposalsRequest proto.InternalMessageInfo

type QueryArchivedProposalsResponse struct {
	Proposals []*types.Proposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
}

func (m *QueryArchivedProposalsResponse) Reset()         { *m = QueryArchivedProposalsResponse{} }
func (m *QueryArchivedProposalsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryArchivedProposalsResponse) ProtoMessage()    {}
func (*QueryArchivedProposalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f1ff4730a29e9dd, []int{3}
}
func (m *QueryArchivedProposalsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryArchivedProposalsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryArchivedProposalsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryArchivedProposalsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryArchivedProposalsResponse.Merge(m, src)
}
func (m *QueryArchivedProposalsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryArchivedProposalsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryArchivedProposalsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryArchivedProposalsResponse proto.InternalMessageInfo

func (m *QueryArchivedProposalsResponse) GetProposals() []*types.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAdminsRequest)(nil), "cosmos.adminmodule.adminmodule.QueryAdminsRequest")
	proto.RegisterType((*QueryAdminsResponse)(nil), "cosmos.adminmodule.adminmodule.QueryAdminsResponse")
	proto.RegisterType((*QueryArchivedProposalsRequest)(nil), "cosmos.adminmodule.adminmodule.QueryArchivedProposalsRequest")
	proto.RegisterType((*QueryArchivedProposalsResponse)(nil), "cosmos.adminmodule.adminmodule.QueryArchivedProposalsResponse")
}

func init() { proto.RegisterFile("adminmodule/query.proto", fileDescriptor_3f1ff4730a29e9dd) }

var fileDescriptor_3f1ff4730a29e9dd = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4c, 0xc9, 0xcd,
	0xcc, 0xcb, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x43, 0x92, 0x47, 0x66,
	0x4b, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5,
	0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x74, 0x4b, 0xc9, 0x40, 0x74, 0xeb, 0xa7,
	0xe7, 0x97, 0xe9, 0x97, 0x19, 0x26, 0xa5, 0x96, 0x24, 0x1a, 0x82, 0xd8, 0x10, 0x59, 0x25, 0x11,
	0x2e, 0xa1, 0x40, 0x90, 0x55, 0x8e, 0x20, 0xf3, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x94, 0x74, 0xb9, 0x84, 0x51, 0x44, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8,
	0xc0, 0xf6, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x79, 0x4a, 0xf2, 0x5c, 0xb2,
	0x10, 0xe5, 0x45, 0xc9, 0x19, 0x99, 0x65, 0xa9, 0x29, 0x01, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0x89,
	0x39, 0x70, 0xf3, 0x62, 0xb8, 0xe4, 0x70, 0x29, 0x80, 0x1a, 0x6d, 0xc5, 0xc5, 0x59, 0x00, 0x13,
	0x04, 0x9b, 0xce, 0x6d, 0x24, 0xa3, 0x07, 0xf5, 0x37, 0xc8, 0xb5, 0x50, 0x97, 0xeb, 0xc1, 0x74,
	0x06, 0x21, 0x94, 0x1b, 0xfd, 0x66, 0xe2, 0x62, 0x05, 0x1b, 0x2f, 0xb4, 0x90, 0x91, 0x8b, 0x0d,
	0xe2, 0x66, 0x21, 0x23, 0x3d, 0xfc, 0xa1, 0xa6, 0x87, 0xe9, 0x6d, 0x29, 0x63, 0x92, 0xf4, 0x40,
	0x5c, 0xae, 0xa4, 0xd7, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x0d, 0x21, 0x35, 0x7d, 0x68, 0x40, 0x23,
	0x47, 0x23, 0x06, 0xbb, 0x58, 0xe8, 0x2c, 0x23, 0x97, 0x20, 0x46, 0x38, 0x08, 0xd9, 0x12, 0x67,
	0x35, 0x8e, 0x00, 0x96, 0xb2, 0x23, 0x57, 0x3b, 0xd4, 0x13, 0x96, 0x60, 0x4f, 0x18, 0x0b, 0x19,
	0x12, 0xf4, 0x04, 0xd4, 0x08, 0x78, 0xe8, 0x3b, 0xf9, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0x8a,
	0xb1, 0xba, 0x50, 0xb3, 0x2a, 0x50, 0x4c, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x27,
	0x4b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x7b, 0x9c, 0x2a, 0x0d, 0x03, 0x00, 0x00,
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
	// Queries a list of admins items.
	Admins(ctx context.Context, in *QueryAdminsRequest, opts ...grpc.CallOption) (*QueryAdminsResponse, error)
	// Queries a list of archived proposals.
	ArchivedProposals(ctx context.Context, in *QueryArchivedProposalsRequest, opts ...grpc.CallOption) (*QueryArchivedProposalsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Admins(ctx context.Context, in *QueryAdminsRequest, opts ...grpc.CallOption) (*QueryAdminsResponse, error) {
	out := new(QueryAdminsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.adminmodule.adminmodule.Query/Admins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ArchivedProposals(ctx context.Context, in *QueryArchivedProposalsRequest, opts ...grpc.CallOption) (*QueryArchivedProposalsResponse, error) {
	out := new(QueryArchivedProposalsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.adminmodule.adminmodule.Query/ArchivedProposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of admins items.
	Admins(context.Context, *QueryAdminsRequest) (*QueryAdminsResponse, error)
	// Queries a list of archived proposals.
	ArchivedProposals(context.Context, *QueryArchivedProposalsRequest) (*QueryArchivedProposalsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Admins(ctx context.Context, req *QueryAdminsRequest) (*QueryAdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admins not implemented")
}
func (*UnimplementedQueryServer) ArchivedProposals(ctx context.Context, req *QueryArchivedProposalsRequest) (*QueryArchivedProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchivedProposals not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Admins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.adminmodule.adminmodule.Query/Admins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admins(ctx, req.(*QueryAdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ArchivedProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryArchivedProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ArchivedProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.adminmodule.adminmodule.Query/ArchivedProposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ArchivedProposals(ctx, req.(*QueryArchivedProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.adminmodule.adminmodule.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Admins",
			Handler:    _Query_Admins_Handler,
		},
		{
			MethodName: "ArchivedProposals",
			Handler:    _Query_ArchivedProposals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adminmodule/query.proto",
}

func (m *QueryAdminsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAdminsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admins) > 0 {
		for iNdEx := len(m.Admins) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Admins[iNdEx])
			copy(dAtA[i:], m.Admins[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Admins[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryArchivedProposalsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryArchivedProposalsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryArchivedProposalsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryArchivedProposalsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryArchivedProposalsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryArchivedProposalsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryAdminsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAdminsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Admins) > 0 {
		for _, s := range m.Admins {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryArchivedProposalsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryArchivedProposalsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAdminsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAdminsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admins", wireType)
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
			m.Admins = append(m.Admins, string(dAtA[iNdEx:postIndex]))
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
func (m *QueryArchivedProposalsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryArchivedProposalsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryArchivedProposalsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryArchivedProposalsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryArchivedProposalsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryArchivedProposalsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, &types.Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
