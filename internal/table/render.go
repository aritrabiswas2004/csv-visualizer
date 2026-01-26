/************* SPDX License Identifier: MIT ************
 * Visualizer - Render dynamic table
 *
 * Copyright (C) 2026 Aritra Biswas
 *
 * Author: Aritra Biswas <aritra.biswas@ru.nl>
 *
 * NOTE: The git commits are not tied to this email.
********************************************************/

package table

import (
	"fmt"
	"strings"

	"github.com/aritrabiswas2004/csv-visualizer/internal/csvutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func repeat(char string, count int) string {
	result := strings.Builder{}

	for i := 0; i < count; i++ {
		result.WriteString(char)
	}

	return result.String()
}

func sliceToDashString(lengths []int) string {
	formatString := strings.Builder{}

	formatString.WriteString("+")

	for _, length := range lengths {
		formatString.WriteString(fmt.Sprintf("-%s-+", repeat("-", length)))
	}

	return formatString.String()
}

func itemSliceToDashString(items []string, itemWidths []int) string {
	formatString := strings.Builder{}

	formatString.WriteString("|")

	for i, item := range items {
		formatString.WriteString(fmt.Sprintf(" %-*s |", itemWidths[i], item))
	}

	return formatString.String()
}

func collectHeaderLengths(hdrs []string) []int {
	sizes := make([]int, 0, len(hdrs))

	for _, hdr := range hdrs {
		sizes = append(sizes, len(hdr))
	}

	return sizes
}

func capitalize(arr []string) []string {
	capitalized := make([]string, 0, len(arr))

	for _, str := range arr {
		capitalized = append(capitalized, cases.Title(language.English, cases.Compact).String(str))
	}

	return capitalized
}

func getDataHeadersAndMaxWidths(lines []string) ([]string, []int, error) {
	headers := capitalize(strings.Split(lines[0], ","))
	maxColumnWidths := collectHeaderLengths(headers)

	for _, v := range lines {
		row := csvutils.CleanRow(strings.Split(v, ","))

		for idx, val := range row {
			maxColumnWidths[idx] = max(maxColumnWidths[idx], len(val))
		}
	}

	if len(headers) != len(maxColumnWidths) {
		return []string{}, []int{}, fmt.Errorf("length of headers doesn't match length of widths array: %d != %d", len(headers), len(maxColumnWidths))
	}

	return headers, maxColumnWidths, nil
}

func Print(data string) error {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		return err
	}

	// top border
	fmt.Println(sliceToDashString(headerLengths))

	// headers
	fmt.Println(itemSliceToDashString(headers, headerLengths))

	// sep
	fmt.Println(sliceToDashString(headerLengths))

	for _, v := range lines[1:] {
		row := csvutils.CleanRow(strings.Split(v, ","))

		// data item
		fmt.Println(itemSliceToDashString(row, headerLengths))
	}

	// bottom border
	fmt.Println(sliceToDashString(headerLengths))

	return nil
}

func ReturnText(data string) (string, error) {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		return "", err
	}

	tableString := strings.Builder{}

	// top
	tableString.WriteString(fmt.Sprintf("%s\n", sliceToDashString(headerLengths)))

	// headers
	tableString.WriteString(fmt.Sprintf("%s\n", itemSliceToDashString(headers, headerLengths)))

	// sep
	tableString.WriteString(fmt.Sprintf("%s\n", sliceToDashString(headerLengths)))

	for _, v := range lines[1:] {
		row := csvutils.CleanRow(strings.Split(v, ","))

		// data item
		tableString.WriteString(fmt.Sprintf("%s\n", itemSliceToDashString(row, headerLengths)))
	}

	// end
	tableString.WriteString(fmt.Sprintf("%s\n", sliceToDashString(headerLengths)))

	return tableString.String(), err
}
