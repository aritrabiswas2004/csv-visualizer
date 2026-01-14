/*
All the program is in main.go only for now
The module structure will come later
*/

/*
Main Working of Visualizer

Copyright (C) 2026 Aritra Biswas

Author: Aritra Biswas <aritrabb@gmail.com>
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func repeat(char string, count int) string {
	result := ""

	for i := 0; i < count; i++ {
		result += char
	}

	return result
}

func sliceToDashString(lengths []int) string {
	formatString := "+"

	for _, length := range lengths {
		formatString += fmt.Sprintf("-%s-+", repeat("-", length))
	}

	return formatString
}

func itemSliceToDashString(items []string, itemWidths []int) string {
	formatString := "|"

	for i, item := range items {
		formatString += fmt.Sprintf(" %-*s |", itemWidths[i], item)
	}

	return formatString
}

func collectHeaderLengths(hdrs []string) []int {
	var sizes []int

	for _, hdr := range hdrs {
		sizes = append(sizes, len(hdr))
	}

	return sizes
}

func capitalize(arr []string) []string {
	var capitalized []string

	for _, str := range arr {
		capitalized = append(capitalized, fmt.Sprintf("%s%s", strings.ToUpper(string(str[0])), str[1:]))
	}

	return capitalized
}

// Basically returns without "" and stuff
func cleanCSVRow(row []string) []string {
	var cleaned []string

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

func main() {
	data, err := os.ReadFile("test.csv")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

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
