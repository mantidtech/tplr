#!/bin/bash

TASK=build
if [[ "${1}" = "install" ]]; then
  TASK=install
fi

go generate -x

MODULE="$(grep module go.mod | cut -d ' ' -f2)"
VERSION=$(git describe --always --dirty)

for p in cmd/*; do
  PROJECT="$(basename "${p}")"
  if [[ -f "cmd/${PROJECT}/.noinstall" ]]; then
    continue
  fi
  FILES="cmd/${PROJECT}/${PROJECT}.go"
  go "${TASK}" -i -ldflags="-X ${MODULE}.version=${VERSION}" "${FILES}" \
    && echo "${TASK} of ${PROJECT} ${VERSION} successful"
done
