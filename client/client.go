package client

import (
	"errors"
	"io"
)

// Client ...
type Client struct{}

// NewClient ...
func NewClient() *Client {
	return &Client{}
}

func (c *Client) Push(r io.ReadCloser) error {
	return errors.New("not implemented")
}

func (c *Client) Get(w io.WriteCloser) (n int64, err error) {
	return 0, errors.New("not implemented")
}
