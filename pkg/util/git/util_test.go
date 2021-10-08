package util_test

import (
	"fmt"
	"os"
	"testing"

	util "github.com/jonatan5524/side-projects-manager/pkg/util/git"
	"github.com/jonatan5524/side-projects-manager/pkg/util/testingUtils"
	"github.com/stretchr/testify/assert"
)

func TestIsInVersionControl_False(t *testing.T) {
	tempFolder := testingUtils.CreateTempDirectory(t)
	defer os.Remove(tempFolder)

	actual, err := util.IsInVersionControl(tempFolder)

	assert.Nil(t, err)
	assert.False(t, actual)
}

func TestIsInVersionControl_True(t *testing.T) {
	tempFolder := testingUtils.CreateTempDirectory(t)
	defer os.Remove(tempFolder)

	err := os.Mkdir(tempFolder+"/.git", os.ModePerm)

	if err != nil {
		t.Fatal("error creating folder")
	}

	actual, err := util.IsInVersionControl(tempFolder)

	assert.Nil(t, err)
	assert.True(t, actual)
}

func TestParseRemoteURL_HTTPS(t *testing.T) {
	expected := "github.com"
	path := fmt.Sprintf("https://%s/madeup/repo", expected)

	actual := util.ParseRemoteURL(path)

	assert.Equal(t, expected, actual)
}

func TestParseRemoteURL_SSH(t *testing.T) {
	expected := "github.com"
	path := fmt.Sprintf("git@%s:madeup/repo.git", expected)

	actual := util.ParseRemoteURL(path)

	assert.Equal(t, expected, actual)
}

func TestParseRemoteURL_Unknown(t *testing.T) {
	expected := "-"
	path := "ftp://github.com:madeup/repo.git"

	actual := util.ParseRemoteURL(path)

	assert.Equal(t, expected, actual)
}
