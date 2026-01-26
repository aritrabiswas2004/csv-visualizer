# Visualizer Docs

## Options

- `-h` or `--help`
  - Prints the GNU type of help message on the terminal but no man page exists yet.
- `-t` or `--text`
  - Prints a general text (MySQL-like) table on standard output. This is enabled by default.
- `-m` or `--markdown`
  - Prints a markdown ASCII table to stdout.
- `--export=FILENAME`
  - Exports the CSV table and rendered ASCII version to a file with any name. The extensions must be `.txt` or `.md` and
  any other will throw an error.

## Usage

The program must take in one CSV file only, the extension
is checked at runtime.

```shell
viz -t test.csv
```

If exporting occurs, nothing will be printed to stdout,
due to the possibility of lag in large CSV tables.

```shell
viz test.csv --export=hello.md
```
