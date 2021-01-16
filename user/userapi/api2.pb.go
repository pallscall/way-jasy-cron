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

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" form:"username" validate:"required"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty" form:"passwd" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_592802a36db5ccc7, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return m.Size()
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

type LoginResp struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_592802a36db5ccc7, []int{1}
}
func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return m.Size()
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoginReq)(nil), "demo.service.v2.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "demo.service.v2.LoginResp")
}

func init() { proto.RegisterFile("api2.proto", fileDescriptor_592802a36db5ccc7) }

var fileDescriptor_592802a36db5ccc7 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0xb3, 0xe1, 0x23, 0x7f, 0x26, 0x7c, 0x08, 0x6b, 0x13, 0x57, 0xdd, 0xc4, 0x91, 0xa0,
	0x8d, 0xbb, 0xb0, 0x76, 0x01, 0x9b, 0xa4, 0xb5, 0x0a, 0xc4, 0xc2, 0x42, 0x98, 0x64, 0x6f, 0xd6,
	0x81, 0xdd, 0xb9, 0x93, 0x99, 0xd9, 0x88, 0xad, 0x0f, 0x60, 0x63, 0xe3, 0x23, 0xa5, 0x14, 0xec,
	0x83, 0x46, 0x2b, 0xcb, 0x3c, 0x81, 0x64, 0x77, 0xa3, 0x20, 0xda, 0xcd, 0xdc, 0xdf, 0x39, 0x87,
	0x7b, 0x2e, 0x21, 0x4c, 0xf2, 0xc0, 0x93, 0x0a, 0x0d, 0xda, 0x5b, 0x21, 0x24, 0xe8, 0x69, 0x50,
	0x33, 0x3e, 0x06, 0x6f, 0x16, 0x38, 0x27, 0x11, 0x37, 0xd7, 0xe9, 0xc8, 0x1b, 0x63, 0xe2, 0x47,
	0x18, 0xa1, 0x9f, 0xe9, 0x46, 0xe9, 0x24, 0xfb, 0x65, 0x9f, 0xec, 0x95, 0xfb, 0x9d, 0xdd, 0x08,
	0x31, 0x8a, 0xe1, 0x5b, 0x05, 0x89, 0x34, 0xb7, 0x05, 0xdc, 0x2b, 0x20, 0x93, 0xdc, 0x67, 0x42,
	0xa0, 0x61, 0x86, 0xa3, 0xd0, 0x39, 0xa5, 0xf7, 0x16, 0xa9, 0x9d, 0x63, 0xc4, 0xc5, 0x00, 0xa6,
	0x76, 0x9f, 0xd4, 0x52, 0x0d, 0x4a, 0xb0, 0x04, 0x9a, 0x56, 0xdb, 0x3a, 0xae, 0xf7, 0x8e, 0x56,
	0x8b, 0xd6, 0xe1, 0x04, 0x55, 0xd2, 0xa5, 0x1b, 0x42, 0xdb, 0x33, 0x16, 0xf3, 0x90, 0x19, 0xe8,
	0x52, 0x05, 0xd3, 0x94, 0x2b, 0x08, 0xe9, 0xe0, 0xcb, 0x68, 0x9f, 0x91, 0x8a, 0x64, 0x5a, 0xdf,
	0x84, 0xcd, 0x72, 0x16, 0xd1, 0x59, 0x2d, 0x5a, 0x07, 0x79, 0x44, 0x3e, 0xff, 0x3d, 0xa0, 0x30,
	0xd1, 0x80, 0xd4, 0x8b, 0x7d, 0xb4, 0xb4, 0x3b, 0xa4, 0xda, 0x47, 0x61, 0x40, 0x98, 0x62, 0x9f,
	0xc6, 0xc7, 0xa2, 0x55, 0x1d, 0xe7, 0xa3, 0xc1, 0x86, 0x05, 0x57, 0xe4, 0xdf, 0x50, 0x83, 0xb2,
	0x2f, 0x8a, 0x2e, 0x43, 0x15, 0xdb, 0x3b, 0xde, 0x8f, 0xa3, 0x7a, 0x9b, 0x9a, 0x8e, 0xf3, 0x17,
	0xd2, 0x92, 0x6e, 0xdf, 0x3d, 0xbf, 0x3f, 0x94, 0xff, 0xdb, 0x0d, 0x7f, 0x5d, 0xc8, 0x8f, 0xd7,
	0xa0, 0xb7, 0x3f, 0x7f, 0x75, 0x4b, 0xf3, 0xa5, 0x6b, 0x3d, 0x2d, 0x5d, 0xeb, 0x65, 0xe9, 0x5a,
	0x8f, 0x6f, 0x6e, 0xe9, 0xb2, 0xba, 0x56, 0x30, 0xc9, 0x47, 0x95, 0xec, 0x94, 0xa7, 0x9f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0xba, 0xd8, 0x76, 0xd3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	LoginUrl(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) LoginUrl(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/demo.service.v2.User/LoginUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	LoginUrl(context.Context, *LoginReq) (*LoginResp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) LoginUrl(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUrl not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_LoginUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.service.v2.User/LoginUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginUrl(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.service.v2.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUrl",
			Handler:    _User_LoginUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api2.proto",
}

func (m *LoginReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Passwd) > 0 {
		i -= len(m.Passwd)
		copy(dAtA[i:], m.Passwd)
		i = encodeVarintApi2(dAtA, i, uint64(len(m.Passwd)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintApi2(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoginResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintApi2(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
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
func (m *LoginReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovApi2(uint64(l))
	}
	l = len(m.Passwd)
	if l > 0 {
		n += 1 + l + sovApi2(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoginResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovApi2(uint64(l))
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
func (m *LoginReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LoginReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
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
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Passwd", wireType)
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
			m.Passwd = string(dAtA[iNdEx:postIndex])
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
func (m *LoginResp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LoginResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
			m.Content = string(dAtA[iNdEx:postIndex])
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