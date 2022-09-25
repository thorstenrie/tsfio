package tsfio

import (
	"fmt"
	"os"
	"path/filepath"
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
		return fmt.Errorf("file or directory name cannot be empty")
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
		if filepath.Clean(string(i)) == filepath.Clean(string(f)) {
			return fmt.Errorf("%v not allowed", f)
		}
	}
	return nil
}

func Sprintf[T fio](format string, a ...any) T {
	return T(fmt.Sprintf(format, a...))
}

func errChk[T fio](f T, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("check %v failed: %w", f, err)
}

func errFio[T fio](op string, f T, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%v %v failed: %w", op, f, err)
}
