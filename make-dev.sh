#!/bin/bash

VERSION="-$(date '+%Y%m%d%H%M%S')-dev" make
./build/linux/amd64/prompt --version
#make install
