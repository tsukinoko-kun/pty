//go:build !zos
// +build !zos

package pty

import (
	"testing"
)

func getNonBlockingFile(t *testing.T, file Pty, _ string) Pty {
	t.Helper()
	return file
}
