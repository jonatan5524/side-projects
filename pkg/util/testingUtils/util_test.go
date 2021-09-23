package testingUtils_test

import (
	"os"
	"testing"

	"github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTempDirectory(t *testing.T) {
	path := testingUtils.CreateTempDirectory(t)
	defer os.Remove(path)

	_, err := os.Stat(path)

	assert.Nil(t, err)
}

func TestCreateTempFile(t *testing.T) {
	path := testingUtils.CreateTempFile(t)
	defer os.Remove(path)

	_, err := os.Stat(path)

	assert.Nil(t, err)
}

func TestCreateMultipleTempFile(t *testing.T) {
	amount := 4
	files := testingUtils.CreateMultipleTempFile(t, amount)
	defer testingUtils.RemoveTempFileSlice(files)

	assert.Len(t, files, amount)
	for _, path := range files {
		_, err := os.Stat(path)

		assert.Nil(t, err)
	}
}

func TestCreateMultipleTempDirectories(t *testing.T) {
	amount := 4
	files := testingUtils.CreateMultipleTempDirectories(t, amount)
	defer testingUtils.RemoveTempFileSlice(files)

	assert.Len(t, files, amount)
	for _, path := range files {
		_, err := os.Stat(path)

		assert.Nil(t, err)
	}
}
