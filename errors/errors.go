package errors

import (
	"syscall"

	"github.com/pkg/errors"
)

var (
	// ErrMissingFile provides the syscall.ENOENT error.
	ErrMissingFile error = syscall.Errno(syscall.ENOENT)

	// ErrKeyNotFound is used when you request a Key that is not registered.
	ErrKeyNotFound error = errors.New("key not found")
)
