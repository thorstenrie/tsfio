// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio_test

import (
	"testing"

	"github.com/thorstenrie/tserr"
	"github.com/thorstenrie/tsfio"
)

func TestGoldenFile(t *testing.T) {
	tc := &tsfio.Testcase{Name: testcase, Data: testcase}
	e := tsfio.CreateGoldenFile(tc)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateGoldenFile", Fn: tc.Name, Err: e}))
	}
	e = tsfio.EvalGoldenFile(tc)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "EvalGoldenFile", Fn: tc.Name, Err: e}))
	}
	fn, e := tsfio.GoldenFilePath(tc.Name)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "GoldenFilePath", Fn: tc.Name, Err: e}))
	}
	b, e := tsfio.ExistsFile(fn)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ExistsFile", Fn: string(fn), Err: e}))
	}
	if !b {
		t.Error(tserr.NotExistent(string(fn)))
	}
	if e := tsfio.RemoveFile(fn); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "RemoveFile", Fn: string(fn), Err: e}))
	}
}
