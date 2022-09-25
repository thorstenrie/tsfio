//go:build !windows

package tsfio

var (
	invalDir [7]Directory = [7]Directory{
		"/",
		"/boot",
		"/dev",
		"/lost+found",
		"/media",
		"/mnt",
		"/proc",
	}
)
