package tsfio

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/thorstenrie/tserr"
)

func TestOpenFile1(t *testing.T) {
	testOpenFile(t, false)
}

func TestOpenFile2(t *testing.T) {
	testOpenFile(t, true)
}

func testOpenFile(t *testing.T, r bool) {
	fn := tmpFile(t)
	if r {
		rm(t, fn)
	}
	f, err := OpenFile(fn)
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: err}))
	}
	if e := CloseFile(f); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CloseFile", Fn: string(fn), Err: e}))
	}
	rm(t, fn)
}

func TestOpenFileEmpty(t *testing.T) {
	if _, err := OpenFile(""); err == nil {
		t.Error(tserr.NilFailed("OpenFile"))
	}
}

func TestCreateDirEmpty(t *testing.T) {
	if err := CreateDir(""); err == nil {
		t.Error(tserr.NilFailed("CreateDir"))
	}
}

func TestCreateDir1(t *testing.T) {
	d := tmpDir(t)
	rm(t, d)
	if err := CreateDir(d); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateDir", Fn: string(d), Err: err}))
	}
	_, e := os.Stat(string(d))
	if os.IsNotExist(e) {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateDir", Fn: string(d), Err: e}))
	}
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "FileInfo (Stat) of", Fn: string(d), Err: e}))
	}
	rm(t, d)
}

func TestCreateDir2(t *testing.T) {
	d := tmpDir(t)
	if err := CreateDir(d); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateDir", Fn: string(d), Err: err}))
	}
	rm(t, d)
}

func TestCloseFileNil(t *testing.T) {
	if err := CloseFile(nil); err == nil {
		t.Error(tserr.NilFailed("CloseFile"))
	}
}

func TestCloseFile(t *testing.T) {
	fn := tmpFile(t)
	f, e := OpenFile(fn)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "OpenFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	if err := CloseFile(f); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "CloseFile",
			Fn:  f.Name(),
			Err: err,
		}))
	}
	rm(t, fn)
}

func TestCloseFileErr(t *testing.T) {
	fn := tmpFile(t)
	f, e := OpenFile(fn)
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{
			Op:  "OpenFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	f.Close()
	if err := CloseFile(f); err == nil {
		t.Error(tserr.NilFailed("CloseFile"))
	}
	rm(t, fn)
}

func TestWriteStr(t *testing.T) {
	rep := 2
	seq := ""
	fn := tmpFile(t)
	for i := 0; i < rep; i++ {
		if e := WriteStr(fn, testcase); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{
				Op:  fmt.Sprintf("WriteStr %v to file", testcase),
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
	if string(b) != seq {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: seq}))
	}
	rm(t, fn)
}

func TestWriteStrErr(t *testing.T) {
	if e := WriteStr("", testcase); e == nil {
		t.Error(tserr.NilFailed("WriteStr"))
	}
}

func TestTouchFileEmpty(t *testing.T) {
	if e := TouchFile(""); e == nil {
		t.Error(tserr.NilFailed("TouchFile"))
	}
}

func TestTouchFile1(t *testing.T) {
	testTouchFile(t, false)
}

func TestTouchFile2(t *testing.T) {
	testTouchFile(t, true)
}

func testTouchFile(t *testing.T, r bool) {
	fn := tmpFile(t)
	t1 := modTime(t, fn)
	if r {
		rm(t, fn)
	}
	time.Sleep(time.Second)
	if e := TouchFile(fn); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{
			Op:  "TouchFile",
			Fn:  string(fn),
			Err: e,
		}))
	}
	t2 := modTime(t, fn)
	d := t2.Sub(t1)
	if d.Round(time.Second) < time.Second {
		t.Error(tserr.Higher(&tserr.HigherArgs{
			Var:        "t2 - t1",
			Actual:     int64(d),
			LowerBound: int64(time.Second),
		}))
	}
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
