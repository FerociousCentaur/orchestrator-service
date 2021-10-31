package main

import (
	"context"
	"errors"
	"net"
	"strconv"

	proto "RPC/logic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterMockUserServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
func (s *server) GetMockUserData(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	name := request.GetName()
	if len(name) < 6 {
		return nil, errors.New("Error!! Name cannot be less than 6 characters")
	} else {
		stlen := strconv.Itoa(len(name))
		return &proto.Response{Name: name, Roll: int64(len(name) * 10), Class: stlen}, nil
	}

}
