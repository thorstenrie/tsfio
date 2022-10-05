package tsfio

import (
	"os"
	"testing"

	"github.com/thorstenrie/tserr"
)

const (
	testprefix string   = "tsfio_*"
	testcase   string   = "test"
	testfile   Filename = "test"
)

func tmpDir(t *testing.T) Directory {
	d, err := os.MkdirTemp("", testprefix)
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "create temp dir", Fn: d, Err: err}))
	}
	return Directory(d)
}

func tmpFile(t *testing.T) Filename {
	f, err := os.CreateTemp("", testprefix)
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "create temp file", Fn: f.Name(), Err: err}))
	}
	return Filename(f.Name())
}

func rm[T fio](t *testing.T, a T) {
	if err := os.Remove(string(a)); err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "remove", Fn: string(a), Err: err}))
	}
}
