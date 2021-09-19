package testingUtils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTempDirectory(t *testing.T) {
	path := CreateTempDirectory(t)
	defer os.Remove(path)

	_, err := os.Stat(path)

	assert.Nil(t, err)
}

func TestCreateTempFile(t *testing.T) {
	path := CreateTempFile(t)
	defer os.Remove(path)

	_, err := os.Stat(path)

	assert.Nil(t, err)
}
