// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bucket/where_is.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// WhereIsVisibility is the visibility of the bucket to authorized users of an application
type WhereIsVisibility int32

const (
	// Bucket does not have visibility set.
	WhereIsVisibility_UNSPECIFIED WhereIsVisibility = 0
	// Bucket is visible to anyone.
	WhereIsVisibility_PUBLIC WhereIsVisibility = 1
	// Bucket is visible to anyone who has access token.
	WhereIsVisibility_PRIVATE WhereIsVisibility = 2
)

var WhereIsVisibility_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "PUBLIC",
	2: "PRIVATE",
}

var WhereIsVisibility_value = map[string]int32{
	"UNSPECIFIED": 0,
	"PUBLIC":      1,
	"PRIVATE":     2,
}

func (x WhereIsVisibility) String() string {
	return proto.EnumName(WhereIsVisibility_name, int32(x))
}

func (WhereIsVisibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{0}
}

// WhereIsRole is the role of the creator of the bucket.
type WhereIsRole int32

const (
	// Bucket does not have role set.
	WhereIsRole_NONE WhereIsRole = 0
	// Bucket is visible to anyone.
	WhereIsRole_APPLICATION WhereIsRole = 1
	// Bucket is visible to anyone who has access token.
	WhereIsRole_USER WhereIsRole = 2
)

var WhereIsRole_name = map[int32]string{
	0: "NONE",
	1: "APPLICATION",
	2: "USER",
}

var WhereIsRole_value = map[string]int32{
	"NONE":        0,
	"APPLICATION": 1,
	"USER":        2,
}

func (x WhereIsRole) String() string {
	return proto.EnumName(WhereIsRole_name, int32(x))
}

func (WhereIsRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{1}
}

type WhereIs struct {
	// DID of the created bucket.
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Creator of the new bucket
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Label of the new bucket.
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// Visibility of the new bucket.
	Visibility WhereIsVisibility `protobuf:"varint,4,opt,name=visibility,proto3,enum=sonrio.sonr.bucket.WhereIsVisibility" json:"visibility,omitempty"`
	// Role of the creator of the new bucket.
	Role WhereIsRole `protobuf:"varint,5,opt,name=role,proto3,enum=sonrio.sonr.bucket.WhereIsRole" json:"role,omitempty"`
	// IsActive flag of the new bucket.
	IsActive bool `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Content of the new bucket map of DIDs to CIDs.
	Content map[string]string `protobuf:"bytes,7,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// JWKs of the new bucket.
	Jwks map[string][]byte `protobuf:"bytes,8,rep,name=jwks,proto3" json:"jwks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Timestamp of the new bucket.
	Timestamp int64 `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *WhereIs) Reset()         { *m = WhereIs{} }
func (m *WhereIs) String() string { return proto.CompactTextString(m) }
func (*WhereIs) ProtoMessage()    {}
func (*WhereIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{0}
}
func (m *WhereIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhereIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhereIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhereIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhereIs.Merge(m, src)
}
func (m *WhereIs) XXX_Size() int {
	return m.Size()
}
func (m *WhereIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhereIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhereIs proto.InternalMessageInfo

func (m *WhereIs) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *WhereIs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *WhereIs) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WhereIs) GetVisibility() WhereIsVisibility {
	if m != nil {
		return m.Visibility
	}
	return WhereIsVisibility_UNSPECIFIED
}

func (m *WhereIs) GetRole() WhereIsRole {
	if m != nil {
		return m.Role
	}
	return WhereIsRole_NONE
}

func (m *WhereIs) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *WhereIs) GetContent() map[string]string {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *WhereIs) GetJwks() map[string][]byte {
	if m != nil {
		return m.Jwks
	}
	return nil
}

func (m *WhereIs) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("sonrio.sonr.bucket.WhereIsVisibility", WhereIsVisibility_name, WhereIsVisibility_value)
	proto.RegisterEnum("sonrio.sonr.bucket.WhereIsRole", WhereIsRole_name, WhereIsRole_value)
	proto.RegisterType((*WhereIs)(nil), "sonrio.sonr.bucket.WhereIs")
	proto.RegisterMapType((map[string]string)(nil), "sonrio.sonr.bucket.WhereIs.ContentEntry")
	proto.RegisterMapType((map[string][]byte)(nil), "sonrio.sonr.bucket.WhereIs.JwksEntry")
}

func init() { proto.RegisterFile("bucket/where_is.proto", fileDescriptor_9d51d2bd3c178d08) }

var fileDescriptor_9d51d2bd3c178d08 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdf, 0x6b, 0xd3, 0x40,
	0x1c, 0xef, 0x35, 0x5d, 0xd3, 0x7c, 0x3b, 0x34, 0x1e, 0x0a, 0xc7, 0x94, 0x18, 0x04, 0x25, 0x0c,
	0x4c, 0xa1, 0x7b, 0x50, 0x07, 0x82, 0x6d, 0x8d, 0x10, 0x19, 0x5d, 0xc8, 0xd6, 0x09, 0xbe, 0x8c,
	0x24, 0x3b, 0xdc, 0x99, 0xb4, 0x57, 0x72, 0xd7, 0xd6, 0xfc, 0x17, 0xfe, 0x59, 0x3e, 0xee, 0xd1,
	0x47, 0x69, 0xff, 0x07, 0x9f, 0x25, 0x97, 0x76, 0x16, 0x06, 0xdd, 0xd3, 0xe5, 0xf3, 0xcd, 0xe7,
	0x07, 0x77, 0x9f, 0x2f, 0x3c, 0x89, 0x67, 0x49, 0x4a, 0x65, 0x67, 0x71, 0x4d, 0x73, 0x7a, 0xc9,
	0x84, 0x3b, 0xcd, 0xb9, 0xe4, 0x18, 0x0b, 0x3e, 0xc9, 0x19, 0x77, 0xcb, 0xc3, 0xad, 0x28, 0x2f,
	0xfe, 0x6a, 0xa0, 0x7f, 0x29, 0x69, 0xbe, 0xc0, 0x26, 0x68, 0x57, 0xec, 0x8a, 0x20, 0x1b, 0x39,
	0x46, 0x58, 0x7e, 0x62, 0x02, 0x7a, 0x92, 0xd3, 0x48, 0xf2, 0x9c, 0xd4, 0xd5, 0x74, 0x03, 0xf1,
	0x63, 0xd8, 0xcb, 0xa2, 0x98, 0x66, 0x44, 0x53, 0xf3, 0x0a, 0x60, 0x0f, 0x60, 0xce, 0x04, 0x8b,
	0x59, 0xc6, 0x64, 0x41, 0x1a, 0x36, 0x72, 0x1e, 0x74, 0x5f, 0xba, 0x77, 0x63, 0xdd, 0x75, 0xe4,
	0xc5, 0x2d, 0x39, 0xdc, 0x12, 0xe2, 0x23, 0x68, 0xe4, 0x3c, 0xa3, 0x64, 0x4f, 0x19, 0x3c, 0xdf,
	0x61, 0x10, 0xf2, 0x8c, 0x86, 0x8a, 0x8c, 0x9f, 0x82, 0xc1, 0xc4, 0x65, 0x94, 0x48, 0x36, 0xa7,
	0xa4, 0x69, 0x23, 0xa7, 0x15, 0xb6, 0x98, 0xe8, 0x29, 0x8c, 0xfb, 0xa0, 0x27, 0x7c, 0x22, 0xe9,
	0x44, 0x12, 0xdd, 0xd6, 0x9c, 0x76, 0xd7, 0xd9, 0x61, 0xea, 0x0e, 0x2a, 0xaa, 0x37, 0x91, 0x79,
	0x11, 0x6e, 0x84, 0xf8, 0x1d, 0x34, 0xbe, 0x2f, 0x52, 0x41, 0x5a, 0xca, 0x60, 0xd7, 0xb5, 0xdc,
	0xcf, 0x8b, 0x54, 0x54, 0x6a, 0x25, 0xc1, 0xcf, 0xc0, 0x90, 0x6c, 0x4c, 0x85, 0x8c, 0xc6, 0x53,
	0x62, 0xd8, 0xc8, 0xd1, 0xc2, 0xff, 0x83, 0x83, 0x63, 0xd8, 0xdf, 0x4e, 0x2c, 0x7b, 0x48, 0x69,
	0xb1, 0xe9, 0x21, 0xa5, 0x45, 0xf9, 0xda, 0xf3, 0x28, 0x9b, 0xd1, 0x75, 0x0b, 0x15, 0x38, 0xae,
	0xbf, 0x45, 0x07, 0x6f, 0xc0, 0xb8, 0x0d, 0xbb, 0x4f, 0xb8, 0xbf, 0x25, 0x3c, 0x7c, 0x0f, 0x8f,
	0xee, 0x94, 0x80, 0x1f, 0x42, 0x7b, 0x34, 0x3c, 0x0b, 0xbc, 0x81, 0xff, 0xc9, 0xf7, 0x3e, 0x9a,
	0x35, 0x0c, 0xd0, 0x0c, 0x46, 0xfd, 0x13, 0x7f, 0x60, 0x22, 0xdc, 0x06, 0x3d, 0x08, 0xfd, 0x8b,
	0xde, 0xb9, 0x67, 0xd6, 0x0f, 0xbb, 0xd0, 0xde, 0xaa, 0x00, 0xb7, 0xa0, 0x31, 0x3c, 0x1d, 0x7a,
	0x66, 0xad, 0xb4, 0xe8, 0x05, 0xc1, 0x89, 0x3f, 0xe8, 0x9d, 0xfb, 0xa7, 0x43, 0x13, 0x95, 0xbf,
	0x46, 0x67, 0x5e, 0x68, 0xd6, 0xfb, 0x1f, 0x7e, 0x2d, 0x2d, 0x74, 0xb3, 0xb4, 0xd0, 0x9f, 0xa5,
	0x85, 0x7e, 0xae, 0xac, 0xda, 0xcd, 0xca, 0xaa, 0xfd, 0x5e, 0x59, 0xb5, 0xaf, 0xaf, 0xbe, 0x31,
	0x79, 0x3d, 0x8b, 0xdd, 0x84, 0x8f, 0x3b, 0xe5, 0x7b, 0xbe, 0x66, 0x5c, 0x9d, 0x9d, 0x1f, 0x9d,
	0xf5, 0x2a, 0xcb, 0x62, 0x4a, 0x45, 0xdc, 0x54, 0x8b, 0x7c, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0xa4, 0x17, 0x2a, 0xe1, 0x02, 0x00, 0x00,
}

func (m *WhereIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhereIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhereIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Jwks) > 0 {
		for k := range m.Jwks {
			v := m.Jwks[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintWhereIs(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWhereIs(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWhereIs(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Content) > 0 {
		for k := range m.Content {
			v := m.Content[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWhereIs(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWhereIs(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWhereIs(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Role != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x28
	}
	if m.Visibility != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Visibility))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhereIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhereIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhereIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	if m.Visibility != 0 {
		n += 1 + sovWhereIs(uint64(m.Visibility))
	}
	if m.Role != 0 {
		n += 1 + sovWhereIs(uint64(m.Role))
	}
	if m.IsActive {
		n += 2
	}
	if len(m.Content) > 0 {
		for k, v := range m.Content {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWhereIs(uint64(len(k))) + 1 + len(v) + sovWhereIs(uint64(len(v)))
			n += mapEntrySize + 1 + sovWhereIs(uint64(mapEntrySize))
		}
	}
	if len(m.Jwks) > 0 {
		for k, v := range m.Jwks {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovWhereIs(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovWhereIs(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovWhereIs(uint64(mapEntrySize))
		}
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhereIs(uint64(m.Timestamp))
	}
	return n
}

func sovWhereIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhereIs(x uint64) (n int) {
	return sovWhereIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhereIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhereIs
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
			return fmt.Errorf("proto: WhereIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhereIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
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
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
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
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
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
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Visibility", wireType)
			}
			m.Visibility = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Visibility |= WhereIsVisibility(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= WhereIsRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
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
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWhereIs
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWhereIs
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWhereIs(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWhereIs
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Content[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jwks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
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
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Jwks == nil {
				m.Jwks = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWhereIs
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthWhereIs
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWhereIs(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWhereIs
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Jwks[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWhereIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhereIs
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
func skipWhereIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhereIs
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
					return 0, ErrIntOverflowWhereIs
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
					return 0, ErrIntOverflowWhereIs
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
				return 0, ErrInvalidLengthWhereIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhereIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhereIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhereIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhereIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhereIs = fmt.Errorf("proto: unexpected end of group")
)
