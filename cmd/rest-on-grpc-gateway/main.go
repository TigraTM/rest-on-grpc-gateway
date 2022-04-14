package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

type server struct {
	userpb.UnimplementedUserAPIServer
}

func (*server) CreateUser(_ context.Context, _ *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return &userpb.CreateUserResponse{Id: "hello Andrey"}, nil
}

func main() {
	// nolint:gosec // copy in tutorial.
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("net.Listen")
	}

	s := grpc.NewServer()

	userpb.RegisterUserAPIServer(s, &server{})

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed connect to dial server: ", err)
	}

	gwmx := runtime.NewServeMux()

	err = userpb.RegisterUserAPIHandler(context.Background(), gwmx, conn)
	if err != nil {
		log.Fatal("failed register user api handler: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmx,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
