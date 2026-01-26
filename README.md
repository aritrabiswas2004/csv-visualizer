# CSV Visualizer

Renders a MySQL-like dynamic terminal table but for CSV files. 

Great for tables ASCII-only notes, or just a lightweight visualizer without a crazy
GUI program.

Good code standards are a myth; code works fine if it works at all.

## Installation

> [!NOTE]
> As of now the install script does not move the binary
> to `/usr/bin` or `/usr/local` so you have to be in the
> directory to run the program. This is subjected to change
> on first stable release.

Get the required files using `git` and then build using
`make`. Run the installation script for latest build.

```shell
./install-git.sh
```

The script does require `git` and `make` to be installed.

## Build Instructions

Check `Makefile` comments before running anything. Uses Go version `1.25.4`.

Build the executable in current directory with
```shell
make build
```

Test all available tests with. It basically runs all test files in the `internal/`
directory. But remember...real devs test in production. And users are just advanced
testers.

```shell
make test
```

And cleanup with

```shell
make clean
```

A `test.csv` is given for your convenience in the repo for testing.

## Terminal Output

See [rendering docs](./documentation/rendering.md)

_Copyright (C) 2026 [Aritra Biswas](mailto:aritrabb@gmail.com)_
