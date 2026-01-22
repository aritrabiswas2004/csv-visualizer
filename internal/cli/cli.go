package cli

import (
	"fmt"
)

// Making a CLI for general purpose use

const HelpMessage string = "CSV Visualizer for the ANSI terminal - v0.0.1\n\n" +
	"USAGE: ./viz [OPTIONS] [FILE].csv\n\n" +
	"Options:\n" +
	"	-h	--help Prints this help message" +
	"\n\nCheck the up-to-date development information on the Mailing List. All are welcome to contribute." +
	"\nMirror repository: https://github.com/aritrabiswas2004/csv-visualizer\n" +
	"Development Mailing List: https://groups.google.com/g/csv-coreutils-dev"

func parseOptions(argv []string) {
	for _, opt := range argv {
		if opt == "-h" || opt == "--help" {
			fmt.Printf("%v\n", HelpMessage)
		}
	}
}
