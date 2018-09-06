package client

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/shoebillk/sbs/blob"
	"google.golang.org/grpc"
)

// Client ...
type Client struct {
	host string
	port int
}

// NewClient ...
func NewClient(host string, port int) *Client {
	return &Client{host, port}
}

func (c *Client) getClient(ID string) (blob.BlobService_GetClient, error) {
	return nil, errors.New("not implemented")
}

// Push ...
func (c *Client) Push(ID string, r io.ReadCloser) error {
	return errors.New("not implemented")
}

// Get ...
func (c *Client) Get(ID string, w io.Writer) (n int64, err error) {
	n = 0
	target := fmt.Sprintf("%s:%d", c.host, c.port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return n, err
	}
	defer conn.Close()

	client := blob.NewBlobServiceClient(conn)

	req := blob.GetRequest{Id: ID}

	var callopts []grpc.CallOption
	getClient, err := client.Get(context.Background(), &req, callopts...)
	if err != nil {
		return n, err
	}

	for {
		chunk, err := getClient.Recv()

		if err == io.EOF {
			// DONE
			err = nil
			break
		} else if err != nil {
			return n, err
		}

		n += int64(len(chunk.Content))
		_, err = w.Write(chunk.Content)
		if err != nil {
			return n, err
		}
	}

	return n, err
}
