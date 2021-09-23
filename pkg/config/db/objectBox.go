package config

import (
	"errors"
	"testing"

	coreErrors "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	testingUtils "github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils"
	"github.com/objectbox/objectbox-go/objectbox"
)

var (
	ErrInitDB = errors.New("error initalize db")
)

func InitDB() (*objectbox.ObjectBox, error) {
	objectbox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()

	if err != nil {
		return nil, coreErrors.NewDBError("Initialize", ErrInitDB)
	}

	return objectbox, nil
}

func InitTestDB(t *testing.T) *objectbox.ObjectBox {
	tempPath := testingUtils.CreateTempDirectory(t)
	objectbox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Directory(tempPath).Build()

	if err != nil {
		t.Fatal("unable to create test db")
	}

	return objectbox
}
