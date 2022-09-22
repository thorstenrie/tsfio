package tsfio

import (
	"fmt"
	"io/fs"
	"os"
)

const (
	flags int         = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	fperm fs.FileMode = 0644
	dperm fs.FileMode = 0755
)

// done
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
	fileAppend Filename
	fileIn     Filename
}

// done
func AppendFile(a *Append) error {
	if a == nil {
		return fmt.Errorf("nil pointer")
	}

	if e := CheckFile(a.fileAppend); e != nil {
		return errChk(a.fileAppend, e)
	}
	if e := CheckFile(a.fileIn); e != nil {
		return errChk(a.fileIn, e)
	}
	f, erro := OpenFile(a.fileAppend)
	if erro != nil {
		return errFio("open", a.fileAppend, erro)
	}
	out, errr := ReadFile(a.fileIn)
	if errr != nil {
		return errFio("read", a.fileIn, errr)
	}
	if _, e := f.Write(out); e != nil {
		f.Close()
		return errFio(fmt.Sprintf("append file %v to", a.fileIn), a.fileAppend, e)
	}
	if e := f.Close(); e != nil {
		return errFio("close", a.fileAppend, e)
	}
	return nil
}

// done
func existsFile(fn Filename) (bool, error) {
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

	b, err := existsFile(f)
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
		return errDir("make", d, err)
	}
	return nil
}

// done
func ResetFile(f Filename) error {
	b, err := existsFile(f)
	if err != nil {
		return errFio("check if exists", f, err)
	}
	if b {
		if e := touchFile(f); e != nil {
			return errFio("touch", f, e)
		}
	}
	err = os.Truncate(string(f), 0)
	if err != nil {
		return errFio("truncate", f, err)
	}
	return nil
}

// done
func touchFile(fn Filename) error {
	if e := CheckFile(fn); e != nil {
		return errChk(fn, e)
	}
	f, err := OpenFile(fn)
	if err != nil {
		return errFio("open", fn, err)
	}
	if e := f.Close(); e != nil {
		return errFio("close", fn, e)
	}
	return nil
}

// done
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
