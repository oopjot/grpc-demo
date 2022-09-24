package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/oopjot/grpc-demo/fibonacci/fibonacci"
	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedFibonacciServer
}

func (s *server) Number(ctx context.Context, in *pb.FibRequest) (*pb.FibResponse, error) {
    log.Printf("Received Number call: %v", in.GetNumber())
    n := in.GetNumber()
    a, b := 0, 1
    for i := int64(0); i < n; i++ {
        a, b = b, a + b
    }
    return &pb.FibResponse{Result: int64(a), Position: n}, nil
}


func (s *server) Sequence(in *pb.FibRequest, stream pb.Fibonacci_SequenceServer) error {
    log.Printf("Received Sequence call: %v", in.GetNumber())
    n := in.GetNumber()
    a, b := 0, 1
    for i := int64(0); i < n; i++ {
        stream.Send(&pb.FibResponse{Result: int64(a), Position: i + 1})
        a, b = b, a + b
    }
    return nil
}

func main() {
    port := 50001
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterFibonacciServer(s, &server{})
    
    log.Printf("Fibonacci listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
    
}
