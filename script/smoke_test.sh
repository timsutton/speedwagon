#!/bin/bash

set -eux

function cleanup() {
    rm -f tvOS*
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

    # # smallest 'diskImage' type
    ./speedwagon download 'tvOS 16 Simulator Runtime'
}

build
test
