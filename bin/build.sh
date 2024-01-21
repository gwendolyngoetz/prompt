#! /usr/bin/env bash

pushd "$(git rev-parse --show-toplevel)" >/dev/null || exit

if [[ -d ./build ]]; then
    echo "Removing ./build folder"
    rm -rf ./build
fi

readonly ARTIFACT_NAME="prompt"
readonly PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

for PLATFORM in "${PLATFORMS[@]}"; do
    echo "Building ${PLATFORM}..."
    IFS="/" read -r -a SPLIT <<< "${PLATFORM}"
	GOOS=${SPLIT[0]}
	GOARCH=${SPLIT[1]}
	name="${ARTIFACT_NAME}"

	if [[ "${GOOS}" == "windows" ]]; then
		name+="win.exe"
	fi

    GOARCH="${GOARCH}" \
    GOOS="${GOOS}" \
    go build -ldflags="-X 'gwendolyngoetz/prompt/cmd.Version=v${VERSION}'" -o "build/${PLATFORM}/${name}" ./main.go
done
