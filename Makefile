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
	rm toothserver

.PHONY: build clean test vet 
