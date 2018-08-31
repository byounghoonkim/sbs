package server

import (
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shoebillk/sbs/blob"
	"github.com/shoebillk/sbs/mock_blob"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	fb := NewFileBase(".").WithFs(afero.NewMemMapFs())
	s := NewServer(fb)
	_, checked := interface{}(s).(blob.BlobServiceServer)
	if checked == false {
		t.Fatal("Not implemented interface.")
	}
}

func TestPush(t *testing.T) {
	fb := NewFileBase(".").WithFs(afero.NewMemMapFs())
	s := NewServer(fb)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	chunk := blob.Chunk{
		Id:      "test",
		Content: []byte("aaaaaaaaaaaaaaaaaaaa"),
	}

	mockPushserver := mock_blob.NewMockBlobService_PushServer(ctrl)
	mockPushserver.EXPECT().Recv().Return(&chunk, nil)
	mockPushserver.EXPECT().Recv().Return(nil, io.EOF)
	mockPushserver.EXPECT().SendAndClose(&blob.PushStatus{
		Message: "OK",
		Code:    blob.PushStatusCode_Ok,
	}).Return(nil)

	err := s.Push(mockPushserver)

	assert.Equal(t, err, nil)

}

func TestGet(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "./test", []byte("aaaaaaaaaaaaaaaaaaaaaaaaaa"), 0644)

	fb := NewFileBase(".").WithFs(fs)
	s := NewServer(fb)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetserver := mock_blob.NewMockBlobService_GetServer(ctrl)

	req := blob.GetRequest{Id: "test"}
	mockGetserver.EXPECT().Send(gomock.Any()).Return(nil)

	err := s.Get(&req, mockGetserver)

	assert.Equal(t, err, nil)
}
