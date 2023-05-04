#!/bin/bash

# used for testing unreadable files
function before() {
  note "setting testdata/secret to unreadable"
  chmod 000 testdata/secret
}

function after() {
  note "setting testdata/secret back to readable"
  chmod 644 testdata/secret
}

# remove vendoring from the default dep task
function process_stage_dep() {
  header "dep" "Ensuring dependencies are clean"
  go mod tidy
  go mod verify
  if grep -qcE ^replace go.mod; then
    warn "go.mod contains 'replace' directives"
  fi
}
