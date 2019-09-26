default:
	@echo Must specify rule


init:
	go mod init go-server

test:
	go test

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build 