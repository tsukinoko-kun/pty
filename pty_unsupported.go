//go:build !linux && !darwin && !freebsd && !dragonfly && !netbsd && !openbsd && !solaris && !windows && !zos
// +build !linux,!darwin,!freebsd,!dragonfly,!netbsd,!openbsd,!solaris,!windows,!zos

package pty

import (
	"os"
)

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}
