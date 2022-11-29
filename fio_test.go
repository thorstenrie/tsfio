package tsfio

// Import standard library packages and tserr
import (
	"fmt"     // fmt
	"os"      // os
	"testing" // testing
	"time"    // time

	"github.com/thorstenrie/tserr" // tserr
)

// TestOpenFile1 tests OpenFile to open an existing temporary
// file and to return *os.File. If OpenFile returns an error or if *os.File is nil,
// the test fails.
func TestOpenFile1(t *testing.T) {
	testOpenFile(t, false)
}

// TestOpenFile2 tests OpenFile to open a new temporary file, not
// existing yet and to return *os.File. If OpenFile returns an error or
// if *os.File is nil, the test fails.
func TestOpenFile2(t *testing.T) {
	testOpenFile(t, true)
}

// testOpenFile is called by Test functions to test OpenFile. If r is true,
// OpenFile opens a non-existing, new temporary file. If r is false,
// OpenFile opens an existing temporary file. If OpenFile returns an error or
// if *os.File is nil, the test fails.
func testOpenFile(t *testing.T, r bool) {
	// Create a temporary file with Filename fn
	fn := tmpFile(t)
	// If r is true, remove file fn
	if r {
		rm(t, fn)
	}
	// Open file fn
	f, err := OpenFile(fn)
	// If OpenFile returns an error, the test fails
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: err}))
	}
	// If *os.File is nil, the test fails
	if f == nil {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "OpenFile", Actual: "nil", Want: "*os.File"}))
	}
	// Close file fn
	if e := CloseFile(f); e != nil {
		// If CloseFile fails, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CloseFile", Fn: string(fn), Err: e}))
	}
	// Remove file fn
	rm(t, fn)
}

// TestOpenFileEmpty tests OpenFile in case it retrieves the empty string "".
// The test fails if OpenFile returns nil instead of an error.
func TestOpenFileEmpty(t *testing.T) {
	// If OpenFile returns nil, the test fails
	if _, err := OpenFile(""); err == nil {
		t.Error(tserr.NilFailed("OpenFile"))
	}
}

// TestCreateDirEmpty tests CreateDor in case it retrieves the empty string "".
// The test fails if CreateDir returns nil instead of an error.
func TestCreateDirEmpty(t *testing.T) {
	// If CreateDir returns nil, the test fails
	if err := CreateDir(""); err == nil {
		t.Error(tserr.NilFailed("CreateDir"))
	}
}

// TestCreateDir1 tests CreateDir to create a temporary directory. The test fails,
// if CreateDir returns a Directory that does not exist or if CreateDir returns
// an error.
func TestCreateDir1(t *testing.T) {
	// Create temporary Directory d
	d := tmpDir(t)
	// Remove temporary Directory d
	rm(t, d)
	// Create Directory d with CreateDir
	if err := CreateDir(d); err != nil {
		// If CreateDir returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateDir", Fn: string(d), Err: err}))
	}
	// Retrieve FileInfo from Directory d with Stat
	_, e := os.Stat(string(d))
	// If Stat returns that d does not exist, the test fails
	if os.IsNotExist(e) {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateDir", Fn: string(d), Err: e}))
	}
	// In case of any other error of Stat, the test fails
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "FileInfo (Stat) of", Fn: string(d), Err: e}))
	}
	// Remove Directory d
	rm(t, d)
}

// TestCreateDir2 tests if CreateDir returns nil, in case it retrieves
// a Directory d which already exists. If CreateDir returns an error,
// the test fails.
func TestCreateDir2(t *testing.T) {
	// Create temporary Directory d
	d := tmpDir(t)
	// Call CreateDir with Directory d, which already exists
	if err := CreateDir(d); err != nil {
		// If CreateDir returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateDir", Fn: string(d), Err: err}))
	}
	// Remove Directory d
	rm(t, d)
}

// TestCloseFileNil tests if CloseFile returns nil, in case it retrieves nil.
// If CloseFile returns nil, the test fails.
func TestCloseFileNil(t *testing.T) {
	// If CloseFile returns nil, the test fails
	if err := CloseFile(nil); err == nil {
		t.Error(tserr.NilFailed("CloseFile"))
	}
}

// TestCloseFile tests CloseFile closing a temporary file. If CloseFile
// returns an error, the test fails.
func TestCloseFile(t *testing.T) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Open file fn and get *os.File f
	f, e := OpenFile(fn)
	// If OpenFile returns an error, the test fails
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "OpenFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	// Close file f with CloseFile
	if err := CloseFile(f); err != nil {
		// If ClosFile returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "CloseFile",
			Fn:  f.Name(),
			Err: err,
		}))
	}
	// Remove temporary file fn
	rm(t, fn)
}

// TestCloseFileErr tests if CloseFile returns an error in case
// it is called for a file already closed. The test fails if
// CloseFile returns nil.
func TestCloseFileErr(t *testing.T) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Open file fn and get *os.File f
	f, e := OpenFile(fn)
	// If OpenFile returns an error, the test fails
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{
			Op:  "OpenFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	// Close f
	f.Close()
	// Close file f using CloseFile
	if err := CloseFile(f); err == nil {
		// If CloseFile returns nil, the test fails
		t.Error(tserr.NilFailed("CloseFile"))
	}
	// Remove file fn
	rm(t, fn)
}

// TestWriteStr tests WriteStr to write two times a string into a new temporary file.
// If the contents of the temporary file does not equal the expected contents,
// the test fails.
func TestWriteStr(t *testing.T) {
	// Define how often WriteStr is called in rep
	rep := 2
	// Set expected test string seq to empty string
	seq := ""
	// Create temporary file fn
	fn := tmpFile(t)
	// Iterate as defined by rep
	for i := 0; i < rep; i++ {
		// If writing the testcase string to fn returns an error, the test fails
		if e := WriteStr(fn, testcase); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{
				Op:  fmt.Sprintf("WriteStr %v to file", testcase),
				Fn:  string(fn),
				Err: e,
			}))
		}
		// Extend expected test string by the testcase string
		seq = seq + testcase
	}
	// Read file fn in b
	b, err := os.ReadFile(string(fn))
	// Remove file fn
	rm(t, fn)
	// If ReadFile returns an error, the test fails
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	// If actual b does not match expected seq, the test fails
	if string(b) != seq {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: seq}))
	}
}

// TestWriteStrEmpty tests WriteStr to return an error for an empty filename.
// If WriteStr returns nil, the test fails.
func TestWriteStrEmpty(t *testing.T) {
	// Write string testcase to empty filename
	if e := WriteStr("", testcase); e == nil {
		// If WriteStr returns nil, the test fails
		t.Error(tserr.NilFailed("WriteStr"))
	}
}

// TestTouchFileEmpty tests TouchFile to return an error for an empty filename.
// If TouchFile returns nil, the test fails.
func TestTouchFileEmpty(t *testing.T) {
	// Touch file with empty filename
	if e := TouchFile(""); e == nil {
		// If Touchfile returns nil, the test fails
		t.Error(tserr.NilFailed("TouchFile"))
	}
}

// TestTouchFile1 tests TouchFile to touch an existing temporary
// file. If TouchFile returns an error or if the modification time
// of the temporary file is not later than before, the test fails.
func TestTouchFile1(t *testing.T) {
	testTouchFile(t, false)
}

// TestTouchFile2 tests TouchFile to touch a temporary
// file which was created and removed. If TouchFile returns
// an error or if the modification time of the temporary file
// is not later than before, the test fails.
func TestTouchFile2(t *testing.T) {
	testTouchFile(t, true)
}

// testTouchFile is called by Test functions to test TouchFile. If r is true,
// TouchFile opens a non-existing, new temporary file. If r is false,
// TouchFile opens a temporary file, which was created and removed.
// If TouchFile returns an error or if the modification time of the temporary file
// is not later than before, the test fails.
func testTouchFile(t *testing.T, r bool) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Retrieve modification time of fn in t1
	t1 := modTime(t, fn)
	// Remove fn if r is true
	if r {
		rm(t, fn)
	}
	// Pause for at least one second
	time.Sleep(time.Second)
	// Touch file fn
	if e := TouchFile(fn); e != nil {
		// If TouchFile returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "TouchFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	// Retrieve modification time of fn in t2
	t2 := modTime(t, fn)
	// Calculate duration d = t2 - t1
	d := t2.Sub(t1)
	// If rounded d is smaller than a second, the test fails
	if d.Round(time.Second) < time.Second {
		t.Error(tserr.Higher(&tserr.HigherArgs{
			Var:        "t2 - t1",
			Actual:     int64(d),
			LowerBound: int64(time.Second),
		}))
	}
	// Remove file fn
	rm(t, fn)
}

func TestWriteSingleStr(t *testing.T) {
	rep := 2
	seq := ""
	fn := tmpFile(t)
	for i := 0; i < rep; i++ {
		if e := WriteSingleStr(fn, testcase); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{
				Op:  fmt.Sprintf("WriteSingleStr %v to file", testcase),
				Fn:  string(fn),
				Err: e,
			}))
		}
		seq = seq + testcase
	}
	b, err := os.ReadFile(string(fn))
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	if string(b) != testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase}))
	}
	rm(t, fn)
}

func TestWriteSingleStrErr(t *testing.T) {
	if e := WriteSingleStr("", testcase); e == nil {
		t.Error(tserr.NilFailed("WriteSingleStr"))
	}
}

func TestReadFile(t *testing.T) {
	fn := tmpFile(t)
	if e := WriteStr(fn, testcase); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  fmt.Sprintf("WriteStr %v to file", testcase),
			Fn:  string(fn),
			Err: e,
		}))
	}
	b, err := ReadFile(fn)
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	if string(b) != testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase}))
	}
	rm(t, fn)
}

func TestReadFileErr2(t *testing.T) {
	fn := tmpFile(t)
	if e := WriteStr(fn, testcase); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  fmt.Sprintf("WriteStr %v to file", testcase),
			Fn:  string(fn),
			Err: e,
		}))
	}
	rm(t, fn)
	_, err := ReadFile(fn)
	if err == nil {
		t.Error(tserr.NilFailed("ReadFile"))
	}
}

func TestReadFileErr1(t *testing.T) {
	if _, e := ReadFile(""); e == nil {
		t.Error(tserr.NilFailed("ReadFile"))
	}
}

func TestAppendFileNil(t *testing.T) {
	if e := AppendFile(nil); e == nil {
		t.Error(tserr.NilFailed("AppendFile"))
	}
}

func TestAppendFileEmpty1(t *testing.T) {
	fn := tmpFile(t)
	if e := AppendFile(&Append{fileA: "", fileI: fn}); e == nil {
		t.Error(tserr.NilFailed("AppendFile"))
	}
	rm(t, fn)
}

func TestAppendFileEmpty2(t *testing.T) {
	fn := tmpFile(t)
	if e := AppendFile(&Append{fileI: "", fileA: fn}); e == nil {
		t.Error(tserr.NilFailed("AppendFile"))
	}
	rm(t, fn)
}

func TestAppendFile(t *testing.T) {
	fn := [2]Filename{tmpFile(t), tmpFile(t)}
	for _, i := range fn {
		if e := WriteStr(i, testcase); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{
				Op:  fmt.Sprintf("WriteStr %v to file", testcase),
				Fn:  string(i),
				Err: e,
			}))
		}
	}
	if e := AppendFile(&Append{fileA: fn[0], fileI: fn[1]}); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  fmt.Sprintf("AppendFile %v to file", fn[1]),
			Fn:  string(fn[0]),
			Err: e,
		}))
	}
	b, err := os.ReadFile(string(fn[0]))
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{
			Op:  "ReadFile",
			Fn:  string(fn[0]),
			Err: err,
		}))
	}
	if string(b) != testcase+testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase + testcase}))
	}
	for _, i := range fn {
		rm(t, i)
	}
}

func TestExistsFileEmpty(t *testing.T) {
	_, e := ExistsFile("")
	if e == nil {
		t.Error(tserr.NilFailed("ExistsFile"))
	}
}

func TestExistsFile1(t *testing.T) {
	testExistsFileWrapper(t, true)
}

func TestExistsFile2(t *testing.T) {
	testExistsFileWrapper(t, false)
}

func testExistsFileWrapper(t *testing.T, r bool) {
	fn := tmpFile(t)
	if r {
		rm(t, fn)
	}
	b, e := ExistsFile(fn)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "ExistsFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	if b == r {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     "ExistsFile",
			Actual: fmt.Sprintf("%t", b),
			Want:   fmt.Sprintf("%t", !b),
		}))
	}
	if !r {
		rm(t, fn)
	}
}

func TestRemoveFileEmpty(t *testing.T) {
	if err := RemoveFile(""); err == nil {
		t.Error(tserr.NilFailed("RemoveFile"))
	}
}

func TestRemoveFile1(t *testing.T) {
	fn := tmpFile(t)
	if e := RemoveFile(fn); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "RemoveFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	b, err := ExistsFile(fn)
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "ExistsFile",
			Fn:  string(fn),
			Err: err,
		}))
	}
	if b {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     "ExistsFile",
			Actual: fmt.Sprintf("%t", b),
			Want:   fmt.Sprintf("%t", !b),
		}))
	}
}

func TestRemoveFile2(t *testing.T) {
	fn := tmpFile(t)
	rm(t, fn)
	if e := RemoveFile(fn); e == nil {
		t.Error(tserr.NilFailed("RemoveFile"))
	}
}

func TestResetFileEmpty(t *testing.T) {
	if err := ResetFile(""); err == nil {
		t.Error(tserr.NilFailed("ResetFile"))
	}
}

func TestResetFile1(t *testing.T) {
	testResetFile(t, false)
}

func TestResetFile2(t *testing.T) {
	testResetFile(t, true)
}

func testResetFile(t *testing.T, r bool) {
	fn := tmpFile(t)
	WriteSingleStr(fn, testcase)
	if r {
		rm(t, fn)
	}
	if err := ResetFile(fn); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "ResetFile",
			Fn:  string(fn),
			Err: err,
		}))
	}
	fi, e := os.Stat(string(fn))
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "FileInfo (Stat) of",
			Fn:  string(fn),
			Err: e,
		}))
	}
	if fi.Size() > 0 {
		t.Error(tserr.Equal(&tserr.EqualArgs{
			Var:    fmt.Sprintf("Size of %v", fn),
			Actual: fi.Size(),
			Want:   0,
		}))
	}
	rm(t, fn)
}
