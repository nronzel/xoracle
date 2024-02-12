#!/bin/bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o xoracle-macos-amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o xoracle-macos-arm64
