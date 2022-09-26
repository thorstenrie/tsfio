//go:build windows

package tsfio

var (
	invalDir [4]Directory = [4]Directory{
		"C:\\Windows\\System32",
		"C:\\System Volume Information",
		"C:\\Windows\\WinSxS",
		"C:\\Windows\\SysWOW64",
	}
	invalFile [5]Filename = [5]Filename{
		"C:\\Program Files",
		"C:\\Program Files (x86)",
		"C:\\",
		"C:\\Windows",
		"C:\\pagefile.sys",
	}
)
