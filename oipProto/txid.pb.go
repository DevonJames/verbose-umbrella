// Code generated by protoc-gen-go. DO NOT EDIT.
// source: txid.proto

/*
Package oipProto is a generated protocol buffer package.

It is generated from these files:
	txid.proto
	multipart.proto
	signedMessage.proto
	historian.proto

It has these top-level messages:
	Txid
	MultiPart
	SignedMessage
	HistorianDataPoint
	HistorianPayout
*/
package oipProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Txid struct {
	Raw []byte `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (m *Txid) Reset()                    { *m = Txid{} }
func (m *Txid) String() string            { return proto.CompactTextString(m) }
func (*Txid) ProtoMessage()               {}
func (*Txid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Txid) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterType((*Txid)(nil), "oipProto.Txid")
}

func init() { proto.RegisterFile("txid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xa9, 0xc8, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcf, 0x2c, 0x08, 0x00, 0xb1, 0x94, 0x24,
	0xb8, 0x58, 0x42, 0x2a, 0x32, 0x53, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x12, 0xcb, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x82, 0x40, 0x4c, 0x27, 0xae, 0x28, 0xb8, 0xaa, 0x24, 0x36, 0xb0, 0x36, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0xd1, 0xb0, 0x96, 0x44, 0x00, 0x00, 0x00,
}