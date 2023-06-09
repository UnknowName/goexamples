package main

import (
    "context"
    "log"
    "net"

    "goexamples/grpc/proto"
    "google.golang.org/grpc"

)

func main() {
    listen, err := net.Listen("tcp","0.0.0.0:7575")
    if err != nil {
        log.Fatal(err)
    }
    s := grpc.NewServer()
    proto.RegisterGreeterServer(s, &HelloServer{})
    log.Println("server listen on", listen.Addr())
    if err := s.Serve(listen); err != nil {
        log.Fatal(err)
    }
}

type HelloServer struct {
    proto.UnimplementedGreeterServer
}

func (hs *HelloServer) SayHello(ctx context.Context, helloRequest *proto.HelloRequest) (*proto.HelloResponse, error) {
    log.Println("received ", helloRequest.String())
    return &proto.HelloResponse{Msg: "hello" + helloRequest.GetName()}, nil
}