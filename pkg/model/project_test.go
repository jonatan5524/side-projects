package model_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	coreError "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	coreIO "github.com/jonatan5524/side-projects-manager/pkg/core/io"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
	"github.com/jonatan5524/side-projects-manager/pkg/util/io/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewProject_ValidPath(t *testing.T) {
	path := os.TempDir()
	retFileInfo, err := os.Stat(path)

	if err != nil {
		t.Fatal(err)
	}

	mockDirectoryGetter := new(mocks.DirectoryGetter)
	mockDirectoryGetter.On("Execute", mock.Anything).Return(retFileInfo, nil)
	expected := model.Project{
		Path:               path,
		Name:               retFileInfo.Name(),
		LastUpdated:        retFileInfo.ModTime(),
		HaveVersionControl: false,
	}

	project, err := model.NewProject(path, mockDirectoryGetter.Execute)

	assert.Nil(t, err)
	assert.Equal(t, expected, project)
}

func TestNewProject_InvalidPath(t *testing.T) {
	path := "madeup"
	retErr := coreError.NewIOError(path, util.ErrFileNotExists)

	mockDirectoryGetter := new(mocks.DirectoryGetter)
	mockDirectoryGetter.On("Execute", mock.Anything).Return(nil, retErr)

	project, err := model.NewProject(path, mockDirectoryGetter.Execute)

	assert.Equal(t, project, model.NilProject)
	assert.ErrorIs(t, err, retErr)
}

func TestConvertProjectToTablerSlice(t *testing.T) {
	projects := []*model.Project{
		{
			Name:               "project",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project2",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project3",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
	}

	tablers := model.ConvertProjectToTablerSlice(projects)

	assert.IsType(t, []coreIO.Tabler{}, tablers)
	assert.Len(t, tablers, len(projects))
}
