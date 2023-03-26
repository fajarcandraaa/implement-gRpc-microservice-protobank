package client

import "github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"

type UserClient struct {
	User pb.UserServiceClient
}

type BookClient struct {
	Book pb.BookServiceClient
}
