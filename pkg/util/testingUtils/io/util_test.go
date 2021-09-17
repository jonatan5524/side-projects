package testingUtils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTempDirectory_success(t *testing.T) {
	path := CreateTempDirectory()
	defer os.Remove(path)

	_, err := os.Stat(path);

	assert.Nil(t, err)
}

func TestCreateTempFile_success(t *testing.T) {
	path := CreateTempFile()
	defer os.Remove(path)

	_, err := os.Stat(path);

	assert.Nil(t, err)
}
