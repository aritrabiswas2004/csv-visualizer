# CSV Visualizer

Renders a MySQL-like dynamic terminal table but for CSV files. 

Great for tables ASCII-only notes, or just a lightweight visualizer without a crazy
GUI program.

Good code standards are a myth; code works fine if it works at all.

### Sample Output

For a random file like [`test.csv`](./test.csv), an output
like the one below can be seen.

```text
$ ./viz test.csv 
+---------------------+----------------+------+-----------+----------+
| Name                | Location       | Age  | Elevation | Good     |
+---------------------+----------------+------+-----------+----------+
| Taj Mahal           | Agra           | 500  | 40.56     | yes      |
| Pyramids of Giza    | Giza           | 4000 | 46.77     | no       |
| Colosseum           | Rome           | 5000 | 10.02     | bearable |
| Christ The Redeemer | Rio de Janeiro | 400  | 100.22    | great    |
| Machu Picchu        | Peru           | 1000 | 1932.22   | nice     |
+---------------------+----------------+------+-----------+----------+
```

Check out further information in [rendering docs.](./documentation/rendering.md)

## Installation

The easiest way is to check the [latest releases](https://github.com/aritrabiswas2004/csv-visualizer/releases) and download a binary
pertaining to your operating system and processor
architecture.

If you are planning to build from source, check out
the [building from source](./documentation/build-from-source.md) documentation.

You can also install with any of the installation scripts
in the [`scripts/`](./scripts) directory.

- [`install.sh`](./scripts/install.sh) is the general installation
script for all UNIX-based systems. This fetches the latest binary
from the latest release automatically.

- [`install-git.sh`](./scripts/install-git.sh) requires you to
have `git`, `make` and `go` version `1.25.4` installed. It is
intended for installation checks during development. So unless
you know what you are doing, this is not the way to install.

> [!NOTE]
> As of now the install script does not move the binary
> to `/usr/bin` or `/usr/local` so you have to be in the
> directory to run the program. This is subjected to change
> on first stable release.

_Copyright (C) 2026 [Aritra Biswas](mailto:aritrabb@gmail.com)_
