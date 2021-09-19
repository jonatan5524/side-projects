package testingUtils

import (
	"fmt"
	"os"
	"testing"

	"github.com/thanhpk/randstr"
)

const TEMP_NAME_LENGTH = 5

func CreateTempDirectory(t *testing.T) string {
	path := tempName()
	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		t.Fatal("testingUtils: error accured while creating temp dir")
	}

	return path
}

func CreateTempFile(t *testing.T) string {
	file, err := os.Create(tempName())

	if err != nil {
		t.Fatal("testingUtils: error accured while creating temp file")
	}

	return file.Name()
}

func tempName() string {
	tempDir := os.TempDir()
	return fmt.Sprintf("%s/%s", tempDir, randstr.String(TEMP_NAME_LENGTH))
}
