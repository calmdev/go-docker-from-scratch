#!/bin/bash

# Build web binary for linux.
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static"' -o ./bin/web ./cmd/web/main.go

# Build docker scratch image for web binary.
docker build -t web:latest .
