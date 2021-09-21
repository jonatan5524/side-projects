package util_test

import (
	"os"
	"testing"

	"github.com/jonatan5524/side-projects-manager/pkg/core"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
	testingUtils "github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils"
	"github.com/stretchr/testify/assert"
)

func TestGetDirectory_DirectoryExists(t *testing.T) {
	tempDirPath := testingUtils.CreateTempDirectory(t)
	defer os.Remove(tempDirPath)

	directory, err := util.GetDirectory(tempDirPath)

	assert.Nil(t, err)
	assert.NotNil(t, directory)
}

func TestGetDirectory_PathNotExists(t *testing.T) {
	path := "/MadeUpPath"
	expectedError := core.NewIOError(path, util.ErrFileNotExists)

	directory, err := util.GetDirectory("/MadeUpPath")

	assert.Equal(t, expectedError, err)
	assert.Nil(t, directory)
}

func TestGetDirectory_DirectoryIsFile(t *testing.T) {
	tempFilePath := testingUtils.CreateTempFile(t)
	defer os.Remove(tempFilePath)
	expectedError := core.NewIOError(tempFilePath, util.ErrDirInvalidType)

	directory, err := util.GetDirectory(tempFilePath)

	assert.Equal(t, expectedError, err)
	assert.Nil(t, directory)
}

func TestFilterDirectories_EmptyList(t *testing.T) {
	list := []os.FileInfo{}

	filteredList := util.FilterDirectories(list, util.FilterByDirectories)

	assert.Empty(t, filteredList)
}

func TestFilterDirectories_AllFilesList(t *testing.T) {
	const AMOUNT int = 4
	tempFilesPath := testingUtils.CreateMultipleTempFile(t, AMOUNT)

	defer testingUtils.RemoveTempFileSlice(tempFilesPath)

	list := convertPathSliceToFilesInfo(t, tempFilesPath)

	filteredList := util.FilterDirectories(list, util.FilterByDirectories)

	assert.Empty(t, filteredList)
}

func TestFilterDirectories_FilesAndDirectories(t *testing.T) {
	const FILE_AMOUNT int = 4
	const FOLDERS_AMOUNT int = 3
	tempFilesPath := testingUtils.CreateMultipleTempFile(t, FILE_AMOUNT)
	tempDirectoryPath := testingUtils.CreateMultipleTempDirectories(t, FOLDERS_AMOUNT)

	defer testingUtils.RemoveTempFileSlice(tempFilesPath)
	defer testingUtils.RemoveTempFileSlice(tempDirectoryPath)

	pathList := append(tempFilesPath, tempDirectoryPath...)
	filesList := convertPathSliceToFilesInfo(t, pathList)

	filteredList := util.FilterDirectories(filesList, util.FilterByDirectories)

	assert.Len(t, filteredList, FOLDERS_AMOUNT)
}

func convertPathSliceToFilesInfo(t *testing.T, paths []string) []os.FileInfo {
	files := []os.FileInfo{}

	for _, path := range paths {
		file, err := os.Stat(path)

		if err != nil {
			t.Fatal("utils_test: error accured while converting paths to file info")
		}

		files = append(files, file)
	}

	return files
}

func TestListDirectory_PathNotExists(t *testing.T) {
	path := "/madeUpPath"
	filterMethodTrue := func(os.FileInfo) bool {
		return true
	}
	expectedError := core.NewIOError(path, util.ErrFileNotExists)

	directories, err := util.ListDirectory(path, filterMethodTrue)

	assert.Equal(t, expectedError, err)
	assert.Empty(t, directories)
}

func TestListDirectory_PathExists(t *testing.T) {
	path := testingUtils.CreateTempDirectory(t)
	defer os.Remove(path)
	filterMethodTrue := func(os.FileInfo) bool {
		return true
	}

	directories, err := util.ListDirectory(path, filterMethodTrue)

	assert.Nil(t, err)
	assert.Empty(t, directories)
}
