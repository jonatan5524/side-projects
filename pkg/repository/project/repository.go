package repository

import (
	"errors"

	core "github.com/jonatan5524/side-projects-manager/pkg/core"
	coreErrors "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/objectbox/objectbox-go/objectbox"
)

var (
	ERR_PROJECT_NOT_FOUND = errors.New("project not found")
)

type ProjectObjectBoxRepository struct {
	box *model.ProjectBox
}

func NewProjectObjectBoxRepository(objectbox *objectbox.ObjectBox) ProjectRepository {
	return &ProjectObjectBoxRepository{model.BoxForProject(objectbox)}
}

func (repo *ProjectObjectBoxRepository) GetAllFilteredGit() ([]*model.Project, error) {
	projects, err := repo.box.Query(model.Project_.HaveVersionControl.Equals(true)).Find()

	if err != nil {
		return []*model.Project{}, coreErrors.NewDBError("GetAll", err)
	}

	return projects, nil
}

func (repo *ProjectObjectBoxRepository) GetAll() ([]*model.Project, error) {
	projects, err := repo.box.GetAll()

	if err != nil {
		return []*model.Project{}, coreErrors.NewDBError("GetAll", err)
	}

	return projects, nil
}

func (repo *ProjectObjectBoxRepository) GetRecent(amount int) ([]*model.Project, error) {
	projects, err := repo.box.Query(model.Project_.LastUpdated.OrderDesc()).Limit(uint64(amount)).Find()

	if err != nil {
		return []*model.Project{}, coreErrors.NewDBError("GetRecent", err)
	}

	return projects, nil
}

func (repo *ProjectObjectBoxRepository) Put(project model.Project) (uint64, error) {
	id, err := repo.box.Put(&project)

	if err != nil {
		return 0, coreErrors.NewDBError("PUT", err)
	}

	return id, nil
}

func (repo *ProjectObjectBoxRepository) DeleteMany(projects ...*model.Project) error {
	_, err := repo.box.RemoveMany(projects...)

	if err != nil {
		return coreErrors.NewDBError("RemoveMany", err)
	}

	return nil
}

func (repo ProjectObjectBoxRepository) Delete(project model.Project) error {
	err := repo.box.Remove(&project)

	if err != nil {
		return coreErrors.NewDBError("Remove", err)
	}

	return nil
}

func (repo ProjectObjectBoxRepository) DeleteByPath(path string) error {
	projects, err := repo.box.Query(model.Project_.Path.Equals(path, core.CASE_SENSATIVE)).Limit(1).Find()

	if err != nil {
		return coreErrors.NewDBError("DeleteByPath", err)
	}

	if len(projects) == 0 {
		return coreErrors.NewDBError("DeleteByPath", ERR_PROJECT_NOT_FOUND)
	}

	err = repo.Delete(*projects[0])

	if err != nil {
		return err
	}

	return nil
}

func (repo ProjectObjectBoxRepository) Get(path string) (model.Project, error) {
	projects, err := repo.box.Query(model.Project_.Path.Equals(path, core.CASE_SENSATIVE)).Limit(1).Find()

	if err != nil {
		return model.NilProject, coreErrors.NewDBError("Get", err)
	}

	if len(projects) == 0 {
		return model.NilProject, coreErrors.NewDBError("Get", ERR_PROJECT_NOT_FOUND)
	}

	return *projects[0], nil
}
