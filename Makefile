install:
	go test -race ./...
	go install ./cmd/bostongo