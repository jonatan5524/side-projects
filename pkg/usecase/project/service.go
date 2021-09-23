package usecase

import (
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
)

type ProjectService struct {
	repository repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectUsecase {
	return &ProjectService{repo}
}

func (service *ProjectService) GetAll() ([]*model.Project, error) {
	return service.repository.GetAll()
}
