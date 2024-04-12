// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/notification/pb/notification.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	SendCommentedNotification(ctx context.Context, in *CommentedNotification, opts ...grpc.CallOption) (*NotificationResponse, error)
	SendFollowedNotification(ctx context.Context, in *FollowedNotification, opts ...grpc.CallOption) (*NotificationResponse, error)
	SendKafkaNotification(ctx context.Context, in *KafkaNotification, opts ...grpc.CallOption) (*NotificationResponse, error)
	SendLikeNotification(ctx context.Context, in *LikeNotification, opts ...grpc.CallOption) (*NotificationResponse, error)
	ConsumeKafkaMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NotificationService_ConsumeKafkaMessagesClient, error)
	ConsumeKafkaCommentMessages(ctx context.Context, in *ConsumeKafkaCommentMessagesRequest, opts ...grpc.CallOption) (*ConsumeKafkaCommentMessagesResponse, error)
	ConsumeKafkaLikeMessages(ctx context.Context, in *ConsumeKafkaLikeMessagesRequest, opts ...grpc.CallOption) (*ConsumeKafkaLikeMessagesResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) SendCommentedNotification(ctx context.Context, in *CommentedNotification, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/SendCommentedNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendFollowedNotification(ctx context.Context, in *FollowedNotification, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/SendFollowedNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendKafkaNotification(ctx context.Context, in *KafkaNotification, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/SendKafkaNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendLikeNotification(ctx context.Context, in *LikeNotification, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/SendLikeNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ConsumeKafkaMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NotificationService_ConsumeKafkaMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[0], "/notification.NotificationService/ConsumeKafkaMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceConsumeKafkaMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_ConsumeKafkaMessagesClient interface {
	Recv() (*NotificationMessage, error)
	grpc.ClientStream
}

type notificationServiceConsumeKafkaMessagesClient struct {
	grpc.ClientStream
}

func (x *notificationServiceConsumeKafkaMessagesClient) Recv() (*NotificationMessage, error) {
	m := new(NotificationMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) ConsumeKafkaCommentMessages(ctx context.Context, in *ConsumeKafkaCommentMessagesRequest, opts ...grpc.CallOption) (*ConsumeKafkaCommentMessagesResponse, error) {
	out := new(ConsumeKafkaCommentMessagesResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/ConsumeKafkaCommentMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ConsumeKafkaLikeMessages(ctx context.Context, in *ConsumeKafkaLikeMessagesRequest, opts ...grpc.CallOption) (*ConsumeKafkaLikeMessagesResponse, error) {
	out := new(ConsumeKafkaLikeMessagesResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/ConsumeKafkaLikeMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	SendCommentedNotification(context.Context, *CommentedNotification) (*NotificationResponse, error)
	SendFollowedNotification(context.Context, *FollowedNotification) (*NotificationResponse, error)
	SendKafkaNotification(context.Context, *KafkaNotification) (*NotificationResponse, error)
	SendLikeNotification(context.Context, *LikeNotification) (*NotificationResponse, error)
	ConsumeKafkaMessages(*Empty, NotificationService_ConsumeKafkaMessagesServer) error
	ConsumeKafkaCommentMessages(context.Context, *ConsumeKafkaCommentMessagesRequest) (*ConsumeKafkaCommentMessagesResponse, error)
	ConsumeKafkaLikeMessages(context.Context, *ConsumeKafkaLikeMessagesRequest) (*ConsumeKafkaLikeMessagesResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) SendCommentedNotification(context.Context, *CommentedNotification) (*NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommentedNotification not implemented")
}
func (UnimplementedNotificationServiceServer) SendFollowedNotification(context.Context, *FollowedNotification) (*NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFollowedNotification not implemented")
}
func (UnimplementedNotificationServiceServer) SendKafkaNotification(context.Context, *KafkaNotification) (*NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKafkaNotification not implemented")
}
func (UnimplementedNotificationServiceServer) SendLikeNotification(context.Context, *LikeNotification) (*NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLikeNotification not implemented")
}
func (UnimplementedNotificationServiceServer) ConsumeKafkaMessages(*Empty, NotificationService_ConsumeKafkaMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeKafkaMessages not implemented")
}
func (UnimplementedNotificationServiceServer) ConsumeKafkaCommentMessages(context.Context, *ConsumeKafkaCommentMessagesRequest) (*ConsumeKafkaCommentMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeKafkaCommentMessages not implemented")
}
func (UnimplementedNotificationServiceServer) ConsumeKafkaLikeMessages(context.Context, *ConsumeKafkaLikeMessagesRequest) (*ConsumeKafkaLikeMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeKafkaLikeMessages not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_SendCommentedNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentedNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendCommentedNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/SendCommentedNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendCommentedNotification(ctx, req.(*CommentedNotification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendFollowedNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowedNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendFollowedNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/SendFollowedNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendFollowedNotification(ctx, req.(*FollowedNotification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendKafkaNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KafkaNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendKafkaNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/SendKafkaNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendKafkaNotification(ctx, req.(*KafkaNotification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendLikeNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendLikeNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/SendLikeNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendLikeNotification(ctx, req.(*LikeNotification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ConsumeKafkaMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).ConsumeKafkaMessages(m, &notificationServiceConsumeKafkaMessagesServer{stream})
}

type NotificationService_ConsumeKafkaMessagesServer interface {
	Send(*NotificationMessage) error
	grpc.ServerStream
}

type notificationServiceConsumeKafkaMessagesServer struct {
	grpc.ServerStream
}

func (x *notificationServiceConsumeKafkaMessagesServer) Send(m *NotificationMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_ConsumeKafkaCommentMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeKafkaCommentMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ConsumeKafkaCommentMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/ConsumeKafkaCommentMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ConsumeKafkaCommentMessages(ctx, req.(*ConsumeKafkaCommentMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ConsumeKafkaLikeMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeKafkaLikeMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ConsumeKafkaLikeMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/ConsumeKafkaLikeMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ConsumeKafkaLikeMessages(ctx, req.(*ConsumeKafkaLikeMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommentedNotification",
			Handler:    _NotificationService_SendCommentedNotification_Handler,
		},
		{
			MethodName: "SendFollowedNotification",
			Handler:    _NotificationService_SendFollowedNotification_Handler,
		},
		{
			MethodName: "SendKafkaNotification",
			Handler:    _NotificationService_SendKafkaNotification_Handler,
		},
		{
			MethodName: "SendLikeNotification",
			Handler:    _NotificationService_SendLikeNotification_Handler,
		},
		{
			MethodName: "ConsumeKafkaCommentMessages",
			Handler:    _NotificationService_ConsumeKafkaCommentMessages_Handler,
		},
		{
			MethodName: "ConsumeKafkaLikeMessages",
			Handler:    _NotificationService_ConsumeKafkaLikeMessages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeKafkaMessages",
			Handler:       _NotificationService_ConsumeKafkaMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/notification/pb/notification.proto",
}
