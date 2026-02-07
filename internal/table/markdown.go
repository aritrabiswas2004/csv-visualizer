/************* SPDX License Identifier: MIT ************
 * Visualizer - Render markdown compatible dynamic table
 *
 * Copyright (C) 2026 Aritra Biswas
 *
 * Author: Aritra Biswas <aritrabb@gmail.com>
 *
********************************************************/

package table

import (
	"fmt"
	"strings"

	"github.com/aritrabiswas2004/csv-visualizer/internal/csvutils"
)

func markdownSliceToDashString(lengths []int) string {
	formatString := strings.Builder{}

	formatString.WriteString("|")

	for _, length := range lengths {
		formatString.WriteString(fmt.Sprintf("-%s-|", repeat("-", length)))
	}

	return formatString.String()
}

func markdownItemSliceToDashString(items []string, itemWidths []int) string {
	formatString := strings.Builder{}

	formatString.WriteString("|")

	for i, item := range items {
		formatString.WriteString(fmt.Sprintf(" %-*s |", itemWidths[i], item))
	}

	return formatString.String()
}

func PrintMarkdown(data string) error {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		return err
	}

	// top
	fmt.Println(markdownItemSliceToDashString(headers, headerLengths))

	// div
	fmt.Println(markdownSliceToDashString(headerLengths))

	for _, v := range lines[1:] {
		row := csvutils.CleanRow(strings.Split(v, ","))

		// data item
		fmt.Println(markdownItemSliceToDashString(row, headerLengths))
	}

	return nil
}

func ReturnMarkdown(data string) (string, error) {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		return "", err
	}

	tableString := strings.Builder{}

	// top
	tableString.Write([]byte(fmt.Sprintf("%s\n", itemSliceToDashString(headers, headerLengths))))

	// div
	tableString.Write([]byte(fmt.Sprintf("%s\n", markdownSliceToDashString(headerLengths))))

	for _, v := range lines[1:] {
		row := csvutils.CleanRow(strings.Split(v, ","))

		// data item
		tableString.Write([]byte(fmt.Sprintf("%s\n", itemSliceToDashString(row, headerLengths))))
	}

	return tableString.String(), err
}
