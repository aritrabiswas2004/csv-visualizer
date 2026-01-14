# CSV Visualizer

Visualize CSV files in the terminal. Written in Go.

Right now reads the default `test.csv` (but it can work with any CSV file)

Good code standards are a myth; code works fine if it works at all.

## Build Instructions

Check `Makefile` comments before running anything. Uses Go version `1.25.4`.

Build the executable in current directory with
```shell
make build
```

Test all available tests with
```shell
make test
```

(Possible soon deprecation) Run with no perma-executable
```shell
make run
```

And cleanup with

```shell
make clean
```

## Terminal Output

Given an ANSI standard terminal (basically anything except Windows) an output can be like
below. It resembles terminal output like MySQL.

```text
+---------------------+----------------+------+-----------+----------+
| Name                | Location       | Age  | Elevation | Good     |
+---------------------+----------------+------+-----------+----------+
| Taj Mahal           | Agra           | 500  | 40.56     | yes      |
| Pyramids of Giza    | Giza           | 4000 | 46.77     | no       |
| Colosseum           | Rome           | 5000 | 10.02     | bearable |
| Christ The Redeemer | Rio de Janeiro | 400  | 100.22    | great    |
+---------------------+----------------+------+-----------+----------+
```

_Copyright (C) 2026 [Aritra Biswas](mailto:aritrabb@gmail.com)_
