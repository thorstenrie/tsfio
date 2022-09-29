package tsfio

import (
	"os"
	"testing"
)

const (
	testprefix string   = "tsfio_*"
	testcase   string   = "test"
	testfile   Filename = "test"
)

func tmpDir(t *testing.T) Directory {
	d, err := os.MkdirTemp("", testprefix)
	if err != nil {
		t.Fatal(errOp("create temp dir", d, err))
	}
	return Directory(d)
}

func tmpFile(t *testing.T) Filename {
	f, err := os.CreateTemp("", testprefix)
	if err != nil {
		t.Fatal(errOp("create temp file", f.Name(), err))
	}
	return Filename(f.Name())
}

func rm[T fio](t *testing.T, a T) {
	if err := os.Remove(string(a)); err != nil {
		t.Fatal(errOp("remove", string(a), err))
	}
}
