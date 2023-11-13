// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shard/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type SubmitShardTxRequest struct {
	// sender is the address of the sender. this will be set to the module address.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// namespace is the namespace of the world the transactions originated from.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// epoch is an arbitrary interval that this transaction was executed in.
	// for loop driven games, this is likely a tick. for event driven games,
	// this could be some general period of time.
	Epoch uint64 `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// txs are the transactions that occurred in this tick.
	Txs []*Transaction `protobuf:"bytes,4,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *SubmitShardTxRequest) Reset()         { *m = SubmitShardTxRequest{} }
func (m *SubmitShardTxRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitShardTxRequest) ProtoMessage()    {}
func (*SubmitShardTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ea9067d7c94eab8, []int{0}
}
func (m *SubmitShardTxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitShardTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitShardTxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitShardTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitShardTxRequest.Merge(m, src)
}
func (m *SubmitShardTxRequest) XXX_Size() int {
	return m.Size()
}
func (m *SubmitShardTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitShardTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitShardTxRequest proto.InternalMessageInfo

func (m *SubmitShardTxRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SubmitShardTxRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SubmitShardTxRequest) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *SubmitShardTxRequest) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type SubmitShardTxResponse struct {
}

func (m *SubmitShardTxResponse) Reset()         { *m = SubmitShardTxResponse{} }
func (m *SubmitShardTxResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitShardTxResponse) ProtoMessage()    {}
func (*SubmitShardTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ea9067d7c94eab8, []int{1}
}
func (m *SubmitShardTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitShardTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitShardTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitShardTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitShardTxResponse.Merge(m, src)
}
func (m *SubmitShardTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *SubmitShardTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitShardTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitShardTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SubmitShardTxRequest)(nil), "shard.v1.SubmitShardTxRequest")
	proto.RegisterType((*SubmitShardTxResponse)(nil), "shard.v1.SubmitShardTxResponse")
}

func init() { proto.RegisterFile("shard/v1/tx.proto", fileDescriptor_2ea9067d7c94eab8) }

var fileDescriptor_2ea9067d7c94eab8 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x4e, 0xea, 0x40,
	0x14, 0x66, 0x6e, 0x81, 0x5c, 0x86, 0xdc, 0xc5, 0x6d, 0x4a, 0xa8, 0x8d, 0xa9, 0x84, 0x8d, 0x84,
	0x84, 0x19, 0xc1, 0x9d, 0x3b, 0x59, 0xb9, 0x31, 0x31, 0x85, 0x95, 0x0b, 0x4d, 0x69, 0x27, 0xa5,
	0xd1, 0xce, 0xd4, 0x39, 0x03, 0xd6, 0x9d, 0xf1, 0x09, 0x7c, 0x14, 0x62, 0x7c, 0x08, 0x97, 0xc4,
	0x95, 0x4b, 0x03, 0x0b, 0x5e, 0xc3, 0xf4, 0xc7, 0x10, 0x8d, 0xee, 0xe6, 0xfb, 0x39, 0xdf, 0x99,
	0x73, 0x0e, 0xfe, 0x0f, 0x53, 0x57, 0xfa, 0x74, 0xde, 0xa7, 0x2a, 0x21, 0xb1, 0x14, 0x4a, 0xe8,
	0x7f, 0x33, 0x8a, 0xcc, 0xfb, 0xd6, 0x8e, 0x27, 0x20, 0x12, 0x70, 0x99, 0xf1, 0x34, 0x07, 0xb9,
	0xc9, 0x6a, 0xe6, 0x88, 0x46, 0x10, 0xa4, 0xc5, 0x11, 0x04, 0x85, 0x60, 0x6c, 0x03, 0xef, 0x62,
	0x56, 0xd8, 0xdb, 0x4f, 0x08, 0x1b, 0xa3, 0xd9, 0x24, 0x0a, 0xd5, 0x28, 0x95, 0xc7, 0x89, 0xc3,
	0x6e, 0x66, 0x0c, 0x94, 0x7e, 0x80, 0xab, 0xc0, 0xb8, 0xcf, 0xa4, 0x89, 0x5a, 0xa8, 0x53, 0x1b,
	0x9a, 0xaf, 0xcf, 0x3d, 0xa3, 0xe8, 0x74, 0xec, 0xfb, 0x92, 0x01, 0x8c, 0x94, 0x0c, 0x79, 0xe0,
	0x14, 0x3e, 0x7d, 0x17, 0xd7, 0xb8, 0x1b, 0x31, 0x88, 0x5d, 0x8f, 0x99, 0x7f, 0xd2, 0x22, 0x67,
	0x4b, 0xe8, 0x06, 0xae, 0xb0, 0x58, 0x78, 0x53, 0x53, 0x6b, 0xa1, 0x4e, 0xd9, 0xc9, 0x81, 0xbe,
	0x8f, 0x35, 0x95, 0x80, 0x59, 0x6e, 0x69, 0x9d, 0xfa, 0xa0, 0x41, 0x3e, 0x07, 0x24, 0x63, 0xe9,
	0x72, 0x70, 0x3d, 0x15, 0x0a, 0xee, 0xa4, 0x8e, 0xa3, 0xfa, 0xc3, 0x66, 0xd1, 0x2d, 0x3a, 0xb5,
	0x9b, 0xb8, 0xf1, 0xed, 0xcf, 0x10, 0x0b, 0x0e, 0x6c, 0x70, 0x81, 0xb5, 0x53, 0x08, 0xf4, 0x33,
	0xfc, 0xef, 0x8b, 0xae, 0xdb, 0xdb, 0xe4, 0x9f, 0x86, 0xb5, 0xf6, 0x7e, 0xd5, 0xf3, 0x60, 0xab,
	0x72, 0xbf, 0x59, 0x74, 0xd1, 0xf0, 0xe4, 0x65, 0x65, 0xa3, 0xe5, 0xca, 0x46, 0xef, 0x2b, 0x1b,
	0x3d, 0xae, 0xed, 0xd2, 0x72, 0x6d, 0x97, 0xde, 0xd6, 0x76, 0xe9, 0x9c, 0xc4, 0x57, 0x01, 0xb9,
	0x15, 0xf2, 0xda, 0x27, 0x3e, 0x9b, 0xd3, 0xec, 0xd5, 0x63, 0x3c, 0x08, 0x39, 0xa3, 0xde, 0xd4,
	0x0d, 0x39, 0x4d, 0x68, 0x7e, 0x80, 0x6c, 0xfb, 0x93, 0x6a, 0xb6, 0xfe, 0xc3, 0x8f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd9, 0xdb, 0xfd, 0x11, 0xe7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SubmitShardTx(ctx context.Context, in *SubmitShardTxRequest, opts ...grpc.CallOption) (*SubmitShardTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitShardTx(ctx context.Context, in *SubmitShardTxRequest, opts ...grpc.CallOption) (*SubmitShardTxResponse, error) {
	out := new(SubmitShardTxResponse)
	err := c.cc.Invoke(ctx, "/shard.v1.Msg/SubmitShardTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitShardTx(context.Context, *SubmitShardTxRequest) (*SubmitShardTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitShardTx(ctx context.Context, req *SubmitShardTxRequest) (*SubmitShardTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitShardTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitShardTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitShardTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitShardTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shard.v1.Msg/SubmitShardTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitShardTx(ctx, req.(*SubmitShardTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shard.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitShardTx",
			Handler:    _Msg_SubmitShardTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shard/v1/tx.proto",
}

func (m *SubmitShardTxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitShardTxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitShardTxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubmitShardTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitShardTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitShardTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SubmitShardTxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Epoch != 0 {
		n += 1 + sovTx(uint64(m.Epoch))
	}
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *SubmitShardTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SubmitShardTxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: SubmitShardTxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitShardTxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, &Transaction{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *SubmitShardTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: SubmitShardTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitShardTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)