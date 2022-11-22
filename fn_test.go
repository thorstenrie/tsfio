package tsfio

import (
	"fmt"
	"os"
	"testing"

	"github.com/thorstenrie/tserr"
)

// TestBlockedDir tests if CheckFile returns an error for all
// blocked directories in invalDir. If it returns nil for
// a blocked directory, the test fails.
func TestBlockedDir(t *testing.T) {
	// Iterate test over all directories in invalDir
	for _, d := range invalDir {
		// Create test Filename p containing the blocked directory
		p := Filename(d) + Filename(os.PathSeparator) + testfile
		// If CheckFile returns nil, then test fails
		if CheckFile(p) == nil {
			t.Error(tserr.NilFailed(fmt.Sprintf("CheckFile of %v", d)))
		}
	}
}

func TestInvalFile(t *testing.T) {
	for _, d := range invalFile {
		if CheckFile(d) == nil {
			t.Error(tserr.NilFailed(fmt.Sprintf("CheckFile of %v", d)))
		}
	}
}

func TestEmptyDir(t *testing.T) {
	if err := CheckDir(""); err == nil {
		t.Error(tserr.NilFailed("CheckDir"))
	}
}

func TestEmptyFile(t *testing.T) {
	if err := CheckFile(""); err == nil {
		t.Error(tserr.NilFailed("CheckFile"))
	}
}

func TestDir1(t *testing.T) {
	d := tmpDir(t)
	if err := CheckDir(d); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckDir of %v", d),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
	rm(t, d)
}

func TestFile1(t *testing.T) {
	f := tmpFile(t)
	if err := CheckFile(f); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckFile of %v", f),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
	rm(t, f)
}

func TestDir2(t *testing.T) {
	d := tmpDir(t)
	rm(t, d)
	if err := CheckDir(d); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckDir of %v", d),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
}

func TestFile2(t *testing.T) {
	f := tmpFile(t)
	rm(t, f)
	if err := CheckFile(f); err != nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     fmt.Sprintf("CheckFile of %v", f),
			Actual: fmt.Sprint(err),
			Want:   "nil",
		}))
	}
}

func TestDir3(t *testing.T) {
	f := tmpFile(t)
	if CheckDir(Directory(f)) == nil {
		t.Error(tserr.NilFailed(fmt.Sprintf("CheckDir of file %v", f)))
	}
	rm(t, f)
}

func TestFile3(t *testing.T) {
	d := tmpDir(t)
	if CheckFile(Filename(d)) == nil {
		t.Error(tserr.NilFailed(fmt.Sprintf("CheckFile of directory %v", d)))
	}
	rm(t, d)
}

func TestSprintf(t *testing.T) {
	a := "tsfio_1234"
	swant := os.TempDir() + string(os.PathSeparator) + a
	d := Sprintf[Directory]("%v%v%v", os.TempDir(), string(os.PathSeparator), a)
	f := Sprintf[Filename]("%v%v%v", os.TempDir(), string(os.PathSeparator), a)
	if Directory(swant) != d {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{
			X: swant,
			Y: string(d),
		}))
	}
	if Filename(swant) != f {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{
			X: swant,
			Y: string(f),
		}))
	}
}
