/************* SPDX License Identifier: MIT ************
 * Visualizer - Utils and helpers related to CSV
 *
 * Copyright (C) 2026 Aritra Biswas
 *
 * Author: Aritra Biswas <aritrabb@gmail.com>
 *
********************************************************/

package csvutils

import "strings"

func CleanRow(row []string) []string {
	out := make([]string, 0, len(row))
	for _, v := range row {
		out = append(out, strings.ReplaceAll(v, `"`, ""))
	}
	return out
}
