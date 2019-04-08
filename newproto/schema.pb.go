// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

/*
	Package exampleproto is a generated protocol buffer package.

	It is generated from these files:
		schema.proto

	It has these top-level messages:
		Something
		StringBackedMessage
		BytesBackedMessage
		NewSchema
*/
package exampleproto

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Something struct {
	FieldX  int64 `protobuf:"varint,1,opt,name=fieldX,proto3" json:"fieldX,omitempty"`
	FieldYY int64 `protobuf:"varint,2,opt,name=fieldYY,proto3" json:"fieldYY,omitempty"`
	FieldZ  int64 `protobuf:"varint,3,opt,name=fieldZ,proto3" json:"fieldZ,omitempty"`
}

func (m *Something) Reset()                    { *m = Something{} }
func (m *Something) String() string            { return proto.CompactTextString(m) }
func (*Something) ProtoMessage()               {}
func (*Something) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{0} }

func (m *Something) GetFieldX() int64 {
	if m != nil {
		return m.FieldX
	}
	return 0
}

func (m *Something) GetFieldYY() int64 {
	if m != nil {
		return m.FieldYY
	}
	return 0
}

func (m *Something) GetFieldZ() int64 {
	if m != nil {
		return m.FieldZ
	}
	return 0
}

type StringBackedMessage struct {
	FieldX  string `protobuf:"bytes,1,opt,name=fieldX,proto3" json:"fieldX,omitempty"`
	FieldYY string `protobuf:"bytes,2,opt,name=fieldYY,proto3" json:"fieldYY,omitempty"`
	FieldZ  string `protobuf:"bytes,3,opt,name=fieldZ,proto3" json:"fieldZ,omitempty"`
}

func (m *StringBackedMessage) Reset()                    { *m = StringBackedMessage{} }
func (m *StringBackedMessage) String() string            { return proto.CompactTextString(m) }
func (*StringBackedMessage) ProtoMessage()               {}
func (*StringBackedMessage) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{1} }

func (m *StringBackedMessage) GetFieldX() string {
	if m != nil {
		return m.FieldX
	}
	return ""
}

func (m *StringBackedMessage) GetFieldYY() string {
	if m != nil {
		return m.FieldYY
	}
	return ""
}

func (m *StringBackedMessage) GetFieldZ() string {
	if m != nil {
		return m.FieldZ
	}
	return ""
}

type BytesBackedMessage struct {
	FieldX  []byte `protobuf:"bytes,1,opt,name=fieldX,proto3" json:"fieldX,omitempty"`
	FieldYY []byte `protobuf:"bytes,2,opt,name=fieldYY,proto3" json:"fieldYY,omitempty"`
	FieldZ  []byte `protobuf:"bytes,3,opt,name=fieldZ,proto3" json:"fieldZ,omitempty"`
}

func (m *BytesBackedMessage) Reset()                    { *m = BytesBackedMessage{} }
func (m *BytesBackedMessage) String() string            { return proto.CompactTextString(m) }
func (*BytesBackedMessage) ProtoMessage()               {}
func (*BytesBackedMessage) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{2} }

func (m *BytesBackedMessage) GetFieldX() []byte {
	if m != nil {
		return m.FieldX
	}
	return nil
}

func (m *BytesBackedMessage) GetFieldYY() []byte {
	if m != nil {
		return m.FieldYY
	}
	return nil
}

func (m *BytesBackedMessage) GetFieldZ() []byte {
	if m != nil {
		return m.FieldZ
	}
	return nil
}

type NewSchema struct {
	FieldX  int64 `protobuf:"varint,1,opt,name=fieldX,proto3" json:"fieldX,omitempty"`
	FieldY  int64 `protobuf:"varint,2,opt,name=fieldY,proto3" json:"fieldY,omitempty"`
	FieldZ  int64 `protobuf:"varint,3,opt,name=fieldZ,proto3" json:"fieldZ,omitempty"`
	FieldZZ int64 `protobuf:"varint,4,opt,name=fieldZZ,proto3" json:"fieldZZ,omitempty"`
}

func (m *NewSchema) Reset()                    { *m = NewSchema{} }
func (m *NewSchema) String() string            { return proto.CompactTextString(m) }
func (*NewSchema) ProtoMessage()               {}
func (*NewSchema) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{3} }

func (m *NewSchema) GetFieldX() int64 {
	if m != nil {
		return m.FieldX
	}
	return 0
}

func (m *NewSchema) GetFieldY() int64 {
	if m != nil {
		return m.FieldY
	}
	return 0
}

func (m *NewSchema) GetFieldZ() int64 {
	if m != nil {
		return m.FieldZ
	}
	return 0
}

func (m *NewSchema) GetFieldZZ() int64 {
	if m != nil {
		return m.FieldZZ
	}
	return 0
}

func init() {
	proto.RegisterType((*Something)(nil), "exampleproto.Something")
	proto.RegisterType((*StringBackedMessage)(nil), "exampleproto.StringBackedMessage")
	proto.RegisterType((*BytesBackedMessage)(nil), "exampleproto.BytesBackedMessage")
	proto.RegisterType((*NewSchema)(nil), "exampleproto.NewSchema")
}
func (m *Something) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Something) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FieldX != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldX))
	}
	if m.FieldYY != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldYY))
	}
	if m.FieldZ != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldZ))
	}
	return i, nil
}

func (m *StringBackedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringBackedMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FieldX) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.FieldX)))
		i += copy(dAtA[i:], m.FieldX)
	}
	if len(m.FieldYY) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.FieldYY)))
		i += copy(dAtA[i:], m.FieldYY)
	}
	if len(m.FieldZ) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.FieldZ)))
		i += copy(dAtA[i:], m.FieldZ)
	}
	return i, nil
}

func (m *BytesBackedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BytesBackedMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FieldX) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.FieldX)))
		i += copy(dAtA[i:], m.FieldX)
	}
	if len(m.FieldYY) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.FieldYY)))
		i += copy(dAtA[i:], m.FieldYY)
	}
	if len(m.FieldZ) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.FieldZ)))
		i += copy(dAtA[i:], m.FieldZ)
	}
	return i, nil
}

func (m *NewSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewSchema) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FieldX != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldX))
	}
	if m.FieldY != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldY))
	}
	if m.FieldZ != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldZ))
	}
	if m.FieldZZ != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.FieldZZ))
	}
	return i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Something) Size() (n int) {
	var l int
	_ = l
	if m.FieldX != 0 {
		n += 1 + sovSchema(uint64(m.FieldX))
	}
	if m.FieldYY != 0 {
		n += 1 + sovSchema(uint64(m.FieldYY))
	}
	if m.FieldZ != 0 {
		n += 1 + sovSchema(uint64(m.FieldZ))
	}
	return n
}

func (m *StringBackedMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.FieldX)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.FieldYY)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.FieldZ)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	return n
}

func (m *BytesBackedMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.FieldX)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.FieldYY)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.FieldZ)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	return n
}

func (m *NewSchema) Size() (n int) {
	var l int
	_ = l
	if m.FieldX != 0 {
		n += 1 + sovSchema(uint64(m.FieldX))
	}
	if m.FieldY != 0 {
		n += 1 + sovSchema(uint64(m.FieldY))
	}
	if m.FieldZ != 0 {
		n += 1 + sovSchema(uint64(m.FieldZ))
	}
	if m.FieldZZ != 0 {
		n += 1 + sovSchema(uint64(m.FieldZZ))
	}
	return n
}

func sovSchema(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Something) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
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
			return fmt.Errorf("proto: Something: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Something: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldX", wireType)
			}
			m.FieldX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldX |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldYY", wireType)
			}
			m.FieldYY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldYY |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldZ", wireType)
			}
			m.FieldZ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldZ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *StringBackedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
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
			return fmt.Errorf("proto: StringBackedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringBackedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
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
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldYY", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
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
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldYY = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldZ", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
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
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldZ = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *BytesBackedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
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
			return fmt.Errorf("proto: BytesBackedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BytesBackedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldX", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldX = append(m.FieldX[:0], dAtA[iNdEx:postIndex]...)
			if m.FieldX == nil {
				m.FieldX = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldYY", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldYY = append(m.FieldYY[:0], dAtA[iNdEx:postIndex]...)
			if m.FieldYY == nil {
				m.FieldYY = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldZ", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldZ = append(m.FieldZ[:0], dAtA[iNdEx:postIndex]...)
			if m.FieldZ == nil {
				m.FieldZ = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *NewSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
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
			return fmt.Errorf("proto: NewSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldX", wireType)
			}
			m.FieldX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldX |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldY", wireType)
			}
			m.FieldY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldY |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldZ", wireType)
			}
			m.FieldZ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldZ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldZZ", wireType)
			}
			m.FieldZZ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldZZ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
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
				return 0, ErrInvalidLengthSchema
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchema
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
				next, err := skipSchema(dAtA[start:])
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
	ErrInvalidLengthSchema = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("schema.proto", fileDescriptorSchema) }

var fileDescriptorSchema = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0x05, 0xf3, 0x94, 0x42, 0xb9, 0x38, 0x83, 0xf3, 0x73, 0x53, 0x4b, 0x32, 0x32, 0xf3, 0xd2,
	0x85, 0xc4, 0xb8, 0xd8, 0xd2, 0x32, 0x53, 0x73, 0x52, 0x22, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98,
	0x83, 0xa0, 0x3c, 0x21, 0x09, 0x2e, 0x76, 0x30, 0x2b, 0x32, 0x52, 0x82, 0x09, 0x2c, 0x01, 0xe3,
	0xc2, 0x75, 0x44, 0x49, 0x30, 0x23, 0xe9, 0x88, 0x52, 0x8a, 0xe7, 0x12, 0x0e, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0x77, 0x4a, 0x4c, 0xce, 0x4e, 0x4d, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x45,
	0xb3, 0x80, 0x13, 0x97, 0x05, 0x9c, 0xb8, 0x2c, 0xe0, 0x84, 0x5b, 0x10, 0xc7, 0x25, 0xe4, 0x54,
	0x59, 0x92, 0x5a, 0x8c, 0xcf, 0x7c, 0x1e, 0x5c, 0xe6, 0xf3, 0xe0, 0x32, 0x9f, 0x07, 0x6e, 0x7e,
	0x2e, 0x17, 0xa7, 0x5f, 0x6a, 0x79, 0x30, 0x38, 0xe0, 0x70, 0x86, 0x0b, 0x4c, 0x1c, 0x16, 0x2c,
	0x50, 0x1e, 0xae, 0x50, 0x81, 0x3b, 0x23, 0x2a, 0x4a, 0x82, 0x05, 0x29, 0x1c, 0xa3, 0xa2, 0x9c,
	0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f,
	0xe5, 0x18, 0x92, 0xd8, 0xc0, 0xf1, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xac, 0x7f,
	0x97, 0xbd, 0x01, 0x00, 0x00,
}
