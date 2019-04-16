// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RBCType int32

const (
	RBC_VAL   RBCType = 0
	RBC_ECHO  RBCType = 1
	RBC_READY RBCType = 2
)

var RBCType_name = map[int32]string{
	0: "VAL",
	1: "ECHO",
	2: "READY",
}

var RBCType_value = map[string]int32{
	"VAL":   0,
	"ECHO":  1,
	"READY": 2,
}

func (x RBCType) String() string {
	return proto.EnumName(RBCType_name, int32(x))
}

func (RBCType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1, 0}
}

type BBAType int32

const (
	BBA_BVAL BBAType = 0
	BBA_AUX  BBAType = 1
)

var BBAType_name = map[int32]string{
	0: "BVAL",
	1: "AUX",
}

var BBAType_value = map[string]int32{
	"BVAL": 0,
	"AUX":  1,
}

func (x BBAType) String() string {
	return proto.EnumName(BBAType_name, int32(x))
}

func (BBAType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2, 0}
}

type Message struct {
	// The signature of source node who sent this message
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The time when the source sends
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Message_Rbc
	//	*Message_Bba
	Payload              isMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Message) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Rbc struct {
	Rbc *RBC `protobuf:"bytes,3,opt,name=rbc,proto3,oneof"`
}

type Message_Bba struct {
	Bba *BBA `protobuf:"bytes,4,opt,name=bba,proto3,oneof"`
}

func (*Message_Rbc) isMessage_Payload() {}

func (*Message_Bba) isMessage_Payload() {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRbc() *RBC {
	if x, ok := m.GetPayload().(*Message_Rbc); ok {
		return x.Rbc
	}
	return nil
}

func (m *Message) GetBba() *BBA {
	if x, ok := m.GetPayload().(*Message_Bba); ok {
		return x.Bba
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Rbc)(nil),
		(*Message_Bba)(nil),
	}
}

type RBC struct {
	// marshaled data by type
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBC) Reset()         { *m = RBC{} }
func (m *RBC) String() string { return proto.CompactTextString(m) }
func (*RBC) ProtoMessage()    {}
func (*RBC) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *RBC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBC.Unmarshal(m, b)
}
func (m *RBC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBC.Marshal(b, m, deterministic)
}
func (m *RBC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBC.Merge(m, src)
}
func (m *RBC) XXX_Size() int {
	return xxx_messageInfo_RBC.Size(m)
}
func (m *RBC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBC.DiscardUnknown(m)
}

var xxx_messageInfo_RBC proto.InternalMessageInfo

func (m *RBC) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type BBA struct {
	// marshaled data by type
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BBA) Reset()         { *m = BBA{} }
func (m *BBA) String() string { return proto.CompactTextString(m) }
func (*BBA) ProtoMessage()    {}
func (*BBA) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *BBA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BBA.Unmarshal(m, b)
}
func (m *BBA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BBA.Marshal(b, m, deterministic)
}
func (m *BBA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BBA.Merge(m, src)
}
func (m *BBA) XXX_Size() int {
	return xxx_messageInfo_BBA.Size(m)
}
func (m *BBA) XXX_DiscardUnknown() {
	xxx_messageInfo_BBA.DiscardUnknown(m)
}

var xxx_messageInfo_BBA proto.InternalMessageInfo

func (m *BBA) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.RBCType", RBCType_name, RBCType_value)
	proto.RegisterEnum("pb.BBAType", BBAType_name, BBAType_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*RBC)(nil), "pb.RBC")
	proto.RegisterType((*BBA)(nil), "pb.BBA")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0xa5, 0xda, 0xf5, 0x9d, 0x83, 0x92, 0x53, 0x9d, 0x82, 0xa3, 0x78, 0xe8, 0x29,
	0x95, 0x79, 0x11, 0x4f, 0x36, 0xb5, 0xb0, 0x83, 0x22, 0x64, 0x2a, 0x7a, 0x4c, 0x66, 0x2c, 0x85,
	0xd5, 0x86, 0x36, 0x13, 0xf6, 0x85, 0xfc, 0x9c, 0xd2, 0x3f, 0x6b, 0x3d, 0x79, 0xcc, 0xef, 0x79,
	0x9e, 0x1f, 0xe1, 0x85, 0x59, 0xae, 0xaa, 0x4a, 0xa4, 0x8a, 0xea, 0xb2, 0x30, 0x05, 0x19, 0x6b,
	0x39, 0xbf, 0x48, 0x8b, 0x22, 0xdd, 0xaa, 0xb0, 0x21, 0x72, 0xf7, 0x19, 0x9a, 0x2c, 0x57, 0x95,
	0x11, 0xb9, 0x6e, 0x4b, 0xfe, 0x0f, 0x02, 0xfb, 0xb1, 0x9d, 0x91, 0x73, 0x70, 0xaa, 0x2c, 0xfd,
	0x12, 0x66, 0x57, 0x2a, 0x0f, 0x2d, 0x50, 0x70, 0xc2, 0x07, 0x40, 0x6e, 0xc0, 0xe9, 0xc7, 0xde,
	0x78, 0x81, 0x82, 0xe9, 0x72, 0x4e, 0x5b, 0x3d, 0x3d, 0xe8, 0xe9, 0xf3, 0xa1, 0xc1, 0x87, 0x32,
	0x39, 0x03, 0x5c, 0xca, 0x8d, 0x87, 0x9b, 0x8d, 0x4d, 0xb5, 0xa4, 0x9c, 0xc5, 0xab, 0x11, 0xaf,
	0x69, 0x1d, 0x4a, 0x29, 0x3c, 0x6b, 0x08, 0x19, 0x8b, 0xea, 0x50, 0x4a, 0xc1, 0x1c, 0xb0, 0xb5,
	0xd8, 0x6f, 0x0b, 0xf1, 0xe1, 0x27, 0x80, 0x39, 0x8b, 0x89, 0xd7, 0x93, 0xee, 0x87, 0x7d, 0xe1,
	0x12, 0x2c, 0xb3, 0xd7, 0x8a, 0xd8, 0x80, 0x5f, 0xa3, 0x07, 0x77, 0x44, 0x26, 0x60, 0x25, 0xf1,
	0xea, 0xc9, 0x45, 0xc4, 0x81, 0x23, 0x9e, 0x44, 0xf7, 0xef, 0xee, 0xd8, 0xbf, 0x05, 0xcc, 0x58,
	0xf4, 0x8f, 0xe6, 0xb4, 0xd3, 0x4c, 0xc0, 0x62, 0xad, 0xc7, 0x06, 0x1c, 0xbd, 0xbc, 0xb9, 0x68,
	0x79, 0x07, 0xb3, 0xb5, 0x29, 0x95, 0xc8, 0xd7, 0xaa, 0xfc, 0xce, 0x36, 0x8a, 0x84, 0x30, 0xeb,
	0x6e, 0xd7, 0x72, 0x32, 0xad, 0xff, 0xdf, 0xa1, 0xf9, 0xdf, 0x87, 0x3f, 0x0a, 0xd0, 0x15, 0x92,
	0xc7, 0xcd, 0xa1, 0xae, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x1a, 0x31, 0x7b, 0xaa, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	MessageStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_MessageStreamClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) MessageStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/pb.StreamService/MessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceMessageStreamClient{stream}
	return x, nil
}

type StreamService_MessageStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type streamServiceMessageStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceMessageStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceMessageStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	MessageStream(StreamService_MessageStreamServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) MessageStream(srv StreamService_MessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStream not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).MessageStream(&streamServiceMessageStreamServer{stream})
}

type StreamService_MessageStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type streamServiceMessageStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceMessageStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceMessageStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageStream",
			Handler:       _StreamService_MessageStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
