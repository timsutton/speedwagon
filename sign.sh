#!/bin/bash

# WIP standalone script to manage signing and notarization

set -eux

declare -r exe_path="${1}"

# eventually call this with goreleaser as a hook, so that this asset is uploaded?

codesign \
    --options=runtime \
    --deep \
    --strict \
    --timestamp \
    --sign 'Developer ID Application: Timothy Sutton (43Y295X5WU)' \
    "${exe_path}"

# zip it
ditto -c -k \
    "${exe_path}" \
    /tmp/speedwagon.zip

# notarytool upload it
xcrun notarytool submit --wait --keychain-profile 'tim@macops.ca' /tmp/speedwagon.zip

# not sure if 'install' type makes sense for a zip
spctl --assess -vv --type install "${exe_path}"

rm /tmp/speedwagon.zip
