// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/proto/payment/v1/payment.proto

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_PaymentExternalAPI_CreatePayment_0(ctx context.Context, marshaler runtime.Marshaler, client PaymentExternalAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreatePaymentRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreatePayment(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PaymentExternalAPI_CreatePayment_0(ctx context.Context, marshaler runtime.Marshaler, server PaymentExternalAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreatePaymentRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreatePayment(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_PaymentExternalAPI_GetAccountByAccountNumber_0 = &utilities.DoubleArray{Encoding: map[string]int{"account_number": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_PaymentExternalAPI_GetAccountByAccountNumber_0(ctx context.Context, marshaler runtime.Marshaler, client PaymentExternalAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAccountByUserIDRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["account_number"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "account_number")
	}

	protoReq.AccountNumber, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "account_number", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PaymentExternalAPI_GetAccountByAccountNumber_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetAccountByAccountNumber(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PaymentExternalAPI_GetAccountByAccountNumber_0(ctx context.Context, marshaler runtime.Marshaler, server PaymentExternalAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAccountByUserIDRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["account_number"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "account_number")
	}

	protoReq.AccountNumber, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "account_number", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PaymentExternalAPI_GetAccountByAccountNumber_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetAccountByAccountNumber(ctx, &protoReq)
	return msg, metadata, err

}

func request_PaymentExternalAPI_TransferBetweenUsers_0(ctx context.Context, marshaler runtime.Marshaler, client PaymentExternalAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TransferBetweenUsersRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.TransferBetweenUsers(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PaymentExternalAPI_TransferBetweenUsers_0(ctx context.Context, marshaler runtime.Marshaler, server PaymentExternalAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TransferBetweenUsersRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.TransferBetweenUsers(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0 = &utilities.DoubleArray{Encoding: map[string]int{"account_number": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0(ctx context.Context, marshaler runtime.Marshaler, client PaymentExternalAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetPaymentsHistoryByAccountIDRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["account_number"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "account_number")
	}

	protoReq.AccountNumber, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "account_number", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetPaymentsHistoryByAccountNumber(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0(ctx context.Context, marshaler runtime.Marshaler, server PaymentExternalAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetPaymentsHistoryByAccountIDRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["account_number"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "account_number")
	}

	protoReq.AccountNumber, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "account_number", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetPaymentsHistoryByAccountNumber(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_PaymentExternalAPI_GetAccountsByUserID_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_PaymentExternalAPI_GetAccountsByUserID_0(ctx context.Context, marshaler runtime.Marshaler, client PaymentExternalAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAccountsByUserIDRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PaymentExternalAPI_GetAccountsByUserID_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetAccountsByUserID(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PaymentExternalAPI_GetAccountsByUserID_0(ctx context.Context, marshaler runtime.Marshaler, server PaymentExternalAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAccountsByUserIDRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PaymentExternalAPI_GetAccountsByUserID_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetAccountsByUserID(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterPaymentExternalAPIHandlerServer registers the http handlers for service PaymentExternalAPI to "mux".
// UnaryRPC     :call PaymentExternalAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterPaymentExternalAPIHandlerFromEndpoint instead.
func RegisterPaymentExternalAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PaymentExternalAPIServer) error {

	mux.Handle("POST", pattern_PaymentExternalAPI_CreatePayment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/CreatePayment", runtime.WithHTTPPathPattern("/payment/api/v1/payment"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PaymentExternalAPI_CreatePayment_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_CreatePayment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PaymentExternalAPI_GetAccountByAccountNumber_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/GetAccountByAccountNumber", runtime.WithHTTPPathPattern("/payment/api/v1/payment/accounts/{account_number}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PaymentExternalAPI_GetAccountByAccountNumber_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_GetAccountByAccountNumber_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_PaymentExternalAPI_TransferBetweenUsers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/TransferBetweenUsers", runtime.WithHTTPPathPattern("/payment/api/v1/payment/transfer"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PaymentExternalAPI_TransferBetweenUsers_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_TransferBetweenUsers_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/GetPaymentsHistoryByAccountNumber", runtime.WithHTTPPathPattern("/payment/api/v1/payment/{account_number}/history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PaymentExternalAPI_GetAccountsByUserID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/GetAccountsByUserID", runtime.WithHTTPPathPattern("/payment/api/v1/payment/accounts"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PaymentExternalAPI_GetAccountsByUserID_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_GetAccountsByUserID_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterPaymentExternalAPIHandlerFromEndpoint is same as RegisterPaymentExternalAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPaymentExternalAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPaymentExternalAPIHandler(ctx, mux, conn)
}

// RegisterPaymentExternalAPIHandler registers the http handlers for service PaymentExternalAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPaymentExternalAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPaymentExternalAPIHandlerClient(ctx, mux, NewPaymentExternalAPIClient(conn))
}

// RegisterPaymentExternalAPIHandlerClient registers the http handlers for service PaymentExternalAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PaymentExternalAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PaymentExternalAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PaymentExternalAPIClient" to call the correct interceptors.
func RegisterPaymentExternalAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PaymentExternalAPIClient) error {

	mux.Handle("POST", pattern_PaymentExternalAPI_CreatePayment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/CreatePayment", runtime.WithHTTPPathPattern("/payment/api/v1/payment"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PaymentExternalAPI_CreatePayment_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_CreatePayment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PaymentExternalAPI_GetAccountByAccountNumber_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/GetAccountByAccountNumber", runtime.WithHTTPPathPattern("/payment/api/v1/payment/accounts/{account_number}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PaymentExternalAPI_GetAccountByAccountNumber_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_GetAccountByAccountNumber_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_PaymentExternalAPI_TransferBetweenUsers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/TransferBetweenUsers", runtime.WithHTTPPathPattern("/payment/api/v1/payment/transfer"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PaymentExternalAPI_TransferBetweenUsers_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_TransferBetweenUsers_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/GetPaymentsHistoryByAccountNumber", runtime.WithHTTPPathPattern("/payment/api/v1/payment/{account_number}/history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_PaymentExternalAPI_GetAccountsByUserID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/api.proto.payment.v1.PaymentExternalAPI/GetAccountsByUserID", runtime.WithHTTPPathPattern("/payment/api/v1/payment/accounts"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PaymentExternalAPI_GetAccountsByUserID_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PaymentExternalAPI_GetAccountsByUserID_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PaymentExternalAPI_CreatePayment_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 0}, []string{"payment", "api", "v1"}, ""))

	pattern_PaymentExternalAPI_GetAccountByAccountNumber_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 0, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"payment", "api", "v1", "accounts", "account_number"}, ""))

	pattern_PaymentExternalAPI_TransferBetweenUsers_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 0, 2, 3}, []string{"payment", "api", "v1", "transfer"}, ""))

	pattern_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 0, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"payment", "api", "v1", "account_number", "history"}, ""))

	pattern_PaymentExternalAPI_GetAccountsByUserID_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 0, 2, 3}, []string{"payment", "api", "v1", "accounts"}, ""))
)

var (
	forward_PaymentExternalAPI_CreatePayment_0 = runtime.ForwardResponseMessage

	forward_PaymentExternalAPI_GetAccountByAccountNumber_0 = runtime.ForwardResponseMessage

	forward_PaymentExternalAPI_TransferBetweenUsers_0 = runtime.ForwardResponseMessage

	forward_PaymentExternalAPI_GetPaymentsHistoryByAccountNumber_0 = runtime.ForwardResponseMessage

	forward_PaymentExternalAPI_GetAccountsByUserID_0 = runtime.ForwardResponseMessage
)
