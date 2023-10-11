
.PHONY=all clean check

all:
	go build -trimpath -ldflags='-s -w'

clean:
	go clean

check:
	go vet ./...
	staticcheck ./...
