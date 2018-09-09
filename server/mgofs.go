package server

import (
	"io"

	"github.com/globalsign/mgo"
)

// MgoFS ...
type MgoFS struct {
	uri    string
	db     string
	prefix string
}

// NewMgoFS ...
func NewMgoFS(uri string, db string, prefix string) *MgoFS {
	return &MgoFS{uri, db, prefix}
}

// Open ...
func (m *MgoFS) Open(ID string) (io.ReadCloser, error) {
	session, err := mgo.Dial(m.uri)
	if err != nil {
		return nil, err
	}

	file, err := session.DB(m.db).GridFS(m.prefix).Open(ID)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Create ...
func (m *MgoFS) Create(ID string) (io.WriteCloser, error) {
	session, err := mgo.Dial(m.uri)
	if err != nil {
		return nil, err
	}

	file, err := session.DB(m.db).GridFS(m.prefix).Create(ID)
	if err != nil {
		return nil, err
	}
	return file, nil
}
