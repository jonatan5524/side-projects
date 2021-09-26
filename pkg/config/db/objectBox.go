package config

import (
	"errors"
	"os"
	"path/filepath"

	coreErrors "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/objectbox/objectbox-go/objectbox"
)

var (
	ERR_INIT_DB = errors.New("error initalize db")
)

func InitDB() (*objectbox.ObjectBox, error) {
	objectbox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()

	if err != nil {
		return nil, coreErrors.NewDBError("Initialize", ERR_INIT_DB)
	}

	return objectbox, nil
}

func InitTestDB() (*objectbox.ObjectBox, string) {
	tempPath := filepath.Join(os.TempDir(), "testdb")
	objectbox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Directory(tempPath).Build()

	if err != nil {
		panic("unable to create test db")
	}

	return objectbox, tempPath
}
