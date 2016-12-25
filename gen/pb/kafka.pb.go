// Code generated by protoc-gen-gogo.
// source: kafka.proto
// DO NOT EDIT!

/*
Package kafkapb is a generated protocol buffer package.

It is generated from these files:
	kafka.proto

It has these top-level messages:
	ConsumerRequest
	ConsumerResponse
	ProducerRequest
	ProducerResponse
	ConsumerStreamRequest
	ConsumerStreamResponse
*/
package kafkapb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConsumerRequest struct {
	Topics   []string `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
	ClientId string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *ConsumerRequest) Reset()                    { *m = ConsumerRequest{} }
func (m *ConsumerRequest) String() string            { return proto.CompactTextString(m) }
func (*ConsumerRequest) ProtoMessage()               {}
func (*ConsumerRequest) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{0} }

func (m *ConsumerRequest) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *ConsumerRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type ConsumerResponse struct {
	Value  string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (m *ConsumerResponse) Reset()                    { *m = ConsumerResponse{} }
func (m *ConsumerResponse) String() string            { return proto.CompactTextString(m) }
func (*ConsumerResponse) ProtoMessage()               {}
func (*ConsumerResponse) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{1} }

func (m *ConsumerResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ConsumerResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ProducerRequest struct {
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *ProducerRequest) Reset()                    { *m = ProducerRequest{} }
func (m *ProducerRequest) String() string            { return proto.CompactTextString(m) }
func (*ProducerRequest) ProtoMessage()               {}
func (*ProducerRequest) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{2} }

func (m *ProducerRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ProducerRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ProducerRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ProducerResponse struct {
	Partition int32  `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset    int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	ErrMsg    string `protobuf:"bytes,3,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (m *ProducerResponse) Reset()                    { *m = ProducerResponse{} }
func (m *ProducerResponse) String() string            { return proto.CompactTextString(m) }
func (*ProducerResponse) ProtoMessage()               {}
func (*ProducerResponse) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{3} }

func (m *ProducerResponse) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *ProducerResponse) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ProducerResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ConsumerStreamRequest struct {
	Topics   []string `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
	ClientId string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *ConsumerStreamRequest) Reset()                    { *m = ConsumerStreamRequest{} }
func (m *ConsumerStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*ConsumerStreamRequest) ProtoMessage()               {}
func (*ConsumerStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{4} }

func (m *ConsumerStreamRequest) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *ConsumerStreamRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type ConsumerStreamResponse struct {
	Value  string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (m *ConsumerStreamResponse) Reset()                    { *m = ConsumerStreamResponse{} }
func (m *ConsumerStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*ConsumerStreamResponse) ProtoMessage()               {}
func (*ConsumerStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{5} }

func (m *ConsumerStreamResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ConsumerStreamResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*ConsumerRequest)(nil), "kafka.ConsumerRequest")
	proto.RegisterType((*ConsumerResponse)(nil), "kafka.ConsumerResponse")
	proto.RegisterType((*ProducerRequest)(nil), "kafka.ProducerRequest")
	proto.RegisterType((*ProducerResponse)(nil), "kafka.ProducerResponse")
	proto.RegisterType((*ConsumerStreamRequest)(nil), "kafka.ConsumerStreamRequest")
	proto.RegisterType((*ConsumerStreamResponse)(nil), "kafka.ConsumerStreamResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KafkaService service

type KafkaServiceClient interface {
	// unaries
	Consumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error)
	Producer(ctx context.Context, in *ProducerRequest, opts ...grpc.CallOption) (*ProducerResponse, error)
	// streams
	ConsumerStream(ctx context.Context, in *ConsumerStreamRequest, opts ...grpc.CallOption) (KafkaService_ConsumerStreamClient, error)
}

type kafkaServiceClient struct {
	cc *grpc.ClientConn
}

func NewKafkaServiceClient(cc *grpc.ClientConn) KafkaServiceClient {
	return &kafkaServiceClient{cc}
}

func (c *kafkaServiceClient) Consumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*ConsumerResponse, error) {
	out := new(ConsumerResponse)
	err := grpc.Invoke(ctx, "/kafka.KafkaService/Consumer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaServiceClient) Producer(ctx context.Context, in *ProducerRequest, opts ...grpc.CallOption) (*ProducerResponse, error) {
	out := new(ProducerResponse)
	err := grpc.Invoke(ctx, "/kafka.KafkaService/Producer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaServiceClient) ConsumerStream(ctx context.Context, in *ConsumerStreamRequest, opts ...grpc.CallOption) (KafkaService_ConsumerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_KafkaService_serviceDesc.Streams[0], c.cc, "/kafka.KafkaService/ConsumerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &kafkaServiceConsumerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KafkaService_ConsumerStreamClient interface {
	Recv() (*ConsumerStreamResponse, error)
	grpc.ClientStream
}

type kafkaServiceConsumerStreamClient struct {
	grpc.ClientStream
}

func (x *kafkaServiceConsumerStreamClient) Recv() (*ConsumerStreamResponse, error) {
	m := new(ConsumerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for KafkaService service

type KafkaServiceServer interface {
	// unaries
	Consumer(context.Context, *ConsumerRequest) (*ConsumerResponse, error)
	Producer(context.Context, *ProducerRequest) (*ProducerResponse, error)
	// streams
	ConsumerStream(*ConsumerStreamRequest, KafkaService_ConsumerStreamServer) error
}

func RegisterKafkaServiceServer(s *grpc.Server, srv KafkaServiceServer) {
	s.RegisterService(&_KafkaService_serviceDesc, srv)
}

func _KafkaService_Consumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).Consumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafka.KafkaService/Consumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).Consumer(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaService_Producer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProducerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).Producer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafka.KafkaService/Producer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).Producer(ctx, req.(*ProducerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaService_ConsumerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KafkaServiceServer).ConsumerStream(m, &kafkaServiceConsumerStreamServer{stream})
}

type KafkaService_ConsumerStreamServer interface {
	Send(*ConsumerStreamResponse) error
	grpc.ServerStream
}

type kafkaServiceConsumerStreamServer struct {
	grpc.ServerStream
}

func (x *kafkaServiceConsumerStreamServer) Send(m *ConsumerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _KafkaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kafka.KafkaService",
	HandlerType: (*KafkaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Consumer",
			Handler:    _KafkaService_Consumer_Handler,
		},
		{
			MethodName: "Producer",
			Handler:    _KafkaService_Producer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumerStream",
			Handler:       _KafkaService_ConsumerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kafka.proto",
}

func init() { proto.RegisterFile("kafka.proto", fileDescriptorKafka) }

var fileDescriptorKafka = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x86, 0x29, 0x4d, 0x81, 0xce, 0xf7, 0x45, 0xc8, 0x06, 0xa1, 0x41, 0x4c, 0xc8, 0x9e, 0x38,
	0x11, 0xa3, 0x67, 0x0f, 0x6a, 0xa2, 0x31, 0x6a, 0xd4, 0x72, 0xf3, 0x42, 0x96, 0x32, 0x90, 0x0d,
	0xd0, 0xad, 0xbb, 0x5b, 0x12, 0x7f, 0xad, 0x7f, 0xc5, 0x74, 0xdb, 0x52, 0x5a, 0xe3, 0x45, 0x6f,
	0x7d, 0xa7, 0x3b, 0xcf, 0xcc, 0xbc, 0x33, 0xf0, 0x6f, 0xcd, 0x96, 0x6b, 0x36, 0x89, 0xa4, 0xd0,
	0x82, 0x38, 0x46, 0xd0, 0x5b, 0x68, 0xdf, 0x88, 0x50, 0xc5, 0x5b, 0x94, 0x3e, 0xbe, 0xc7, 0xa8,
	0x34, 0xe9, 0x41, 0x43, 0x8b, 0x88, 0x07, 0xca, 0xb3, 0x46, 0xf6, 0xd8, 0xf5, 0x33, 0x45, 0x4e,
	0xc0, 0x0d, 0x36, 0x1c, 0x43, 0x3d, 0xe3, 0x0b, 0xaf, 0x3e, 0xb2, 0xc6, 0xae, 0xdf, 0x4a, 0x03,
	0xf7, 0x0b, 0x7a, 0x05, 0x9d, 0x82, 0xa3, 0x22, 0x11, 0x2a, 0x24, 0x5d, 0x70, 0x76, 0x6c, 0x13,
	0xa3, 0x67, 0x99, 0xc7, 0xa9, 0x20, 0x7d, 0x68, 0xa2, 0x94, 0xb3, 0xad, 0x5a, 0x65, 0x90, 0x06,
	0x4a, 0xf9, 0xa4, 0x56, 0xf4, 0x19, 0xda, 0x2f, 0x52, 0x2c, 0xe2, 0xa0, 0x68, 0xa5, 0x0b, 0x8e,
	0x29, 0x9e, 0x13, 0x8c, 0x28, 0xb8, 0xf5, 0x43, 0x6e, 0x07, 0xec, 0x35, 0x7e, 0x78, 0xb6, 0x89,
	0x25, 0x9f, 0x94, 0x41, 0xa7, 0x00, 0x66, 0x3d, 0x0d, 0xc1, 0x8d, 0x98, 0xd4, 0x5c, 0x73, 0x11,
	0x1a, 0xaa, 0xe3, 0x17, 0x81, 0x64, 0x74, 0xb1, 0x5c, 0x2a, 0xd4, 0x06, 0x6d, 0xfb, 0x99, 0x3a,
	0xec, 0xd9, 0x2e, 0xf5, 0xfc, 0x08, 0xc7, 0xf9, 0xd8, 0x53, 0x2d, 0x91, 0x6d, 0xff, 0x64, 0xe2,
	0x1d, 0xf4, 0xaa, 0xb4, 0x5f, 0x59, 0x79, 0xfe, 0x69, 0xc1, 0xff, 0x87, 0x64, 0xbf, 0x53, 0x94,
	0x3b, 0x1e, 0x20, 0xb9, 0x84, 0x56, 0x4e, 0x26, 0xbd, 0x49, 0x7a, 0x07, 0x95, 0xbd, 0x0f, 0xfa,
	0xdf, 0xe2, 0x69, 0x71, 0x5a, 0x4b, 0xd2, 0x73, 0x27, 0xf7, 0xe9, 0x95, 0x5d, 0xed, 0xd3, 0xab,
	0x96, 0xd3, 0x1a, 0x79, 0x85, 0xa3, 0xf2, 0x5c, 0x64, 0x58, 0xa9, 0x55, 0x32, 0x6f, 0x70, 0xfa,
	0xc3, 0xdf, 0x1c, 0x78, 0x66, 0x5d, 0xbb, 0x6f, 0x4d, 0xf3, 0x26, 0x9a, 0xcf, 0x1b, 0xe6, 0xa0,
	0x2f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xce, 0x51, 0x2c, 0xb5, 0xdf, 0x02, 0x00, 0x00,
}
