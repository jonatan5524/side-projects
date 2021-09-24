package repository

import (
	"github.com/jonatan5524/side-projects-manager/pkg/model"
)

type ParentDirectoryRepository interface {
	Put(model.ParentDirectory) (uint64, error)
	GetAll() ([]*model.ParentDirectory, error)
}
