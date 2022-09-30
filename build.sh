#!/bin/bash

set -eux

NAME=speedwagon

PLATFORMS=(
    darwin
    linux
    windows
)

for platform in "${PLATFORMS[@]}"; do
    exe_name="${NAME}_${platform}"

    # Windows should have an .exe at the end
    if [[ "${platform}" = "windows" ]]; then
        exe_name="${exe_name}.exe"
    fi

    # macOS should be a universal binary
    if [[ "${platform}" = "darwin" ]]; then
        for goarch in arm64 amd64; do
            GOARCH="${goarch}" go build -o "${NAME}_${goarch}"
        done
        lipo -create -output "${exe_name}" "${NAME}_arm64" "${NAME}_amd64"
        rm -f "${NAME}_arm64" "${NAME}_amd64"
    else
        GOOS="${platform}" go build -o "${exe_name}"
    fi
done

