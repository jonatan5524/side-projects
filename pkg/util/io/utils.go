package util

import (
	"errors"
	"io/ioutil"
	"os"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
)

var (
	ERR_FILE_NOT_EXISTS  = errors.New("folder of file is not exists")
	ERR_DIR_INVALID_TYPE = errors.New("path is not a directory type")
	FilterByDirectories  = func(curr os.FileInfo) bool {
		return curr.IsDir()
	}
)

type DirectoryGetter func(string) (os.FileInfo, error)
type FilterDirectoriesMethod func(os.FileInfo) bool

func GetDirectory(path string) (os.FileInfo, error) {
	directory, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return nil, core.NewIOError(path, ERR_FILE_NOT_EXISTS)
	} else if !directory.IsDir() {
		return nil, core.NewIOError(path, ERR_DIR_INVALID_TYPE)
	} else if err != nil {
		return nil, core.NewIOError(path, err)
	}

	return directory, nil
}

func ListDirectory(path string, filterByMethod FilterDirectoriesMethod) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if errors.Is(err, os.ErrNotExist) {
		return []os.FileInfo{}, core.NewIOError(path, ERR_FILE_NOT_EXISTS)
	} else if err != nil {
		return []os.FileInfo{}, core.NewIOError(path, err)
	}

	return FilterDirectories(files, filterByMethod), nil
}

func FilterDirectories(list []os.FileInfo, filterFunc FilterDirectoriesMethod) []os.FileInfo {
	filteredSlice := make([]os.FileInfo, 0)

	for _, value := range list {
		if filterFunc(value) {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return filteredSlice
}
