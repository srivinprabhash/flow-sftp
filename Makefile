default: run

run:
	go run main.go config.go flow.go send.go

build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o build/flow-sftp


.PHONY: build