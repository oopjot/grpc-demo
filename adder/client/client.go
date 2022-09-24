package client

import (
	"context"
	"fmt"
	"log"

	rpc "github.com/oopjot/grpc-demo/adder/adder"
	"google.golang.org/grpc"
)

type Client struct {
    stub rpc.AdderClient
}

func New(addr string, port int) *Client {
    client := Client{}
    connection, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Cannot connect: %v", err)
    }
    client.stub = rpc.NewAdderClient(connection)
    return &client
}

func (c *Client) Add(a, b int64) (int64, error) {
    request := &rpc.AddRequest{A: a, B: b}
    res, err := c.stub.Add(context.Background(), request)
    if err != nil {
        return 0, err
    }
    return res.Result, nil
}

