/************* SPDX License Identifier: MIT ************
 * Visualizer - Describes export package with export encoding
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
	"strings"

	"github.com/aritrabiswas2004/csv-visualizer/internal/table"
)

const (
	UTF8 = iota
	UTF16
	ISO
)

type ExportPackage struct {
	filename string // includes extension
	encoding int
}

func ExportFile(fileName string, data []byte) error {
	fmt.Printf("Exporting to %v...\n", fileName)

	var (
		result string
		err    error
	)

	switch {
	case strings.HasSuffix(fileName, ".md"):
		result, err = table.ReturnMarkdown(string(data))
	case strings.HasSuffix(fileName, ".txt"):
		result, err = table.ReturnText(string(data))
	default:
		return fmt.Errorf("unsupported filetype for %s", fileName)
	}

	if err != nil {
		return err
	}

	return appendOrCreateFile(fileName, result)
}

// Adds double newline for better looks
// Come at me if you think this sucks
func appendOrCreateFile(filename string, data string) error {
	filePtr, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		_ = filePtr.Close()
		return err
	}

	bytes, writeErr := filePtr.Write([]byte("\n" + data))
	if writeErr != nil {
		_ = filePtr.Close()
		return writeErr
	}

	fmt.Printf("Wrote %d bytes to %s\n", bytes, filename)

	closeErr := filePtr.Close()
	if closeErr != nil {
		return closeErr
	}

	return nil
}
