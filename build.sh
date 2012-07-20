#!/bin/bash
export GOROOT=/usr/local/go
export GOARCH=amd64
export GOOS=linux
export GOBIN=/usr/local/go/bin
export PATH=$PATH:$GOBIN

go build
