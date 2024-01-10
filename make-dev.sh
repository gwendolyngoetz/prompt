#!/bin/bash

VERSION="-$(date '+%Y%m%d%H%M%S')-dev" make
./build/prompt --version
#make install
