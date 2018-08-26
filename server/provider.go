package server

import (
	"io"
)

type Provider interface {
	Create(Id string) (io.WriteCloser, error)
	Open(Id string) (io.ReadCloser, error)
}
