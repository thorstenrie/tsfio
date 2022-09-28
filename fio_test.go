package tsfio

import (
	"os"
	"testing"
)

func TestWriteStr(t *testing.T) {
	rep := 2
	seq := ""
	fn := tmpFile(t)
	for i := 0; i < rep; i++ {
		if e := WriteStr(fn, testcase); e != nil {
			t.Errorf("WriteStr %v to file %v failed: %v", testcase, fn, e)
		}
		seq = seq + testcase
	}
	b, err := os.ReadFile(string(fn))
	if err != nil {
		t.Errorf("ReadFile %v failed: %v", fn, err)
	}
	if string(b) != seq {
		t.Errorf("%v does not equal %v", string(b), seq)
	}
}

func TestOpenFile(t *testing.T) {
	fn := tmpFile(t)
	f, err := OpenFile(fn)
	if err != nil {
		t.Errorf("OpenFile %v failed: %v", fn, err)
	}
	if e := f.Close(); e != nil {
		t.Errorf("Close %v failed: %v", fn, err)
	}
}

func TestOpenFileRm(t *testing.T) {
	fn := tmpFile(t)
	rm(t, fn)
	f, err := OpenFile(fn)
	if err != nil {
		t.Errorf("OpenFile %v failed: %v", fn, err)
	}
	if e := f.Close(); e != nil {
		t.Errorf("Close %v failed: %v", fn, err)
	}
}

func TestOpenFileEmpty(t *testing.T) {
	_, err := OpenFile("")
	if err == nil {
		t.Errorf("OpenFile returned nil, but error expected")
	}
}

func TestWriteStrErr(t *testing.T) {
	if e := WriteStr("", testcase); e == nil {
		t.Errorf("WriteStr returned nil, but error expected")
	}
}

func TestWriteSingleStr(t *testing.T) {
	rep := 2
	seq := ""
	fn := tmpFile(t)
	for i := 0; i < rep; i++ {
		if e := WriteSingleStr(fn, testcase); e != nil {
			t.Errorf("WriteSingleStr %v to file %v failed: %v", testcase, fn, e)
		}
		seq = seq + testcase
	}
	b, err := os.ReadFile(string(fn))
	if err != nil {
		t.Errorf("ReadFile %v failed: %v", fn, err)
	}
	if string(b) != testcase {
		t.Errorf("%v does not equal %v", string(b), testcase)
	}
}

func TestWriteSingleStrErr(t *testing.T) {
	if e := WriteSingleStr("", testcase); e == nil {
		t.Errorf("WriteSingleStr returned nil, but error expected")
	}
}

func TestReadFile(t *testing.T) {
	fn := tmpFile(t)
	if e := WriteStr(fn, testcase); e != nil {
		t.Errorf("WriteStr %v to file %v failed: %v", testcase, fn, e)
	}
	b, err := ReadFile(fn)
	if err != nil {
		t.Errorf("ReadFile %v failed: %v", fn, err)
	}
	if string(b) != testcase {
		t.Errorf("%v does not equal %v", string(b), testcase)
	}
}

func TestReadFileErr2(t *testing.T) {
	fn := tmpFile(t)
	if e := WriteStr(fn, testcase); e != nil {
		t.Errorf("WriteStr %v to file %v failed: %v", testcase, fn, e)
	}
	rm(t, fn)
	_, err := ReadFile(fn)
	if err == nil {
		t.Errorf("ReadFile returned nil, but error expected")
	}
}

func TestReadFileErr1(t *testing.T) {
	if _, e := ReadFile(""); e == nil {
		t.Errorf("ReadFile returned nil, but error expected")
	}
}

func TestAppendFileNil(t *testing.T) {
	if e := AppendFile(nil); e == nil {
		t.Errorf("AppendFile returned nil, but error expected")
	}
}

func TestAppendFileEmpty1(t *testing.T) {
	if e := AppendFile(&Append{fileA: "", fileI: tmpFile(t)}); e == nil {
		t.Errorf("AppendFile returned nil, but error expected")
	}
}

func TestAppendFileEmpty2(t *testing.T) {
	if e := AppendFile(&Append{fileI: "", fileA: tmpFile(t)}); e == nil {
		t.Errorf("AppendFile returned nil, but error expected")
	}
}

func TestAppendFile(t *testing.T) {
	fn := [2]Filename{tmpFile(t), tmpFile(t)}
	for _, i := range fn {
		if e := WriteStr(i, testcase); e != nil {
			t.Errorf("WriteStr %v to file %v failed: %v", testcase, i, e)
		}
	}
	if e := AppendFile(&Append{fileA: fn[0], fileI: fn[1]}); e != nil {
		t.Errorf("AppendFile %v to %v failed: %v", fn[1], fn[0], e)
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
