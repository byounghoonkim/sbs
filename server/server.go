package server

import (
	"io"
	"log"

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

	var wc io.WriteCloser
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&blob.PushStatus{
				Message: "OK",
				Code:    blob.PushStatusCode_Ok,
			})
		}
		if err != nil {
			return err
		}

		if wc == nil {
			wc, err = s.provider.Create(chunk.Id)
			if err != nil {
				return err
			}
			defer wc.Close()
		}

		n, err := wc.Write(chunk.Content)
		if err != nil {
			return err
		}

		log.Printf("Id : %s push %d bytes", chunk.Id, n)
	}

}

// Get handles get request from client.
func (s *Server) Get(req *blob.GetRequest, stream blob.BlobService_GetServer) error {

	b, err := s.provider.Open(req.Id)
	if err != nil {
		return err
	}
	defer b.Close()

	chunk := blob.Chunk{
		Id:      req.Id,
		Content: make([]byte, 4048),
	}

	for {
		n, err := b.Read(chunk.Content)
		if err != nil && err != io.EOF {
			return err
		}

		if n > 0 {
			if nil != stream.Send(&chunk) {
				return err
			}
		}

		if err == io.EOF {
			return nil
		}

	}
}
