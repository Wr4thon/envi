package errors

import (
	"syscall"

	"github.com/pkg/errors"
)

var (
	ErrMissingFile error = syscall.Errno(syscall.ENOENT)

	ErrKeyNotFound error = errors.New("key not found")
)
