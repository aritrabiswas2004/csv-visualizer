/************* SPDX License Identifier: MIT ************
 * Visualizer - Main run function of this program
 *
 * Copyright (C) 2026 Aritra Biswas
 *
 * Author: Aritra Biswas <aritrabb@gmail.com>
 *
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

	if opts.Markdown {
		err := table.PrintMarkdown(string(data))
		if err != nil {
			return err
		}
	}

	if opts.Text {
		err := table.Print(string(data))
		if err != nil {
			return err
		}
	}

	if opts.Export.filename != "" {
		err := ExportFile(opts.Export.filename, data)
		if err != nil {
			return err
		}
	}

	return nil
}
