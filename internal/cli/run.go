/************* SPDX License Identifier: MIT ************
 * Visualizer - Main run function of this program
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
	"os"

	"github.com/aritrabiswas2004/csv-visualizer/internal/table"
)

func Run(args []string) error {
	opts, err := ParseArgs(args)
	if err != nil {
		return err
	}

	if opts.Help {
		fmt.Println(HelpMessage)
		return nil
	}

	data, err := os.ReadFile(opts.File)
	if err != nil {
		return err
	}

	return table.Print(string(data))
}
