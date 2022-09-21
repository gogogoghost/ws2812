#!/bin/sh

export GOPROXY=https://goproxy.cn,direct
export env GOOS=linux
export GOARCH=arm
export CGO_ENABLED=1
export CC=arm-linux-gnueabihf-musl-gcc
go build -p 12 -ldflags "-s -w" -o dist/spi