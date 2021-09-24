package usecase

import "github.com/jonatan5524/side-projects-manager/pkg/model"

type ParentDirectoryUsecase interface {
	Put(model.ParentDirectory) (model.ParentDirectory, error)
	GetAll() ([]*model.ParentDirectory, error)
}
