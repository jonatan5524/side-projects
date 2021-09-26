package usecase

import "github.com/jonatan5524/side-projects-manager/pkg/model"

type ProjectUsecase interface {
	GetAll() ([]*model.Project, error)
	DeleteByPath(string) error
	Delete(model.Project) error
}
