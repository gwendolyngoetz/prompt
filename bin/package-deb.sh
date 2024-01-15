#! /usr/bin/env bash

pushd "$(git rev-parse --show-toplevel)" >/dev/null || exit

VERSION=0.0.1

mkdir -p ./package/prompt_${VERSION}/usr/local/bin
mkdir -p ./package/prompt_${VERSION}/DEBIAN

cp ./build/linux/amd64/prompt ./package/prompt_${VERSION}/usr/local/bin
touch ./package/prompt_${VERSION}/DEBIAN/control

CONTROL_FILE_BODY="Package: prompt
Version: ${VERSION}
Section: base
Priority: optional
Architecture: amd64
Maintainer: Gwendolyn Goetz
Description: Fancy prompt output"

echo "${CONTROL_FILE_BODY}" > ./package/prompt_${VERSION}/DEBIAN/control

dpkg-deb --build ./package/prompt_${VERSION}


