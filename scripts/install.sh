#!/bin/bash

# Install script but only does minimal
# Does not move binary to /usr/local/bin
# Assumes you have git and make

git clone https://github.com/aritrabiswas2004/csv-visualizer.git

cd csv-visualizer

# Not make run because make run is temporary
make build
