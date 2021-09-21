package testingUtils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thanhpk/randstr"
)

var TempDir string = os.TempDir()

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

func CreateMultipleTempFile(t *testing.T, amount int) []string {
	files := []string{}

	for index := 0; index < amount; index++ {
		files = append(files, CreateTempFile(t))
	}

	return files
}

func CreateMultipleTempDirectories(t *testing.T, amount int) []string {
	files := []string{}

	for index := 0; index < amount; index++ {
		files = append(files, CreateTempDirectory(t))
	}

	return files
}

func RemoveTempFileSlice(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}

func tempName() string {
	return filepath.Join(TempDir, randstr.String(TEMP_NAME_LENGTH))
}
