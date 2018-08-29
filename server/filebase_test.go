package server

import (
	"testing"

	"github.com/spf13/afero"
)

func TestNew(t *testing.T) {
	fs := afero.NewMemMapFs()
	fb := NewFileBase(".").WithFs(fs)
	_, checked := interface{}(fb).(Provider)
	if checked == false {
		t.Fatal("FileBase must be a kind of Provider")
	}
}

func TestOpen(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "./test", []byte("aaa"), 0644)
	fb := NewFileBase(".").WithFs(fs)
	r, err := fb.Open("test")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
}

func TestCreate(t *testing.T) {
	fs := afero.NewMemMapFs()
	fb := NewFileBase(".").WithFs(fs)
	r, err := fb.Create("test")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
}
