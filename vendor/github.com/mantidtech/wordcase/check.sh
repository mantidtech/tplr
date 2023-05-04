#!/bin/bash

function bold() {
  printf "\x1b[1m%s\x1b[0m\n" "$*"
}

bold go fmt
go fmt ./...

bold golint
golint -set_exit_status ./...

bold go vet
go vet ./...

bold errcheck
errcheck -ignoretests ./...

bold staticcheck
staticcheck ./...

bold gosec
gosec -exclude=G304 -quiet ./...

bold goocyclo
gocyclo -over 10 -avg .


bold go test
go test ./... -cover -coverprofile=coverage.out
go tool cover -html coverage.out -o coverage.html
go tool cover -func coverage.out

# go test -bench=.
