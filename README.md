# tsfio

[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/tsfio)](https://goreportcard.com/report/github.com/thorstenrie/tsfio)
[![CodeFactor](https://www.codefactor.io/repository/github/thorstenrie/tsfio/badge)](https://www.codefactor.io/repository/github/thorstenrie/tsfio)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorstenrie/tsfio)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/tsfio)](https://pkg.go.dev/mod/github.com/thorstenrie/tsfio)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorstenrie/tsfio)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorstenrie/tsfio)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorstenrie/tsfio)
![GitHub last commit](https://img.shields.io/github/last-commit/thorstenrie/tsfio)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorstenrie/tsfio)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorstenrie/tsfio)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorstenrie/tsfio)
![GitHub](https://img.shields.io/github/license/thorstenrie/tsfio)

[Go](https://go.dev/) package with a [simple](https://en.wikipedia.org/wiki/KISS_principle) API for file input output. The tsfio package is a supplement to the standard library and supplies additional functions for file input output operations, e.g., appending one file to another file.

- **Simple**: Without configuration, just function calls, and default flags are used
- **Resilient**: File input output operations on Linux and Windows system directories or files are blocked (see [inval_unix.go](https://github.com/thorstenrie/tsfio/blob/main/inval_unix.go) and [inval_win.go](https://github.com/thorstenrie/tsfio/blob/main/inval_win.go))
- **Tested**: Unit tests with high [code coverage](https://gocover.io/github.com/thorstenrie/tsfio)
- **Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std) and [tserr](https://github.com/thorstenrie/tserr)

## Defaults

All file input output operations on Linux and Windows system directories or
files are blocked (see [inval_unix.go](https://github.com/thorstenrie/tsfio/blob/main/inval_unix.go) and [inval_win.go](https://github.com/thorstenrie/tsfio/blob/main/inval_win.go)) and an error is returned.
All operations expect a directory or a regular file, return an error otherwise.
Default flags and file mode is used when opening files, creating files or directories
and when writing to files (with exceptions documented in the function descriptions)

- Files are opened read-write (os.O_RDWR).
- Data is appended when writing to file (os.O_APPEND).
- A file is created if it does not exist (os.O_CREATE).
- File mode and permission bits are 0644.
- Directory mode and permissions bits are 0755.

If an API call is not successful, a [tserr](https://github.com/thorstenrie/tserr) error in JSON format is returned.

## Usage

In the Go app, the package is imported with

```
import "github.com/thorstenrie/tsfio"
```

A Filename is the name of a regular file and may contain its path. A Directory is the name of a directory and may contain its path

```
type Filename string
type Directory string
```

CheckFile performs checks on Filename f and CheckDir performs checks on Directory d

```
func CheckFile(f Filename) error
func CheckDir(d Directory) error
```

All external functions contain a CheckFile or CheckDir call at the beginning.

```
func OpenFile(fn Filename) (*os.File, error)
func CloseFile(f *os.File) error
func WriteStr(fn Filename, s string) error
func WriteSingleStr(fn Filename, s string) error
func TouchFile(fn Filename) error
func ReadFile(f Filename) ([]byte, error)
func AppendFile(a *Append) error
func ExistsFile(fn Filename) (bool, error)
func RemoveFile(f Filename) error
func ResetFile(fn Filename) error
func CreateDir(d Directory) error
```

## Example

```
package main

import (
	"fmt"
	"os"

	"github.com/thorstenrie/tsfio"
)

func main() {
	f1, _ := os.CreateTemp("", "foo")
	fn1 := tsfio.Filename(f1.Name())

	f2, _ := os.CreateTemp("", "foo")
	fn2 := tsfio.Filename(f2.Name())

	tsfio.RemoveFile(fn1)
	b, _ := tsfio.ExistsFile(fn1)
	fmt.Println(b)

	tsfio.TouchFile(fn1)
	b, _ = tsfio.ExistsFile(fn1)
	fmt.Println(b)

	tsfio.WriteStr(fn1, "foo")
	tsfio.WriteStr(fn1, "foo")
	c, _ := tsfio.ReadFile(fn1)
	fmt.Println(string(c))

	tsfio.WriteStr(fn2, "foo")
	tsfio.AppendFile(&tsfio.Append{FileA: fn1, FileI: fn2})
	c, _ = tsfio.ReadFile(fn1)
	fmt.Println(string(c))

	tsfio.WriteSingleStr(fn1, "foo")
	c, _ = tsfio.ReadFile(fn1)
	fmt.Println(string(c))

	tsfio.ResetFile(fn1)
	c, _ = tsfio.ReadFile(fn1)
	fmt.Println(string(c))
}
```

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/tsfio)

[Gocover.io](https://gocover.io/github.com/thorstenrie/tsfio)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/tsfio)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Ftsfio)
