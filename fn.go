package tsfio

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thorstenrie/tserr"
)

type fio interface {
	Filename | Directory
}

type Filename string

type Directory string

func CheckFile(f Filename) error {
	return checkWrapper(f, false)
}

func CheckDir(d Directory) error {
	return checkWrapper(d, true)
}

func checkWrapper[T fio](f T, dir bool) error {
	if f == "" {
		return tserr.Empty(string(f))
	}
	if err := checkInval(f); err != nil {
		return err
	}
	i, err := os.Stat(string(f))
	w := "file"
	if dir {
		w = "directory"
	}
	if err == nil {
		if i.IsDir() == dir {
			return nil
		} else {
			return tserr.TypeNotMatching(&tserr.TypeNotMatchingArgs{Act: string(f), Want: w})
		}
	} else {
		if os.IsNotExist(err) {
			return nil
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
			return fmt.Errorf("file operation on %v not allowed", f)
		}
	}
	for _, i := range invalDir {
		ic := filepath.Clean(string(i))
		if strings.HasPrefix(fc, ic) {
			return fmt.Errorf("directory %v in %v blocked by default", i, f)
		}
	}
	return nil
}

func Sprintf[T fio](format string, a ...any) T {
	return T(fmt.Sprintf(format, a...))
}
