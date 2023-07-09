//go:build !windows

package tsfio

// Linux blocked Directories and Filenames.
// If a directory or their parents match invalDir,
// tsfio functions will return an error. If a Filename
// matches invalFile, tsfio functions will return an error.
var (
	// blocked directories
	invalDir [4]Directory = [4]Directory{
		"/boot",
		"/dev",
		"/lost+found",
		"/proc",
	}
	// blocked filenames
	invalFile [14]Filename = [14]Filename{
		"/",
		"/bin",
		"/etc",
		"/home",
		"/lib",
		"/media",
		"/mnt",
		"/opt",
		"/root",
		"/sbin",
		"/srv",
		"/tmp",
		"/usr",
		"/var",
	}
)

// InvalDir returns the array of blocked directories. If a directory or their parents match InvalDir, tsfio
// functions will return an error.
func InvalDir() [4]Directory {
	return invalDir
}

// InvalFile returns the array of blocked filenames. If a Filename matches InvalFile, tsfio
// functions will return an error.
func InvalFile() [14]Filename {
	return invalFile
}
