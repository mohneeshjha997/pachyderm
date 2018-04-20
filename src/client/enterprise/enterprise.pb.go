// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/enterprise/enterprise.proto

/*
	Package enterprise is a generated protocol buffer package.

	It is generated from these files:
		client/enterprise/enterprise.proto

	It has these top-level messages:
		EnterpriseRecord
		TokenInfo
		ActivateRequest
		ActivateResponse
		GetStateRequest
		GetStateResponse
		DeactivateRequest
		DeactivateResponse
*/
package enterprise

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type State int32

const (
	State_NONE    State = 0
	State_ACTIVE  State = 1
	State_EXPIRED State = 2
)

var State_name = map[int32]string{
	0: "NONE",
	1: "ACTIVE",
	2: "EXPIRED",
}
var State_value = map[string]int32{
	"NONE":    0,
	"ACTIVE":  1,
	"EXPIRED": 2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{0} }

// EnterpriseRecord is the record we store in etcd for a Pachyderm enterprise
// token that has been provided to a Pachyderm cluster
type EnterpriseRecord struct {
	ActivationCode string `protobuf:"bytes,1,opt,name=activation_code,json=activationCode,proto3" json:"activation_code,omitempty"`
	// expires is a timestamp indicating when this activation code will expire.
	Expires *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expires" json:"expires,omitempty"`
}

func (m *EnterpriseRecord) Reset()                    { *m = EnterpriseRecord{} }
func (m *EnterpriseRecord) String() string            { return proto.CompactTextString(m) }
func (*EnterpriseRecord) ProtoMessage()               {}
func (*EnterpriseRecord) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{0} }

func (m *EnterpriseRecord) GetActivationCode() string {
	if m != nil {
		return m.ActivationCode
	}
	return ""
}

func (m *EnterpriseRecord) GetExpires() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

// TokenInfo contains information about the currently active enterprise token
type TokenInfo struct {
	// expires indicates when the current token expires (unset if there is no
	// current token)
	Expires *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=expires" json:"expires,omitempty"`
}

func (m *TokenInfo) Reset()                    { *m = TokenInfo{} }
func (m *TokenInfo) String() string            { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()               {}
func (*TokenInfo) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{1} }

func (m *TokenInfo) GetExpires() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

type ActivateRequest struct {
	// activation_code is a Pachyderm enterprise activation code. New users can
	// obtain trial activation codes
	ActivationCode string `protobuf:"bytes,1,opt,name=activation_code,json=activationCode,proto3" json:"activation_code,omitempty"`
	// expires is a timestamp indicating when this activation code will expire.
	// This should not generally be set (it's primarily used for testing), and is
	// only applied if it's earlier than the signed expiration time in
	// 'activation_code'.
	Expires *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expires" json:"expires,omitempty"`
}

func (m *ActivateRequest) Reset()                    { *m = ActivateRequest{} }
func (m *ActivateRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivateRequest) ProtoMessage()               {}
func (*ActivateRequest) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{2} }

func (m *ActivateRequest) GetActivationCode() string {
	if m != nil {
		return m.ActivationCode
	}
	return ""
}

func (m *ActivateRequest) GetExpires() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

type ActivateResponse struct {
	Info *TokenInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *ActivateResponse) Reset()                    { *m = ActivateResponse{} }
func (m *ActivateResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivateResponse) ProtoMessage()               {}
func (*ActivateResponse) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{3} }

func (m *ActivateResponse) GetInfo() *TokenInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetStateRequest struct {
}

func (m *GetStateRequest) Reset()                    { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()               {}
func (*GetStateRequest) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{4} }

type GetStateResponse struct {
	State State      `protobuf:"varint,1,opt,name=state,proto3,enum=enterprise.State" json:"state,omitempty"`
	Info  *TokenInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *GetStateResponse) Reset()                    { *m = GetStateResponse{} }
func (m *GetStateResponse) String() string            { return proto.CompactTextString(m) }
func (*GetStateResponse) ProtoMessage()               {}
func (*GetStateResponse) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{5} }

func (m *GetStateResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_NONE
}

func (m *GetStateResponse) GetInfo() *TokenInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type DeactivateRequest struct {
}

func (m *DeactivateRequest) Reset()                    { *m = DeactivateRequest{} }
func (m *DeactivateRequest) String() string            { return proto.CompactTextString(m) }
func (*DeactivateRequest) ProtoMessage()               {}
func (*DeactivateRequest) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{6} }

type DeactivateResponse struct {
}

func (m *DeactivateResponse) Reset()                    { *m = DeactivateResponse{} }
func (m *DeactivateResponse) String() string            { return proto.CompactTextString(m) }
func (*DeactivateResponse) ProtoMessage()               {}
func (*DeactivateResponse) Descriptor() ([]byte, []int) { return fileDescriptorEnterprise, []int{7} }

func init() {
	proto.RegisterType((*EnterpriseRecord)(nil), "enterprise.EnterpriseRecord")
	proto.RegisterType((*TokenInfo)(nil), "enterprise.TokenInfo")
	proto.RegisterType((*ActivateRequest)(nil), "enterprise.ActivateRequest")
	proto.RegisterType((*ActivateResponse)(nil), "enterprise.ActivateResponse")
	proto.RegisterType((*GetStateRequest)(nil), "enterprise.GetStateRequest")
	proto.RegisterType((*GetStateResponse)(nil), "enterprise.GetStateResponse")
	proto.RegisterType((*DeactivateRequest)(nil), "enterprise.DeactivateRequest")
	proto.RegisterType((*DeactivateResponse)(nil), "enterprise.DeactivateResponse")
	proto.RegisterEnum("enterprise.State", State_name, State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	// Provide a Pachyderm enterprise token, enabling Pachyderm enterprise
	// features, such as the Pachyderm Dashboard and Auth system
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
	// Deactivate is a testing API. It removes a cluster's enterprise activation
	// token and sets its enterprise state to NONE (normally, once a cluster has
	// been activated, the only reachable state is EXPIRED).
	//
	// NOTE: This endpoint also calls DeleteAll (and deletes all Pachyderm data in
	// its cluster). This is to avoid dealing with invalid, intermediate states
	// (e.g. auth is activated but enterprise state is NONE)
	Deactivate(ctx context.Context, in *DeactivateRequest, opts ...grpc.CallOption) (*DeactivateResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := grpc.Invoke(ctx, "/enterprise.API/Activate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := grpc.Invoke(ctx, "/enterprise.API/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Deactivate(ctx context.Context, in *DeactivateRequest, opts ...grpc.CallOption) (*DeactivateResponse, error) {
	out := new(DeactivateResponse)
	err := grpc.Invoke(ctx, "/enterprise.API/Deactivate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Provide a Pachyderm enterprise token, enabling Pachyderm enterprise
	// features, such as the Pachyderm Dashboard and Auth system
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	// Deactivate is a testing API. It removes a cluster's enterprise activation
	// token and sets its enterprise state to NONE (normally, once a cluster has
	// been activated, the only reachable state is EXPIRED).
	//
	// NOTE: This endpoint also calls DeleteAll (and deletes all Pachyderm data in
	// its cluster). This is to avoid dealing with invalid, intermediate states
	// (e.g. auth is activated but enterprise state is NONE)
	Deactivate(context.Context, *DeactivateRequest) (*DeactivateResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enterprise.API/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enterprise.API/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enterprise.API/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Deactivate(ctx, req.(*DeactivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "enterprise.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _API_Activate_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _API_GetState_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _API_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/enterprise/enterprise.proto",
}

func (m *EnterpriseRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterpriseRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ActivationCode) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(len(m.ActivationCode)))
		i += copy(dAtA[i:], m.ActivationCode)
	}
	if m.Expires != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(m.Expires.Size()))
		n1, err := m.Expires.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *TokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Expires != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(m.Expires.Size()))
		n2, err := m.Expires.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ActivateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ActivationCode) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(len(m.ActivationCode)))
		i += copy(dAtA[i:], m.ActivationCode)
	}
	if m.Expires != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(m.Expires.Size()))
		n3, err := m.Expires.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *ActivateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Info != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(m.Info.Size()))
		n4, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *GetStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(m.State))
	}
	if m.Info != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnterprise(dAtA, i, uint64(m.Info.Size()))
		n5, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *DeactivateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeactivateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DeactivateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeactivateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintEnterprise(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnterpriseRecord) Size() (n int) {
	var l int
	_ = l
	l = len(m.ActivationCode)
	if l > 0 {
		n += 1 + l + sovEnterprise(uint64(l))
	}
	if m.Expires != nil {
		l = m.Expires.Size()
		n += 1 + l + sovEnterprise(uint64(l))
	}
	return n
}

func (m *TokenInfo) Size() (n int) {
	var l int
	_ = l
	if m.Expires != nil {
		l = m.Expires.Size()
		n += 1 + l + sovEnterprise(uint64(l))
	}
	return n
}

func (m *ActivateRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ActivationCode)
	if l > 0 {
		n += 1 + l + sovEnterprise(uint64(l))
	}
	if m.Expires != nil {
		l = m.Expires.Size()
		n += 1 + l + sovEnterprise(uint64(l))
	}
	return n
}

func (m *ActivateResponse) Size() (n int) {
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovEnterprise(uint64(l))
	}
	return n
}

func (m *GetStateRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetStateResponse) Size() (n int) {
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovEnterprise(uint64(m.State))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovEnterprise(uint64(l))
	}
	return n
}

func (m *DeactivateRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *DeactivateResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovEnterprise(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEnterprise(x uint64) (n int) {
	return sovEnterprise(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnterpriseRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EnterpriseRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterpriseRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivationCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActivationCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expires == nil {
				m.Expires = &google_protobuf.Timestamp{}
			}
			if err := m.Expires.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *TokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expires == nil {
				m.Expires = &google_protobuf.Timestamp{}
			}
			if err := m.Expires.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *ActivateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivationCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActivationCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expires == nil {
				m.Expires = &google_protobuf.Timestamp{}
			}
			if err := m.Expires.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *ActivateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &TokenInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *GetStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *GetStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (State(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnterprise
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &TokenInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *DeactivateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeactivateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeactivateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func (m *DeactivateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnterprise
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeactivateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeactivateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEnterprise(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnterprise
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
func skipEnterprise(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnterprise
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
					return 0, ErrIntOverflowEnterprise
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnterprise
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEnterprise
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnterprise
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEnterprise(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEnterprise = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnterprise   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("client/enterprise/enterprise.proto", fileDescriptorEnterprise) }

var fileDescriptorEnterprise = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x8e, 0xd2, 0x50,
	0x14, 0xc6, 0x7b, 0x91, 0xbf, 0x87, 0x04, 0xca, 0x55, 0x13, 0x52, 0xb5, 0x92, 0xbb, 0x01, 0x59,
	0x94, 0x04, 0xdd, 0xba, 0xa8, 0xd0, 0x90, 0x2e, 0x44, 0x52, 0x89, 0x71, 0x67, 0x4a, 0x39, 0x25,
	0x8d, 0xd0, 0x5b, 0xda, 0x8b, 0xf1, 0x51, 0x7c, 0x24, 0x97, 0x3e, 0x82, 0x61, 0x92, 0x79, 0x8e,
	0x09, 0x2d, 0xa5, 0x85, 0x99, 0xc9, 0xcc, 0x66, 0x76, 0xcd, 0x39, 0x5f, 0xbe, 0xdf, 0x77, 0xbf,
	0x53, 0x60, 0xce, 0xda, 0x43, 0x5f, 0x0c, 0xd0, 0x17, 0x18, 0x06, 0xa1, 0x17, 0x61, 0xee, 0x53,
	0x0b, 0x42, 0x2e, 0x38, 0x85, 0x6c, 0xa2, 0xbc, 0x5d, 0x71, 0xbe, 0x5a, 0xe3, 0x20, 0xde, 0x2c,
	0x76, 0xee, 0x40, 0x78, 0x1b, 0x8c, 0x84, 0xbd, 0x09, 0x12, 0x31, 0xdb, 0x82, 0x6c, 0x9c, 0xe4,
	0x16, 0x3a, 0x3c, 0x5c, 0xd2, 0x2e, 0x34, 0x6d, 0x47, 0x78, 0xbf, 0x6c, 0xe1, 0x71, 0xff, 0x87,
	0xc3, 0x97, 0xd8, 0x26, 0x1d, 0xd2, 0xab, 0x59, 0x8d, 0x6c, 0x3c, 0xe2, 0x4b, 0xa4, 0x1f, 0xa0,
	0x82, 0xbf, 0x03, 0x2f, 0xc4, 0xa8, 0x5d, 0xe8, 0x90, 0x5e, 0x7d, 0xa8, 0x68, 0x09, 0x4f, 0x4b,
	0x79, 0xda, 0x3c, 0xe5, 0x59, 0xa9, 0x94, 0xe9, 0x50, 0x9b, 0xf3, 0x9f, 0xe8, 0x9b, 0xbe, 0xcb,
	0xf3, 0x16, 0xe4, 0xf1, 0x16, 0x01, 0x34, 0xf5, 0x24, 0x0a, 0x5a, 0xb8, 0xdd, 0x61, 0x24, 0x9e,
	0x3a, 0xf4, 0x47, 0x90, 0x33, 0x62, 0x14, 0x70, 0x3f, 0x42, 0xfa, 0x0e, 0x8a, 0x9e, 0xef, 0xf2,
	0x63, 0xf0, 0x97, 0x5a, 0xee, 0x12, 0xa7, 0x07, 0x5a, 0xb1, 0x84, 0xb5, 0xa0, 0x39, 0x41, 0xf1,
	0x55, 0x64, 0x81, 0x99, 0x0b, 0x72, 0x36, 0x3a, 0x3a, 0x76, 0xa1, 0x14, 0x1d, 0x06, 0xb1, 0x65,
	0x63, 0xd8, 0xca, 0x5b, 0x26, 0xca, 0x64, 0x7f, 0x42, 0x17, 0x1e, 0x46, 0x3f, 0x87, 0xd6, 0x18,
	0xed, 0xf3, 0xb6, 0xd8, 0x0b, 0xa0, 0xf9, 0x61, 0x82, 0xef, 0xf7, 0xa1, 0x14, 0x53, 0x68, 0x15,
	0x8a, 0xd3, 0x2f, 0x53, 0x43, 0x96, 0x28, 0x40, 0x59, 0x1f, 0xcd, 0xcd, 0x6f, 0x86, 0x4c, 0x68,
	0x1d, 0x2a, 0xc6, 0xf7, 0x99, 0x69, 0x19, 0x63, 0xb9, 0x30, 0xbc, 0x26, 0xf0, 0x4c, 0x9f, 0x99,
	0x74, 0x02, 0xd5, 0xb4, 0x18, 0xfa, 0x2a, 0x9f, 0xe3, 0xe2, 0x40, 0xca, 0xeb, 0xbb, 0x97, 0x09,
	0x9a, 0x49, 0x07, 0xa3, 0xb4, 0x8f, 0x73, 0xa3, 0x8b, 0xe2, 0xce, 0x8d, 0x2e, 0x2b, 0x64, 0x12,
	0xfd, 0x0c, 0x90, 0xbd, 0x8d, 0xbe, 0xc9, 0xab, 0x6f, 0x15, 0xa1, 0xa8, 0xf7, 0xad, 0x53, 0xbb,
	0x4f, 0xf2, 0xdf, 0xbd, 0x4a, 0xfe, 0xed, 0x55, 0xf2, 0x7f, 0xaf, 0x92, 0x3f, 0x57, 0xaa, 0xb4,
	0x28, 0xc7, 0x3f, 0xca, 0xfb, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x5e, 0xb4, 0x7c, 0x8d,
	0x03, 0x00, 0x00,
}
