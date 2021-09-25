package cmd

import (
	"testing"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
	"github.com/jonatan5524/side-projects-manager/pkg/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteParentDirectory_Success(t *testing.T) {
	path := "tmp"

	mockService := new(mocks.ParentDirectoryUsecase)
	mockService.On("DeleteByPath", path).Return(nil)

	testedFunc := func() {
		deleteDirectory(mockService, path)
	}

	assert.NotPanics(t, testedFunc)
}

func TestDeleteParentDirectory_Fail(t *testing.T) {
	path := "tmp"
	err := core.NewDBError("RemoveByPath", repository.ERR_PATH_NOT_FOUND)

	mockService := new(mocks.ParentDirectoryUsecase)
	mockService.On("DeleteByPath", path).Return(err)

	testedFunc := func() {
		deleteDirectory(mockService, path)
	}

	assert.PanicsWithError(t, err.Error(), testedFunc)
}
