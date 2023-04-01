// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio

// Import standard library packages and tserr
import (
	"testing" // testing

	"github.com/thorstenrie/tserr" // tserr
)

// TestPrintable tests if Printable returns printable characters in a copy and drops
// non-printable characters. If the added non-printable character is not dropped,
// the test fails.
func TestPrintable(t *testing.T) {
	// Add non-printable character to testcase string and retrieve printable characters in p
	p := Printable(testcase + "\n")
	// If testcase does not equal p, the test fails
	if testcase != p {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: testcase, Y: p}))
	}
}
