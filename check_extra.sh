#!/bin/bash

function before() {
  chmod 000 testdata/secret
}

function after() {
  chmod 644 testdata/secret
}
