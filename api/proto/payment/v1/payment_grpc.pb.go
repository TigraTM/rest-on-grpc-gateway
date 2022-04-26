// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/proto/payment/v1/payment.proto

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

// PaymentAPIClient is the client API for PaymentAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentAPIClient interface {
	// Create Payments. Payment can be created with positive or negative value.
	// Payments are made only in rubles(RUB).
	//
	// ```
	// Example request:
	//    sum: '99.99'
	//    company_name: 'AppStore'
	//    category: 'supermarkets'
	//    user_id: 1
	// ```
	//
	// ```
	// Example response:
	//    empty
	// Specific codes:
	//    * codes.InvalidArgument
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	// Get user balance by user ID.
	// By default the balance is given in rubles, if you want to convert the amount into another currency,
	// pass in query parameter ?currency={another_currency} another currency.
	// The name of the other currency should consist of 3 letters
	//
	// ```
	// Example request:
	//    id: 1
	//    currency: USD
	// ```
	//
	// ```
	// Example response:
	//    id: 1
	//    sum: '99.99'
	//    currency: USD
	// ```
	//
	// Specific codes:
	//    * codes.InvalidArgument
	//    * codes.NotFound
	GetBalanceByUserID(ctx context.Context, in *GetBalanceByUserIDRequest, opts ...grpc.CallOption) (*GetBalanceByUserIDResponse, error)
	// Transferring money between users.
	// Transferring money are made only in rubles(RUB).
	//
	// ```
	// Example request:
	//    sum: '99.99'
	//    sender_id: 1
	//    recipient_id: 2
	//    recipient_name: 'Artem'
	// ```
	//
	// ```
	// Example response:
	//    sum: '99.99'
	//    recipient_id: 2
	//    recipient_name: 'Artem'
	// ```
	//
	// Specific codes:
	//    * codes.InvalidArgument
	//    * codes.NotFound
	TransferBetweenUsers(ctx context.Context, in *TransferBetweenUsersRequest, opts ...grpc.CallOption) (*TransferBetweenUsersResponse, error)
	// Get payments history by user id.
	// Query params:
	// - limit (default = 100)
	// - offset (default = 0)
	// - sort (only 'sum' and 'create_at', default = 'creat_at')
	// - order_by (only 'asc' or 'desc', default = 'desc')
	//
	// ```
	// Example request:
	//    user_id: 1
	//    limit: 5
	//    offset: 0
	//    sort: create_at
	//    order_by: desc
	// ```
	//
	// ```
	// Example response:
	//    id: 1
	//    create_at: // FIXME: fix example time
	//    sum: '99.99'
	//    company_name: 2
	//    category: 'Artem'
	//    total: 1
	// ```
	//
	// Specific codes:
	//    * codes.InvalidArgument
	//    * codes.NotFound
	GetPaymentsHistoryByUserID(ctx context.Context, in *GetPaymentsHistoryByUserIDRequest, opts ...grpc.CallOption) (*GetPaymentsHistoryByUserIDResponse, error)
}

type paymentAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentAPIClient(cc grpc.ClientConnInterface) PaymentAPIClient {
	return &paymentAPIClient{cc}
}

func (c *paymentAPIClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/api.proto.payment.v1.PaymentAPI/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentAPIClient) GetBalanceByUserID(ctx context.Context, in *GetBalanceByUserIDRequest, opts ...grpc.CallOption) (*GetBalanceByUserIDResponse, error) {
	out := new(GetBalanceByUserIDResponse)
	err := c.cc.Invoke(ctx, "/api.proto.payment.v1.PaymentAPI/GetBalanceByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentAPIClient) TransferBetweenUsers(ctx context.Context, in *TransferBetweenUsersRequest, opts ...grpc.CallOption) (*TransferBetweenUsersResponse, error) {
	out := new(TransferBetweenUsersResponse)
	err := c.cc.Invoke(ctx, "/api.proto.payment.v1.PaymentAPI/TransferBetweenUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentAPIClient) GetPaymentsHistoryByUserID(ctx context.Context, in *GetPaymentsHistoryByUserIDRequest, opts ...grpc.CallOption) (*GetPaymentsHistoryByUserIDResponse, error) {
	out := new(GetPaymentsHistoryByUserIDResponse)
	err := c.cc.Invoke(ctx, "/api.proto.payment.v1.PaymentAPI/GetPaymentsHistoryByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentAPIServer is the server API for PaymentAPI service.
// All implementations should embed UnimplementedPaymentAPIServer
// for forward compatibility
type PaymentAPIServer interface {
	// Create Payments. Payment can be created with positive or negative value.
	// Payments are made only in rubles(RUB).
	//
	// ```
	// Example request:
	//    sum: '99.99'
	//    company_name: 'AppStore'
	//    category: 'supermarkets'
	//    user_id: 1
	// ```
	//
	// ```
	// Example response:
	//    empty
	// Specific codes:
	//    * codes.InvalidArgument
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	// Get user balance by user ID.
	// By default the balance is given in rubles, if you want to convert the amount into another currency,
	// pass in query parameter ?currency={another_currency} another currency.
	// The name of the other currency should consist of 3 letters
	//
	// ```
	// Example request:
	//    id: 1
	//    currency: USD
	// ```
	//
	// ```
	// Example response:
	//    id: 1
	//    sum: '99.99'
	//    currency: USD
	// ```
	//
	// Specific codes:
	//    * codes.InvalidArgument
	//    * codes.NotFound
	GetBalanceByUserID(context.Context, *GetBalanceByUserIDRequest) (*GetBalanceByUserIDResponse, error)
	// Transferring money between users.
	// Transferring money are made only in rubles(RUB).
	//
	// ```
	// Example request:
	//    sum: '99.99'
	//    sender_id: 1
	//    recipient_id: 2
	//    recipient_name: 'Artem'
	// ```
	//
	// ```
	// Example response:
	//    sum: '99.99'
	//    recipient_id: 2
	//    recipient_name: 'Artem'
	// ```
	//
	// Specific codes:
	//    * codes.InvalidArgument
	//    * codes.NotFound
	TransferBetweenUsers(context.Context, *TransferBetweenUsersRequest) (*TransferBetweenUsersResponse, error)
	// Get payments history by user id.
	// Query params:
	// - limit (default = 100)
	// - offset (default = 0)
	// - sort (only 'sum' and 'create_at', default = 'creat_at')
	// - order_by (only 'asc' or 'desc', default = 'desc')
	//
	// ```
	// Example request:
	//    user_id: 1
	//    limit: 5
	//    offset: 0
	//    sort: create_at
	//    order_by: desc
	// ```
	//
	// ```
	// Example response:
	//    id: 1
	//    create_at: // FIXME: fix example time
	//    sum: '99.99'
	//    company_name: 2
	//    category: 'Artem'
	//    total: 1
	// ```
	//
	// Specific codes:
	//    * codes.InvalidArgument
	//    * codes.NotFound
	GetPaymentsHistoryByUserID(context.Context, *GetPaymentsHistoryByUserIDRequest) (*GetPaymentsHistoryByUserIDResponse, error)
}

// UnimplementedPaymentAPIServer should be embedded to have forward compatible implementations.
type UnimplementedPaymentAPIServer struct {
}

func (UnimplementedPaymentAPIServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentAPIServer) GetBalanceByUserID(context.Context, *GetBalanceByUserIDRequest) (*GetBalanceByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalanceByUserID not implemented")
}
func (UnimplementedPaymentAPIServer) TransferBetweenUsers(context.Context, *TransferBetweenUsersRequest) (*TransferBetweenUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferBetweenUsers not implemented")
}
func (UnimplementedPaymentAPIServer) GetPaymentsHistoryByUserID(context.Context, *GetPaymentsHistoryByUserIDRequest) (*GetPaymentsHistoryByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentsHistoryByUserID not implemented")
}

// UnsafePaymentAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentAPIServer will
// result in compilation errors.
type UnsafePaymentAPIServer interface {
	mustEmbedUnimplementedPaymentAPIServer()
}

func RegisterPaymentAPIServer(s grpc.ServiceRegistrar, srv PaymentAPIServer) {
	s.RegisterService(&PaymentAPI_ServiceDesc, srv)
}

func _PaymentAPI_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.payment.v1.PaymentAPI/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentAPI_GetBalanceByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).GetBalanceByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.payment.v1.PaymentAPI/GetBalanceByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).GetBalanceByUserID(ctx, req.(*GetBalanceByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentAPI_TransferBetweenUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferBetweenUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).TransferBetweenUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.payment.v1.PaymentAPI/TransferBetweenUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).TransferBetweenUsers(ctx, req.(*TransferBetweenUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentAPI_GetPaymentsHistoryByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentsHistoryByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentAPIServer).GetPaymentsHistoryByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.payment.v1.PaymentAPI/GetPaymentsHistoryByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentAPIServer).GetPaymentsHistoryByUserID(ctx, req.(*GetPaymentsHistoryByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentAPI_ServiceDesc is the grpc.ServiceDesc for PaymentAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.proto.payment.v1.PaymentAPI",
	HandlerType: (*PaymentAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentAPI_CreatePayment_Handler,
		},
		{
			MethodName: "GetBalanceByUserID",
			Handler:    _PaymentAPI_GetBalanceByUserID_Handler,
		},
		{
			MethodName: "TransferBetweenUsers",
			Handler:    _PaymentAPI_TransferBetweenUsers_Handler,
		},
		{
			MethodName: "GetPaymentsHistoryByUserID",
			Handler:    _PaymentAPI_GetPaymentsHistoryByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/payment/v1/payment.proto",
}
