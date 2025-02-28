// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epochs/v1/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventEpochStart struct {
	// Epoch number, starting from 1.
	EpochNumber uint64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	// The start timestamp of the epoch.
	EpochStartTime time.Time `protobuf:"bytes,2,opt,name=epoch_start_time,json=epochStartTime,proto3,stdtime" json:"epoch_start_time"`
}

func (m *EventEpochStart) Reset()         { *m = EventEpochStart{} }
func (m *EventEpochStart) String() string { return proto.CompactTextString(m) }
func (*EventEpochStart) ProtoMessage()    {}
func (*EventEpochStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ce08b394f3742c9, []int{0}
}
func (m *EventEpochStart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEpochStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEpochStart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEpochStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEpochStart.Merge(m, src)
}
func (m *EventEpochStart) XXX_Size() int {
	return m.Size()
}
func (m *EventEpochStart) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEpochStart.DiscardUnknown(m)
}

var xxx_messageInfo_EventEpochStart proto.InternalMessageInfo

func (m *EventEpochStart) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *EventEpochStart) GetEpochStartTime() time.Time {
	if m != nil {
		return m.EpochStartTime
	}
	return time.Time{}
}

type EventEpochEnd struct {
	// Epoch number, starting from 1.
	EpochNumber uint64 `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *EventEpochEnd) Reset()         { *m = EventEpochEnd{} }
func (m *EventEpochEnd) String() string { return proto.CompactTextString(m) }
func (*EventEpochEnd) ProtoMessage()    {}
func (*EventEpochEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ce08b394f3742c9, []int{1}
}
func (m *EventEpochEnd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEpochEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEpochEnd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEpochEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEpochEnd.Merge(m, src)
}
func (m *EventEpochEnd) XXX_Size() int {
	return m.Size()
}
func (m *EventEpochEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEpochEnd.DiscardUnknown(m)
}

var xxx_messageInfo_EventEpochEnd proto.InternalMessageInfo

func (m *EventEpochEnd) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*EventEpochStart)(nil), "nibiru.epochs.v1.EventEpochStart")
	proto.RegisterType((*EventEpochEnd)(nil), "nibiru.epochs.v1.EventEpochEnd")
}

func init() { proto.RegisterFile("epochs/v1/event.proto", fileDescriptor_6ce08b394f3742c9) }

var fileDescriptor_6ce08b394f3742c9 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2d, 0xc8, 0x4f,
	0xce, 0x28, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0xc8, 0xcb, 0x4c, 0xca, 0x2c, 0x2a, 0xd5, 0x83, 0xc8, 0xea, 0x95, 0x19, 0x4a,
	0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x25, 0xf5, 0x41, 0x2c, 0x88, 0x3a, 0x29, 0xf9, 0xf4, 0xfc,
	0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0xbf, 0x24, 0x33, 0x37, 0xb5, 0xb8,
	0x24, 0x31, 0xb7, 0x00, 0xa2, 0x40, 0xa9, 0x85, 0x91, 0x8b, 0xdf, 0x15, 0x64, 0xb0, 0x2b, 0xc8,
	0xa4, 0xe0, 0x92, 0xc4, 0xa2, 0x12, 0x21, 0x45, 0x2e, 0x1e, 0xb0, 0xb9, 0xf1, 0x79, 0xa5, 0xb9,
	0x49, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0xdc, 0x60, 0x31, 0x3f, 0xb0, 0x90,
	0x90, 0x1f, 0x97, 0x00, 0x44, 0x49, 0x31, 0x48, 0x47, 0x3c, 0xc8, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x29, 0x3d, 0x88, 0x95, 0x7a, 0x30, 0x2b, 0xf5, 0x42, 0x60, 0x56, 0x3a, 0x71,
	0x9c, 0xb8, 0x27, 0xcf, 0x30, 0xe1, 0xbe, 0x3c, 0x63, 0x10, 0x5f, 0x2a, 0xdc, 0x3a, 0x90, 0xb4,
	0x92, 0x11, 0x17, 0x2f, 0xc2, 0x15, 0xae, 0x79, 0x29, 0x44, 0xb8, 0xc1, 0xc9, 0xed, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0xfd, 0xc0, 0x01, 0xe5, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x0f, 0x09, 0x34,
	0xfd, 0x0a, 0x7d, 0x68, 0xa0, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x5d, 0x6a, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x8b, 0x61, 0x38, 0x6b, 0x01, 0x00, 0x00,
}

func (m *EventEpochStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEpochStart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEpochStart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EpochStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvent(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.EpochNumber != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventEpochEnd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEpochEnd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEpochEnd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventEpochStart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovEvent(uint64(m.EpochNumber))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EpochStartTime)
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventEpochEnd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovEvent(uint64(m.EpochNumber))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventEpochStart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventEpochStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEpochStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventEpochEnd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventEpochEnd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEpochEnd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
