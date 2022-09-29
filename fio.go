package tsfio

import (
	"fmt"
	"io/fs"
	"os"
	"time"
)

const (
	flags int         = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	fperm fs.FileMode = 0644
	dperm fs.FileMode = 0755
)

func WriteStr(fn Filename, s string) error {
	f, err := OpenFile(fn)
	if err != nil {
		return errOp("open", string(fn), err)
	}
	if _, e := f.WriteString(s); e != nil {
		f.Close()
		return errOp("write string to", string(fn), e)
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
		return errChk(string(a.fileA), e)
	}
	if e := CheckFile(a.fileI); e != nil {
		return errChk(string(a.fileI), e)
	}
	f, erro := OpenFile(a.fileA)
	if erro != nil {
		return errOp("open", string(a.fileA), erro)
	}
	out, errr := ReadFile(a.fileI)
	if errr != nil {
		return errOp("read", string(a.fileI), errr)
	}
	if _, e := f.Write(out); e != nil {
		f.Close()
		return errOp(fmt.Sprintf("append file %v to", a.fileI), string(a.fileA), e)
	}
	if e := f.Close(); e != nil {
		return errOp("close", string(a.fileA), e)
	}
	return nil
}

func ExistsFile(fn Filename) (bool, error) {
	if err := CheckFile(fn); err != nil {
		return false, errChk(string(fn), err)
	}
	_, err := os.Stat(string(fn))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errOp("FileInfo of", string(fn), err)
}

// done
func WriteSingleStr(fn Filename, s string) error {
	if e := ResetFile(fn); e != nil {
		return errOp("reset", string(fn), e)
	}
	if e := WriteStr(fn, s); e != nil {
		return errOp(fmt.Sprintf("write string %v to", s), string(fn), e)
	}
	return nil
}

// done
func ReadFile(f Filename) ([]byte, error) {
	if e := CheckFile(f); e != nil {
		return nil, errChk(string(f), e)
	}
	b, err := os.ReadFile(string(f))
	if err != nil {
		return nil, errOp("read", string(f), err)
	}
	return b, nil
}

// done
func RemoveFile(f Filename) error {
	if e := CheckFile(f); e != nil {
		return errChk(string(f), e)
	}

	b, err := ExistsFile(f)
	if err != nil {
		return errOp("check if exists", string(f), err)
	}

	if b {
		e := os.Remove(string(f))
		if e != nil {
			return errOp("remove", string(f), err)
		}
	} else {
		return fmt.Errorf("remove file %v failed, because it does not exist", f)
	}

	return nil
}

// done
func CreateDir(d Directory) error {
	if e := CheckDir(d); e != nil {
		return errChk(string(d), e)
	}
	err := os.MkdirAll(string(d), dperm)
	if err != nil {
		return errOp("make", string(d), err)
	}
	return nil
}

// done
func ResetFile(f Filename) error {
	b, err := ExistsFile(f)
	if err != nil {
		return errOp("check if exists", string(f), err)
	}
	if b {
		if e := TouchFile(f); e != nil {
			return errOp("touch", string(f), e)
		}
	}
	err = os.Truncate(string(f), 0)
	if err != nil {
		return errOp("truncate", string(f), err)
	}
	return nil
}

func TouchFile(fn Filename) error {
	if e := CheckFile(fn); e != nil {
		return errChk(string(fn), e)
	}
	b, erre := ExistsFile(fn)
	if erre != nil {
		return errOp("check if exists", string(fn), erre)
	}
	if b {
		t := time.Now().Local()
		if e := os.Chtimes(string(fn), t, t); e != nil {
			return errOp("chtimes", string(fn), e)
		}
	} else {
		f, erro := OpenFile(fn)
		if erro != nil {
			return errOp("open", string(fn), erro)
		}
		if e := f.Close(); e != nil {
			return errOp("close", string(fn), e)
		}
	}
	return nil
}

func OpenFile(fn Filename) (*os.File, error) {
	if e := CheckFile(fn); e != nil {
		return nil, errChk(string(fn), e)
	}
	f, err := os.OpenFile(string(fn), flags, fperm)
	if err != nil {
		return nil, errOp("open", string(fn), err)
	}
	return f, nil
}
