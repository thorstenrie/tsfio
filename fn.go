package tsfio

import (
	"fmt"
	"os"
)

var (
	invalDir [8]Directory = [8]Directory{
		"",
		"/",
		"/boot",
		"/dev",
		"/lost+found",
		"/media",
		"/mnt",
		"/proc",
	}
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
			return fmt.Errorf("%v is not a %v", f, w)
		}
	} else {
		if os.IsNotExist(err) {
			return nil
		} else {
			return fmt.Errorf("%v check of %v failed: %w", w, f, err)
		}
	}
}

func checkInval[T fio](f T) error {
	for _, i := range invalDir {
		if string(i) == string(f) {
			return fmt.Errorf("%v not allowed", f)
		}
	}
	return nil
}

func Sprintf[T fio](format string, a ...any) T {
	return T(fmt.Sprintf(format, a...))
}

// Todo write test
func errChk[T fio](f T, err error) error {
	return fmt.Errorf("check %v failed: %w", f, err)
}

// Todo write test
func errFio(op string, f Filename, err error) error {
	return fmt.Errorf("%v file %v failed: %w", op, f, err)
}

// Todo write test
func errDir(op string, d Directory, err error) error {
	return fmt.Errorf("%v dir %v failed: %w", op, d, err)
}
