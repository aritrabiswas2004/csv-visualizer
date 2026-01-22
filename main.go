/*
All the program is in main.go only for now
The module structure will come later
*/

/************* SPDX License Identifier: MIT License ************
 * Visualizer - Main source file for the Visualizer
 *
 * Copyright (C) 2026 Aritra Biswas
 *
 * Author: Aritra Biswas <aritra.biswas@ru.nl>
*****************************************************************/

package main

import (
	"fmt"
	"os"
	"strings"

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

// Basically returns without "" and stuff
func cleanCSVRow(row []string) []string {
	cleaned := make([]string, 0, len(row))

	for _, item := range row {
		cleaned = append(cleaned, strings.Replace(item, "\"", "", -1))
	}

	return cleaned
}

func getDataHeadersAndMaxWidths(lines []string) ([]string, []int, error) {
	headers := capitalize(strings.Split(lines[0], ","))
	maxColumnWidths := collectHeaderLengths(headers)

	for _, v := range lines {
		row := cleanCSVRow(strings.Split(v, ","))

		for idx, val := range row {
			maxColumnWidths[idx] = max(maxColumnWidths[idx], len(val))
		}
	}

	if len(headers) != len(maxColumnWidths) {
		return []string{}, []int{}, fmt.Errorf("length of headers doesn't match length of widths array: %d != %d", len(headers), len(maxColumnWidths))
	}

	return headers, maxColumnWidths, nil
}

func printTerminalTable(data string) {
	lines := strings.Split(data, "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		panic(err)
	}

	// top border
	fmt.Println(sliceToDashString(headerLengths))

	// headers
	fmt.Println(itemSliceToDashString(headers, headerLengths))

	// sep
	fmt.Println(sliceToDashString(headerLengths))

	for _, v := range lines[1:] {
		row := cleanCSVRow(strings.Split(v, ","))

		// data item
		fmt.Println(itemSliceToDashString(row, headerLengths))
	}

	// bottom border
	fmt.Println(sliceToDashString(headerLengths))
}

func main() {
	data, err := os.ReadFile("test.csv")
	if err != nil {
		panic(err)
	}

	printTerminalTable(string(data))
}
