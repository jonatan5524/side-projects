package cmd

import (
	"testing"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
	"github.com/jonatan5524/side-projects-manager/pkg/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProject(t *testing.T) {
	path := "tmp"

	mockService := new(mocks.ProjectUsecase)
	mockService.On("DeleteByPath", path).Return(nil)

	testedFunc := func() {
		deleteProject(mockService, path)
	}

	assert.NotPanics(t, testedFunc)
}

func TestDeleteProject_Fail(t *testing.T) {
	path := "tmp"
	err := core.NewDBError("RemoveByPath", repository.ERR_PROJECT_NOT_FOUND)

	mockService := new(mocks.ProjectUsecase)
	mockService.On("DeleteByPath", path).Return(err)

	testedFunc := func() {
		deleteProject(mockService, path)
	}

	assert.PanicsWithError(t, err.Error(), testedFunc)
}
