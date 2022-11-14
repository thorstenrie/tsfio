// Package tsfio provides a simple API for file input output based on the standard library.
//
// The tsfio package is a supplement to the standard library and supplies additional
// functions for file input output operations, e.g., appending one file to another file.
// Also, all file input output operations on Linux and Windows system directories or
// files are blocked (see inval_unix.go and inval_win.go) and an error is returned.
// Default flags and file mode is used when opening files, creating files or directories
// and when writing to files (with exceptions documented in the function descriptions)
//
//   - Files are opened read-write (os.O_RDWR).
//   - Data is appended when writing to file (os.O_APPEND).
//   - A file is created if it does not exist (os.O_CREATE).
//   - File mode and permission bits are 0644.
//   - Directory mode and permissions bits are 0755.
//
// If an API call is not successful, a tserr error in JSON format is returned.
//
// Copyright (c) 2022 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio

// Import standard library packages and tserr
import (
	"fmt"   // fmt
	"io/fs" // ios/fs
	"os"    // os
	"time"  // time

	"github.com/thorstenrie/tserr" // tserr
)

// The constants hold the default flags and file mode
const (
	// Int flags holds the default for files opened read-write, created if it does not exist, data is appended
	flags int = os.O_APPEND | os.O_CREATE | os.O_RDWR
	// FileMode fperm holds the default file mode and permission bits
	fperm fs.FileMode = 0644
	// FileMode dperm holds the default directory mode and permissions bits
	dperm fs.FileMode = 0755
)

// Note: All external functions must contain a CheckFile or CheckDir call at the beginning.
// This can't be tested, because a failed test could break the testing environment.

// OpenFile opens the named file fn with default flags and permission bits. If the file does
// not exist, it is created. If opened successful, the file is returned and can be used for
// file input output and error is nil.
func OpenFile(fn Filename) (*os.File, error) {
	// Return nil and error in case fn contains a blocked directory or filename
	if e := CheckFile(fn); e != nil {
		return nil, tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	// Open file with default flags and permission bits
	f, err := os.OpenFile(string(fn), flags, fperm)
	// In case of an error, return nil and error
	if err != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: err})
	}
	// If successfull, return file for file input output
	return f, nil
}

// CloseFile closes f and f is unusable for file input output. An error is returned
// if f has already been closed.
func CloseFile(f *os.File) error {
	// Return error in case f is nil
	if f == nil {
		return tserr.NilPtr()
	}
	// Close f
	if e := f.Close(); e != nil {
		// Return error if not successful
		return tserr.Op(&tserr.OpArgs{Op: "Close", Fn: f.Name(), Err: e})
	}
	// If successful, return nil
	return nil
}

func WriteStr(fn Filename, s string) error {
	if e := CheckFile(fn); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	f, err := OpenFile(fn)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: err})
	}
	if _, e := f.WriteString(s); e != nil {
		f.Close()
		return tserr.Op(&tserr.OpArgs{Op: "write string to", Fn: string(fn), Err: e})
	}
	return nil
}

func TouchFile(fn Filename) error {
	if e := CheckFile(fn); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	b, erre := ExistsFile(fn)
	if erre != nil {
		return tserr.Op(&tserr.OpArgs{Op: "ExistsFile", Fn: string(fn), Err: erre})
	}
	if b {
		t := time.Now().Local()
		if e := os.Chtimes(string(fn), t, t); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "Chtimes", Fn: string(fn), Err: e})
		}
	} else {
		f, erro := OpenFile(fn)
		if erro != nil {
			return tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: erro})
		}
		if e := f.Close(); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "Close", Fn: string(fn), Err: e})
		}
	}
	return nil
}

func WriteSingleStr(fn Filename, s string) error {
	if e := CheckFile(fn); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	if e := ResetFile(fn); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "ResetFile", Fn: string(fn), Err: e})
	}
	if e := WriteStr(fn, s); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: fmt.Sprintf("write string %v to", s), Fn: string(fn), Err: e})
	}
	return nil
}

func ReadFile(f Filename) ([]byte, error) {
	if e := CheckFile(f); e != nil {
		return nil, tserr.Check(&tserr.CheckArgs{F: string(f), Err: e})
	}
	b, err := os.ReadFile(string(f))
	if err != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(f), Err: err})
	}
	return b, nil
}

type Append struct {
	fileA Filename
	fileI Filename
}

func AppendFile(a *Append) error {
	if a == nil {
		return fmt.Errorf("nil pointer")
	}
	if e := CheckFile(a.fileA); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(a.fileA), Err: e})
	}
	if e := CheckFile(a.fileI); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(a.fileI), Err: e})
	}
	f, erro := OpenFile(a.fileA)
	if erro != nil {
		return tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(a.fileA), Err: erro})
	}
	out, errr := ReadFile(a.fileI)
	if errr != nil {
		return tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(a.fileI), Err: errr})
	}
	if _, e := f.Write(out); e != nil {
		f.Close()
		return tserr.Op(&tserr.OpArgs{Op: fmt.Sprintf("append file %v to", a.fileI), Fn: string(a.fileA), Err: e})
	}
	if e := f.Close(); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "Close", Fn: string(a.fileA), Err: e})
	}
	return nil
}

func ExistsFile(fn Filename) (bool, error) {
	if err := CheckFile(fn); err != nil {
		return false, tserr.Check(&tserr.CheckArgs{F: string(fn), Err: err})
	}
	_, err := os.Stat(string(fn))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, tserr.Op(&tserr.OpArgs{Op: "FileInfo (Stat) of", Fn: string(fn), Err: err})
}

func RemoveFile(f Filename) error {
	if e := CheckFile(f); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(f), Err: e})
	}
	b, err := ExistsFile(f)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "check if exists", Fn: string(f), Err: err})
	}
	if b {
		e := os.Remove(string(f))
		if e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "Remove", Fn: string(f), Err: err})
		}
	} else {
		return tserr.NotExistent(string(f))
	}
	return nil
}

func CreateDir(d Directory) error {
	if e := CheckDir(d); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(d), Err: e})
	}
	err := os.MkdirAll(string(d), dperm)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "make directory", Fn: string(d), Err: err})
	}
	return nil
}

func ResetFile(fn Filename) error {
	if e := CheckFile(fn); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	b, err := ExistsFile(fn)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "check if exists", Fn: string(fn), Err: err})
	}
	if !b {
		if e := TouchFile(fn); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "TouchFile", Fn: string(fn), Err: e})
		}
	}
	err = os.Truncate(string(fn), 0)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "Truncate", Fn: string(fn), Err: err})
	}
	return nil
}
