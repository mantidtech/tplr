#!/bin/bash
set -e

# shellcheck disable=SC2034
PROJECT_NAME="$(basename "${PWD}")"

tasks=(dep fmt revive vet err static sec vuln cyclo test cover)

while [[ "$1" =~ ^- ]]; do
  case $1 in
  -v)
    TESTARGS='-v'
    ;;
  esac
  shift
done

function header() {
  name="${1}"
  shift
  printf "\e[1;32m%-15s\e[0;34m%s\e[0m\n" "${name}" "${*}"
}

function warn() {
  printf "  \e[1;33m*** WARNING\e[0;33m %s\e[0m\n" "${*}"
}

function note() {
  printf "\e[2;36m>> %s\e[0m\n" "${*}"
}

# Things to do before processing anything.  Override in ./check_extra.sh or ./develop.env
function before() {
  # nothing by default
  :
}

# Thing to do after processing everything.  Override in ./check_extra.sh or ./develop.env
function after() {
  # nothing by default
  :
}

if [[ ! -f "go.mod" ]]; then
  note "Current directory is not the root of a go project"
  exit
fi

modules=()
while read -r m; do
  modules+=("$m")
done < <(go list ./... | grep -vE '/cmd/' )

#note "Processing Modules:"
#for i in "${modules[@]}"; do
#  note "$i"
#done

if [[ "${#@}" -gt 0 ]]; then
  tasks=("$@")
fi


### define stages

function process_stage_defined_stages() {
  header "$task" "List of all defined stages"
  declare -F | cut -d' ' -f3 | grep process_stage | cut -d'_' -f3-
}

function process_stage_dep() {
  header "$task" "Ensuring dependencies are clean"
  go mod tidy
  go mod vendor
  go mod verify
  if grep -qcE ^replace go.mod; then
    warn "go.mod contains 'replace' directives"
  fi
}

function process_stage_fmt() {
  header "$task" "Standardising formatting"
  files=()
  while IFS='' read -r filename; do
    files+=("${filename}")
  done < <(find . -name '*.go' -not -name '*.pb.go' -not -path '*/vendor/*')
  for f in "${files[@]}"; do
    sed -i -e '/import (/,/)/{/\/\//,/^$/N;/^$/d;}' "${f}"
    goimports -w -local code.mantid.org "${f}"
  done
}

function process_stage_revive() {
  header "$task" "Checking linting rules"
  shift
  revive -formatter default -exclude vendor/... -exclude mocks/... "$@"
}

function process_stage_vet() {
  header "$task" "Examining code for suspicious constructs"
  shift
  go vet "$@"
}

function process_stage_err() {
  header "$task" "Checking for uncaught error returns"
  shift
  errcheck -ignoretests "$@"
}

function process_stage_static() {
  header "$task" "Static checking of code for common errors"
  shift
  # https://staticcheck.io/docs/checks
  staticcheck "$@"
}

function process_stage_sec() {
  header "$task" "Looking for common programming mistakes that can lead to security problems"
  gosec -exclude=G304 -quiet ./...
}

function process_stage_vuln() {
  header "$task" "Checking code against published vulnerabilities"
  govulncheck ./... || true
}

function process_stage_cyclo() {
  header "$task" "Looking for potential refactoring required for functions with high complexity"
  gocyclo -over 10 -avg .
  true
}

function process_stage_test() {
  header "$task" "Running all unit tests"
  shift
  COVERAGE=$(mktemp)
  go test "$@" -cover -coverprofile="${COVERAGE}" ${TESTARGS}
  grep -v '.pb.go:' "${COVERAGE}" > coverage.out
  rm "${COVERAGE}"
}

function process_stage_cover() {
  header "$task" "Generating coverage report"
  go tool cover -html coverage.out -o coverage.html
  go tool cover -func coverage.out
  gocover-cobertura < coverage.out > coverage.xml
}

# install/update the tools used by this script
function process_stage_updatetools() {
  header "$task" "Updating tool set"
  tools=(
    github.com/boumenot/gocover-cobertura@latest
    github.com/fzipp/gocyclo/cmd/gocyclo@latest
    github.com/mgechev/revive@latest
    github.com/securego/gosec/v2/cmd/gosec@latest
    golang.org/x/tools/cmd/goimports@latest
    golang.org/x/vuln/cmd/govulncheck@latest
  )
  for t in "${tools[@]}"; do
    note "Updating ${t}"
    go install "${t}" &
  done
  wait
  true
}

###

function not_stage_exists() {
  LC_ALL=C [ "$(type -t "process_stage_${1}")" = "function" ]
}

# Add in project specific stuff
if [[ -f "./check_overrides.sh" ]]; then
  source ./check_overrides.sh
fi

# add in any developer specific environment
if [[ -f "./develop.env" ]]; then
  source ./develop.env
fi

# run all of the specified stages
before
for task in "${tasks[@]}"; do
  if [[ $(not_stage_exists "${task}") ]]; then
    warn "task '${task}' doesn't exist"
    continue
  fi
  if ! "process_stage_${task}" "${modules[@]}"; then
    warn "task '${task}' failed"
    break
  fi
done
after

exit 0
