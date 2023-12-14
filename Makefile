.PHONY: build
build:
	go build -o ./bin/scanner -ldflags="-s -w" -trimpath ./cmd/osvscanner/scanner.go

.PHONY: test
test:
	go test -race ./...
