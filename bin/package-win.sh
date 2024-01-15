#! /usr/bin/env bash

pushd "$(git rev-parse --show-toplevel)" >/dev/null || exit

VERSION=0.0.1

mkdir -p ./package
zip -j package/prompt_${VERSION}.zip build/windows/amd64/promptwin.exe
