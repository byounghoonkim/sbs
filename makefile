
test:
	go test ./...

build:
	protoc -I blob/ blob/blob.proto --go_out=plugins=grpc:blob
	go build ./...


mockgen:
	mockgen github.com/shoebillk/sbs/blob BlobService_PushServer,BlobService_GetServer > ./mock_blob/mock_blob.go


