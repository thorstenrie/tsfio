package tsfio

import (
	"fmt"
	"os"

	"github.com/thorstenrie/tslog"
)

func WriteStr(f *os.File, s string) error {
	if f == nil {
		return fmt.Errorf("nil pointer")
	}
	if _, err := f.WriteString(s); err != nil {
		f.Close()
		return fmt.Errorf("write string to file %v failed: %w", f.Name(), err)
	}
	return nil
}

func AppendFile(fileAppend Filename, fileIn Filename) error {
	if err := CheckFile(fileAppend); err != nil {
		return errChk(fileAppend, err)
	}
	if err := CheckFile(fileIn); err != nil {
		return errChk(fileIn, err)
	}
	f := OpenFile(fileAppend)
	out := ReadFile(fileIn)
	if _, err := f.Write(out); err != nil {
		f.Close()
		return fmt.Errorf("append file %v to file %v failed: %w", fileIn, f.Name(), err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close file %v failed: %w", f.Name(), err)
	}
	return nil
}

func existsFile(fn Filename) (bool, error) {
	if err := CheckFile(fn); err != nil {
		return false, errChk(fn, err)
	}
	_, err := os.Stat(string(fn))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("FileInfo of file %v failed: %w", fn, err)
}

func WriteSingleStr(fn Filename, s string) error {
	ResetFile(fn)
	f := OpenFile(fn)
	WriteStr(f, s)
	f.Close()
	return nil
}

func ReadFile(f Filename) []byte {
	CheckFile(f)
	b, err := os.ReadFile(string(f))
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error read file: %w", err))
	}
	return b
}

func RemoveFile(f Filename) {
	CheckFile(f)
	if b, _ := existsFile(f); b {
		err := os.Remove(string(f))
		if err != nil {
			tslog.E.Panic(fmt.Errorf("fatal error with removing file %v: %w", f, err))
		}
	} else {
		tslog.E.Panic(fmt.Errorf("fatal error removing file %v: does not exist", f))
	}
}

func CreateDir(d Directory) {
	CheckDir(d)
	err := os.MkdirAll(string(d), 0755)
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error with creating dir: %w", err))
	}
}

func ResetFile(f Filename) {
	if b, _ := existsFile(f); b {
		touchFile(f)
	}
	err := os.Truncate(string(f), 0)
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error truncate file: %w", err))
	}
}

func touchFile(f Filename) {
	CheckFile(f)
	(OpenFile(f)).Close()
}

/*func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error closing file: %w", err))
	}
}*/

func OpenFile(filename Filename) *os.File {
	CheckFile(filename)
	f, err := os.OpenFile(string(filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error open file: %w", err))
	}
	return f
}
