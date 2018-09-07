package client

import (
	"bytes"
	"io"
	"testing"

	"io/ioutil"

	"github.com/golang/mock/gomock"
	"github.com/shoebillk/sbs/blob"
	"github.com/shoebillk/sbs/mock_blob"
)

func TestNewClient(t *testing.T) {

	ctrl := gomock.NewController(t)
	b := mock_blob.NewMockBlobServiceClient(ctrl)

	_, err := NewClient(b)

	if err != nil {
		t.Fatal(err)
	}

}

func TestGet(t *testing.T) {

	chunk := blob.Chunk{
		Id:      "test",
		Content: []byte("aaaaaaaaaaaaaaaaaaaa"),
	}

	ctrl := gomock.NewController(t)

	mockGetClient := mock_blob.NewMockBlobService_GetClient(ctrl)

	mockGetClient.EXPECT().Recv().Return(&chunk, nil)
	mockGetClient.EXPECT().Recv().Return(nil, io.EOF)

	mockBSC := mock_blob.NewMockBlobServiceClient(ctrl)
	mockBSC.EXPECT().Get(gomock.Any(), &blob.GetRequest{Id: "test"}).Return(mockGetClient, nil)

	c, err := NewClient(mockBSC)

	if err != nil {
		t.Fatal(err)
	}

	n, err := c.Get("test", ioutil.Discard)

	if err != nil {
		t.Fatal(err)
	}

	if n != int64(len(chunk.Content)) {
		t.Fatal("Get recieve failed")
	}

}

func TestPush(t *testing.T) {
	chunk := blob.Chunk{
		Id:      "test",
		Content: []byte("aaaaaaaaaaaaaaaaaaaa"),
	}

	pushStatus := blob.PushStatus{
		Message: "OK",
		Code:    blob.PushStatusCode_Ok,
	}

	ctrl := gomock.NewController(t)

	mockPushClient := mock_blob.NewMockBlobService_PushClient(ctrl)
	mockPushClient.EXPECT().Send(&chunk).Return(nil)
	mockPushClient.EXPECT().CloseAndRecv().Return(&pushStatus, nil)

	mockBSC := mock_blob.NewMockBlobServiceClient(ctrl)
	mockBSC.EXPECT().Push(gomock.Any()).Return(mockPushClient, nil)

	c, err := NewClient(mockBSC)

	if err != nil {
		t.Fatal(err)
	}

	pushStatusResult, err := c.Push("test", bytes.NewReader(chunk.Content))

	if err != nil {
		t.Fatal(err)
	}

	if pushStatusResult.Code != pushStatus.Code {
		t.Fatal(err)
	}

}
