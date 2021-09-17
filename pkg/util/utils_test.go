package util

import (
	"os"
	"testing"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	testingUtils "github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils/io"
	"github.com/stretchr/testify/assert"
)

func TestGetDirectoryIfExists_DirectoryExists(t *testing.T) {
	tempDirPath := testingUtils.CreateTempDirectory()
	defer os.Remove(tempDirPath)

	directory, err := GetDirectoryIfExists(tempDirPath)

	assert.Nil(t, err)
	assert.NotNil(t, directory)
}

func TestGetDirectoryIfExists_PathNotExists(t *testing.T) {
	directory, err := GetDirectoryIfExists("/MadeUpPath")
	
	assert.IsType(t, err, &core.IOError{})
	assert.Nil(t, directory)
}

func TestGetDirectoryIfExists_DirectoryIsFile(t *testing.T) {
	tempFilePath := testingUtils.CreateTempFile()
	defer os.Remove(tempFilePath)

	directory, err := GetDirectoryIfExists(tempFilePath)

	assert.IsType(t, err, &core.IOError{})
	assert.Nil(t, directory)
}
