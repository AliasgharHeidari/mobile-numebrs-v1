build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/mobile-numbers-mine-linux-amd64 cmd/main.go

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/mobile-numbers-mine-darwin-amd64 cmd/main.go

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/mobile-numbers-mine-windows-amd64.exe cmd/main.go

doc:
	swag init -g cmd/main.go -o docs/api

run:
	echo "Starting the server..."
	go run cmd/main.go

build:
	go build cmd/main.go

build-all: build-linux-amd64 build-darwin-amd64 build-windows-amd64

clean:
	rm -rf bin/*
	rm config/config.yaml

.PHONY: doc build run build-linux-amd64 build-darwin-amd64 build-windows-amd64 build-all clean