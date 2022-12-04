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

// TestWriteSingleStr tests WriteSingleStr to write two times a string into a temporary file.
// If the contents of the temporary file does not equal the expected contents, the test fails.
func TestWriteSingleStr(t *testing.T) {
	// Define how often WriteSingleStr is called in rep
	rep := 2
	// Create temporary file fn
	fn := tmpFile(t)
	// Iterate as defined by rep
	for i := 0; i < rep; i++ {
		// If writing the testcase string to fn returns an error, the test fails
		if e := WriteSingleStr(fn, testcase); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{
				Op:  fmt.Sprintf("WriteSingleStr %v to file", testcase),
				Fn:  string(fn),
				Err: e,
			}))
		}
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
	if string(b) != testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase}))
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

// TestWriteSingleStrErr tests WriteSingleStr to return an error for an empty filename.
// If WriteSingleStr returns nil, the test fails.
func TestWriteSingleStrErr(t *testing.T) {
	// Write string testcase to empty filename
	if e := WriteSingleStr("", testcase); e == nil {
		// If WriteSingleStr returns nil, the test fails
		t.Error(tserr.NilFailed("WriteSingleStr"))
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

// TestReadFile tests ReadFile to read from a temporary file. If the returned
// contents of the temporary file does not match the expected contents,
// the test fails.
func TestReadFile(t *testing.T) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Write string testcase to file fn by calling WriteStr
	if e := WriteStr(fn, testcase); e != nil {
		// If WriteStr returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  fmt.Sprintf("WriteStr %v to file", testcase),
			Fn:  string(fn),
			Err: e,
		}))
	}
	// Read fn in b
	b, err := ReadFile(fn)
	// Remove temporary file fn
	rm(t, fn)
	// If ReadFile returns an error, the test fails
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	// If b does not match testcase, the test fails
	if string(b) != testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase}))
	}
}

// TestReadFileEmpty tests ReadFile to return an error for an empty filename.
// If ReadFile returns nil, the test fails.
func TestReadFileEmpty(t *testing.T) {
	// Read file with empty Filename
	if _, e := ReadFile(""); e == nil {
		// If ReadFile returns nil, the test fails
		t.Error(tserr.NilFailed("ReadFile"))
	}
}

// TestReadFileErr tests ReadFile to return an error when reading a file
// which does not exist. If ReadFile returns nil, the test fails.
func TestReadFileErr(t *testing.T) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Write string testcase to fn
	if e := WriteStr(fn, testcase); e != nil {
		// If WriteStr returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  fmt.Sprintf("WriteStr %v to file", testcase),
			Fn:  string(fn),
			Err: e,
		}))
	}
	// Remove file fn
	rm(t, fn)
	// Read from file fn
	_, err := ReadFile(fn)
	// If ReadFile returns nil, the test fails
	if err == nil {
		t.Error(tserr.NilFailed("ReadFile"))
	}
}

// TestAppendFileNil tests AppendFile to return an error if retrieving a nil pointer.
// If AppendFile returns nil, the test fails.
func TestAppendFileNil(t *testing.T) {
	// Call AppendFile with a nil pointer
	if e := AppendFile(nil); e == nil {
		// If AppendFile returns nil, the test fails
		t.Error(tserr.NilFailed("AppendFile"))
	}
}

// TestAppendFileEmptyA tests AppendFile to return an error if Append argument fileA is
// an empty string. If AppendFile returns nil, the test fails.
func TestAppendFileEmptyA(t *testing.T) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Call AppendFile with an empty string as fileA and fn as fileI
	if e := AppendFile(&Append{fileA: "", fileI: fn}); e == nil {
		// If AppendFile returns nil, the test fails
		t.Error(tserr.NilFailed("AppendFile"))
	}
	// Remove temporary file fn
	rm(t, fn)
}

// TestAppendFileEmptyI tests AppendFile to return an error if Append argument fileI is
// an empty string. If AppendFile returns nil, the test fails.
func TestAppendFileEmptyI(t *testing.T) {
	// Create temporary file fn
	fn := tmpFile(t)
	// Call AppendFile with an empty string as fileI and fn as fileA
	if e := AppendFile(&Append{fileI: "", fileA: fn}); e == nil {
		// If AppendFile returns nil, the test fails
		t.Error(tserr.NilFailed("AppendFile"))
	}
	// Remove temporary file fn
	rm(t, fn)
}

// TestAppendFileEmptyIA tests AppendFile to return an error if Append arguments fileI and fileA
// are both an empty string. If AppendFile returns nil, the test fails.
func TestAppendFileEmptyIA(t *testing.T) {
	// Call AppendFile with an empty string as fileI and fileA
	if e := AppendFile(&Append{fileI: "", fileA: ""}); e == nil {
		// If AppendFile returns nil, the test fails
		t.Error(tserr.NilFailed("AppendFile"))
	}
}

// TestAppendFile tests AppendFile to append a temporary file to another temporary file.
// If the resulting file does not match the expected contents, the test fails.
func TestAppendFile(t *testing.T) {
	// Create two temporary files in fn
	fn := [2]Filename{tmpFile(t), tmpFile(t)}
	// Write testcase string to each temporary file in fn
	for _, i := range fn {
		// If WriteStr returns an error, the test fails
		if e := WriteStr(i, testcase); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{
				Op:  fmt.Sprintf("WriteStr %v to file", testcase),
				Fn:  string(i),
				Err: e,
			}))
		}
	}
	// Append fn[1] to fn[0] in fn[0]
	if e := AppendFile(&Append{fileA: fn[0], fileI: fn[1]}); e != nil {
		// If AppendFile returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  fmt.Sprintf("AppendFile %v to file", fn[1]),
			Fn:  string(fn[0]),
			Err: e,
		}))
	}
	// Read contents of fn[0] in b with os.ReadFile
	b, err := os.ReadFile(string(fn[0]))
	// Remove all files of fn
	for _, i := range fn {
		rm(t, i)
	}
	// If os.ReadFile returns an error, the test fails
	if err != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{
			Op:  "ReadFile",
			Fn:  string(fn[0]),
			Err: err,
		}))
	}
	// If b does not match the expected string, the test fails
	if string(b) != testcase+testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase + testcase}))
	}
}

// TestExistsFileEmpty tests ExistsFile to return an error for an empty string as filename.
// If ExistsFile returns the error to be nil, the test fails.
func TestExistsFileEmpty(t *testing.T) {
	// Call ExistsFile for an empty string as filename
	_, e := ExistsFile("")
	// If ExistsFile returns the error to be nil, the test fails
	if e == nil {
		t.Error(tserr.NilFailed("ExistsFile"))
	}
}

// TestExistsFile1 tests ExistsFile for a non-existing file.
// If ExistsFile returns an error or if ExistsFile returns true,
// the test fails.
func TestExistsFile1(t *testing.T) {
	testExistsFile(t, true)
}

// TestExistsFile2 tests ExistsFile for an existing temporary file.
// If ExistsFile returns an error or if ExistsFile returns false,
// the test fails.
func TestExistsFile2(t *testing.T) {
	testExistsFile(t, false)
}

// testExistsFile is called by Test functions to test ExistsFile. If r is true,
// ExistsFile is tested for a non-existing, removed file. If r is false,
// ExistsFile is tested for an existing temporary file. If ExistsFile returns
// an error or if the actual result of ExistsFile does not match the expected
// result, the test fails.
func testExistsFile(t *testing.T, r bool) {
	// Create a temporary file fn
	fn := tmpFile(t)
	// Remove fn, if r is true
	if r {
		rm(t, fn)
	}
	// Call ExistsFile for fn
	b, e := ExistsFile(fn)
	// If ExistsFile returns an error, the test fails
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "ExistsFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	// If the actual result of ExistsFile does not match the expected result,
	// the test fails
	if b == r {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     "ExistsFile",
			Actual: fmt.Sprintf("%t", b),
			Want:   fmt.Sprintf("%t", !b),
		}))
	}
	// Remove fn, if r is true
	if !r {
		rm(t, fn)
	}
}

// TestRemoveFileEmpty tests RemoveFile to return an error for an empty string as filename.
// If RemoveFile returns nil, the test fails.
func TestRemoveFileEmpty(t *testing.T) {
	// Call RemoveFile for an empty string as filename
	if err := RemoveFile(""); err == nil {
		// If RemoveFile returns nil, the test fails
		t.Error(tserr.NilFailed("RemoveFile"))
	}
}

// TestRemoveFile1 tests RemoveFile to remove an existing temporary file.
// If RemoveFile returns an error or the temporary file still exists after
// calling RemoveFile, the test fails.
func TestRemoveFile1(t *testing.T) {
	// Create the temporary file fn
	fn := tmpFile(t)
	// Remove fn with RemoveFile
	if e := RemoveFile(fn); e != nil {
		// If RemoveFile returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "RemoveFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	// Check if fn still exists with ExistsFile
	b, err := ExistsFile(fn)
	// If ExistsFile returns an error, the test fails
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "ExistsFile",
			Fn:  string(fn),
			Err: err,
		}))
	}
	// If fn still exists, the test fails
	if b {
		t.Error(tserr.Return(&tserr.ReturnArgs{
			Op:     "ExistsFile",
			Actual: fmt.Sprintf("%t", b),
			Want:   fmt.Sprintf("%t", !b),
		}))
		// Remove file fn
		rm(t, fn)
	}
}

// TestRemoveFile2 tests RemoveFile for an non-existing file.
// If RemoveFile returns nil, the test fails.
func TestRemoveFile2(t *testing.T) {
	// Create the temporary file fn
	fn := tmpFile(t)
	// Remove the temporary file fn
	rm(t, fn)
	// Call RemoveFile for fn
	if e := RemoveFile(fn); e == nil {
		// If RemoveFile returns nil, the test fails
		t.Error(tserr.NilFailed("RemoveFile"))
	}
}

// TestResetFileEmpty tests ResetFile with an empty string as filename.
// If ResetFiles returns nil, it fails.
func TestResetFileEmpty(t *testing.T) {
	// Call ResetFile with an empty string as filename
	if err := ResetFile(""); err == nil {
		// If ResetFile returns nil, the test fails
		t.Error(tserr.NilFailed("ResetFile"))
	}
}

func TestResetFile1(t *testing.T) {
	testResetFile(t, false)
}

func TestResetFile2(t *testing.T) {
	testResetFile(t, true)
}

// testResetFile is called by Test functions to test ResetFile. If r is true,
// ResetFile is tested for a non-existing, removed file. If r is false,
// ResetFile is tested for an existing temporary file. If ResetFile returns
// an error, if FileInfo of the tested file fails, or if the tested file has
// a size greater than zero, the test fails.
func testResetFile(t *testing.T, r bool) {
	// Create the temporary file fn
	fn := tmpFile(t)
	// Write single string testcase to fn
	if err := WriteSingleStr(fn, testcase); err != nil {
		// If WriteSingleStr returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "WriteSingleStr",
			Fn:  string(fn),
			Err: err,
		}))
	}
	// Remove file fn, if r is true
	if r {
		rm(t, fn)
	}
	// Reset file fn
	if err := ResetFile(fn); err != nil {
		// If ResetFile returns an error, the test fails
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "ResetFile",
			Fn:  string(fn),
			Err: err,
		}))
	}
	// Retrieve FileInfo fi from fn with os.Stat
	fi, e := os.Stat(string(fn))
	// If os.Stat returns an error, the test fails
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "FileInfo (Stat) of",
			Fn:  string(fn),
			Err: e,
		}))
	}
	// If fn has a size greater than zero, the test fails
	if fi.Size() > 0 {
		t.Error(tserr.Equal(&tserr.EqualArgs{
			Var:    fmt.Sprintf("Size of %v", fn),
			Actual: fi.Size(),
			Want:   0,
		}))
	}
	// Remove file fn
	rm(t, fn)
}
