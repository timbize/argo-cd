// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/services/application.proto

package services // import "github.com/argoproj/argo-cd/controller/services"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ResourcesQuery struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcesQuery) Reset()         { *m = ResourcesQuery{} }
func (m *ResourcesQuery) String() string { return proto.CompactTextString(m) }
func (*ResourcesQuery) ProtoMessage()    {}
func (*ResourcesQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_9e06f70e07dd0483, []int{0}
}
func (m *ResourcesQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourcesQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourcesQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourcesQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesQuery.Merge(dst, src)
}
func (m *ResourcesQuery) XXX_Size() int {
	return m.Size()
}
func (m *ResourcesQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesQuery proto.InternalMessageInfo

func (m *ResourcesQuery) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

type ResourceTreeResponse struct {
	Items                []*v1alpha1.ResourceNode `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ResourceTreeResponse) Reset()         { *m = ResourceTreeResponse{} }
func (m *ResourceTreeResponse) String() string { return proto.CompactTextString(m) }
func (*ResourceTreeResponse) ProtoMessage()    {}
func (*ResourceTreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_9e06f70e07dd0483, []int{1}
}
func (m *ResourceTreeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceTreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceTreeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourceTreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceTreeResponse.Merge(dst, src)
}
func (m *ResourceTreeResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResourceTreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceTreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceTreeResponse proto.InternalMessageInfo

func (m *ResourceTreeResponse) GetItems() []*v1alpha1.ResourceNode {
	if m != nil {
		return m.Items
	}
	return nil
}

type ManagedResourcesResponse struct {
	Items                []*v1alpha1.ResourceDiff `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ManagedResourcesResponse) Reset()         { *m = ManagedResourcesResponse{} }
func (m *ManagedResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ManagedResourcesResponse) ProtoMessage()    {}
func (*ManagedResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_9e06f70e07dd0483, []int{2}
}
func (m *ManagedResourcesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManagedResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManagedResourcesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ManagedResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedResourcesResponse.Merge(dst, src)
}
func (m *ManagedResourcesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ManagedResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedResourcesResponse proto.InternalMessageInfo

func (m *ManagedResourcesResponse) GetItems() []*v1alpha1.ResourceDiff {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourcesQuery)(nil), "github.com.argoproj.argo_cd.controller.services.ResourcesQuery")
	proto.RegisterType((*ResourceTreeResponse)(nil), "github.com.argoproj.argo_cd.controller.services.ResourceTreeResponse")
	proto.RegisterType((*ManagedResourcesResponse)(nil), "github.com.argoproj.argo_cd.controller.services.ManagedResourcesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApplicationService service

type ApplicationServiceClient interface {
	ResourceTree(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ResourceTreeResponse, error)
	ManagedResources(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ManagedResourcesResponse, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) ResourceTree(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ResourceTreeResponse, error) {
	out := new(ResourceTreeResponse)
	err := c.cc.Invoke(ctx, "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ResourceTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ManagedResources(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ManagedResourcesResponse, error) {
	out := new(ManagedResourcesResponse)
	err := c.cc.Invoke(ctx, "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ManagedResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationService service

type ApplicationServiceServer interface {
	ResourceTree(context.Context, *ResourcesQuery) (*ResourceTreeResponse, error)
	ManagedResources(context.Context, *ResourcesQuery) (*ManagedResourcesResponse, error)
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_ResourceTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourcesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ResourceTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ResourceTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ResourceTree(ctx, req.(*ResourcesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ManagedResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourcesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ManagedResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ManagedResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ManagedResources(ctx, req.(*ResourcesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.argoproj.argo_cd.controller.services.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResourceTree",
			Handler:    _ApplicationService_ResourceTree_Handler,
		},
		{
			MethodName: "ManagedResources",
			Handler:    _ApplicationService_ManagedResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/services/application.proto",
}

func (m *ResourcesQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourcesQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ApplicationName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(m.ApplicationName)))
		i += copy(dAtA[i:], m.ApplicationName)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResourceTreeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceTreeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintApplication(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ManagedResourcesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManagedResourcesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintApplication(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourcesQuery) Size() (n int) {
	var l int
	_ = l
	l = len(m.ApplicationName)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResourceTreeResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovApplication(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManagedResourcesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovApplication(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourcesQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ResourcesQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourcesQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
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
func (m *ResourceTreeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ResourceTreeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceTreeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &v1alpha1.ResourceNode{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
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
func (m *ManagedResourcesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: ManagedResourcesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManagedResourcesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &v1alpha1.ResourceDiff{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
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
func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
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
				return 0, ErrInvalidLengthApplication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplication
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
				next, err := skipApplication(dAtA[start:])
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
	ErrInvalidLengthApplication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("controller/services/application.proto", fileDescriptor_application_9e06f70e07dd0483)
}

var fileDescriptor_application_9e06f70e07dd0483 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x6d, 0xbe, 0x0f, 0x05, 0xa3, 0xa8, 0x04, 0x17, 0xa5, 0x8b, 0x52, 0x06, 0x84, 0x6e, 0x4c,
	0x68, 0xdd, 0x09, 0x22, 0x8a, 0x22, 0x5d, 0x58, 0x70, 0x74, 0x25, 0x88, 0xa4, 0x99, 0xdb, 0x34,
	0x76, 0x3a, 0x09, 0x49, 0xa6, 0xd0, 0x37, 0x71, 0xe9, 0xe3, 0xb8, 0x14, 0x9f, 0x40, 0xfa, 0x24,
	0x62, 0xeb, 0x38, 0xd3, 0xa2, 0x85, 0x2a, 0xee, 0x2e, 0x21, 0xe7, 0x27, 0xe7, 0xdc, 0xe0, 0x5d,
	0xa1, 0x13, 0x6f, 0x75, 0x1c, 0x83, 0x65, 0x0e, 0xec, 0x50, 0x09, 0x70, 0x8c, 0x1b, 0x13, 0x2b,
	0xc1, 0xbd, 0xd2, 0x09, 0x35, 0x56, 0x7b, 0x4d, 0x98, 0x54, 0xbe, 0x97, 0x76, 0xa8, 0xd0, 0x03,
	0xca, 0xad, 0xd4, 0xc6, 0xea, 0xfb, 0xc9, 0x70, 0x27, 0x22, 0x9a, 0x53, 0xd0, 0x8c, 0xa2, 0xd2,
	0xca, 0x01, 0x2c, 0x03, 0x4c, 0x86, 0x3d, 0x11, 0x31, 0xd3, 0x97, 0x8c, 0x1b, 0x35, 0x23, 0xc4,
	0x86, 0x0d, 0x1e, 0x9b, 0x1e, 0x6f, 0x30, 0x09, 0x09, 0x58, 0xee, 0x21, 0x9a, 0x6a, 0x07, 0x07,
	0x78, 0x33, 0x04, 0xa7, 0x53, 0x2b, 0xc0, 0x5d, 0xa6, 0x60, 0x47, 0xa4, 0x8e, 0xb7, 0x0a, 0xc8,
	0x36, 0x1f, 0x40, 0x19, 0xd5, 0x50, 0x7d, 0x2d, 0x9c, 0x3f, 0x0e, 0x52, 0xbc, 0x93, 0x61, 0xaf,
	0x2d, 0x40, 0x08, 0xce, 0xe8, 0xc4, 0x01, 0xb9, 0xc5, 0x2b, 0xca, 0xc3, 0xc0, 0x95, 0x51, 0xed,
	0x7f, 0x7d, 0xbd, 0x79, 0x4e, 0x17, 0xbd, 0xcf, 0xf4, 0x25, 0x7d, 0xb7, 0x4b, 0x8b, 0xb9, 0x64,
	0x76, 0x69, 0xc6, 0xdf, 0xd6, 0x11, 0x84, 0x53, 0xd6, 0x60, 0x84, 0xcb, 0x17, 0x3c, 0xe1, 0x12,
	0xa2, 0x4f, 0xe7, 0x7f, 0x29, 0x7d, 0xaa, 0xba, 0xdd, 0x0f, 0xe9, 0xe6, 0xcb, 0x3f, 0x4c, 0x8e,
	0xf3, 0xcb, 0x57, 0xd3, 0x42, 0xc8, 0x03, 0xc2, 0x1b, 0xc5, 0x24, 0xc8, 0x11, 0x5d, 0xb2, 0x52,
	0x3a, 0x5b, 0x42, 0xe5, 0xec, 0xc7, 0x04, 0xc5, 0x26, 0x82, 0x12, 0x79, 0x44, 0x78, 0x7b, 0x3e,
	0xad, 0xdf, 0xdb, 0x6b, 0x2d, 0x4d, 0xf0, 0x5d, 0x63, 0x41, 0xe9, 0xe4, 0xf0, 0x69, 0x5c, 0x45,
	0xcf, 0xe3, 0x2a, 0x7a, 0x1d, 0x57, 0xd1, 0x0d, 0x5b, 0xb4, 0xdb, 0x5f, 0xfc, 0xa7, 0xce, 0xea,
	0x64, 0x91, 0xf7, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x5d, 0xbf, 0x7d, 0x6d, 0x03, 0x00,
	0x00,
}