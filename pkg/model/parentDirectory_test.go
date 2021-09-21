package model_test

import (
	"os"
	"testing"
	"time"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
	"github.com/jonatan5524/side-projects-manager/pkg/util/io/mocks"
	"github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewParentDirectory_ValidPath(t *testing.T) {
	path := os.TempDir()
	retFileInfo, err := os.Stat(path)

	if err != nil {
		t.Fatal(err)
	}

	mockDirectoryGetter := new(mocks.DirectoryGetter)
	mockDirectoryGetter.On("Execute", mock.Anything).Return(retFileInfo, nil)
	expected := model.ParentDirectory{Path: path, LastUpdated: retFileInfo.ModTime(), Projects: []*model.Project{}}

	parentDir, err := model.NewParentDirectory(path, mockDirectoryGetter.Execute)

	assert.Nil(t, err)
	assert.Equal(t, expected, parentDir)
}

func TestNewParentDirectory_InvalidPath(t *testing.T) {
	path := "madeup"
	retErr := core.NewIOError(path, util.ErrFileNotExists)

	mockDirectoryGetter := new(mocks.DirectoryGetter)
	mockDirectoryGetter.On("Execute", mock.Anything).Return(nil, retErr)

	parentDir, err := model.NewParentDirectory(path, mockDirectoryGetter.Execute)

	assert.Equal(t, parentDir, model.NilParentDirectory)
	assert.ErrorIs(t, err, retErr)
}

func TestPut_EmptyDirectory(t *testing.T) {
	path := testingUtils.CreateTempDirectory(t)
	defer os.Remove(path)

	parentDir := model.ParentDirectory{
		Path:        path,
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}

	err := parentDir.LoadProjects()

	assert.Nil(t, err)
	assert.Empty(t, parentDir.Projects)
}

func TestPut_WithDirs(t *testing.T) {
	const AMOUNT int = 4

	path := testingUtils.CreateTempDirectory(t)
	testingUtils.TempDir = path
	testingUtils.CreateMultipleTempDirectories(t, AMOUNT)

	defer func() { testingUtils.TempDir = os.TempDir() }()
	defer os.Remove(path)

	parentDir := model.ParentDirectory{
		Path:        path,
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}

	err := parentDir.LoadProjects()

	assert.Nil(t, err)
	assert.Len(t, parentDir.Projects, AMOUNT)
}
