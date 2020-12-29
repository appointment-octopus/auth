#!/bin/bash

set -x

go test -v ./tests/... -coverpkg=$(go list ./... | grep -v "/test" | tr "\n" ",") -coverprofile=coverage.txt -race -covermode=atomic

# go tool cover -html=coverage.out -o coverage.html
