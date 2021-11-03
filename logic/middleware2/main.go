package main

import (
	proto "RPC/logic/proto"
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {

	listener, err := net.Listen("tcp", ":9001")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterOrchestratorsServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
func (s *server) GetUser(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewMockUserServiceClient(conn)
	name := request.GetName()
	req := &proto.Request{Name: string(name)}
	if response, err := client.GetMockUserData(ctx, req); err == nil {
		return &proto.Response{Name: string(response.Name), Roll: response.Roll, Class: response.Class}, nil
	} else {
		return nil, err
	}

	//
	// stlen := strconv.Itoa(len(name))
	// return &proto.Response{Name: name, Roll: int64(len(name) * 10), Class: stlen}, nil
	//return nil, errors.New("not implemented yet. Vivek will implement me")
}
