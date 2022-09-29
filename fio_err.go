package tsfio

import "fmt"

type errm struct {
	ec int
	m  string
}

var (
	chkFailed     = errm{412, "check %v failed: %w"}
	opFailed      = errm{422, "%v %v failed: %w"}
	testNilFailed = errm{500, "%v returned nil, but error expected"}
)

func errorf(e errm, a ...any) error {
	return fmt.Errorf("%d: %w", e.ec, fmt.Errorf(e.m, a...))
}

func errChk(f string, err error) error {
	if err == nil {
		return nil
	}
	return errorf(chkFailed, f, err)
}

func errOp(op string, f string, err error) error {
	if err == nil {
		return nil
	}
	return errorf(opFailed, op, f, err)
}

func testErrNil(op string) error {
	return errorf(testNilFailed, op)
}
