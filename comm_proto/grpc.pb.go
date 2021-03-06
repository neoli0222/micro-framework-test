// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comm_proto/grpc.proto

package comm_proto // import "github.com/lpy-neo/micro-framework-test/comm_proto"

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

type GrpcRequest struct {
	Head                 *GrpcRequestHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body                 []byte           `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GrpcRequest) Reset()         { *m = GrpcRequest{} }
func (m *GrpcRequest) String() string { return proto.CompactTextString(m) }
func (*GrpcRequest) ProtoMessage()    {}
func (*GrpcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_35b1cd9ee426bdfb, []int{0}
}
func (m *GrpcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcRequest.Unmarshal(m, b)
}
func (m *GrpcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcRequest.Marshal(b, m, deterministic)
}
func (dst *GrpcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcRequest.Merge(dst, src)
}
func (m *GrpcRequest) XXX_Size() int {
	return xxx_messageInfo_GrpcRequest.Size(m)
}
func (m *GrpcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcRequest proto.InternalMessageInfo

func (m *GrpcRequest) GetHead() *GrpcRequestHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *GrpcRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type GrpcRequestHead struct {
	Cmd                  int32    `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Encoding             int32    `protobuf:"varint,3,opt,name=encoding,proto3" json:"encoding,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcRequestHead) Reset()         { *m = GrpcRequestHead{} }
func (m *GrpcRequestHead) String() string { return proto.CompactTextString(m) }
func (*GrpcRequestHead) ProtoMessage()    {}
func (*GrpcRequestHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_35b1cd9ee426bdfb, []int{1}
}
func (m *GrpcRequestHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcRequestHead.Unmarshal(m, b)
}
func (m *GrpcRequestHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcRequestHead.Marshal(b, m, deterministic)
}
func (dst *GrpcRequestHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcRequestHead.Merge(dst, src)
}
func (m *GrpcRequestHead) XXX_Size() int {
	return xxx_messageInfo_GrpcRequestHead.Size(m)
}
func (m *GrpcRequestHead) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcRequestHead.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcRequestHead proto.InternalMessageInfo

func (m *GrpcRequestHead) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *GrpcRequestHead) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GrpcRequestHead) GetEncoding() int32 {
	if m != nil {
		return m.Encoding
	}
	return 0
}

type GrpcReply struct {
	Cmd                  int32    `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcReply) Reset()         { *m = GrpcReply{} }
func (m *GrpcReply) String() string { return proto.CompactTextString(m) }
func (*GrpcReply) ProtoMessage()    {}
func (*GrpcReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_35b1cd9ee426bdfb, []int{2}
}
func (m *GrpcReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReply.Unmarshal(m, b)
}
func (m *GrpcReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReply.Marshal(b, m, deterministic)
}
func (dst *GrpcReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReply.Merge(dst, src)
}
func (m *GrpcReply) XXX_Size() int {
	return xxx_messageInfo_GrpcReply.Size(m)
}
func (m *GrpcReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReply.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReply proto.InternalMessageInfo

func (m *GrpcReply) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *GrpcReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GrpcRequest)(nil), "comm_proto.GrpcRequest")
	proto.RegisterType((*GrpcRequestHead)(nil), "comm_proto.GrpcRequestHead")
	proto.RegisterType((*GrpcReply)(nil), "comm_proto.GrpcReply")
}

func init() { proto.RegisterFile("comm_proto/grpc.proto", fileDescriptor_grpc_35b1cd9ee426bdfb) }

var fileDescriptor_grpc_35b1cd9ee426bdfb = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0x5b, 0xc5, 0x9d, 0x5d, 0x50, 0x02, 0x42, 0xd1, 0xcb, 0xd2, 0xd3, 0x5e, 0xda,
	0xe0, 0xea, 0x2f, 0xf0, 0xa2, 0x57, 0x73, 0xf4, 0x22, 0x69, 0x32, 0x76, 0x8b, 0x9b, 0x4e, 0xcc,
	0xa6, 0x48, 0xff, 0xbd, 0x64, 0x0a, 0x56, 0x64, 0x4f, 0xf9, 0xf2, 0xf8, 0x78, 0xc3, 0x83, 0x1b,
	0x43, 0xce, 0xbd, 0xfb, 0x40, 0x91, 0x64, 0x1b, 0xbc, 0xa9, 0x19, 0x05, 0xcc, 0x71, 0xa9, 0x60,
	0xf5, 0x1c, 0xbc, 0x51, 0xf8, 0x35, 0xe0, 0x31, 0x0a, 0x09, 0xf9, 0x1e, 0xb5, 0x2d, 0xb2, 0x4d,
	0xb6, 0x5d, 0xed, 0xee, 0xea, 0xd9, 0xac, 0xff, 0x68, 0x2f, 0xa8, 0xad, 0x62, 0x51, 0x08, 0xc8,
	0x1b, 0xb2, 0x63, 0x71, 0xb6, 0xc9, 0xb6, 0x6b, 0xc5, 0x5c, 0xbe, 0xc2, 0xd5, 0x3f, 0x59, 0x5c,
	0xc3, 0xc2, 0x38, 0xcb, 0xd6, 0xb9, 0x4a, 0x98, 0x92, 0xa1, 0x9b, 0x0e, 0x2d, 0x55, 0x42, 0x71,
	0x0b, 0x97, 0xd8, 0x1b, 0xb2, 0x5d, 0xdf, 0x16, 0x0b, 0x16, 0x7f, 0xff, 0xe5, 0x3d, 0x2c, 0xa7,
	0x4a, 0x7f, 0x18, 0x4f, 0x94, 0x09, 0xc8, 0xad, 0x8e, 0x9a, 0xdb, 0xd6, 0x8a, 0xf9, 0xe9, 0xf1,
	0x6d, 0xd7, 0x76, 0x71, 0x3f, 0x34, 0x69, 0x84, 0x3c, 0xf8, 0xb1, 0xea, 0x91, 0xa4, 0xeb, 0x4c,
	0xa0, 0xea, 0x23, 0x68, 0x87, 0xdf, 0x14, 0x3e, 0xab, 0x88, 0xc7, 0x28, 0xe7, 0x95, 0xcd, 0x05,
	0x3f, 0x0f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x4e, 0xc2, 0xe6, 0x3b, 0x01, 0x00, 0x00,
}
