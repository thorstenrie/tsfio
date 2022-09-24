//go:build !windows

package tsfio

var (
	invalDir [8]Directory = [8]Directory{
		"",
		"/",
		"/boot",
		"/dev",
		"/lost+found",
		"/media",
		"/mnt",
		"/proc",
	}
)
