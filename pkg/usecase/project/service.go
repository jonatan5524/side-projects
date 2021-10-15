package usecase

import (
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
)

const RECENT_AMOUNT = 4

type ProjectService struct {
	repository repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectUsecase {
	return &ProjectService{repo}
}

func (service *ProjectService) GetAll() ([]*model.Project, error) {
	return service.repository.GetAll()
}

func (service ProjectService) DeleteByPath(path string) error {
	return service.repository.DeleteByPath(path)
}

func (service ProjectService) Delete(project model.Project) error {
	return service.repository.Delete(project)
}

func (service ProjectService) Get(path string) (model.Project, error) {
	return service.repository.Get(path)
}

func (service *ProjectService) GetRecent() ([]*model.Project, error) {
	return service.repository.GetRecent(RECENT_AMOUNT)
}
