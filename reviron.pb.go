// Code generated by protoc-gen-go.
// source: reviron.proto
// DO NOT EDIT!

/*
Package reviron is a generated protocol buffer package.

It is generated from these files:
	reviron.proto

It has these top-level messages:
	User
	Item
	Rating
	RecommendationRequest
	Recommendation
*/
package reviron

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

type User struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Item struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Rating struct {
	User   *User   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Item   *Item   `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	Rating float32 `protobuf:"fixed32,3,opt,name=rating" json:"rating,omitempty"`
}

func (m *Rating) Reset()                    { *m = Rating{} }
func (m *Rating) String() string            { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()               {}
func (*Rating) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Rating) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Rating) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Rating) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type RecommendationRequest struct {
	User     *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	ItemType string `protobuf:"bytes,2,opt,name=itemType" json:"itemType,omitempty"`
}

func (m *RecommendationRequest) Reset()                    { *m = RecommendationRequest{} }
func (m *RecommendationRequest) String() string            { return proto.CompactTextString(m) }
func (*RecommendationRequest) ProtoMessage()               {}
func (*RecommendationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RecommendationRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RecommendationRequest) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

type Recommendation struct {
	Item   *Item   `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
	Rating float32 `protobuf:"fixed32,2,opt,name=rating" json:"rating,omitempty"`
}

func (m *Recommendation) Reset()                    { *m = Recommendation{} }
func (m *Recommendation) String() string            { return proto.CompactTextString(m) }
func (*Recommendation) ProtoMessage()               {}
func (*Recommendation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Recommendation) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Recommendation) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "reviron.User")
	proto.RegisterType((*Item)(nil), "reviron.Item")
	proto.RegisterType((*Rating)(nil), "reviron.Rating")
	proto.RegisterType((*RecommendationRequest)(nil), "reviron.RecommendationRequest")
	proto.RegisterType((*Recommendation)(nil), "reviron.Recommendation")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Reviron service

type RevironClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	AddRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Rating, error)
	MakeRecommendation(ctx context.Context, in *RecommendationRequest, opts ...grpc.CallOption) (Reviron_MakeRecommendationClient, error)
}

type revironClient struct {
	cc *grpc.ClientConn
}

func NewRevironClient(cc *grpc.ClientConn) RevironClient {
	return &revironClient{cc}
}

func (c *revironClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/reviron.Reviron/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revironClient) AddItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := grpc.Invoke(ctx, "/reviron.Reviron/AddItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revironClient) AddRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Rating, error) {
	out := new(Rating)
	err := grpc.Invoke(ctx, "/reviron.Reviron/AddRating", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revironClient) MakeRecommendation(ctx context.Context, in *RecommendationRequest, opts ...grpc.CallOption) (Reviron_MakeRecommendationClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Reviron_serviceDesc.Streams[0], c.cc, "/reviron.Reviron/MakeRecommendation", opts...)
	if err != nil {
		return nil, err
	}
	x := &revironMakeRecommendationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Reviron_MakeRecommendationClient interface {
	Recv() (*Recommendation, error)
	grpc.ClientStream
}

type revironMakeRecommendationClient struct {
	grpc.ClientStream
}

func (x *revironMakeRecommendationClient) Recv() (*Recommendation, error) {
	m := new(Recommendation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Reviron service

type RevironServer interface {
	AddUser(context.Context, *User) (*User, error)
	AddItem(context.Context, *Item) (*Item, error)
	AddRating(context.Context, *Rating) (*Rating, error)
	MakeRecommendation(*RecommendationRequest, Reviron_MakeRecommendationServer) error
}

func RegisterRevironServer(s *grpc.Server, srv RevironServer) {
	s.RegisterService(&_Reviron_serviceDesc, srv)
}

func _Reviron_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevironServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviron.Reviron/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevironServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviron_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevironServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviron.Reviron/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevironServer).AddItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviron_AddRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevironServer).AddRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviron.Reviron/AddRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevironServer).AddRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviron_MakeRecommendation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecommendationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RevironServer).MakeRecommendation(m, &revironMakeRecommendationServer{stream})
}

type Reviron_MakeRecommendationServer interface {
	Send(*Recommendation) error
	grpc.ServerStream
}

type revironMakeRecommendationServer struct {
	grpc.ServerStream
}

func (x *revironMakeRecommendationServer) Send(m *Recommendation) error {
	return x.ServerStream.SendMsg(m)
}

var _Reviron_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reviron.Reviron",
	HandlerType: (*RevironServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Reviron_AddUser_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _Reviron_AddItem_Handler,
		},
		{
			MethodName: "AddRating",
			Handler:    _Reviron_AddRating_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeRecommendation",
			Handler:       _Reviron_MakeRecommendation_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reviron.proto",
}

func init() { proto.RegisterFile("reviron.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0xcc, 0xc6, 0x90, 0xd8, 0x27, 0xad, 0xf0, 0xc0, 0x1a, 0x72, 0x90, 0xba, 0xa7, 0xea, 0xa1,
	0x4a, 0xfd, 0x82, 0x1e, 0x45, 0x3c, 0xb8, 0xa8, 0xf7, 0xe8, 0x3e, 0x65, 0x91, 0x24, 0x75, 0xb3,
	0x15, 0xfc, 0x5e, 0x7f, 0x44, 0xf2, 0x1a, 0xd3, 0x6e, 0xb0, 0xd0, 0x5b, 0xe6, 0xcd, 0x30, 0x33,
	0x19, 0x16, 0x86, 0x96, 0xbe, 0x8c, 0xad, 0xca, 0xd9, 0xd2, 0x56, 0xae, 0xc2, 0xa4, 0x85, 0x72,
	0x0c, 0xd1, 0x53, 0x4d, 0x16, 0x47, 0x10, 0x1a, 0x9d, 0x8a, 0x89, 0x98, 0x0e, 0x54, 0x68, 0xb4,
	0xbc, 0x84, 0xe8, 0xd6, 0x51, 0xd1, 0xbf, 0x23, 0x42, 0xe4, 0xbe, 0x97, 0x94, 0x86, 0x7c, 0xe1,
	0x6f, 0xf9, 0x06, 0xb1, 0xca, 0x9d, 0x29, 0xdf, 0xf1, 0x1c, 0xa2, 0x55, 0x4d, 0x96, 0xf5, 0x47,
	0xf3, 0xe1, 0xec, 0x2f, 0xb4, 0x89, 0x50, 0x4c, 0x35, 0x12, 0xe3, 0xa8, 0x60, 0x83, 0x6d, 0x49,
	0x93, 0xa6, 0x98, 0xc2, 0x31, 0xc4, 0x96, 0xfd, 0xd2, 0x83, 0x89, 0x98, 0x86, 0xaa, 0x45, 0xf2,
	0x19, 0x4e, 0x14, 0xbd, 0x56, 0x45, 0x41, 0xa5, 0xce, 0x9d, 0xa9, 0x4a, 0x45, 0x9f, 0x2b, 0xaa,
	0xdd, 0x3e, 0xb1, 0x19, 0x1c, 0x36, 0xde, 0x8f, 0x9b, 0xee, 0x1d, 0x96, 0x77, 0x30, 0xf2, 0x7d,
	0xbb, 0x92, 0x62, 0x9f, 0x92, 0xe1, 0x76, 0xc9, 0xf9, 0x8f, 0x80, 0x44, 0xad, 0xe5, 0x78, 0x01,
	0xc9, 0x42, 0x6b, 0xde, 0xd7, 0x2f, 0x95, 0xf9, 0x50, 0x06, 0xad, 0x94, 0x27, 0xf7, 0xe3, 0x32,
	0x1f, 0xca, 0x00, 0xaf, 0x60, 0xb0, 0xd0, 0xba, 0x5d, 0xfc, 0xb8, 0x63, 0xd7, 0x87, 0xac, 0x7f,
	0x90, 0x01, 0x3e, 0x00, 0xde, 0xe7, 0x1f, 0xd4, 0xfb, 0xc7, 0xb3, 0x8d, 0xf0, 0xbf, 0x51, 0xb3,
	0xd3, 0x1d, 0xbc, 0x0c, 0xae, 0xc5, 0x4b, 0xcc, 0xcf, 0xe8, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0x1f, 0x6b, 0xcb, 0x57, 0x02, 0x00, 0x00,
}
