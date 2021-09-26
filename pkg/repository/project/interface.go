package repository

import (
	"github.com/jonatan5524/side-projects-manager/pkg/model"
)

type ProjectRepository interface {
	GetAll() ([]*model.Project, error)
	Put(model.Project) (uint64, error)
	DeleteMany(...*model.Project) error
	Delete(model.Project) error
	DeleteByPath(string) error
}
