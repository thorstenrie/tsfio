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

func TestInvalFile(t *testing.T) {
	for _, d := range testDir {
		if CheckFile(Filename(d)) == nil {
			t.Errorf("%v should be an invalid file", d)
		}
	}
}

func TestDir1(t *testing.T) {
	d := tmpDir(t)
	if err := CheckDir(d); err != nil {
		t.Errorf("Directory check %v returned error %v, but error expected to be nil", d, err)
	}
}

func TestFile1(t *testing.T) {
	f := tmpFile(t)
	if err := CheckFile(f); err != nil {
		t.Errorf("Filename check %v returned error %v, but error expected to be nil", f, err)
	}
}

func TestDir2(t *testing.T) {
	d := tmpDir(t)
	rm(t, d)
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

func TestSprintf(t *testing.T) {
	a := "tsfio_1234"
	swant := os.TempDir() + string(os.PathSeparator) + a
	d := Sprintf[Directory]("%v%v%v", os.TempDir(), string(os.PathSeparator), a)
	f := Sprintf[Filename]("%v%v%v", os.TempDir(), string(os.PathSeparator), a)
	if (Directory(swant) != d) || (Filename(swant) != f) {
		t.Errorf("string(%v), Directory(%v), Filename(%v) not identical, but expected to be identical", swant, d, f)
	}
}

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
