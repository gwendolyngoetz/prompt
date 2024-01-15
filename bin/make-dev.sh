#!/bin/bash

pushd "$(git rev-parse --show-toplevel)" >/dev/null || exit

VERSION="-$(date '+%Y%m%d%H%M%S')-dev" make
./build/linux/amd64/prompt --version
#make install
