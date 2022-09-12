package tsfio

import (
	"fmt"
	"os"

	"github.com/thorstenrie/tslog"
)

func WriteFile(f *os.File, s string) {
	_, err := f.WriteString(s)
	if err != nil {
		CloseFile(f)
		tslog.E.Panic(fmt.Errorf("fatal error write file: %w", err))
	}
}

func AppendFile(fileAppend Filename, fileIn Filename) {
	CheckFile(fileAppend)
	CheckFile(fileIn)
	f := OpenFile(fileAppend)
	out := ReadFile(fileIn)
	_, err := f.Write(out)
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error write to file: %w", err))
	}
	CloseFile(f)
}

func WriteSingleStringToFile(filename Filename, s string) {
	ResetFile(filename)
	f := OpenFile(filename)
	WriteFile(f, s)
	CloseFile(f)
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
	if existsFile(f) {
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

func existsFile(f Filename) bool {
	CheckFile(f)
	s, err := os.Stat(string(f))
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error with FileInfo of file: %w", err))
	}
	if s.IsDir() {
		tslog.E.Panic(fmt.Errorf("fatal error with file %v: is s directory", f))
	}
	return true
}

func ResetFile(f Filename) {
	if !existsFile(f) {
		touchFile(f)
	}
	err := os.Truncate(string(f), 0)
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error truncate file: %w", err))
	}
}

func touchFile(f Filename) {
	CheckFile(f)
	CloseFile(OpenFile(f))
}

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error closing file: %w", err))
	}
}

func OpenFile(filename Filename) *os.File {
	CheckFile(filename)
	f, err := os.OpenFile(string(filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		tslog.E.Panic(fmt.Errorf("fatal error open file: %w", err))
	}
	return f
}
