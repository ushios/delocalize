package delocalize

import (
	"errors"
	"strings"
)

const (
	// LocalizedFilename has filename
	LocalizedFilename = ".localized"
)

var (
	// ErrThisFileIsNotLocalizedFile is error when delete
	ErrThisFileIsNotLocalizedFile = errors.New("ErrThisFileIsNotLocalizedFile")
)

// IsLocalizedFile check file name
func IsLocalizedFile(path string) bool {
	if strings.HasSuffix(path, LocalizedFilename) {
		return true
	}

	return false
}
