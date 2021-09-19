package util

import (
	"os"
	"testing"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	testingUtils "github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils/io"
	"github.com/stretchr/testify/assert"
)

func TestGetDirectory_DirectoryExists(t *testing.T) {
	tempDirPath := testingUtils.CreateTempDirectory(t)
	defer os.Remove(tempDirPath)

	directory, err := GetDirectory(tempDirPath)

	assert.Nil(t, err)
	assert.NotNil(t, directory)
}

func TestGetDirectory_PathNotExists(t *testing.T) {
	path := "/MadeUpPath"
	expectedError := core.NewIOError(path, ErrFileNotExists)

	directory, err := GetDirectory("/MadeUpPath")

	assert.Equal(t, expectedError, err)
	assert.Nil(t, directory)
}

func TestGetDirectory_DirectoryIsFile(t *testing.T) {
	tempFilePath := testingUtils.CreateTempFile(t)
	defer os.Remove(tempFilePath)
	expectedError := core.NewIOError(tempFilePath, ErrDirInvalidType)

	directory, err := GetDirectory(tempFilePath)

	assert.Equal(t, expectedError, err)
	assert.Nil(t, directory)
}
