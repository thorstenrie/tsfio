//go:build windows

package tsfio

// Windows blocked Directories and Filenames.
// If a directory or their parents match InvalDir,
// tsfio functions will return an error. If a Filename
// matches InvalFile, tsfio functions will return an error.
var (
	InvalDir [4]Directory = [4]Directory{
		"C:\\Windows\\System32",
		"C:\\System Volume Information",
		"C:\\Windows\\WinSxS",
		"C:\\Windows\\SysWOW64",
	}
	InvalFile [5]Filename = [5]Filename{
		"C:\\Program Files",
		"C:\\Program Files (x86)",
		"C:\\",
		"C:\\Windows",
		"C:\\pagefile.sys",
	}
)
