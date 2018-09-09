package server

import (
	"io"
	"log"

	blob "github.com/shoebillk/sbs/blob"
)

// Provider interface defines Blob Provider's Methods.
type Provider interface {
	Create(ID string) (io.WriteCloser, error)
	Open(ID string) (io.ReadCloser, error)
}

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
	recvByte := 0
	var id string
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Id : %s push %d bytes", id, recvByte)
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

			id = chunk.Id
		}

		n, err := wc.Write(chunk.Content)
		if err != nil {
			return err
		}

		recvByte += n
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
		Content: make([]byte, 4096),
	}

	for {
		n, err := b.Read(chunk.Content)
		if err != nil && err != io.EOF {
			return err
		}

		if n > 0 {
			chunk.Content = chunk.Content[:n]
			if nil != stream.Send(&chunk) {
				return err
			}
		}

		if err == io.EOF {
			return nil
		}

	}
}
