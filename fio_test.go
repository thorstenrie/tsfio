package tsfio

import (
	"fmt"
	"os"
	"testing"

	"github.com/thorstenrie/tserr"
)

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
}

func TestOpenFile(t *testing.T) {
	fn := tmpFile(t)
	f, err := OpenFile(fn)
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: err}))
	}
	if e := CloseFile(f); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CloseFile", Fn: string(fn), Err: e}))
	}
}

func TestOpenFileRm(t *testing.T) {
	fn := tmpFile(t)
	rm(t, fn)
	f, err := OpenFile(fn)
	if err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "OpenFile", Fn: string(fn), Err: err}))
	}
	if e := CloseFile(f); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CloseFile", Fn: string(fn), Err: e}))
	}
}

func TestOpenFileEmpty(t *testing.T) {
	if _, err := OpenFile(""); err == nil {
		t.Error(tserr.NilFailed("OpenFile"))
	}
}

func TestTouchFileEmpty(t *testing.T) {
	if e := TouchFile(""); e == nil {
		t.Error(tserr.NilFailed("TouchFile"))
	}
}

func TestWriteStrErr(t *testing.T) {
	if e := WriteStr("", testcase); e == nil {
		t.Error(tserr.NilFailed("WriteStr"))
	}
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
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	if string(b) != testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase}))
	}
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
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: err}))
	}
	if string(b) != testcase {
		t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: string(b), Y: testcase}))
	}
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
	if e := AppendFile(&Append{fileA: "", fileI: tmpFile(t)}); e == nil {
		t.Error(tserr.NilFailed("AppendFile"))
	}
}

func TestAppendFileEmpty2(t *testing.T) {
	if e := AppendFile(&Append{fileI: "", fileA: tmpFile(t)}); e == nil {
		t.Error(tserr.NilFailed("AppendFile"))
	}
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
		t.Errorf("ReadFile %v failed: %v", fn[0], err)
	}
	if string(b) != testcase+testcase {
		t.Errorf("%v does not equal %v", string(b), testcase+testcase)
	}
}

func TestExistsFileEmpty(t *testing.T) {
	_, e := ExistsFile("")
	if e == nil {
		t.Errorf("ExistsFile returned nil, but error expected")
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
		t.Errorf("ExistsFile %v failed: %v", fn, e)
	}
	if b == r {
		t.Errorf("ExistsFile %v returned %t, but %t expected", fn, b, !b)
	}
}
