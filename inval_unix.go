//go:build !windows

package tsfio

// Linux blocked Directories and Filenames
// If Directories or their parents match invalDir,
// tsfio functions will return an error. If Filenames
// match invalFile, tsfio functions will return an error.
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
