package tsfio

import (
	"os"
	"testing"
	"time"

	"github.com/thorstenrie/tserr"
)

// The testcases use these tokens
const (
	testprefix string   = "tsfio_*" // mostly used as prefix for temporary files or directories
	testcase   string   = "test"    // test string
	testfile   Filename = "test"    // test Filename
)

// tmpDir creates a new temporary directory in the default directory for temporary files
// with the prefix testprefix and a random string to the end. In case of an error
// the execution stops.
func tmpDir(t *testing.T) Directory {
	// Create the temporary directory
	d, err := os.MkdirTemp("", testprefix)
	// Stop execution in case of an error
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "create temp dir", Fn: d, Err: err}))
	}
	// Return the temporary Directory
	return Directory(d)
}

// tmpFile creates a new tmporary file in the default directory for temporary files
// with the prefix testprefix and a random string to the end. In case of an error
// the execution stops.
func tmpFile(t *testing.T) Filename {
	// Create the temporary file
	f, err := os.CreateTemp("", testprefix)
	// Stop execution in case of an error
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "create temp file", Fn: f.Name(), Err: err}))
	}
	// Return the temporary Filename
	return Filename(f.Name())
}

// rm removes file named Filename a or empty directory Directory a. In case of an error
// execution stops.
func rm[T fio](t *testing.T, a T) {
	// Remove file or empty directory
	if err := os.Remove(string(a)); err != nil {
		// Stop execution in case of an error
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "Remove", Fn: string(a), Err: err}))
	}
}

// modTime returns the modification time of the file with Filename fn.
// In case of an error, it stops execution.
func modTime(t *testing.T, fn Filename) time.Time {
	// Retrieve the FileInfo structure from fn in fi
	fi, err := os.Stat(string(fn))
	// Stop execution in case of an error
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{
			Op:  "Stat",
			Fn:  string(fn),
			Err: err,
		}))
	}
	// Retrieve modification time from FileInfo fi
	t1 := fi.ModTime()
	// Return the modification time of file with Filename fn
	return t1
}
