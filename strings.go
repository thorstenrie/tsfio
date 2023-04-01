// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio

// Import standard library packages
import (
	"strings" // strings
	"unicode" // unicode
)

// Printable returns a copy of s with printable characters in s as defined by Go. Non-printable characters are dropped.
func Printable(s string) string {
	// Map all printable characters and drop non-printable characters
	return strings.Map(
		func(r rune) rune {
			// Return printable character
			if unicode.IsPrint(r) {
				return r
			}
			// Drop non-printable character
			return -1
		},
		s,
	)
}
