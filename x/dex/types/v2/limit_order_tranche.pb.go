// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/v2/limit_order_tranche.proto

package v2

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	github_com_neutron_org_neutron_v2_utils_math "github.com/neutron-org/neutron/v2/utils/math"
	types "github.com/neutron-org/neutron/v2/x/dex/types"
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

type LimitOrderTrancheKey struct {
	TradePairId           *types.TradePairID `protobuf:"bytes,1,opt,name=trade_pair_id,json=tradePairId,proto3" json:"trade_pair_id,omitempty"`
	TickIndexTakerToMaker int64              `protobuf:"varint,2,opt,name=tick_index_taker_to_maker,json=tickIndexTakerToMaker,proto3" json:"tick_index_taker_to_maker,omitempty"`
	TrancheKey            string             `protobuf:"bytes,3,opt,name=tranche_key,json=trancheKey,proto3" json:"tranche_key,omitempty"`
}

func (m *LimitOrderTrancheKey) Reset()         { *m = LimitOrderTrancheKey{} }
func (m *LimitOrderTrancheKey) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheKey) ProtoMessage()    {}
func (*LimitOrderTrancheKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_e233cc8c93e2baf2, []int{0}
}
func (m *LimitOrderTrancheKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheKey.Merge(m, src)
}
func (m *LimitOrderTrancheKey) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheKey.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheKey proto.InternalMessageInfo

func (m *LimitOrderTrancheKey) GetTradePairId() *types.TradePairID {
	if m != nil {
		return m.TradePairId
	}
	return nil
}

func (m *LimitOrderTrancheKey) GetTickIndexTakerToMaker() int64 {
	if m != nil {
		return m.TickIndexTakerToMaker
	}
	return 0
}

func (m *LimitOrderTrancheKey) GetTrancheKey() string {
	if m != nil {
		return m.TrancheKey
	}
	return ""
}

type LimitOrderTranche struct {
	Key                *LimitOrderTrancheKey                  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ReservesMakerDenom github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=reserves_maker_denom,json=reservesMakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserves_maker_denom" yaml:"reserves_maker_denom"`
	ReservesTakerDenom github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=reserves_taker_denom,json=reservesTakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserves_taker_denom" yaml:"reserves_taker_denom"`
	TotalMakerDenom    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=total_maker_denom,json=totalMakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_maker_denom" yaml:"total_maker_denom"`
	TotalTakerDenom    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_taker_denom,json=totalTakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_taker_denom" yaml:"total_taker_denom"`
	// JIT orders also use goodTilDate to handle deletion but represent a special case
	// All JIT orders have a goodTilDate of 0 and an exception is made to still still treat these orders as live
	// Order deletion still functions the same and the orders will be deleted at the end of the block
	ExpirationTime    *time.Time                                           `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time,omitempty"`
	PriceTakerToMaker github_com_neutron_org_neutron_v2_utils_math.PrecDec `protobuf:"bytes,7,opt,name=price_taker_to_maker,json=priceTakerToMaker,proto3,customtype=github.com/neutron-org/neutron/v2/utils/math.PrecDec" json:"price_taker_to_maker" yaml:"price_taker_to_maker"`
}

func (m *LimitOrderTranche) Reset()         { *m = LimitOrderTranche{} }
func (m *LimitOrderTranche) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTranche) ProtoMessage()    {}
func (*LimitOrderTranche) Descriptor() ([]byte, []int) {
	return fileDescriptor_e233cc8c93e2baf2, []int{1}
}
func (m *LimitOrderTranche) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTranche) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTranche.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTranche) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTranche.Merge(m, src)
}
func (m *LimitOrderTranche) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTranche) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTranche.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTranche proto.InternalMessageInfo

func (m *LimitOrderTranche) GetKey() *LimitOrderTrancheKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LimitOrderTranche) GetExpirationTime() *time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitOrderTrancheKey)(nil), "neutron.dex.v2.LimitOrderTrancheKey")
	proto.RegisterType((*LimitOrderTranche)(nil), "neutron.dex.v2.LimitOrderTranche")
}

func init() {
	proto.RegisterFile("neutron/dex/v2/limit_order_tranche.proto", fileDescriptor_e233cc8c93e2baf2)
}

var fileDescriptor_e233cc8c93e2baf2 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xbd, 0x6e, 0xd3, 0x50,
	0x14, 0xc7, 0x73, 0x09, 0xb4, 0xf4, 0x56, 0xb4, 0xaa, 0x15, 0x24, 0xb7, 0x48, 0x76, 0x65, 0x21,
	0x94, 0xa5, 0xd7, 0x28, 0x20, 0x84, 0x10, 0x53, 0x94, 0x25, 0x40, 0x44, 0x65, 0x79, 0x62, 0xb1,
	0x1c, 0xfb, 0xe2, 0x5c, 0x25, 0xf6, 0xb5, 0xae, 0x4f, 0xa2, 0xe4, 0x19, 0x58, 0xca, 0xc4, 0x0b,
	0xb0, 0xf3, 0x1a, 0x19, 0x3b, 0x22, 0x06, 0x83, 0x92, 0x8d, 0xb1, 0x4f, 0x80, 0xae, 0xed, 0x34,
	0x0e, 0x31, 0x20, 0xd4, 0xc9, 0x27, 0xe7, 0x2b, 0xbf, 0xf3, 0xf7, 0x39, 0xc6, 0xcd, 0x88, 0x8e,
	0x41, 0xf0, 0xc8, 0xf4, 0xe9, 0xd4, 0x9c, 0xb4, 0xcc, 0x11, 0x0b, 0x19, 0x38, 0x5c, 0xf8, 0x54,
	0x38, 0x20, 0xdc, 0xc8, 0x1b, 0x50, 0x12, 0x0b, 0x0e, 0x5c, 0x39, 0x28, 0x32, 0x89, 0x4f, 0xa7,
	0x64, 0xd2, 0x3a, 0xd1, 0x03, 0xce, 0x83, 0x11, 0x35, 0xb3, 0x68, 0x7f, 0xfc, 0xde, 0x04, 0x16,
	0xd2, 0x04, 0xdc, 0x30, 0xce, 0x0b, 0x4e, 0xf4, 0x72, 0x6b, 0x10, 0xae, 0x4f, 0x9d, 0xd8, 0x65,
	0xc2, 0x61, 0x7e, 0x91, 0xd0, 0x08, 0x78, 0xc0, 0x33, 0xd3, 0x94, 0x56, 0xe1, 0x3d, 0x2e, 0x97,
	0x6d, 0x14, 0x18, 0x5f, 0x10, 0x6e, 0xbc, 0x91, 0x80, 0x6f, 0x25, 0x9f, 0x9d, 0xe3, 0xbd, 0xa6,
	0x33, 0xe5, 0x25, 0xbe, 0xb7, 0xf1, 0x07, 0x2a, 0x3a, 0x45, 0xcd, 0xfd, 0x96, 0x4a, 0xca, 0xcc,
	0xb6, 0xcc, 0x38, 0x77, 0x99, 0xe8, 0x76, 0xac, 0x7d, 0xb8, 0xfe, 0xe1, 0x2b, 0xcf, 0xf1, 0x31,
	0x30, 0x6f, 0xe8, 0xb0, 0xc8, 0xa7, 0x53, 0x07, 0xdc, 0xa1, 0x9c, 0x9d, 0x3b, 0xa1, 0x34, 0xd4,
	0x5b, 0xa7, 0xa8, 0x59, 0xb7, 0xee, 0xcb, 0x84, 0xae, 0x8c, 0xdb, 0xd2, 0x6b, 0xf3, 0x9e, 0x7c,
	0x28, 0x3a, 0xde, 0x2f, 0x44, 0x72, 0x86, 0x74, 0xa6, 0xd6, 0x4f, 0x51, 0x73, 0xcf, 0xc2, 0x70,
	0x0d, 0x66, 0x7c, 0xdc, 0xc5, 0x47, 0x5b, 0xc4, 0xca, 0x33, 0x5c, 0x97, 0xe9, 0x39, 0xe4, 0x43,
	0xb2, 0x29, 0x2c, 0xa9, 0x9a, 0xd0, 0x92, 0x05, 0xca, 0x27, 0x84, 0x1b, 0x82, 0x26, 0x54, 0x4c,
	0x68, 0x92, 0xe3, 0x39, 0x3e, 0x8d, 0x78, 0x98, 0x41, 0xee, 0xb5, 0xe9, 0x3c, 0xd5, 0x6b, 0xdf,
	0x52, 0xfd, 0x51, 0xc0, 0x60, 0x30, 0xee, 0x13, 0x8f, 0x87, 0xa6, 0xc7, 0x93, 0x90, 0x27, 0xc5,
	0xe3, 0x2c, 0xf1, 0x87, 0x26, 0xcc, 0x62, 0x9a, 0x90, 0x6e, 0x04, 0x3f, 0x53, 0xbd, 0xb2, 0xdb,
	0x55, 0xaa, 0x3f, 0x98, 0xb9, 0xe1, 0xe8, 0x85, 0x51, 0x15, 0x35, 0x2c, 0x65, 0xe5, 0xce, 0x34,
	0xe8, 0x48, 0xe7, 0x26, 0x19, 0x94, 0xc8, 0xea, 0x37, 0x26, 0x83, 0xbf, 0x92, 0x41, 0x25, 0x99,
	0xbd, 0x26, 0xfb, 0x80, 0xf0, 0x11, 0x70, 0x70, 0x47, 0x1b, 0x82, 0xdd, 0xce, 0xb0, 0x9c, 0xff,
	0xc6, 0xda, 0x6e, 0x75, 0x95, 0xea, 0x6a, 0xce, 0xb4, 0x15, 0x32, 0xac, 0xc3, 0xcc, 0xd7, 0xab,
	0xa2, 0x29, 0x8b, 0x74, 0xe7, 0x66, 0x34, 0xf0, 0x67, 0x1a, 0xd8, 0xa6, 0x29, 0x69, 0xd3, 0xc3,
	0x87, 0x74, 0x1a, 0x33, 0xe1, 0x02, 0xe3, 0x91, 0x23, 0xef, 0x57, 0xdd, 0xc9, 0x76, 0xf2, 0x84,
	0xe4, 0xc7, 0x4d, 0x56, 0xc7, 0x4d, 0xec, 0xd5, 0x71, 0xb7, 0xef, 0xce, 0x53, 0x1d, 0x5d, 0x7c,
	0xd7, 0x91, 0x75, 0xb0, 0x2e, 0x96, 0x61, 0xe5, 0x33, 0xc2, 0x8d, 0x58, 0x30, 0x8f, 0xfe, 0x7e,
	0x43, 0xbb, 0xd9, 0x7c, 0x49, 0x31, 0xdf, 0xd3, 0xd2, 0x7c, 0xc5, 0xea, 0x9f, 0x71, 0x11, 0xac,
	0x6c, 0xf9, 0x15, 0x1a, 0x03, 0x1b, 0x25, 0x66, 0xe8, 0xc2, 0x80, 0x9c, 0x0b, 0xea, 0x75, 0xa8,
	0x27, 0x57, 0xa2, 0xaa, 0xf7, 0x7a, 0x25, 0xaa, 0xa2, 0x86, 0x75, 0x94, 0xb9, 0xcb, 0x47, 0xdb,
	0x7e, 0x35, 0x5f, 0x68, 0xe8, 0x72, 0xa1, 0xa1, 0x1f, 0x0b, 0x0d, 0x5d, 0x2c, 0xb5, 0xda, 0xe5,
	0x52, 0xab, 0x7d, 0x5d, 0x6a, 0xb5, 0x77, 0x8f, 0xff, 0x4d, 0x36, 0xcd, 0xbf, 0x66, 0xf2, 0x25,
	0x98, 0x93, 0x56, 0x7f, 0x27, 0x13, 0xe8, 0xc9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0xd6,
	0x40, 0xdd, 0x47, 0x05, 0x00, 0x00,
}

func (m *LimitOrderTrancheKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TrancheKey) > 0 {
		i -= len(m.TrancheKey)
		copy(dAtA[i:], m.TrancheKey)
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(len(m.TrancheKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TickIndexTakerToMaker != 0 {
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(m.TickIndexTakerToMaker))
		i--
		dAtA[i] = 0x10
	}
	if m.TradePairId != nil {
		{
			size, err := m.TradePairId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LimitOrderTranche) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTranche) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTranche) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PriceTakerToMaker.Size()
		i -= size
		if _, err := m.PriceTakerToMaker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.ExpirationTime != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.ExpirationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.TotalTakerDenom.Size()
		i -= size
		if _, err := m.TotalTakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TotalMakerDenom.Size()
		i -= size
		if _, err := m.TotalMakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ReservesTakerDenom.Size()
		i -= size
		if _, err := m.ReservesTakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ReservesMakerDenom.Size()
		i -= size
		if _, err := m.ReservesMakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTranche(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTranche(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradePairId != nil {
		l = m.TradePairId.Size()
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	if m.TickIndexTakerToMaker != 0 {
		n += 1 + sovLimitOrderTranche(uint64(m.TickIndexTakerToMaker))
	}
	l = len(m.TrancheKey)
	if l > 0 {
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	return n
}

func (m *LimitOrderTranche) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	l = m.ReservesMakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.ReservesTakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.TotalMakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.TotalTakerDenom.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	if m.ExpirationTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime)
		n += 1 + l + sovLimitOrderTranche(uint64(l))
	}
	l = m.PriceTakerToMaker.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	return n
}

func sovLimitOrderTranche(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTranche(x uint64) (n int) {
	return sovLimitOrderTranche(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTranche
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
			return fmt.Errorf("proto: LimitOrderTrancheKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradePairId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TradePairId == nil {
				m.TradePairId = &types.TradePairID{}
			}
			if err := m.TradePairId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndexTakerToMaker", wireType)
			}
			m.TickIndexTakerToMaker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndexTakerToMaker |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrancheKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrancheKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTranche(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTranche
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
func (m *LimitOrderTranche) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTranche
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
			return fmt.Errorf("proto: LimitOrderTranche: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTranche: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &LimitOrderTrancheKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesMakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesMakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesTakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesTakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalMakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceTakerToMaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceTakerToMaker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTranche(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTranche
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
func skipLimitOrderTranche(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTranche
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
					return 0, ErrIntOverflowLimitOrderTranche
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
					return 0, ErrIntOverflowLimitOrderTranche
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
				return 0, ErrInvalidLengthLimitOrderTranche
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTranche
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTranche
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTranche        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTranche          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTranche = fmt.Errorf("proto: unexpected end of group")
)
