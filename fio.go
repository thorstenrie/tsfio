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
		return errFio("open", fn, err)
	}
	if _, e := f.WriteString(s); e != nil {
		f.Close()
		return errFio("write string to", fn, e)
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
		return errChk(a.fileA, e)
	}
	if e := CheckFile(a.fileI); e != nil {
		return errChk(a.fileI, e)
	}
	f, erro := OpenFile(a.fileA)
	if erro != nil {
		return errFio("open", a.fileA, erro)
	}
	out, errr := ReadFile(a.fileI)
	if errr != nil {
		return errFio("read", a.fileI, errr)
	}
	if _, e := f.Write(out); e != nil {
		f.Close()
		return errFio(fmt.Sprintf("append file %v to", a.fileI), a.fileA, e)
	}
	if e := f.Close(); e != nil {
		return errFio("close", a.fileA, e)
	}
	return nil
}

func ExistsFile(fn Filename) (bool, error) {
	if err := CheckFile(fn); err != nil {
		return false, errChk(fn, err)
	}
	_, err := os.Stat(string(fn))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errFio("FileInfo of", fn, err)
}

// done
func WriteSingleStr(fn Filename, s string) error {
	if e := ResetFile(fn); e != nil {
		return errFio("reset", fn, e)
	}
	if e := WriteStr(fn, s); e != nil {
		return errFio(fmt.Sprintf("write string %v to", s), fn, e)
	}
	return nil
}

// done
func ReadFile(f Filename) ([]byte, error) {
	if e := CheckFile(f); e != nil {
		return nil, errChk(f, e)
	}
	b, err := os.ReadFile(string(f))
	if err != nil {
		return nil, errFio("read", f, err)
	}
	return b, nil
}

// done
func RemoveFile(f Filename) error {
	if e := CheckFile(f); e != nil {
		return errChk(f, e)
	}

	b, err := ExistsFile(f)
	if err != nil {
		return errFio("check if exists", f, err)
	}

	if b {
		e := os.Remove(string(f))
		if e != nil {
			return errFio("remove", f, err)
		}
	} else {
		return fmt.Errorf("remove file %v failed, because it does not exist", f)
	}

	return nil
}

// done
func CreateDir(d Directory) error {
	if e := CheckDir(d); e != nil {
		return errChk(d, e)
	}
	err := os.MkdirAll(string(d), dperm)
	if err != nil {
		return errFio("make", d, err)
	}
	return nil
}

// done
func ResetFile(f Filename) error {
	b, err := ExistsFile(f)
	if err != nil {
		return errFio("check if exists", f, err)
	}
	if b {
		if e := TouchFile(f); e != nil {
			return errFio("touch", f, e)
		}
	}
	err = os.Truncate(string(f), 0)
	if err != nil {
		return errFio("truncate", f, err)
	}
	return nil
}

func TouchFile(fn Filename) error {
	if e := CheckFile(fn); e != nil {
		return errChk(fn, e)
	}
	b, erre := ExistsFile(fn)
	if erre != nil {
		return errFio("check if exists", fn, erre)
	}
	if b {
		t := time.Now().Local()
		if e := os.Chtimes(string(fn), t, t); e != nil {
			return errFio("chtimes", fn, e)
		}
	} else {
		f, erro := OpenFile(fn)
		if erro != nil {
			return errFio("open", fn, erro)
		}
		if e := f.Close(); e != nil {
			return errFio("close", fn, e)
		}
	}
	return nil
}

func OpenFile(fn Filename) (*os.File, error) {
	if e := CheckFile(fn); e != nil {
		return nil, errChk(fn, e)
	}
	f, err := os.OpenFile(string(fn), flags, fperm)
	if err != nil {
		return nil, errFio("open", fn, err)
	}
	return f, nil
}
