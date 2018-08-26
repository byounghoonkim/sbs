package server

import (
	blob "github.com/shoebillk/sbs/blob"
)

type Server struct{}

func (s *Server) Push(stream blob.BlobService_PushServer) error {
	return nil
}
func (s *Server) Get(req *blob.GetRequest, stream blob.BlobService_GetServer) error {
	return nil
}
