package repository

import (
	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/objectbox/objectbox-go/objectbox"
)

type ParentDirectoryObjectBoxRepository struct {
	box *model.ParentDirectoryBox
}

func NewParentDirectoryObjectBoxRepository(objectbox *objectbox.ObjectBox) ParentDirectoryRepository {
	return &ParentDirectoryObjectBoxRepository{model.BoxForParentDirectory(objectbox)}
}

func (parentDirRepo *ParentDirectoryObjectBoxRepository) Put(parentDir model.ParentDirectory) (uint64, error) {
	id, err := parentDirRepo.box.Put(&parentDir)

	if err != nil {
		return 0, core.NewDBError("PUT", err)
	}

	return id, nil
}
