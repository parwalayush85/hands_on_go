#Test
compile:
	go build -o server cmd/server.go
format:
	go fmt ./...
build:
	format compile
run:
	./server
