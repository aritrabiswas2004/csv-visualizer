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

import (
	"fmt"
	"strings"
)

type Options struct {
	File     string
	Help     bool
	Markdown bool
	Text     bool
	Export   ExportPackage
}

func ParseArgs(args []string) (Options, error) {
	opts := Options{Text: true}

	if len(args) < 2 {
		return opts, fmt.Errorf("no input file provided")
	}

	for _, arg := range args[1:] {
		switch arg {
		case "-h", "--help":
			opts.Help = true
		case "-m", "--markdown":
			opts.Markdown = true
			opts.Text = false
		case "-t", "--text":
			opts.Text = true
		default:
			if strings.HasPrefix(arg, "--export=") {
				opts.Text = false
				filename := strings.TrimPrefix(arg, "--export=")
				opts.Export = ExportPackage{
					filename: filename,
					encoding: UTF8,
				}
			} else {
				if strings.HasSuffix(arg, ".csv") {
					opts.File = arg
				} else {
					return opts, fmt.Errorf("invalid file format given, only .csv files permitted")
				}
			}
		}
	}

	return opts, nil
}
