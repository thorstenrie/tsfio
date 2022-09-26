package tsfio

import (
	"os"
	"testing"
)

const (
	testcase string   = "test"
	testfile Filename = "test"
)

func tmpDir(t *testing.T) Directory {
	d, err := os.MkdirTemp("", "tsfio_*")
	if err != nil {
		t.Fatalf("could not create temp dir %v: %v", d, err)
	}
	return Directory(d)
}

func tmpFile(t *testing.T) Filename {
	f, err := os.CreateTemp("", "tsfio_*")
	if err != nil {
		t.Fatalf("could not create temp file %v: %v", f.Name(), err)
	}
	return Filename(f.Name())
}

func rm[T fio](t *testing.T, a T) {
	if err := os.Remove(string(a)); err != nil {
		t.Fatalf("could not remove %v: %v", a, err)
	}
}
