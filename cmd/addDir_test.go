package cmd

import (
	"errors"
	"testing"
	"time"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	modelMock "github.com/jonatan5524/side-projects-manager/pkg/model/mocks"
	usecaseMock "github.com/jonatan5524/side-projects-manager/pkg/usecase/mocks"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
	"github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddParentDirectoryToDB_Success(t *testing.T) {
	path := testingUtils.CreateTempDirectory(t)
	parentDir := model.ParentDirectory{
		Path:        path,
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}
	parentDirWithId := parentDir
	parentDirWithId.Id = 1

	mockParentDirCtor := new(modelMock.ParentDirectoryConstructor)
	mockParentDirCtor.On("Execute", mock.Anything, mock.Anything).Return(parentDir, nil)

	mockService := new(usecaseMock.ParentDirectoryUsecase)
	mockService.On("Put", parentDir).Return(parentDirWithId, nil)

	testedFunc := func() {
		addParentDirectoryToDB(mockService, path, mockParentDirCtor.Execute)
	}

	assert.NotPanics(t, testedFunc)
}

func TestAddParentDirectoryToDB_InvalidPath(t *testing.T) {
	path := "/madeUp"
	expectedError := core.NewIOError(path, util.ErrFileNotExists)

	mockParentDirCtor := new(modelMock.ParentDirectoryConstructor)
	mockParentDirCtor.On("Execute", mock.Anything, mock.Anything).Return(model.NilParentDirectory, expectedError)

	mockService := new(usecaseMock.ParentDirectoryUsecase)

	testedFunc := func() {
		addParentDirectoryToDB(mockService, path, mockParentDirCtor.Execute)
	}

	assert.PanicsWithError(t, expectedError.Error(), testedFunc)
}

func TestAddParentDirectoryToDB_PutError(t *testing.T) {
	path := testingUtils.CreateTempDirectory(t)
	parentDir := model.ParentDirectory{
		Path:        path,
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}
	parentDirWithId := parentDir
	parentDirWithId.Id = 1
	expectedError := core.NewDBError("PUT", errors.New("db error"))

	mockParentDirCtor := new(modelMock.ParentDirectoryConstructor)
	mockParentDirCtor.On("Execute", mock.Anything, mock.Anything).Return(parentDir, nil)

	mockService := new(usecaseMock.ParentDirectoryUsecase)
	mockService.On("Put", parentDir).Return(model.NilParentDirectory, expectedError)

	testedFunc := func() {
		addParentDirectoryToDB(mockService, path, mockParentDirCtor.Execute)
	}

	assert.PanicsWithError(t, expectedError.Error(), testedFunc)
}
