// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chat.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Types
type Message struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Register             bool     `protobuf:"varint,3,opt,name=register" json:"register,omitempty"`
	Group                int32    `protobuf:"varint,4,opt,name=group" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_f90c2c89aa72a752, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetRegister() bool {
	if m != nil {
		return m.Register
	}
	return false
}

func (m *Message) GetGroup() int32 {
	if m != nil {
		return m.Group
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	TransferMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_TransferMessageClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) TransferMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_TransferMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/proto.Chat/TransferMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatTransferMessageClient{stream}
	return x, nil
}

type Chat_TransferMessageClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatTransferMessageClient struct {
	grpc.ClientStream
}

func (x *chatTransferMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatTransferMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	TransferMessage(Chat_TransferMessageServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_TransferMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).TransferMessage(&chatTransferMessageServer{stream})
}

type Chat_TransferMessageServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatTransferMessageServer struct {
	grpc.ServerStream
}

func (x *chatTransferMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatTransferMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferMessage",
			Handler:       _Chat_TransferMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chat.proto",
}

func init() { proto.RegisterFile("proto/chat.proto", fileDescriptor_chat_f90c2c89aa72a752) }

var fileDescriptor_chat_f90c2c89aa72a752 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x48, 0x2c, 0xd1, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x52, 0x3a, 0x17,
	0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x71, 0x6a, 0x5e, 0x4a,
	0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0x27, 0x24, 0xc4, 0xc5, 0x52, 0x92,
	0x5a, 0x51, 0x22, 0xc1, 0x04, 0x16, 0x05, 0xb3, 0x85, 0xa4, 0xb8, 0x38, 0x8a, 0x52, 0xd3, 0x33,
	0x8b, 0x4b, 0x52, 0x8b, 0x24, 0x98, 0x15, 0x18, 0x35, 0x38, 0x82, 0xe0, 0x7c, 0x21, 0x11, 0x2e,
	0xd6, 0xf4, 0xa2, 0xfc, 0xd2, 0x02, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0xc7, 0xc8,
	0x9e, 0x8b, 0xc5, 0x39, 0x23, 0xb1, 0x44, 0xc8, 0x9c, 0x8b, 0x3f, 0xa4, 0x28, 0x31, 0xaf, 0x38,
	0x2d, 0xb5, 0x08, 0x66, 0x31, 0x1f, 0xc4, 0x49, 0x7a, 0x50, 0xbe, 0x14, 0x1a, 0x5f, 0x89, 0x41,
	0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x2c, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x12, 0xf0,
	0x16, 0x64, 0xcb, 0x00, 0x00, 0x00,
}