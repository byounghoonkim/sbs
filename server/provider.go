package server

import "io"

// Provider interface defines Blob Provider's Methods.
type Provider interface {
	Create(ID string) (io.WriteCloser, error)
	Open(ID string) (io.ReadCloser, error)
}
