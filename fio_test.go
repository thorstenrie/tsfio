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
