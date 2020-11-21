export GOFLAGS=-mod=vendor
export GO111MODULE=on

test:
	go test -v -count=1 -race ./...

format:
	go fmt ./...
	goimports -l -w *.go
