package model

import (
	"github.com/jonatan5524/side-projects-manager/pkg/util"
	"time"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type ParentDirectory struct {
	Id uint64
	Path string
	LastUpdated time.Time `objectbox:"date"`
	Projects []*Project
}

// TODO: making tests
func NewParentDirectory(path string) (*ParentDirectory, error) {
	directoryInfo, err := util.GetDirectoryIfExists(path)
	
	if err != nil {
		return nil, err
	}

	return &ParentDirectory{
		Path: path,
		LastUpdated: directoryInfo.ModTime(),
		Projects: []*Project{},
	}, nil
}

type ParentDirectoryRepository interface {
	Put(*ParentDirectory) (uint64, error)
}
