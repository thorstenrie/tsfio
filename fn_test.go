package tsfio

import (
	"fmt"
	"os"
	"testing"
)

func TestBlockedDir(t *testing.T) {
	for _, d := range invalDir {
		p := Filename(d) + Filename(os.PathSeparator) + testfile
		if CheckFile(p) == nil {
			t.Errorf("%v should be an invalid dir", d)
		}
	}
}

func TestInvalFile(t *testing.T) {
	for _, d := range invalFile {
		if CheckFile(d) == nil {
			t.Errorf("%v should be an invalid file", d)
		}
	}
}

func TestEmptyDir(t *testing.T) {
	if err := CheckDir(""); err == nil {
		t.Errorf("Directory check returned nil, but error expected for empty directory name")
	}
}

func TestEmptyFile(t *testing.T) {
	if err := CheckFile(""); err == nil {
		t.Errorf("File check returned nil, but error expected for empty file name")
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

func TestFile2(t *testing.T) {
	f := tmpFile(t)
	rm(t, f)
	if CheckFile(f) != nil {
		t.Errorf("File check %v returned error, but error expected to be nil", f)
	}
}

func TestDir3(t *testing.T) {
	f := tmpFile(t)
	if CheckDir(Directory(f)) == nil {
		t.Errorf("Directory check %v returned nil, but error expected", f)
	}
}

func TestFile3(t *testing.T) {
	d := tmpDir(t)
	if CheckFile(Filename(d)) == nil {
		t.Errorf("File check %v returned nil, but error expected", d)
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

func TestErrChkNil(t *testing.T) {
	if errChk(testcase, nil) != nil {
		t.Errorf("errChk returned error, but error expected to be nil")
	}
}

func TestErrChk(t *testing.T) {
	if errChk(testcase, fmt.Errorf("test")) == nil {
		t.Errorf("errChk returned nil, but error expected")
	}
}

func TestErrFioNil(t *testing.T) {
	if errOp(testcase, testcase, nil) != nil {
		t.Errorf("errFio returned error, but error expected to be nil")
	}
}

func TestErrFio(t *testing.T) {
	if errOp(testcase, testcase, fmt.Errorf(testcase)) == nil {
		t.Errorf("errFio returned nil, but error expected")
	}
}
