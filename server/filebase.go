package server

import (
	"errors"
	"io"
)

// FileBase implements Provider interface using file system.
type FileBase struct {
	root string
}

// NewFileBase ...
func NewFileBase(root string) *FileBase {
	return &FileBase{root}
}

// Open return ReadCloser interface for read the contents.
func (fb *FileBase) Open(ID string) (io.ReadCloser, error) {
	return nil, errors.New("not implemented")
}

// Create return WriteCloser interface for write the contents.
func (fb *FileBase) Create(ID string) (io.WriteCloser, error) {
	return nil, errors.New("not implemented")

}
