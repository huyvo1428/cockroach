// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/sqlbase/privilege.proto

package sqlbase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

// UserPrivileges describes the list of privileges available for a given user.
type UserPrivileges struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user"`
	// privileges is a bitfield of 1<<Privilege values.
	Privileges           uint32   `protobuf:"varint,2,opt,name=privileges" json:"privileges"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPrivileges) Reset()         { *m = UserPrivileges{} }
func (m *UserPrivileges) String() string { return proto.CompactTextString(m) }
func (*UserPrivileges) ProtoMessage()    {}
func (*UserPrivileges) Descriptor() ([]byte, []int) {
	return fileDescriptor_privilege_a8444a679d17e9af, []int{0}
}
func (m *UserPrivileges) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserPrivileges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *UserPrivileges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPrivileges.Merge(dst, src)
}
func (m *UserPrivileges) XXX_Size() int {
	return m.Size()
}
func (m *UserPrivileges) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPrivileges.DiscardUnknown(m)
}

var xxx_messageInfo_UserPrivileges proto.InternalMessageInfo

// PrivilegeDescriptor describes a list of users and attached
// privileges. The list should be sorted by user for fast access.
type PrivilegeDescriptor struct {
	Users                []UserPrivileges `protobuf:"bytes,1,rep,name=users" json:"users"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PrivilegeDescriptor) Reset()         { *m = PrivilegeDescriptor{} }
func (m *PrivilegeDescriptor) String() string { return proto.CompactTextString(m) }
func (*PrivilegeDescriptor) ProtoMessage()    {}
func (*PrivilegeDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_privilege_a8444a679d17e9af, []int{1}
}
func (m *PrivilegeDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrivilegeDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *PrivilegeDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivilegeDescriptor.Merge(dst, src)
}
func (m *PrivilegeDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *PrivilegeDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivilegeDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_PrivilegeDescriptor proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserPrivileges)(nil), "cockroach.sql.sqlbase.UserPrivileges")
	proto.RegisterType((*PrivilegeDescriptor)(nil), "cockroach.sql.sqlbase.PrivilegeDescriptor")
}
func (m *UserPrivileges) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserPrivileges) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintPrivilege(dAtA, i, uint64(len(m.User)))
	i += copy(dAtA[i:], m.User)
	dAtA[i] = 0x10
	i++
	i = encodeVarintPrivilege(dAtA, i, uint64(m.Privileges))
	return i, nil
}

func (m *PrivilegeDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrivilegeDescriptor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, msg := range m.Users {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPrivilege(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintPrivilege(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserPrivileges) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	n += 1 + l + sovPrivilege(uint64(l))
	n += 1 + sovPrivilege(uint64(m.Privileges))
	return n
}

func (m *PrivilegeDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovPrivilege(uint64(l))
		}
	}
	return n
}

func sovPrivilege(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPrivilege(x uint64) (n int) {
	return sovPrivilege(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserPrivileges) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivilege
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
			return fmt.Errorf("proto: UserPrivileges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserPrivileges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivilege
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
				return ErrInvalidLengthPrivilege
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			m.Privileges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivilege
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Privileges |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrivilege(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivilege
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
func (m *PrivilegeDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivilege
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
			return fmt.Errorf("proto: PrivilegeDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivilegeDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivilege
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
				return ErrInvalidLengthPrivilege
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, UserPrivileges{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivilege(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivilege
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
func skipPrivilege(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrivilege
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
					return 0, ErrIntOverflowPrivilege
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
					return 0, ErrIntOverflowPrivilege
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
				return 0, ErrInvalidLengthPrivilege
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPrivilege
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
				next, err := skipPrivilege(dAtA[start:])
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
	ErrInvalidLengthPrivilege = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrivilege   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/sqlbase/privilege.proto", fileDescriptor_privilege_a8444a679d17e9af)
}

var fileDescriptor_privilege_a8444a679d17e9af = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x2e, 0xcc, 0xd1,
	0x2f, 0x2e, 0xcc, 0x49, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2c, 0xcb, 0xcc, 0x49, 0x4d,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0xce, 0x4f, 0xce, 0x2e, 0xca, 0x4f,
	0x4c, 0xce, 0xd0, 0x2b, 0x2e, 0xcc, 0xd1, 0x83, 0x2a, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xab, 0xd0, 0x07, 0xb1, 0x20, 0x8a, 0x95, 0x02, 0xb8, 0xf8, 0x42, 0x8b, 0x53, 0x8b, 0x02, 0x60,
	0x66, 0x14, 0x0b, 0x49, 0x70, 0xb1, 0x94, 0x16, 0xa7, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x3a, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0x16, 0x11, 0x52, 0xe1, 0xe2, 0x82, 0xdb, 0x55,
	0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x0b, 0x95, 0x47, 0x12, 0x57, 0x8a, 0xe0, 0x12, 0x86, 0x9b,
	0xe6, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x5f, 0x24, 0xe4, 0xc8, 0xc5, 0x0a, 0x32,
	0xa4, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x55, 0x0f, 0xab, 0x2b, 0xf5, 0x50, 0x1d,
	0x03, 0x35, 0x1e, 0xa2, 0xd3, 0x49, 0xf1, 0xc4, 0x43, 0x39, 0x86, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0xbc, 0xf1, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0xa2, 0xd8, 0xa1, 0xda, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0xd3, 0x05, 0x13, 0x19, 0x01,
	0x00, 0x00,
}
