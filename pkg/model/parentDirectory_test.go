package model

import (
	"os"
	"testing"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	"github.com/jonatan5524/side-projects-manager/pkg/model/mocks"
	"github.com/jonatan5524/side-projects-manager/pkg/util"
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
	expected := ParentDirectory{Path: path, LastUpdated: retFileInfo.ModTime(), Projects: []*Project{}}

	parentDir, err := NewParentDirectory(path, mockDirectoryGetter.Execute)

	assert.Nil(t, err)
	assert.Equal(t, expected, parentDir)
}

func TestNewParentDirectory_InvalidPath(t *testing.T) {
	path := "tmp"
	retErr := core.NewIOError(path, util.ErrFileNotExists)

	mockDirectoryGetter := new(mocks.DirectoryGetter)
	mockDirectoryGetter.On("Execute", mock.Anything).Return(nil, retErr)

	parentDir, err := NewParentDirectory(path, mockDirectoryGetter.Execute)

	assert.Equal(t, parentDir, NilParentDirectory)
	assert.ErrorIs(t, err, retErr)
}
