#!/bin/bash

function bold() {
  printf "\033[34;1m%s\033[0m\n" "$*"
}

tasks=(fmt lint vet err static sec cyclo test cover)

if [[ "${#@}" -gt 0 ]]; then
  tasks=("$@")
fi

for task in "${tasks[@]}"; do
  bold $task
  case $task in
    fmt)
      go fmt ./...
      ;;
    lint)
      golint ./...
      ;;
    vet)
      go vet ./...
      ;;
    err)
      errcheck -ignoretests ./...
      ;;
    static)
      # https://staticcheck.io/docs/checks
      staticcheck ./...
      ;;
    sec)
      gosec -exclude=G304 -quiet ./...
      ;;
    cyclo)
      gocyclo -over 10 -avg .
      ;;
    test)
      chmod 000 testdata/secret
      go test ./... -cover -coverprofile=coverage.out
      chmod 644 testdata/secret
      ;;
    cover)
      go tool cover -html coverage.out -o coverage.html
      go tool cover -func coverage.out
      ;;
  esac
done
