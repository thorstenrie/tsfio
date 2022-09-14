package tsfio

import (
	"os"
	"testing"
)

var (
	testDir [8]Directory = invalDir
)

func TestInvalDir(t *testing.T) {
	for _, d := range testDir {
		if CheckDir(d) == nil {
			t.Errorf("%v should be an invalid dir", d)
		}
	}
}

func TestDir1(t *testing.T) {
	d := tmpDir(t)
	if CheckDir(d) != nil {
		t.Errorf("Directory check %v returned error, but error expected to be nil", d)
	}
}

func TestDir2(t *testing.T) {
	d := tmpDir(t)
	if err := os.Remove(string(d)); err != nil {
		t.Fatalf("could not remove temp dir %v", d)
	}
	if CheckDir(d) != nil {
		t.Errorf("Directory check %v returned error, but error expected to be nil", d)
	}
}

func TestDir3(t *testing.T) {
	f := tmpFile(t)
	if CheckDir(Directory(f)) == nil {
		t.Errorf("Directory check %v returned nil, but error expected", f)
	}
}

func tmpDir(t *testing.T) Directory {
	d, err := os.MkdirTemp("", "tsfio_*")
	if err != nil {
		t.Fatalf("could not create temp dir %v", d)
	}
	return Directory(d)
}

func tmpFile(t *testing.T) Filename {
	f, err := os.CreateTemp("", "tsfio_*")
	if err != nil {
		t.Fatalf("could not create temp file %v", f.Name())
	}
	return Filename(f.Name())
}
