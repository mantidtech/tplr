#!/bin/bash

PROJECT_NAME="$(basename "${PWD}")"

modules=$(go list ./...)

tasks=(dep fmt revive vet err static sec cyclo test cover)


while [[ "$1" =~ ^- ]]; do
  case $1 in
  -v)
    TESTARGS='-v'
    ;;
  esac
  shift
done


# Things to do before processing anything.  Override in ./check_extra.sh or .develop.env
function before() {
  header tests running: ${tasks}
}

# Thing to do after processing everything.  Override in ./check_extra.sh or .develop.env
function after() {
  header done
}

function header() {
  name="${1}"
  shift
  printf "\e[1;32m%-15s\e[0;34m%s\e[0m\n" "${name}" "${*}"
}

function warn() {
  printf "  \e[1;33m*** WARNING\e[0;33m%s\e[0m\n" "${*}"
}

# Add in project specific stuff
if [[ -f "./check_extra.sh" ]]; then
  source ./check_extra.sh
fi

# add in any developer specific environment
if [[ -f "./develop.env" ]]; then
  source ./develop.env
fi

if [[ "${#@}" -gt 0 ]]; then
  tasks=("$@")
fi

before
for task in "${tasks[@]}"; do
  case $task in
    dep)
      header $task "Ensuring dependencies are clean"
      go mod tidy
      go mod download
      if grep -qcE ^replace go.mod; then
        warn "go.mod contains 'replace' directives"
      fi
      ;;
    fmt)
      header $task "Standardising formatting"
      files=()
      while IFS='' read -r filename; do
        files+=("${filename}")
      done < <(find . -name '*.go' -not -name '*.pb.go' -not -path '*/vendor/*')
      for f in "${files[@]}"; do
        sed -i "" -e '/import (/,/)/{/\/\//,/^$/N;/^$/d;}' "${f}"
        goimports -w -local code.mantid.org "${f}"
      done
      go fmt $modules
      ;;
    revive)
      header $task "Checking linting rules"
      revive -formatter friendly -exclude vendor/... -exclude mocks/... $modules
      ;;
    vet)
      header $task "Examining code for suspicious constructs"
      go vet $modules
      ;;
    err)
      header $task "Checking for uncaught error returns"
      errcheck -ignoretests $modules
      ;;
    static)
      header $task "Static checking of code for common errors"
      # https://staticcheck.io/docs/checks
      staticcheck $modules
      ;;
    sec)
      header $task "Looking for common programming mistakes that can lead to security problems."
      gosec -exclude=G304 -quiet ./...
      ;;
    cyclo)
      header $task "Looking for potential refactoring required for functions with high complexity"
      gocyclo -over 10 -avg .
      ;;
    test)
      header $task "Running all unit tests"
      go test $modules -cover -coverprofile=coverage.out ${TESTARGS}
      ;;
    cover)
      header $task "Generating coverage report"
      go tool cover -html coverage.out -o coverage.html
      go tool cover -func coverage.out
      ;;
  esac
done
after
