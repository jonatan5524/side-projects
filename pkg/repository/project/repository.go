package repository

import (
	"errors"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/objectbox/objectbox-go/objectbox"
)

var (
	ERR_PATH_NOT_FOUND = errors.New("path not found")
)

type ProjectObjectBoxRepository struct {
	box *model.ProjectBox
}

func NewProjectObjectBoxRepository(objectbox *objectbox.ObjectBox) ProjectRepository {
	return &ProjectObjectBoxRepository{model.BoxForProject(objectbox)}
}

func (repo *ProjectObjectBoxRepository) GetAll() ([]*model.Project, error) {
	projects, err := repo.box.GetAll()

	if err != nil {
		return []*model.Project{}, core.NewDBError("GetAll", err)
	}

	return projects, nil
}

func (repo *ProjectObjectBoxRepository) Put(project model.Project) (uint64, error) {
	id, err := repo.box.Put(&project)

	if err != nil {
		return 0, core.NewDBError("PUT", err)
	}

	return id, nil
}

func (repo *ProjectObjectBoxRepository) DeleteMany(projects ...*model.Project) error {
	_, err := repo.box.RemoveMany(projects...)

	if err != nil {
		return core.NewDBError("RemoveMany", err)
	}

	return nil
}

func (repo ProjectObjectBoxRepository) Delete(project model.Project) error {
	err := repo.box.Remove(&project)

	if err != nil {
		return core.NewDBError("Remove", err)
	}

	return nil
}

func (repo ProjectObjectBoxRepository) DeleteByPath(path string) error {
	projects, err := repo.box.Query(model.Project_.Path.Equals(path, true)).Limit(1).Find()

	if err != nil {
		return core.NewDBError("DeleteByPath", err)
	}

	if len(projects) == 0 {
		return core.NewDBError("DeleteByPath", ERR_PATH_NOT_FOUND)
	}

	err = repo.Delete(*projects[0])

	if err != nil {
		return err
	}

	return nil
}
