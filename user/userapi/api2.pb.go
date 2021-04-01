// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api2.proto

// package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

package userapi

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VerifyReq struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty" json:"access_key" validate:"required"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" json:"token" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReq) Reset()         { *m = VerifyReq{} }
func (m *VerifyReq) String() string { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()    {}
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_592802a36db5ccc7, []int{0}
}
func (m *VerifyReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReq.Merge(m, src)
}
func (m *VerifyReq) XXX_Size() int {
	return m.Size()
}
func (m *VerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReq proto.InternalMessageInfo

type VerifyReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReply) Reset()         { *m = VerifyReply{} }
func (m *VerifyReply) String() string { return proto.CompactTextString(m) }
func (*VerifyReply) ProtoMessage()    {}
func (*VerifyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_592802a36db5ccc7, []int{1}
}
func (m *VerifyReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReply.Merge(m, src)
}
func (m *VerifyReply) XXX_Size() int {
	return m.Size()
}
func (m *VerifyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReply.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VerifyReq)(nil), "user.v2.VerifyReq")
	proto.RegisterType((*VerifyReply)(nil), "user.v2.VerifyReply")
}

func init() { proto.RegisterFile("api2.proto", fileDescriptor_592802a36db5ccc7) }

var fileDescriptor_592802a36db5ccc7 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x6b, 0x44, 0x0b, 0x75, 0x37, 0x8b, 0xa1, 0x2a, 0x25, 0xad, 0x2c, 0x90, 0x40, 0x88,
	0x54, 0x0a, 0x5b, 0x37, 0xba, 0x30, 0xb0, 0x55, 0x88, 0x81, 0x05, 0xb9, 0xe9, 0x35, 0x98, 0xa6,
	0x39, 0x37, 0x71, 0x2a, 0xf9, 0x2f, 0x18, 0xf9, 0xa4, 0x8e, 0x7c, 0x41, 0x04, 0x61, 0x63, 0xec,
	0x17, 0x20, 0xec, 0x40, 0x07, 0x58, 0xac, 0xbb, 0x7b, 0x7e, 0xa7, 0xf3, 0x99, 0x52, 0xa1, 0x64,
	0xe0, 0xab, 0x14, 0x35, 0xb2, 0xbd, 0x3c, 0x83, 0xd4, 0x5f, 0x05, 0x9d, 0x8b, 0x48, 0xea, 0xc7,
	0x7c, 0xe2, 0x87, 0xb8, 0x18, 0x44, 0x18, 0xe1, 0xc0, 0xf2, 0x49, 0x3e, 0xb3, 0x99, 0x4d, 0x6c,
	0xe4, 0xbc, 0xce, 0x61, 0x84, 0x18, 0xc5, 0xb0, 0xbd, 0x05, 0x0b, 0xa5, 0x4d, 0x05, 0xbb, 0x15,
	0x14, 0x4a, 0x0e, 0x44, 0x92, 0xa0, 0x16, 0x5a, 0x62, 0x92, 0x39, 0xca, 0x9f, 0x09, 0x6d, 0xde,
	0x41, 0x2a, 0x67, 0x66, 0x0c, 0x4b, 0x76, 0x4d, 0x9b, 0x22, 0x0c, 0x21, 0xcb, 0x6e, 0xc0, 0xb4,
	0x49, 0x9f, 0x9c, 0x36, 0x47, 0x67, 0x9b, 0xa2, 0x77, 0xf2, 0x94, 0x61, 0x32, 0xe4, 0x0e, 0x3d,
	0xcc, 0xc1, 0xf0, 0xfe, 0x4a, 0xc4, 0x72, 0x2a, 0x34, 0x0c, 0x79, 0x0a, 0xcb, 0x5c, 0xa6, 0x30,
	0xe5, 0xe3, 0xad, 0xcb, 0x86, 0xb4, 0xae, 0x71, 0x0e, 0x49, 0x7b, 0xc7, 0x36, 0x39, 0xde, 0x14,
	0xbd, 0xbe, 0x6b, 0x62, 0xcb, 0xff, 0xfb, 0x4e, 0xe1, 0xe7, 0xb4, 0xf5, 0x33, 0x91, 0x8a, 0x0d,
	0xeb, 0xd2, 0xdd, 0x10, 0xa7, 0x60, 0xc7, 0xa9, 0x8f, 0xf6, 0x3f, 0x8b, 0x9e, 0xcd, 0xc7, 0xf6,
	0x0c, 0xae, 0x68, 0xeb, 0xf6, 0xdb, 0x72, 0x06, 0x0b, 0x68, 0xa3, 0x8a, 0x98, 0x5f, 0x2d, 0xd3,
	0xff, 0x7d, 0x5e, 0xe7, 0xe0, 0x4f, 0x4d, 0xc5, 0x66, 0x74, 0xb4, 0x7e, 0xf7, 0x6a, 0xeb, 0xd2,
	0x23, 0xaf, 0xa5, 0x47, 0xde, 0x4a, 0x8f, 0xbc, 0x7c, 0x78, 0xb5, 0x7b, 0xfb, 0x17, 0x42, 0xc9,
	0x49, 0xc3, 0x2e, 0xea, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x5e, 0xf4, 0x69, 0xa9, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenVerifyClient is the client API for TokenVerify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenVerifyClient interface {
	Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyReply, error)
}

type tokenVerifyClient struct {
	cc *grpc.ClientConn
}

func NewTokenVerifyClient(cc *grpc.ClientConn) TokenVerifyClient {
	return &tokenVerifyClient{cc}
}

func (c *tokenVerifyClient) Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyReply, error) {
	out := new(VerifyReply)
	err := c.cc.Invoke(ctx, "/user.v2.TokenVerify/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenVerifyServer is the server API for TokenVerify service.
type TokenVerifyServer interface {
	Verify(context.Context, *VerifyReq) (*VerifyReply, error)
}

// UnimplementedTokenVerifyServer can be embedded to have forward compatible implementations.
type UnimplementedTokenVerifyServer struct {
}

func (*UnimplementedTokenVerifyServer) Verify(ctx context.Context, req *VerifyReq) (*VerifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterTokenVerifyServer(s *grpc.Server, srv TokenVerifyServer) {
	s.RegisterService(&_TokenVerify_serviceDesc, srv)
}

func _TokenVerify_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenVerifyServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v2.TokenVerify/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenVerifyServer).Verify(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenVerify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.v2.TokenVerify",
	HandlerType: (*TokenVerifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _TokenVerify_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api2.proto",
}

func (m *VerifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintApi2(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccessKey) > 0 {
		i -= len(m.AccessKey)
		copy(dAtA[i:], m.AccessKey)
		i = encodeVarintApi2(dAtA, i, uint64(len(m.AccessKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VerifyReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintApi2(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi2(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VerifyReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccessKey)
	if l > 0 {
		n += 1 + l + sovApi2(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovApi2(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VerifyReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovApi2(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi2(x uint64) (n int) {
	return sovApi2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi2
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
			return fmt.Errorf("proto: VerifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi2
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
				return ErrInvalidLengthApi2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi2
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
				return ErrInvalidLengthApi2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerifyReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi2
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
			return fmt.Errorf("proto: VerifyReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi2
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
					return 0, ErrIntOverflowApi2
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
					return 0, ErrIntOverflowApi2
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
				return 0, ErrInvalidLengthApi2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi2 = fmt.Errorf("proto: unexpected end of group")
)
