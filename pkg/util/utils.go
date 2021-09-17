package util

import (
	"errors"
	"os"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
)

func GetDirectoryIfExists(path string) (os.FileInfo, error) {
	directory, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return nil, core.NewIOError(path, errors.New("folder of file is not exists"))
	} else if !directory.IsDir() {
		return nil, core.NewIOError(path, errors.New("the directory is a file"))
	}

	return directory, nil
}

