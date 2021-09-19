package model

import (
	"os"
	"time"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen -out objectbox

type ParentDirectory struct {
	Id          uint64
	Path        string
	LastUpdated time.Time `objectbox:"date"`
	Projects    []*Project
}

var NilParentDirectory = ParentDirectory{}

type DirectoryGetter func(string) (os.FileInfo, error)

func NewParentDirectory(path string, directoryGetter DirectoryGetter) (ParentDirectory, error) {
	directoryInfo, err := directoryGetter(path)

	if err != nil {
		return NilParentDirectory, err
	}

	return ParentDirectory{
		Path:        path,
		LastUpdated: directoryInfo.ModTime(),
		Projects:    []*Project{},
	}, nil
}
