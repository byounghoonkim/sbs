package server

import (
	"io"
	"path/filepath"

	"github.com/spf13/afero"
)

// FileBase implements Provider interface using file system.
type FileBase struct {
	root string
	fs   afero.Fs
}

// NewFileBase ...
func NewFileBase(root string) *FileBase {
	return &FileBase{root, afero.NewOsFs()}
}

func (fb *FileBase) WithFs(fs afero.Fs) *FileBase {
	fb.fs = fs
	return fb
}

// Open return ReadCloser interface for read the contents.
func (fb *FileBase) Open(ID string) (io.ReadCloser, error) {
	return fb.fs.Open(filepath.Join(fb.root, ID))
}

// Create return WriteCloser interface for write the contents.
func (fb *FileBase) Create(ID string) (io.WriteCloser, error) {
	return fb.fs.Create(filepath.Join(fb.root, ID))
}
