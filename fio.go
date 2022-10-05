package tsfio

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/thorstenrie/tserr"
)

const (
	flags int         = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	fperm fs.FileMode = 0644
	dperm fs.FileMode = 0755
)

func WriteStr(fn Filename, s string) error {
	f, err := OpenFile(fn)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "open", Fn: string(fn), Err: err})
	}
	if _, e := f.WriteString(s); e != nil {
		f.Close()
		return tserr.Op(&tserr.OpArgs{Op: "write string to", Fn: string(fn), Err: e})
	}
	return nil
}

type Append struct {
	fileA Filename
	fileI Filename
}

func AppendFile(a *Append) error {
	if a == nil {
		return fmt.Errorf("nil pointer")
	}

	if e := CheckFile(a.fileA); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(a.fileA), Err: e})
	}
	if e := CheckFile(a.fileI); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(a.fileI), Err: e})
	}
	f, erro := OpenFile(a.fileA)
	if erro != nil {
		return tserr.Op(&tserr.OpArgs{Op: "open", Fn: string(a.fileA), Err: erro})
	}
	out, errr := ReadFile(a.fileI)
	if errr != nil {
		return tserr.Op(&tserr.OpArgs{Op: "read", Fn: string(a.fileI), Err: errr})
	}
	if _, e := f.Write(out); e != nil {
		f.Close()
		return tserr.Op(&tserr.OpArgs{Op: fmt.Sprintf("append file %v to", a.fileI), Fn: string(a.fileA), Err: e})
	}
	if e := f.Close(); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "close", Fn: string(a.fileA), Err: e})
	}
	return nil
}

func ExistsFile(fn Filename) (bool, error) {
	if err := CheckFile(fn); err != nil {
		return false, tserr.Check(&tserr.CheckArgs{F: string(fn), Err: err})
	}
	_, err := os.Stat(string(fn))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, tserr.Op(&tserr.OpArgs{Op: "FileInfo of", Fn: string(fn), Err: err})
}

func WriteSingleStr(fn Filename, s string) error {
	if e := ResetFile(fn); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "reset", Fn: string(fn), Err: e})
	}
	if e := WriteStr(fn, s); e != nil {
		return tserr.Op(&tserr.OpArgs{Op: fmt.Sprintf("write string %v to", s), Fn: string(fn), Err: e})
	}
	return nil
}

func ReadFile(f Filename) ([]byte, error) {
	if e := CheckFile(f); e != nil {
		return nil, tserr.Check(&tserr.CheckArgs{F: string(f), Err: e})
	}
	b, err := os.ReadFile(string(f))
	if err != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "read", Fn: string(f), Err: err})
	}
	return b, nil
}

func RemoveFile(f Filename) error {
	if e := CheckFile(f); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(f), Err: e})
	}

	b, err := ExistsFile(f)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "check if exists", Fn: string(f), Err: err})
	}

	if b {
		e := os.Remove(string(f))
		if e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "remove", Fn: string(f), Err: err})
		}
	} else {
		return tserr.NotExistent(string(f))
	}

	return nil
}

func CreateDir(d Directory) error {
	if e := CheckDir(d); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(d), Err: e})
	}
	err := os.MkdirAll(string(d), dperm)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "make", Fn: string(d), Err: err})
	}
	return nil
}

func ResetFile(f Filename) error {
	b, err := ExistsFile(f)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "check if exists", Fn: string(f), Err: err})
	}
	if b {
		if e := TouchFile(f); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "touch", Fn: string(f), Err: e})
		}
	}
	err = os.Truncate(string(f), 0)
	if err != nil {
		return tserr.Op(&tserr.OpArgs{Op: "truncate", Fn: string(f), Err: err})
	}
	return nil
}

func TouchFile(fn Filename) error {
	if e := CheckFile(fn); e != nil {
		return tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	b, erre := ExistsFile(fn)
	if erre != nil {
		return tserr.Op(&tserr.OpArgs{Op: "check if exists", Fn: string(fn), Err: erre})
	}
	if b {
		t := time.Now().Local()
		if e := os.Chtimes(string(fn), t, t); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "chtimes", Fn: string(fn), Err: e})
		}
	} else {
		f, erro := OpenFile(fn)
		if erro != nil {
			return tserr.Op(&tserr.OpArgs{Op: "open", Fn: string(fn), Err: erro})
		}
		if e := f.Close(); e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "close", Fn: string(fn), Err: e})
		}
	}
	return nil
}

func OpenFile(fn Filename) (*os.File, error) {
	if e := CheckFile(fn); e != nil {
		return nil, tserr.Check(&tserr.CheckArgs{F: string(fn), Err: e})
	}
	f, err := os.OpenFile(string(fn), flags, fperm)
	if err != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "open", Fn: string(fn), Err: err})
	}
	return f, nil
}
