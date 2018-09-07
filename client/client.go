package client

import (
	"context"
	"fmt"
	"io"

	"github.com/shoebillk/sbs/blob"
	"google.golang.org/grpc"
)

// Client ...
type Client struct {
	blobClient blob.BlobServiceClient
}

// NewClient ...
func NewClient(b blob.BlobServiceClient) (*Client, error) {
	return &Client{b}, nil
}

// NewBlobServiceClient ...
func NewBlobServiceClient(host string, port int) (blob.BlobServiceClient, error) {
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := blob.NewBlobServiceClient(conn)
	return c, nil
}

// Push ...
func (c *Client) Push(ID string, r io.Reader) (*blob.PushStatus, error) {
	var callopts []grpc.CallOption
	pushClient, err := c.blobClient.Push(context.Background(), callopts...)
	if err != nil {
		return nil, err
	}

	chunkContentSize := 0x1000
	chunk := blob.Chunk{
		Id:      ID,
		Content: make([]byte, chunkContentSize),
	}

	for {
		n, err := r.Read(chunk.Content)

		if err == io.EOF {
			// DONE
			err = nil
			break
		} else if err != nil {
			return nil, err
		}

		chunk.Content = chunk.Content[:n]

		err = pushClient.Send(&chunk)
		if err != nil {
			return nil, err
		}
	}

	return pushClient.CloseAndRecv()
}

// Get ...
func (c *Client) Get(ID string, w io.Writer) (n int64, err error) {
	n = 0
	req := blob.GetRequest{Id: ID}

	var callopts []grpc.CallOption
	getClient, err := c.blobClient.Get(context.Background(), &req, callopts...)
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
