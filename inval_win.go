//go:build windows

package tsfio

var (
	invalDir [8]Directory = [8]Directory{
		"C:\\Program Files",
		"C:\\Program Files (x86)",
		"C:\\Windows",
		"C:\\Windows\\System32",
		"C:\\",
		"C:\\System Volume Information",
		"C:\\Windows\\WinSxS",
		"C:\\Windows\\SysWOW64",
	}
)
