package delocalize

import (
	"errors"
	"os"
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

// Delocalize directory with searching
func Delocalize(path string) error {
	// list, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	return err
	// }
	//
	// for _, fi := range list {
	//
	// }

	return nil
}

// delete path name file
func delete(path string) error {
	if !IsLocalizedFile(path) {
		return ErrThisFileIsNotLocalizedFile
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func IsLocalizedFile(path string) bool {
	if strings.HasSuffix(path, LocalizedFilename) {
		return true
	}

	return false
}
