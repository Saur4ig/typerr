.PHONY: lint, test, fmtall, all, o

lint:
	golangci-lint run

test:
	go test -count=1 ./... -cover

fmtall:
	go fmt ./...

all:
	make fmtall
	make lint
	make test