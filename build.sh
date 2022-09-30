#!/bin/bash

set -eux

NAME=speedwagon
PLATFORMS=( linux windows )
ARCHS=( amd64 arm64)
BUILD_DIR=build

version=$(go run main.go version)

rm -rf "${BUILD_DIR}" && mkdir "${BUILD_DIR}"

for platform in "${PLATFORMS[@]}"; do
    for goarch in "${ARCHS[@]}"; do
        exe_name="${NAME}-${version}-${platform}-${goarch}"

        # Windows should have an .exe at the end
        if [[ "${platform}" = "windows" ]]; then
            exe_name="${exe_name}.exe"
        fi

        GOARCH="${goarch}" GOOS="${platform}" go build -o "${BUILD_DIR}/${exe_name}"
    done
done

# macOS should be a universal binary, so we just handle it separately
platform=darwin
exe_name="${NAME}-${version}-${platform}-universal"
for goarch in arm64 amd64; do
    GOARCH="${goarch}" go build -o "${BUILD_DIR}/${NAME}_${goarch}"
done
lipo -create -output "${BUILD_DIR}/${exe_name}" "${BUILD_DIR}/${NAME}_arm64" "${BUILD_DIR}/${NAME}_amd64"
rm -f "${BUILD_DIR}/${NAME}_arm64" "${BUILD_DIR}/${NAME}_amd64"
