#!/bin/bash

set -eux

function cleanup() {
    rm -f tvOS* com.apple.pkg*
}

trap cleanup EXIT

function build() {
    go build -v
}

function test() {
    ./speedwagon list

    (
        version=$(./speedwagon version)
        grep -q $version util/version.go
    )

    # # smallest 'package' type
    ./speedwagon download 'tvOS 12.4 Simulator'
    file com.apple.pkg* | grep -q 'zlib compressed data'

    # # smallest 'diskImage' type
    ./speedwagon download 'tvOS 16 Simulator Runtime'
}

build
test
