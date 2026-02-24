# Building the Visualizer from Source

You must first clone the repository

```shell
git clone https://github.com/aritrabiswas2004/csv-visualizer.git
```
A `test.csv` is given for your convenience in the repo for testing.

## Build Instructions

Check `Makefile` comments before running anything. Uses Go version `1.25.4`.

### Build the Executable

Build the executable in current directory with
```shell
make build
```

### Installation

Check if installation works correct and the [instructions for it](../README.md#installation) are
also correct. Executes `install.sh` shell script.

```shell
make install
```

### Tests

Test all available tests with. It basically runs all test files in the `internal/`
directory. But remember...real devs test in production. And users are just advanced
testers.

```shell
make test
```

### Cleanup

Remove any file matching the
regexp pattern `viz*` or Go RegExp `viz.*`. Made for clean up after compiling multiple
binaries.

> [!WARNING]
> Do NOT name any important file containing the pattern `viz` in the filename. If you do
> this command will delete it and banish you to the kingdom of viz.

```shell
make clean
```

### Build Static Binary

Build a static binary for a particular OS and arch, the
default values are the system specifics on whichever system
the command is being run.

As of now it only works on POSIX compliant shells.

```shell
# For example for the M-series MacBooks
make static-binary OS="darwin" ARCH="arm64"
```

### Build All Binaries

This is mainly for the final, drag-and-drop binaries
before publishing a release.

```shell
make all-static-bin
```
