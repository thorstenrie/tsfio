//go:build windows

package tsfio

// Windows blocked Directories and Filenames.
// If a directory or their parents match invalDir,
// tsfio functions will return an error. If a Filename
// matches invalFile, tsfio functions will return an error.
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

// InvalDir returns the array of blocked directories. If a directory or their parents match InvalDir, tsfio
// functions will return an error.
func InvalDir() [4]Directory {
	return invalDir
}

// InvalFile returns the array of blocked filenames. If a Filename matches InvalFile, tsfio
// functions will return an error.
func InvalFile() [5]Filename {
	return invalFile
}
