PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(shell go env | grep PATH | cut -d "=" -f 2 | sed -E 's/"//g')
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')


go-build: clean
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES) test

clean:
	@rm -f $(GOBIN)/$(PROJECTNAME)
