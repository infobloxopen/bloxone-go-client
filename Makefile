default: test

fmt:
	go fmt ./...

vet:
	go vet ./...

test: fmt vet
	go test -coverprofile cover.out ./...
