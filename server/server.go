package server

import (
	"io"

	blob "github.com/shoebillk/sbs/blob"
)

// Server structure implements BlobServiceServer.
type Server struct {
	provider Provider
}

// NewServer return Server object.
func NewServer(provider Provider) *Server {
	return &Server{provider}
}

// Push handles push call from client.
func (s *Server) Push(stream blob.BlobService_PushServer) error {

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&blob.PushStatus{})
		}
		if err != nil {
			return err
		}

		//TODO handle content
		chunk = chunk

	}

}

// Get handles get request from client.
func (s *Server) Get(req *blob.GetRequest, stream blob.BlobService_GetServer) error {
	return nil
}
