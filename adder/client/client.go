package client

import (
	"context"
	"fmt"

	rpc "github.com/oopjot/grpc-demo/adder/adder"
	"google.golang.org/grpc"
)

type Client struct {
    stub rpc.AdderClient
}

func New(addr string, port int) (*Client, error) {
    client := Client{}
    connection, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithInsecure())
    if err != nil {
        return &Client{}, err
    }
    client.stub = rpc.NewAdderClient(connection)
    return &client, nil
}

func (c *Client) Add(a, b int64) (int64, error) {
    request := &rpc.AddRequest{A: a, B: b}
    res, err := c.stub.Add(context.Background(), request)
    if err != nil {
        return 0, err
    }
    return res.Result, nil
}

