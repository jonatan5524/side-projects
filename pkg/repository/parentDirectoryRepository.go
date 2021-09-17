package repository

import (
	"fmt"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/objectbox/objectbox-go/objectbox"
)

type ParentDirectoryRepositoryObjectBox struct {
	box *model.ParentDirectoryBox
}

func NewParentDirectoryRepositoryObjectBox(objectbox *objectbox.ObjectBox) *ParentDirectoryRepositoryObjectBox {
	return &ParentDirectoryRepositoryObjectBox{model.BoxForParentDirectory(objectbox)}
}

func (parentDirRepo *ParentDirectoryRepositoryObjectBox) Put(parentDir *model.ParentDirectory) (uint64, error) {
	id, err := parentDirRepo.box.Put(parentDir)

	if err != nil {
		return 0, &core.DBError{ActionRequested: fmt.Sprintf("Put: %v", parentDir), Err: err}
	}

	return id, nil
}
