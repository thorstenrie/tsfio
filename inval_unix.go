//go:build !windows

package tsfio

// Linux blocked Directories and Filenames.
// If a directory or their parents match InvalDir,
// tsfio functions will return an error. If a Filename
// matches InvalFile, tsfio functions will return an error.
var (
	// blocked directories
	InvalDir [4]Directory = [4]Directory{
		"/boot",
		"/dev",
		"/lost+found",
		"/proc",
	}
	// blocked filenames
	InvalFile [14]Filename = [14]Filename{
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
