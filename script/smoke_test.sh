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
    file com.apple.pkg.AppleTVSimulatorSDK12_4*.dmg | grep -q 'zlib compressed data'

    # # smallest 'diskImage' type
    ./speedwagon download 'tvOS 16 Simulator Runtime'
    file tvOS_16_Simulator_Runtime.dmg | grep -q 'lzfse encoded, lzvn compressed'
}

build
test
