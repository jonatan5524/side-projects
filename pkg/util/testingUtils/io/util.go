package testingUtils

import (
	"fmt"
	"os"

	"github.com/thanhpk/randstr"
)

const TEMP_NAME_LENGTH = 5

func CreateTempDirectory() string {
	tempDir := os.TempDir()
	path := fmt.Sprintf("%s/%s", tempDir, randstr.String(TEMP_NAME_LENGTH))
	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		panic("testingUtils: error accured while creating temp dir")
	}

	return path
}

func CreateTempFile() string {
	tempDir := os.TempDir()
	path := fmt.Sprintf("%s/%s", tempDir, randstr.String(TEMP_NAME_LENGTH))
	file, err := os.Create(path)

	if err != nil {
		panic("testingUtils: error accured while creating temp file")
	}

	return file.Name()
}
