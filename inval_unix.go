//go:build !windows

package tsfio

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
