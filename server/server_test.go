package server

import (
	"testing"

	"github.com/shoebillk/sbs/blob"
	"github.com/spf13/afero"
)

func TestNewServer(t *testing.T) {
	fb := NewFileBase(".").WithFs(afero.NewMemMapFs())
	s := NewServer(fb)
	_, checked := interface{}(s).(blob.BlobServiceServer)
	if checked == false {
		t.Fatal("Not implemented interface.")
	}
}
