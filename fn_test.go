// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio_test

// Import standard library packages as well as tserr and tsfio
import (
	"fmt" // fmt
	"os"  // os
	"path/filepath"
	"testing" // testing

	"github.com/thorstenrie/tserr" // tserr
	"github.com/thorstenrie/tsfio" // tsfio
)

// TestBlockedDir tests if CheckFile returns an error for all
// blocked directories in invalDir. If it returns nil for
// a blocked directory, the test fails.
func TestBlockedDir(t *testing.T) {
	// Iterate test over all directories in invalDir
	for _, d := range tsfio.InvalDir {
		// Create test Filename p containing the blocked directory
		p := tsfio.Filename(d) + tsfio.Filename(os.PathSeparator) + testfile
		// If CheckFile returns nil, then the test fails
		if tsfio.CheckFile(p) == nil {
			t.Error(tserr.NilFailed(fmt.Sprintf("CheckFile of %v", d)))
		}
	}
}

// TestInvalFile tests if CheckFile returns an error for all
// blocked files in invalFile. If it returns nil for a blocked
// file, the test fails.
func TestInvalFile(t *testing.T) {
	// Iterate test over all files in invalFile
	for _, d := range tsfio.InvalFile {
		// If CheckFile returns nil, then the test fails
		if tsfio.CheckFile(d) == nil {
			t.Error(tserr.NilFailed(fmt.Sprintf("CheckFile of %v", d)))
		}
	}
}

// TestEmptyDir tests if CheckDir returns an error for an empty string
// as Directory. If it returns nil for the empty string as Directory,
// the test fails.
func TestEmptyDir(t *testing.T) {
	// If CheckDir of the empty string as Directory returns nil, then the test fails
	if err := tsfio.CheckDir(""); err == nil {
		t.Error(tserr.NilFailed("CheckDir"))
	}
}

// TestEmptyFile tests if CheckFile returns an error for an empty string
// as Filename. If it returns nil for the empty string as Filename,
// the test fails.
func TestEmptyFile(t *testing.T) {
	// If CheckFile of the empty string as Filename returns nil, then the test fails
	if err := tsfio.CheckFile(""); err == nil {
		t.Error(tserr.NilFailed("CheckFile"))
	}
}

// TestDir1 tests if CheckDir returns nil for a newly created temporary directory.
// If it returns an error, the test fails.
func TestDir1(t *testing.T) {
	// Create a temporary directory with name d
	d := tmpDir(t)
	// Test fails, if Checkdir returns an error for d
	if err := tsfio.CheckDir(d); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckDir of %v", d),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
	// Remove the temporary directory d
	rm(t, d)
}

// TestDir2 tests if CheckDir returns nil for a newly created temporary directory
// which is removed before the check. If it returns an error, the test fails.
func TestDir2(t *testing.T) {
	// Create a temporary directory with name d
	d := tmpDir(t)
	// Remove the temporary directory d before the check
	rm(t, d)
	// Test fails, if Checkdir returns an error for d
	if err := tsfio.CheckDir(d); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckDir of %v", d),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
}

// TestFile1 tests if CheckFile returns nil for a newly created temporary file.
// If it returns an error, the test fails.
func TestFile1(t *testing.T) {
	// Create a temporary file with name f
	f := tmpFile(t)
	// Test fails, if CheckFile returns an error for f
	if err := tsfio.CheckFile(f); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckFile of %v", f),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
	// Remove the temporary file f
	rm(t, f)
}

// TestFile2 tests if CheckFile returns nil for a newly created temporary file
// which is removed before the check. If it returns an error, the test fails.
func TestFile2(t *testing.T) {
	// Create a temporary file with name f
	f := tmpFile(t)
	// Remove the temporary file f before the check
	rm(t, f)
	// Test fails, if CheckFile returns an error for f
	if err := tsfio.CheckFile(f); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckFile of %v", f),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
}

// TestDir3 tests if CheckDir returns an error when it receives
// a temporary file. If it returns nil, the test fails.
func TestDir3(t *testing.T) {
	// Create a temporary file with name f
	f := tmpFile(t)
	// If CheckDir returns nil, the test fails
	if tsfio.CheckDir(tsfio.Directory(f)) == nil {
		t.Error(tserr.NilFailed(fmt.Sprintf("CheckDir of file %v", f)))
	}
	// Remove the temporary file f
	rm(t, f)
}

// TestFile3 tests if CheckFile returns an error when it receives a
// temporary directory. If it returns nil, the test fails.
func TestFile3(t *testing.T) {
	// Create a temporary directory with name d
	d := tmpDir(t)
	// If CheckFile returns nil, the test fails
	if tsfio.CheckFile(tsfio.Filename(d)) == nil {
		t.Error(tserr.NilFailed(fmt.Sprintf("CheckFile of directory %v", d)))
	}
	// Remove the temporary directory d
	rm(t, d)
}

// TestSprintf tests if Sprintf formats according to the format specifier and
// returns the resulting Filename or Directory. The test fails if not.
func TestSprintf(t *testing.T) {
	// Wanted string swant
	swant := os.TempDir() + string(os.PathSeparator) + testcase
	// Use Sprintf for a Directory d
	d := tsfio.Sprintf[tsfio.Directory]("%v%v%v", os.TempDir(), string(os.PathSeparator), testcase)
	// Use Sprintf for a Filename f
	f := tsfio.Sprintf[tsfio.Filename]("%v%v%v", os.TempDir(), string(os.PathSeparator), testcase)
	// If d does not equal swant, the test fails
	if tsfio.Directory(swant) != d {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{
			X: swant,
			Y: string(d),
		}))
	}
	// If f does not equal swant, the test fails
	if tsfio.Filename(swant) != f {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{
			X: swant,
			Y: string(f),
		}))
	}
}

func TestEmptyJoin(t *testing.T) {
	p, e := tsfio.Path("", "")
	if p != "" || e == nil {
		t.Error(tserr.NilFailed("Path"))
	}
}

func TestBlockedDirJoin(t *testing.T) {
	if len(tsfio.InvalDir) == 0 {
		t.Fatal(tserr.NilPtr())
	}
	d := tsfio.InvalDir[0]
	f := testfile
	p, e := tsfio.Path(d, f)
	if p != "" || e == nil {
		t.Error(tserr.NilFailed("Path"))
	}
}

func TestBlockedFileJoin(t *testing.T) {
	if len(tsfio.InvalFile) == 0 {
		t.Fatal(tserr.NilPtr())
	}
	d := testdir
	f := tsfio.InvalFile[0]
	p, e := tsfio.Path(d, f)
	if p != "" || e == nil {
		t.Error(tserr.NilFailed("Path"))
	}
}

func TestJoin(t *testing.T) {
	d := testdir
	f := testfile
	p, e := tsfio.Path(d, f)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Path", Fn: fmt.Sprintf("%v,%v", d, f), Err: e}))
	}
	r := tsfio.Filename(filepath.Join(string(d), string(f)))
	if p != r {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(p), Y: string(r)}))
	}
}
