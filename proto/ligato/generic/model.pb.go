// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/generic/model.proto

package generic

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

// ModelSpec defines a model specification to identify a model.
type ModelSpec struct {
	// Module describes grouping for the model.
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// Version describes version of the model schema.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Type describes name of type described by this model.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Class describes purpose for the model.
	Class                string   `protobuf:"bytes,4,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelSpec) Reset()         { *m = ModelSpec{} }
func (m *ModelSpec) String() string { return proto.CompactTextString(m) }
func (*ModelSpec) ProtoMessage()    {}
func (*ModelSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d85de19425b7f8, []int{0}
}

func (m *ModelSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelSpec.Unmarshal(m, b)
}
func (m *ModelSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelSpec.Marshal(b, m, deterministic)
}
func (m *ModelSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelSpec.Merge(m, src)
}
func (m *ModelSpec) XXX_Size() int {
	return xxx_messageInfo_ModelSpec.Size(m)
}
func (m *ModelSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ModelSpec proto.InternalMessageInfo

func (m *ModelSpec) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ModelSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ModelSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ModelSpec) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

// ModelDetail represents info about model details.
type ModelDetail struct {
	// Spec is a specificaiton the model was registered with.
	Spec *ModelSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// ProtoName is a name of protobuf message representing the model.
	ProtoName            string                `protobuf:"bytes,2,opt,name=proto_name,json=protoName,proto3" json:"proto_name,omitempty"`
	Options              []*ModelDetail_Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ModelDetail) Reset()         { *m = ModelDetail{} }
func (m *ModelDetail) String() string { return proto.CompactTextString(m) }
func (*ModelDetail) ProtoMessage()    {}
func (*ModelDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d85de19425b7f8, []int{1}
}

func (m *ModelDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelDetail.Unmarshal(m, b)
}
func (m *ModelDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelDetail.Marshal(b, m, deterministic)
}
func (m *ModelDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelDetail.Merge(m, src)
}
func (m *ModelDetail) XXX_Size() int {
	return xxx_messageInfo_ModelDetail.Size(m)
}
func (m *ModelDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ModelDetail proto.InternalMessageInfo

func (m *ModelDetail) GetSpec() *ModelSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ModelDetail) GetProtoName() string {
	if m != nil {
		return m.ProtoName
	}
	return ""
}

func (m *ModelDetail) GetOptions() []*ModelDetail_Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type ModelDetail_Option struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelDetail_Option) Reset()         { *m = ModelDetail_Option{} }
func (m *ModelDetail_Option) String() string { return proto.CompactTextString(m) }
func (*ModelDetail_Option) ProtoMessage()    {}
func (*ModelDetail_Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d85de19425b7f8, []int{1, 0}
}

func (m *ModelDetail_Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelDetail_Option.Unmarshal(m, b)
}
func (m *ModelDetail_Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelDetail_Option.Marshal(b, m, deterministic)
}
func (m *ModelDetail_Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelDetail_Option.Merge(m, src)
}
func (m *ModelDetail_Option) XXX_Size() int {
	return xxx_messageInfo_ModelDetail_Option.Size(m)
}
func (m *ModelDetail_Option) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelDetail_Option.DiscardUnknown(m)
}

var xxx_messageInfo_ModelDetail_Option proto.InternalMessageInfo

func (m *ModelDetail_Option) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ModelDetail_Option) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*ModelSpec)(nil), "ligato.generic.ModelSpec")
	proto.RegisterType((*ModelDetail)(nil), "ligato.generic.ModelDetail")
	proto.RegisterType((*ModelDetail_Option)(nil), "ligato.generic.ModelDetail.Option")
}

func init() { proto.RegisterFile("ligato/generic/model.proto", fileDescriptor_13d85de19425b7f8) }

var fileDescriptor_13d85de19425b7f8 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0x65, 0xb7, 0xfd, 0xba, 0x74, 0x0a, 0x1f, 0x12, 0x44, 0xe2, 0x82, 0x50, 0x7a, 0xda, 0xcb,
	0x26, 0xd2, 0xbd, 0x7a, 0x12, 0xaf, 0x2a, 0xd4, 0x9b, 0x17, 0x89, 0xdd, 0xa1, 0x04, 0xd3, 0x26,
	0x34, 0xd9, 0xc2, 0xfe, 0x48, 0xff, 0x93, 0x74, 0xda, 0x0a, 0x0b, 0xde, 0xe6, 0xbd, 0x37, 0xc9,
	0x7b, 0x33, 0x03, 0x5b, 0xa3, 0x1b, 0x15, 0xac, 0x6c, 0xb0, 0xc3, 0x5e, 0xd7, 0xb2, 0xb5, 0x47,
	0x34, 0xc2, 0xf5, 0x36, 0x58, 0xf6, 0x7f, 0xd2, 0xc4, 0xac, 0x15, 0x0d, 0xa4, 0xcf, 0xa3, 0xfc,
	0xe6, 0xb0, 0x66, 0x37, 0x90, 0xb4, 0xf6, 0x78, 0x32, 0xc8, 0x57, 0xf9, 0x6a, 0x97, 0x56, 0x33,
	0x62, 0x1c, 0x36, 0x03, 0xf6, 0x5e, 0xdb, 0x8e, 0xaf, 0x49, 0x58, 0x20, 0x63, 0x10, 0x87, 0xb3,
	0x43, 0x1e, 0x11, 0x4d, 0x35, 0xbb, 0x86, 0x7f, 0xb5, 0x51, 0xde, 0xf3, 0x98, 0xc8, 0x09, 0x14,
	0xdf, 0x2b, 0xc8, 0xc8, 0xe9, 0x09, 0x83, 0xd2, 0x86, 0xed, 0x21, 0xf6, 0x0e, 0x6b, 0x72, 0xca,
	0xca, 0x5b, 0x71, 0x99, 0x4b, 0xfc, 0x86, 0xaa, 0xa8, 0x8d, 0xdd, 0x01, 0xd0, 0x00, 0x1f, 0x9d,
	0x6a, 0x71, 0x4e, 0x91, 0x12, 0xf3, 0xa2, 0x5a, 0x64, 0x0f, 0xb0, 0xb1, 0x2e, 0x68, 0xdb, 0x79,
	0x1e, 0xe5, 0xd1, 0x2e, 0x2b, 0x8b, 0x3f, 0x3f, 0x9c, 0xbc, 0xc5, 0x2b, 0xb5, 0x56, 0xcb, 0x93,
	0x6d, 0x09, 0xc9, 0x44, 0xb1, 0x2b, 0x88, 0xbe, 0xf0, 0x3c, 0x8f, 0x3f, 0x96, 0xe3, 0x4e, 0x06,
	0x65, 0x4e, 0xe8, 0xf9, 0x3a, 0x8f, 0xc6, 0x9d, 0x4c, 0xe8, 0xf1, 0xfe, 0x5d, 0x34, 0x76, 0x31,
	0xd1, 0x56, 0x0e, 0xce, 0xed, 0x55, 0x83, 0x5d, 0x90, 0xc3, 0x41, 0x52, 0x30, 0x79, 0x79, 0x86,
	0xcf, 0x84, 0xd8, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x06, 0xe4, 0xbe, 0x9f, 0x01,
	0x00, 0x00,
}