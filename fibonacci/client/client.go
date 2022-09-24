package client

import (
	"context"
	"fmt"
	"io"
	"log"

	rpc "github.com/oopjot/grpc-demo/fibonacci/fibonacci"
	"google.golang.org/grpc"
)

type Client struct {
    stub rpc.FibonacciClient
}

func New(addr string, port int) *Client {
    client := Client{}
    connection, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Cannot connect: %v", err)
    }
    client.stub = rpc.NewFibonacciClient(connection)
    return &client
}

func (c *Client) Number(n int64) (int64, error) {
    request := &rpc.FibRequest{Number: n}
    res, err := c.stub.Number(context.Background(), request)
    if err != nil {
        return 0, err
    }
    return res.Result, nil
}

func (c *Client) Sequence(n int64, results chan *rpc.FibResponse) error {
    request := &rpc.FibRequest{Number: n}
    stream, err := c.stub.Sequence(context.Background(), request)
    if err != nil {
        return err
    }

    for {
        res, err := stream.Recv()
        if err != nil {
            return err
        }
        if err == io.EOF {
            return nil
        }

        results <- res
    }
}

