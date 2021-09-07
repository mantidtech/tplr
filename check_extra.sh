#!/bin/bash

function before() {
  note "setting testdata/secret to unreadable"
  chmod 000 testdata/secret
}

function after() {
  note "setting testdata/secret back to readable"
  chmod 644 testdata/secret
}
