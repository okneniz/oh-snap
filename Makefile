default: test

test:
	go test -v -count=1 ./...

install-linter:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.13.2

lint: fmt
	golangci-lint run ./...

fmt:
	golangci-lint fmt

race:
	go test -race -count=1 ./...

coverage:
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

build:
	go build ./...

pub:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/okneniz/oh-snap
