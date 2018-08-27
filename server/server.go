package server

import (
	blob "github.com/shoebillk/sbs/blob"
)

type Server struct {
}

// Push handles push call from client.
func (s *Server) Push(stream blob.BlobService_PushServer) error {
	return nil
}

// Get handles get request from client.
func (s *Server) Get(req *blob.GetRequest, stream blob.BlobService_GetServer) error {
	return nil
}
