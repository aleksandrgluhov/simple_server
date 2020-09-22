# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

SRC_DIR=src

BINARY_NAME=simple_server
BINARY_UNIX=$(BINARY_NAME)_unix

all: clean deps test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

build-prod:
	$(GOBUILD) -ldflags '-s -w' -o $(BINARY_NAME) -v 

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/joho/godotenv
	$(GOGET) github.com/kelseyhightower/envconfig
	$(GOGET) github.com/dchest/uniuri

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
