package server

import (
	"errors"
	"io"
)

// Mgo ...
type Mgo struct {
}

// NewMgo ...
func NewMgo(uri string) *Mgo {
	return nil

}

// Open ...
func (m *Mgo) Open(ID string) (io.ReadCloser, error) {
	return nil, errors.New("no impl")
}

// Create ...
func (m *Mgo) Create(ID string) (io.WriteCloser, error) {
	return nil, errors.New("no impl")
}
