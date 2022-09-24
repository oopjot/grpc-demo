package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/oopjot/grpc-demo/adder/adder"
)

type server struct {
    pb.UnimplementedAdderServer

}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
    log.Printf("Received: a: %v b: %v", in.GetA(), in.GetB())

    result := in.GetA() + in.GetB()


    return &pb.AddResponse{Result: result}, nil
}

func main() {
    port := 50000
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterAdderServer(s, &server{})
    log.Printf("Adder listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failes to serve: %v", err)
    }
}
