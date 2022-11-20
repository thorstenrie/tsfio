package tsfio

// Import standard library packages and tserr
import (
	"fmt"           // fmt
	"os"            // os
	"path/filepath" // path/filepath
	"strings"       // strings

	"github.com/thorstenrie/tserr" //tserr
)

// Interface fio is constrained to type Filename and Directory
type fio interface {
	Filename | Directory
}

// A Filename is the name of a file and may contain its path
type Filename string

// A Directory is the name of a directory and may contain its path
type Directory string

// CheckFile performs checks on file f and returns an error if
//   - f is an empty string
//   - f contains a blocked directory or filename
//   - f is an existing directory, not a file
//   - os.Stat returns an error when retrieving FileInfo
//
// Otherwise it returns nil.
func CheckFile(f Filename) error {
	return checkWrapper(f, false)
}

// CheckDir performs checks on directory d and returns an error if
//   - d is an empty string
//   - d contains a blocked directory or filename
//   - d is an existing file, not a directory
//   - os.Stat returns an error when retrieving FileInfo
//
// Otherwise it returns nil.
func CheckDir(d Directory) error {
	return checkWrapper(d, true)
}

// checkWrapper performs checks on a file or directory using fio as type parameter.
// It returns an error, if any check fails. Otherwise it returns nil.
func checkWrapper[T fio](f T, dir bool) error {
	// Return an error if f is an empty string
	if f == "" {
		return tserr.Empty(string(f))
	}
	// Return an error if f contains a blocked directory or filename
	if err := checkInval(f); err != nil {
		return err
	}
	// Retrieve FileInfo of f
	i, err := os.Stat(string(f))
	// Set w to file or directory
	w := "file"
	if dir {
		w = "directory"
	}
	// If Stat returns no error, then check if expected type matches type of f
	if err == nil {
		if i.IsDir() == dir {
			return nil
		} else {
			return tserr.TypeNotMatching(&tserr.TypeNotMatchingArgs{Act: string(f), Want: w})
		}
	} else {
		// If Stat returns an error reporting f does not exist, return nil
		if os.IsNotExist(err) {
			return nil
			// If Stat returns any other error, return the error
		} else {
			return tserr.Check(&tserr.CheckArgs{F: string(f), Err: err})
		}
	}
}

func checkInval[T fio](f T) error {
	fc := filepath.Clean(string(f))
	for _, i := range invalFile {
		ic := filepath.Clean(string(i))
		if ic == fc {
			return tserr.Forbidden(string(f))
		}
	}
	for _, i := range invalDir {
		ic := filepath.Clean(string(i))
		if strings.HasPrefix(fc, ic) {
			return tserr.Forbidden(string(i))
		}
	}
	return nil
}

func Sprintf[T fio](format string, a ...any) T {
	return T(fmt.Sprintf(format, a...))
}
