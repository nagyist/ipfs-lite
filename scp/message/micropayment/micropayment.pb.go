// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micropayment.proto

package micropayment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type SignedTxn struct {
	Receiver             string    `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount               float64   `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	BillingCycle         int32     `protobuf:"varint,3,opt,name=billing_cycle,json=billingCycle,proto3" json:"billing_cycle,omitempty"`
	TxnHash              string    `protobuf:"bytes,4,opt,name=txn_hash,json=txnHash,proto3" json:"txn_hash,omitempty"`
	Mtdt                 *Metadata `protobuf:"bytes,5,opt,name=mtdt,proto3" json:"mtdt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SignedTxn) Reset()         { *m = SignedTxn{} }
func (m *SignedTxn) String() string { return proto.CompactTextString(m) }
func (*SignedTxn) ProtoMessage()    {}
func (*SignedTxn) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1606e132466fc89, []int{0}
}

func (m *SignedTxn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTxn.Unmarshal(m, b)
}
func (m *SignedTxn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTxn.Marshal(b, m, deterministic)
}
func (m *SignedTxn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTxn.Merge(m, src)
}
func (m *SignedTxn) XXX_Size() int {
	return xxx_messageInfo_SignedTxn.Size(m)
}
func (m *SignedTxn) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTxn.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTxn proto.InternalMessageInfo

func (m *SignedTxn) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *SignedTxn) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SignedTxn) GetBillingCycle() int32 {
	if m != nil {
		return m.BillingCycle
	}
	return 0
}

func (m *SignedTxn) GetTxnHash() string {
	if m != nil {
		return m.TxnHash
	}
	return ""
}

func (m *SignedTxn) GetMtdt() *Metadata {
	if m != nil {
		return m.Mtdt
	}
	return nil
}

type Metadata struct {
	Vals                 map[string]*Metadata_MtdtVal `protobuf:"bytes,1,rep,name=vals,proto3" json:"vals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1606e132466fc89, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetVals() map[string]*Metadata_MtdtVal {
	if m != nil {
		return m.Vals
	}
	return nil
}

type Metadata_MtdtVal struct {
	IncludeSignature bool `protobuf:"varint,1,opt,name=include_signature,json=includeSignature,proto3" json:"include_signature,omitempty"`
	// Types that are valid to be assigned to Val:
	//	*Metadata_MtdtVal_IntVal
	//	*Metadata_MtdtVal_StrVal
	Val                  isMetadata_MtdtVal_Val `protobuf_oneof:"Val"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Metadata_MtdtVal) Reset()         { *m = Metadata_MtdtVal{} }
func (m *Metadata_MtdtVal) String() string { return proto.CompactTextString(m) }
func (*Metadata_MtdtVal) ProtoMessage()    {}
func (*Metadata_MtdtVal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1606e132466fc89, []int{1, 0}
}

func (m *Metadata_MtdtVal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata_MtdtVal.Unmarshal(m, b)
}
func (m *Metadata_MtdtVal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata_MtdtVal.Marshal(b, m, deterministic)
}
func (m *Metadata_MtdtVal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata_MtdtVal.Merge(m, src)
}
func (m *Metadata_MtdtVal) XXX_Size() int {
	return xxx_messageInfo_Metadata_MtdtVal.Size(m)
}
func (m *Metadata_MtdtVal) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata_MtdtVal.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata_MtdtVal proto.InternalMessageInfo

func (m *Metadata_MtdtVal) GetIncludeSignature() bool {
	if m != nil {
		return m.IncludeSignature
	}
	return false
}

type isMetadata_MtdtVal_Val interface {
	isMetadata_MtdtVal_Val()
}

type Metadata_MtdtVal_IntVal struct {
	IntVal int32 `protobuf:"varint,2,opt,name=int_val,json=intVal,proto3,oneof"`
}

type Metadata_MtdtVal_StrVal struct {
	StrVal string `protobuf:"bytes,3,opt,name=str_val,json=strVal,proto3,oneof"`
}

func (*Metadata_MtdtVal_IntVal) isMetadata_MtdtVal_Val() {}

func (*Metadata_MtdtVal_StrVal) isMetadata_MtdtVal_Val() {}

func (m *Metadata_MtdtVal) GetVal() isMetadata_MtdtVal_Val {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Metadata_MtdtVal) GetIntVal() int32 {
	if x, ok := m.GetVal().(*Metadata_MtdtVal_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (m *Metadata_MtdtVal) GetStrVal() string {
	if x, ok := m.GetVal().(*Metadata_MtdtVal_StrVal); ok {
		return x.StrVal
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metadata_MtdtVal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metadata_MtdtVal_IntVal)(nil),
		(*Metadata_MtdtVal_StrVal)(nil),
	}
}

func init() {
	proto.RegisterType((*SignedTxn)(nil), "micropayment.SignedTxn")
	proto.RegisterType((*Metadata)(nil), "micropayment.Metadata")
	proto.RegisterMapType((map[string]*Metadata_MtdtVal)(nil), "micropayment.Metadata.ValsEntry")
	proto.RegisterType((*Metadata_MtdtVal)(nil), "micropayment.Metadata.MtdtVal")
}

func init() {
	proto.RegisterFile("micropayment.proto", fileDescriptor_e1606e132466fc89)
}

var fileDescriptor_e1606e132466fc89 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0xaf, 0xda, 0x30,
	0x14, 0x85, 0x6b, 0x42, 0x48, 0x72, 0xa1, 0x12, 0xf5, 0x80, 0x02, 0x43, 0x15, 0xd1, 0x25, 0x6a,
	0x25, 0x06, 0xca, 0x50, 0x75, 0x6c, 0x55, 0x89, 0x85, 0xc5, 0x54, 0xe9, 0x18, 0x99, 0xc4, 0x02,
	0xab, 0x8e, 0x83, 0xec, 0x9b, 0x88, 0xfc, 0x89, 0xfe, 0x91, 0xfe, 0xc9, 0x2a, 0x6e, 0x40, 0xef,
	0x49, 0xef, 0x6d, 0xbe, 0xdf, 0x3d, 0x3e, 0x3a, 0xc7, 0x06, 0x5a, 0xc9, 0xc2, 0xd4, 0x57, 0xde,
	0x55, 0x42, 0xe3, 0xe6, 0x6a, 0x6a, 0xac, 0xe9, 0xec, 0x29, 0x5b, 0xff, 0x25, 0x10, 0x1d, 0xe5,
	0x59, 0x8b, 0xf2, 0xe7, 0x4d, 0xd3, 0x15, 0x84, 0x46, 0x14, 0x42, 0xb6, 0xc2, 0xc4, 0x24, 0x21,
	0x69, 0xc4, 0x1e, 0x33, 0x5d, 0xc0, 0x84, 0x57, 0x75, 0xa3, 0x31, 0x1e, 0x25, 0x24, 0x25, 0x6c,
	0x98, 0xe8, 0x07, 0x78, 0x7b, 0x92, 0x4a, 0x49, 0x7d, 0xce, 0x8b, 0xae, 0x50, 0x22, 0xf6, 0x12,
	0x92, 0xfa, 0x6c, 0x36, 0xc0, 0xef, 0x3d, 0xa3, 0x4b, 0x08, 0xf1, 0xa6, 0xf3, 0x0b, 0xb7, 0x97,
	0x78, 0xec, 0x8c, 0x03, 0xbc, 0xe9, 0x3d, 0xb7, 0x17, 0xfa, 0x11, 0xc6, 0x15, 0x96, 0x18, 0xfb,
	0x09, 0x49, 0xa7, 0xdb, 0xc5, 0xe6, 0x59, 0xe4, 0x83, 0x40, 0x5e, 0x72, 0xe4, 0xcc, 0x69, 0xd6,
	0x7f, 0x46, 0x10, 0xde, 0x11, 0xdd, 0xc1, 0xb8, 0xe5, 0xca, 0xc6, 0x24, 0xf1, 0xd2, 0xe9, 0x36,
	0x79, 0xf9, 0xe2, 0x26, 0xe3, 0xca, 0xfe, 0xd0, 0x68, 0x3a, 0xe6, 0xd4, 0x2b, 0x0b, 0xc1, 0x01,
	0x4b, 0xcc, 0xb8, 0xa2, 0x9f, 0xe0, 0x9d, 0xd4, 0x85, 0x6a, 0x4a, 0x91, 0x5b, 0x79, 0xd6, 0x1c,
	0x1b, 0x23, 0x5c, 0xed, 0x90, 0xcd, 0x87, 0xc5, 0xf1, 0xce, 0xe9, 0x12, 0x02, 0xa9, 0x31, 0x6f,
	0xb9, 0x72, 0xfd, 0xfd, 0xfd, 0x1b, 0x36, 0x91, 0xda, 0xf9, 0x2c, 0x21, 0xb0, 0x68, 0xdc, 0xaa,
	0xef, 0x1e, 0xf5, 0x2b, 0x8b, 0x26, 0xe3, 0xea, 0x9b, 0x0f, 0x5e, 0xc6, 0xd5, 0xea, 0x17, 0x44,
	0x8f, 0x1c, 0x74, 0x0e, 0xde, 0x6f, 0xd1, 0x0d, 0xef, 0xdb, 0x1f, 0xe9, 0x0e, 0xfc, 0x96, 0xab,
	0x46, 0x38, 0xe7, 0xe9, 0xf6, 0xfd, 0x2b, 0x55, 0x86, 0xdc, 0xec, 0xbf, 0xf8, 0xeb, 0xe8, 0x0b,
	0x39, 0x4d, 0xdc, 0x9f, 0x7e, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x62, 0x85, 0x77, 0xe9,
	0x01, 0x00, 0x00,
}