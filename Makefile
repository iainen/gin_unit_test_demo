# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test

run:
	$(GOCMD) run cmd/demo/main.go
test:
	$(GOTEST) -v ./...

cover:
	$(GOTEST) ./... -coverprofile=tmp/cover.out

html:
	$(GOCMD) tool cover -html=tmp/cover.out -o tmp/cover.html
