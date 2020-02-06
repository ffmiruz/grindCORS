build:
	mkdir -p functions
	go get ./...
	go build -o functions/grindCORS cmd/main.go
