// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

/*
	Package meta is a generated protocol buffer package.

	It is generated from these files:
		pb.proto

	It has these top-level messages:
		Label
		Peer
		Epoch
		VectorDB
		Store
*/
package meta

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StoreState store state
type StoreState int32

const (
	UP        StoreState = 0
	Down      StoreState = 1
	Tombstone StoreState = 2
)

var StoreState_name = map[int32]string{
	0: "UP",
	1: "Down",
	2: "Tombstone",
}
var StoreState_value = map[string]int32{
	"UP":        0,
	"Down":      1,
	"Tombstone": 2,
}

func (x StoreState) Enum() *StoreState {
	p := new(StoreState)
	*p = x
	return p
}
func (x StoreState) String() string {
	return proto.EnumName(StoreState_name, int32(x))
}
func (x *StoreState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StoreState_value, data, "StoreState")
	if err != nil {
		return err
	}
	*x = StoreState(value)
	return nil
}
func (StoreState) EnumDescriptor() ([]byte, []int) { return fileDescriptorPb, []int{0} }

// DBState db state
type DBState int32

const (
	// RU read and update
	RU DBState = 0
	// RWU read, append write and update
	RWU DBState = 1
)

var DBState_name = map[int32]string{
	0: "RU",
	1: "RWU",
}
var DBState_value = map[string]int32{
	"RU":  0,
	"RWU": 1,
}

func (x DBState) Enum() *DBState {
	p := new(DBState)
	*p = x
	return p
}
func (x DBState) String() string {
	return proto.EnumName(DBState_name, int32(x))
}
func (x *DBState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DBState_value, data, "DBState")
	if err != nil {
		return err
	}
	*x = DBState(value)
	return nil
}
func (DBState) EnumDescriptor() ([]byte, []int) { return fileDescriptorPb, []int{1} }

type Label struct {
	Key              string `protobuf:"bytes,1,opt,name=key" json:"key"`
	Value            string `protobuf:"bytes,2,opt,name=value" json:"value"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{0} }

func (m *Label) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Peer struct {
	ID               uint64 `protobuf:"varint,1,opt,name=id" json:"id"`
	StoreID          uint64 `protobuf:"varint,2,opt,name=storeID" json:"storeID"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{1} }

func (m *Peer) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Peer) GetStoreID() uint64 {
	if m != nil {
		return m.StoreID
	}
	return 0
}

type Epoch struct {
	Version          uint64 `protobuf:"varint,1,opt,name=version" json:"version"`
	ConfVersion      uint64 `protobuf:"varint,2,opt,name=confVersion" json:"confVersion"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Epoch) Reset()                    { *m = Epoch{} }
func (m *Epoch) String() string            { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()               {}
func (*Epoch) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{2} }

func (m *Epoch) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Epoch) GetConfVersion() uint64 {
	if m != nil {
		return m.ConfVersion
	}
	return 0
}

// VectorDB is a vectorDB
type VectorDB struct {
	ID               uint64  `protobuf:"varint,1,opt,name=id" json:"id"`
	State            DBState `protobuf:"varint,2,opt,name=state,enum=meta.DBState" json:"state"`
	Start            uint64  `protobuf:"varint,3,opt,name=start" json:"start"`
	Epoch            Epoch   `protobuf:"bytes,4,opt,name=epoch" json:"epoch"`
	Peers            []*Peer `protobuf:"bytes,5,rep,name=peers" json:"peers,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VectorDB) Reset()                    { *m = VectorDB{} }
func (m *VectorDB) String() string            { return proto.CompactTextString(m) }
func (*VectorDB) ProtoMessage()               {}
func (*VectorDB) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{3} }

func (m *VectorDB) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *VectorDB) GetState() DBState {
	if m != nil {
		return m.State
	}
	return RU
}

func (m *VectorDB) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *VectorDB) GetEpoch() Epoch {
	if m != nil {
		return m.Epoch
	}
	return Epoch{}
}

func (m *VectorDB) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Store container of vectordb
type Store struct {
	ID               uint64     `protobuf:"varint,1,opt,name=id" json:"id"`
	Address          string     `protobuf:"bytes,2,opt,name=address" json:"address"`
	ClientAddress    string     `protobuf:"bytes,3,opt,name=clientAddress" json:"clientAddress"`
	Lables           []Label    `protobuf:"bytes,4,rep,name=lables" json:"lables"`
	State            StoreState `protobuf:"varint,5,opt,name=state,enum=meta.StoreState" json:"state"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{4} }

func (m *Store) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Store) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Store) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (m *Store) GetLables() []Label {
	if m != nil {
		return m.Lables
	}
	return nil
}

func (m *Store) GetState() StoreState {
	if m != nil {
		return m.State
	}
	return UP
}

func init() {
	proto.RegisterType((*Label)(nil), "meta.Label")
	proto.RegisterType((*Peer)(nil), "meta.Peer")
	proto.RegisterType((*Epoch)(nil), "meta.Epoch")
	proto.RegisterType((*VectorDB)(nil), "meta.VectorDB")
	proto.RegisterType((*Store)(nil), "meta.Store")
	proto.RegisterEnum("meta.StoreState", StoreState_name, StoreState_value)
	proto.RegisterEnum("meta.DBState", DBState_name, DBState_value)
}
func (m *Label) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Label) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintPb(dAtA, i, uint64(len(m.Key)))
	i += copy(dAtA[i:], m.Key)
	dAtA[i] = 0x12
	i++
	i = encodeVarintPb(dAtA, i, uint64(len(m.Value)))
	i += copy(dAtA[i:], m.Value)
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.ID))
	dAtA[i] = 0x10
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.StoreID))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Epoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Epoch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.Version))
	dAtA[i] = 0x10
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.ConfVersion))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *VectorDB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VectorDB) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.ID))
	dAtA[i] = 0x10
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.State))
	dAtA[i] = 0x18
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.Start))
	dAtA[i] = 0x22
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.Epoch.Size()))
	n1, err := m.Epoch.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Peers) > 0 {
		for _, msg := range m.Peers {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintPb(dAtA, i, uint64(msg.Size()))
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

func (m *Store) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Store) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.ID))
	dAtA[i] = 0x12
	i++
	i = encodeVarintPb(dAtA, i, uint64(len(m.Address)))
	i += copy(dAtA[i:], m.Address)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintPb(dAtA, i, uint64(len(m.ClientAddress)))
	i += copy(dAtA[i:], m.ClientAddress)
	if len(m.Lables) > 0 {
		for _, msg := range m.Lables {
			dAtA[i] = 0x22
			i++
			i = encodeVarintPb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x28
	i++
	i = encodeVarintPb(dAtA, i, uint64(m.State))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Label) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovPb(uint64(l))
	l = len(m.Value)
	n += 1 + l + sovPb(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Peer) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPb(uint64(m.ID))
	n += 1 + sovPb(uint64(m.StoreID))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Epoch) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPb(uint64(m.Version))
	n += 1 + sovPb(uint64(m.ConfVersion))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VectorDB) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPb(uint64(m.ID))
	n += 1 + sovPb(uint64(m.State))
	n += 1 + sovPb(uint64(m.Start))
	l = m.Epoch.Size()
	n += 1 + l + sovPb(uint64(l))
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Store) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPb(uint64(m.ID))
	l = len(m.Address)
	n += 1 + l + sovPb(uint64(l))
	l = len(m.ClientAddress)
	n += 1 + l + sovPb(uint64(l))
	if len(m.Lables) > 0 {
		for _, e := range m.Lables {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	n += 1 + sovPb(uint64(m.State))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Label) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: Label: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Label: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *Epoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: Epoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfVersion", wireType)
			}
			m.ConfVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfVersion |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *VectorDB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: VectorDB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VectorDB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (DBState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Epoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *Store) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: Store: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Store: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lables = append(m.Lables, Label{})
			if err := m.Lables[len(m.Lables)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (StoreState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
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
					return 0, ErrIntOverflowPb
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
					return 0, ErrIntOverflowPb
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
				return 0, ErrInvalidLengthPb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPb
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
				next, err := skipPb(dAtA[start:])
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
	ErrInvalidLengthPb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb.proto", fileDescriptorPb) }

var fileDescriptorPb = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x18, 0xc4, 0xbd, 0xf6, 0x6e, 0x93, 0x7e, 0x51, 0x90, 0xb5, 0xaa, 0x90, 0x95, 0x83, 0x1b, 0xf9,
	0x50, 0xd2, 0x08, 0x82, 0x94, 0x2b, 0x27, 0x2c, 0x73, 0xa8, 0x84, 0x44, 0xe5, 0x92, 0x72, 0x76,
	0x92, 0x8f, 0x12, 0xe1, 0x7a, 0xa3, 0xdd, 0xa5, 0x88, 0x37, 0xe1, 0x69, 0x38, 0xe7, 0xd8, 0x27,
	0xa8, 0x20, 0xbc, 0x08, 0xda, 0x3f, 0x16, 0x8e, 0x90, 0xe8, 0xcd, 0x9a, 0xf9, 0xed, 0x78, 0x67,
	0x6c, 0xe8, 0x6f, 0x97, 0xb3, 0xad, 0x14, 0x5a, 0x70, 0x7a, 0x8b, 0xba, 0x1a, 0x9d, 0xdc, 0x88,
	0x1b, 0x61, 0x85, 0x97, 0xe6, 0xc9, 0x79, 0xd9, 0x2b, 0x60, 0x6f, 0xab, 0x25, 0xd6, 0xfc, 0x29,
	0x44, 0x9f, 0xf1, 0x5b, 0x42, 0xc6, 0x64, 0x72, 0x9c, 0xd3, 0xdd, 0xc3, 0x69, 0x50, 0x1a, 0x81,
	0x8f, 0x80, 0xdd, 0x55, 0xf5, 0x17, 0x4c, 0xc2, 0x8e, 0xe3, 0xa4, 0x2c, 0x07, 0x7a, 0x89, 0x28,
	0xf9, 0x08, 0xc2, 0xcd, 0xda, 0x1e, 0xa5, 0x39, 0x18, 0x60, 0xff, 0x70, 0x1a, 0x5e, 0x14, 0x65,
	0xb8, 0x59, 0xf3, 0x14, 0x7a, 0x4a, 0x0b, 0x89, 0x17, 0x85, 0x4d, 0xa0, 0x3e, 0xa1, 0x15, 0xb3,
	0x77, 0xc0, 0xde, 0x6c, 0xc5, 0xea, 0x93, 0x01, 0xef, 0x50, 0xaa, 0x8d, 0x68, 0x7c, 0x92, 0x07,
	0xbd, 0xc8, 0xcf, 0x60, 0xb0, 0x12, 0xcd, 0xc7, 0x6b, 0xcf, 0x74, 0xc3, 0xba, 0x46, 0xf6, 0x83,
	0x40, 0xff, 0x1a, 0x57, 0x5a, 0xc8, 0x22, 0xff, 0xef, 0xcd, 0xce, 0x81, 0x29, 0x5d, 0x69, 0xd7,
	0xec, 0xc9, 0x7c, 0x38, 0x33, 0x33, 0xcd, 0x8a, 0xfc, 0xca, 0x88, 0x6d, 0x51, 0x4b, 0x98, 0x11,
	0x94, 0xae, 0xa4, 0x4e, 0xa2, 0xce, 0x5b, 0x9d, 0xc4, 0x9f, 0x01, 0x43, 0x53, 0x20, 0xa1, 0x63,
	0x32, 0x19, 0xcc, 0x07, 0x2e, 0xc6, 0x76, 0x6a, 0x41, 0xeb, 0xf3, 0x33, 0x60, 0x5b, 0x44, 0xa9,
	0x12, 0x36, 0x8e, 0x26, 0x83, 0x39, 0x38, 0xd0, 0x0c, 0x68, 0x39, 0x52, 0x3a, 0x3b, 0xdb, 0x11,
	0x60, 0x57, 0x66, 0x9d, 0xc7, 0x76, 0xad, 0xd6, 0x6b, 0x89, 0x4a, 0x1d, 0x7c, 0x99, 0x56, 0xe4,
	0x53, 0x18, 0xae, 0xea, 0x0d, 0x36, 0xfa, 0xb5, 0xa7, 0xa2, 0x0e, 0x75, 0x68, 0xf1, 0x73, 0x38,
	0xaa, 0xab, 0x65, 0x8d, 0x2a, 0xa1, 0xf6, 0x6a, 0xbe, 0x83, 0xfd, 0x31, 0xfc, 0x09, 0x0f, 0xf0,
	0xe7, 0xed, 0x68, 0xcc, 0x8e, 0x16, 0x3b, 0xd2, 0x5e, 0xf7, 0xdf, 0xdd, 0xa6, 0x2f, 0x00, 0xfe,
	0x5a, 0xfc, 0x08, 0xc2, 0xc5, 0x65, 0x1c, 0xf0, 0x3e, 0xd0, 0x42, 0x7c, 0x6d, 0x62, 0xc2, 0x87,
	0x70, 0xfc, 0x5e, 0xdc, 0x2e, 0x95, 0x16, 0x0d, 0xc6, 0xe1, 0x74, 0x04, 0x3d, 0x3f, 0xbf, 0x61,
	0xcb, 0x45, 0x1c, 0xf0, 0x1e, 0x44, 0xe5, 0x87, 0x45, 0x4c, 0xf2, 0x93, 0xfb, 0x5f, 0x69, 0xb0,
	0xdb, 0xa7, 0xe4, 0x7e, 0x9f, 0x92, 0x9f, 0xfb, 0x94, 0x7c, 0xff, 0x9d, 0x06, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xce, 0xef, 0xcd, 0x64, 0xe5, 0x02, 0x00, 0x00,
}
