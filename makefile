
all: clean test build


test:
	go test ./... -cover
	./integration_test.sh


build: mockgen
	protoc -I blob/ blob/blob.proto --go_out=plugins=grpc:blob
	go build ./...


mockgen:
	mockgen github.com/shoebillk/sbs/blob \
		BlobServiceClient,BlobService_PushClient,BlobService_GetClient,BlobServiceServer,BlobService_PushServer,BlobService_GetServer \
		> ./mock_blob/mock_blob.go


clean:
	go clean 
	go clean -testcache

