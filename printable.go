// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio

// Import Go standard packages and tserr
import (
	"strings" // strings
	"unicode" // unicode

	"github.com/thorstenrie/tserr" // tserr
)

// Printable returns all printable runes of string a as defined by Go.
// All non-printable runes of string a are dropped.
func Printable(a string) string {
	// Return a copy of string a with all non-printable characters dropped
	return strings.Map(func(r rune) rune {
		// If rune r is printable, return r, otherwise drop
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, a)
}

// IsPrintable returns true, if slice of strings a only consists of
// printable strings. If one or more strings of slice contain
// a non-printable character, it returns false. If a is nil or
// has zero length, it returns false and an error.
func IsPrintable(a []string) (bool, error) {
	// Return false and an an error, if a is nil or has zero length
	if (a == nil) || (len(a) == 0) {
		return false, tserr.Empty("slice of strings a")
	}
	// Iterate all elements of the slice
	for _, s := range a {
		// Compare element s of slice a with a copy of s only containing printable runes of s
		if s != Printable(s) {
			// Return false if element s of slice a contains non-printable characters
			return false, nil
		}
	}
	// Otherwise, return true
	return true, nil
}

// RuneToPrintable returns rune r as a string, if it is printable. Otherwise, it returns an empty string.
func RuneToPrintable(r rune) string {
	return Printable(string(r))
}
