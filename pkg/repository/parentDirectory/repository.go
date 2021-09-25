package repository

import (
	"errors"

	coreErrors "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
	"github.com/objectbox/objectbox-go/objectbox"
)

var (
	ERR_PATH_NOT_FOUND = errors.New("path not found")
)

type ParentDirectoryObjectBoxRepository struct {
	box         *model.ParentDirectoryBox
	projectRepo repository.ProjectRepository
}

func NewParentDirectoryObjectBoxRepository(objectbox *objectbox.ObjectBox) ParentDirectoryRepository {
	return &ParentDirectoryObjectBoxRepository{
		box:         model.BoxForParentDirectory(objectbox),
		projectRepo: repository.NewProjectObjectBoxRepository(objectbox),
	}
}

func (parentDirRepo ParentDirectoryObjectBoxRepository) Put(parentDir model.ParentDirectory) (uint64, error) {
	id, err := parentDirRepo.box.Put(&parentDir)

	if err != nil {
		return 0, coreErrors.NewDBError("PUT", err)
	}

	return id, nil
}

func (repo ParentDirectoryObjectBoxRepository) GetAll() ([]*model.ParentDirectory, error) {
	directories, err := repo.box.GetAll()

	if err != nil {
		return []*model.ParentDirectory{}, coreErrors.NewDBError("GetAll", err)
	}

	return directories, nil
}

func (repo ParentDirectoryObjectBoxRepository) Delete(dir model.ParentDirectory) error {
	err := repo.box.Remove(&dir)

	if err != nil {
		return coreErrors.NewDBError("Remove", err)
	}

	return nil
}

func (repo ParentDirectoryObjectBoxRepository) DeleteByPath(path string) error {
	dir, err := repo.box.Query(model.ParentDirectory_.Path.Equals(path, true)).Limit(1).Find()

	if err != nil {
		return coreErrors.NewDBError("DeleteByPath", err)
	}

	if len(dir) == 0 {
		return coreErrors.NewDBError("DeleteByPath", ERR_PATH_NOT_FOUND)
	}

	err = repo.Delete(*dir[0])

	if err != nil {
		return err
	}

	repo.projectRepo.DeleteMany(dir[0].Projects...)

	return nil
}
