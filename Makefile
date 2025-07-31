default: test

test:
	go test -v -vet=off ./...

install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

lint: fmt
	golangci-lint run ./...

fmt:
	gofmt -w -s .

coverage:
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

build:
	go build ./...

pub:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/okneniz/ohsnap
