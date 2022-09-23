package main

import (
    "fmt"

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
    
}
