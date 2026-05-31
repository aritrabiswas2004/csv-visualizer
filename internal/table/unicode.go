/************* SPDX License Identifier: MIT ************
 * Visualizer - Render unicode version of table like TUIs
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

func rowBorderString(strBuilder *strings.Builder, lengths []int, startChar string, midChar string, endChar string) {
	strBuilder.WriteString(startChar)

	for i, length := range lengths {
		if i == len(lengths)-1 {
			strBuilder.WriteString(fmt.Sprintf("─%s─%s", repeat("─", length), endChar))
		} else {
			strBuilder.WriteString(fmt.Sprintf("─%s─%s", repeat("─", length), midChar))
		}
	}
}

func unicodeSliceToDashString(lengths []int, pos string) string {
	formatString := strings.Builder{}

	// I am not a fan of the rounded corners
	if pos == "top" {
		rowBorderString(&formatString, lengths, "╭", "┬", "╮")
	} else if pos == "sep" {
		rowBorderString(&formatString, lengths, "├", "┼", "┤")
	} else {
		rowBorderString(&formatString, lengths, "╰", "┴", "╯")
	}

	return formatString.String()
}

func unicodeItemSliceToDashString(items []string, itemWidths []int) string {
	formatString := strings.Builder{}

	formatString.WriteString("│")

	for i, item := range items {
		formatString.WriteString(fmt.Sprintf(" %-*s │", itemWidths[i], item))
	}

	return formatString.String()
}

func PrintUnicode(data string) error {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		return err
	}

	// top border
	fmt.Println(unicodeSliceToDashString(headerLengths, "top"))

	// headers
	fmt.Println(unicodeItemSliceToDashString(headers, headerLengths))

	// sep
	fmt.Println(unicodeSliceToDashString(headerLengths, "sep"))

	for _, v := range lines[1:] {
		row := csvutils.CleanRow(strings.Split(v, ","))

		// data item
		fmt.Println(unicodeItemSliceToDashString(row, headerLengths))
	}

	// bottom border
	fmt.Println(unicodeSliceToDashString(headerLengths, "bottom"))

	return nil
}

func ReturnUnicode(data string) (string, error) {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	headers, headerLengths, err := getDataHeadersAndMaxWidths(lines)
	if err != nil {
		return "", err
	}

	tableString := strings.Builder{}

	// top
	tableString.WriteString(fmt.Sprintf("%s\n", unicodeSliceToDashString(headerLengths, "top")))

	// headers
	tableString.WriteString(fmt.Sprintf("%s\n", unicodeItemSliceToDashString(headers, headerLengths)))

	// sep
	tableString.WriteString(fmt.Sprintf("%s\n", unicodeSliceToDashString(headerLengths, "sep")))

	for _, v := range lines[1:] {
		row := csvutils.CleanRow(strings.Split(v, ","))

		// data item
		tableString.WriteString(fmt.Sprintf("%s\n", unicodeItemSliceToDashString(row, headerLengths)))
	}

	// end
	tableString.WriteString(fmt.Sprintf("%s\n", unicodeSliceToDashString(headerLengths, "bottom")))

	return tableString.String(), err
}
