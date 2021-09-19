package cmd

import (
	"errors"
	"testing"
	"time"

	cmdMock "github.com/jonatan5524/side-projects-manager/cmd/mocks"
	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	usecaseMock "github.com/jonatan5524/side-projects-manager/pkg/usecase/mocks"
	"github.com/jonatan5524/side-projects-manager/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddParentDirectoryToDB_Success(t *testing.T) {
	path := "/tmp"
	parentDir := model.ParentDirectory{
		Path:        path,
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}
	parentDirWithId := parentDir
	parentDirWithId.Id = 1

	mockParentDirCtor := new(cmdMock.ParentDirectoryConstructor)
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

	mockParentDirCtor := new(cmdMock.ParentDirectoryConstructor)
	mockParentDirCtor.On("Execute", mock.Anything, mock.Anything).Return(model.NilParentDirectory, expectedError)

	mockService := new(usecaseMock.ParentDirectoryUsecase)

	testedFunc := func() {
		addParentDirectoryToDB(mockService, path, mockParentDirCtor.Execute)
	}

	assert.PanicsWithError(t, expectedError.Error(), testedFunc)
}

func TestAddParentDirectoryToDB_PutError(t *testing.T) {
	path := "/tmp"
	parentDir := model.ParentDirectory{
		Path:        path,
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}
	parentDirWithId := parentDir
	parentDirWithId.Id = 1
	expectedError := core.NewDBError("PUT", errors.New("db error"))

	mockParentDirCtor := new(cmdMock.ParentDirectoryConstructor)
	mockParentDirCtor.On("Execute", mock.Anything, mock.Anything).Return(parentDir, nil)

	mockService := new(usecaseMock.ParentDirectoryUsecase)
	mockService.On("Put", parentDir).Return(model.NilParentDirectory, expectedError)

	testedFunc := func() {
		addParentDirectoryToDB(mockService, path, mockParentDirCtor.Execute)
	}

	assert.PanicsWithError(t, expectedError.Error(), testedFunc)
}
