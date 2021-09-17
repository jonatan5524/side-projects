package config

import (
	"errors"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	testingUtils "github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils/io"
	"github.com/objectbox/objectbox-go/objectbox"
)

func InitDB() (*objectbox.ObjectBox, error) {
	objectbox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()

	if err != nil {
		return nil, core.NewDBError("Initialize", errors.New("trying to initizling db"))
	}

	return objectbox, nil
}

func InitTestDB() *objectbox.ObjectBox {
	tempPath := testingUtils.CreateTempDirectory()
	objectbox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Directory(tempPath).Build()

	if err != nil {
		panic("unable to create test db")
	}

	return objectbox
}
