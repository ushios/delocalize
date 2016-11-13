package delocalize

import (
	"errors"
	"io/ioutil"
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

// IsLocalizedFile check file name
func IsLocalizedFile(path string) bool {
	if strings.HasSuffix(path, LocalizedFilename) {
		return true
	}

	return false
}

// directories from path
func directories(path string) ([]os.FileInfo, error) {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dl := []os.FileInfo{}
	for _, fi := range list {
		if fi.Mode() != os.ModeSymlink {
			if fi.IsDir() {
				dl = append(dl, fi)
			}
		}
	}

	return dl, nil
}
