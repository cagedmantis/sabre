all: test vet build clean

test:
	@echo ">> running tests"
	go test ./...

vet:
	@echo ">> vetting code"
	go vet ./...

build: 
	@echo ">> building binaries"
	go build ./...

clean:
	@echo ">> cleaning"

proto:
	protoc -I toothapi/ toothapi/tooth.proto --go_out=plugins=grpc:toothapi
	protoc -I gumapi/ gumapi/gum.proto --go_out=plugins=grpc:gumapi

.PHONY: build clean test vet 
