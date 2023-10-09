# speedwagon

A very simple CLI tool for downloading Apple's Xcode simulator runtimes without needing a copy of Xcode, or a Mac.

<img width="234" alt="Screenshot 2023-10-09 at 8 50 51 AM" src="https://github.com/timsutton/speedwagon/assets/119358/f4db63dd-b56f-41a1-8266-40e1f8668fd8">

## Why?

It is often handy to have local copies of additional simulator runtimes for the purposes of setting up macOS build machines for continuous integration. For example, they can be hosted on nearby mirror servers and retrieved more quickly and reliably.

Also, depending on the simulator runtime version there are two different mechanisms required to download them, and this tool handles both of these.

**Note:** This is a scratch-my-own-itch usecase I picked as a first project to write something in Go. I can only assume the Go code and structure is not idiomatic, and so I don't recommend looking at any of it.


## Installation

A few options:

  * Download the [latest pre-built release](https://github.com/timsutton/speedwagon/releases/latest)
  * Install it with homebrew:
    * `brew install timsutton/formulae/speedwagon` to install from the above pre-built releases
    * `brew install --head timsutton/formulae/speedwagon` to build from the tip of `main`
  * Build it yourself using `go build`


## Usage

List available runtimes:

```
# output below is abbreviated
$ speedwagon list
┌──────────────────────────────────────┬─────────┬──────────┬───────────┬────────┐
│ NAME                                 │ VERSION │ BUILD    │ TYPE      │ SIZE   │
├──────────────────────────────────────┼─────────┼──────────┼───────────┼────────┤
│ iOS 12.4 Simulator                   │ 12.4    │ 16G73    │ package   │ 2.6 GB │
├──────────────────────────────────────┼─────────┼──────────┼───────────┼────────┤
│ tvOS 12.4 Simulator                  │ 12.4    │ 16M567   │ package   │ 1.2 GB │
├──────────────────────────────────────┼─────────┼──────────┼───────────┼────────┤
│ iOS 16 Simulator Runtime             │ 16.0    │ 20A360   │ diskImage │ 6.3 GB │
├──────────────────────────────────────┼─────────┼──────────┼───────────┼────────┤
│ tvOS 16.1 Simulator Runtime Beta     │ 16.1    │ 20K5041d │ diskImage │ 3.3 GB │
├──────────────────────────────────────┼─────────┼──────────┼───────────┼────────┤
│ watchOS 9.1 Simulator Runtime Beta   │ 9.1     │ 20S5044e │ diskImage │ 3.6 GB │
└──────────────────────────────────────┴─────────┴──────────┴───────────┴────────┘
```

Download one:

```
$ speedwagon download 'watchOS 9.1'
Downloading 'watchOS 9.1 Simulator Runtime Beta.dmg'...
```

Use the `-h` flag for more usage help.

## Simulator types

Pre-iOS-16-era simulator runtimes are distributed using installer packages (wrapped in a disk image) that copy all the runtime files directly onto the filesystem. Current runtimes contain the files directly on the (now [LZFSE-compressed](https://en.wikipedia.org/wiki/LZFSE)) disk image, and Xcode's simulator daemon keeps these images mounted at special volume paths.


## Scope

This tool will *download* both of the above types of simulators, but it doesn't currently offer any functionality to *install* them on a macOS system. As the newer disk-image-based type supports CLI installation via `xcrun simctl runtime` (see [Apple's docs](https://developer.apple.com/documentation/xcode/installing-additional-simulator-runtimes)), and all new simulators will use this format going forward, there's less and less need to wrap that functionality into this tool. This may be added in the future, though.


## More

There are more details about this newer simulator distribution format and how they're retrievable in [this blog post](https://macops.ca/xcode-14-new-platforms-packaging-format/).


## Future feature ideas

* `list` command output in machine-parseable formats (to aid with automation)
* Nicer progress output, better implementation for downloads given the large file sizes
* Configurable output destinations and file naming template
* Error handling
* Support rewriting older packages' installation location metadata-rewrite process [as described here](https://macops.ca/xcode-deployment-the-dvtdownloadableindex-and-ios-simulators/), so that downloaded packages are immediately useful
