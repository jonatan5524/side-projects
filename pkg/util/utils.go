package util

import (
	"errors"
	"os"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
)

var (
	ErrFileNotExists  = errors.New("folder of file is not exists")
	ErrDirInvalidType = errors.New("path is not a directory type")
)

func GetDirectory(path string) (os.FileInfo, error) {
	directory, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return nil, core.NewIOError(path, ErrFileNotExists)
	} else if !directory.IsDir() {
		return nil, core.NewIOError(path, ErrDirInvalidType)
	}

	return directory, nil
}
