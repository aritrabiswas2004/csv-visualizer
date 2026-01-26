/************* SPDX License Identifier: MIT ************
 * Visualizer - Parse command line arguments
 *
 * Copyright (C) 2026 Aritra Biswas
 *
 * Author: Aritra Biswas <aritra.biswas@ru.nl>
 *
 * NOTE: The git commits are not tied to this email.
********************************************************/

package cli

import "fmt"

type Options struct {
	File string
	Help bool
}

func ParseArgs(args []string) (Options, error) {
	opts := Options{}

	if len(args) < 2 {
		return opts, fmt.Errorf("no input file provided")
	}

	for _, arg := range args[1:] {
		switch arg {
		case "-h", "--help":
			opts.Help = true
		default:
			opts.File = arg
		}
	}

	return opts, nil
}
