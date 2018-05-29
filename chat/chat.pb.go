// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

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

type Text struct {
	Txt                  string   `protobuf:"bytes,1,opt,name=txt" json:"txt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_a03c197333cb330d, []int{0}
}
func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (dst *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(dst, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

func init() {
	proto.RegisterType((*Text)(nil), "chat.Text")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatAppClient is the client API for ChatApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatAppClient interface {
	ContChat(ctx context.Context, opts ...grpc.CallOption) (ChatApp_ContChatClient, error)
}

type chatAppClient struct {
	cc *grpc.ClientConn
}

func NewChatAppClient(cc *grpc.ClientConn) ChatAppClient {
	return &chatAppClient{cc}
}

func (c *chatAppClient) ContChat(ctx context.Context, opts ...grpc.CallOption) (ChatApp_ContChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatApp_serviceDesc.Streams[0], "/chat.ChatApp/ContChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatAppContChatClient{stream}
	return x, nil
}

type ChatApp_ContChatClient interface {
	Send(*Text) error
	Recv() (*Text, error)
	grpc.ClientStream
}

type chatAppContChatClient struct {
	grpc.ClientStream
}

func (x *chatAppContChatClient) Send(m *Text) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatAppContChatClient) Recv() (*Text, error) {
	m := new(Text)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatAppServer is the server API for ChatApp service.
type ChatAppServer interface {
	ContChat(ChatApp_ContChatServer) error
}

func RegisterChatAppServer(s *grpc.Server, srv ChatAppServer) {
	s.RegisterService(&_ChatApp_serviceDesc, srv)
}

func _ChatApp_ContChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatAppServer).ContChat(&chatAppContChatServer{stream})
}

type ChatApp_ContChatServer interface {
	Send(*Text) error
	Recv() (*Text, error)
	grpc.ServerStream
}

type chatAppContChatServer struct {
	grpc.ServerStream
}

func (x *chatAppContChatServer) Send(m *Text) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatAppContChatServer) Recv() (*Text, error) {
	m := new(Text)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatApp",
	HandlerType: (*ChatAppServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ContChat",
			Handler:       _ChatApp_ContChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_a03c197333cb330d) }

var fileDescriptor_chat_a03c197333cb330d = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x24, 0xb8, 0x58, 0x42, 0x52,
	0x2b, 0x4a, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x2a, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x40, 0x4c, 0x23, 0x63, 0x2e, 0x76, 0xe7, 0x8c, 0xc4, 0x12, 0xc7, 0x82, 0x02, 0x21, 0x0d, 0x2e,
	0x0e, 0xe7, 0xfc, 0xbc, 0x12, 0x10, 0x57, 0x88, 0x4b, 0x0f, 0x6c, 0x06, 0x48, 0x93, 0x14, 0x12,
	0x5b, 0x89, 0x41, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x6c, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x90, 0x0a, 0xe7, 0xb6, 0x69, 0x00, 0x00, 0x00,
}