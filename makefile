
test:
	go test ./...

build:
	protoc -I blob/ blob/blob.proto --go_out=plugins=grpc:blob
	go build ./...
