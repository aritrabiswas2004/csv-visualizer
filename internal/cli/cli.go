package cli

import (
	"fmt"
)

// Making a CLI for general purpose use

const HELP_MESSAGE string = "CSV Visualizer for the ANSI terminal - v0.0.1\n\n" +
	"USAGE: ./viz [OPTIONS] [FILE]\n\n" +
	"Options:\n" +
	"	-h	--help Prints this help message" +
	"\n\nCheck development information on https://github.com/aritrabiswas2004/csv-visualizer\n" +
	"More options are to be added."

func parseOptions(argv []string) {
	for _, opt := range argv {
		if opt == "-h" || opt == "--help" {
			fmt.Printf("%v\n", HELP_MESSAGE)
		}
	}
}
