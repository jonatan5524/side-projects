package repository_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()

	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)
	parentDir := model.ParentDirectory{Path: os.TempDir(), LastUpdated: time.Now(), Projects: []*model.Project{}}

	id, err := repo.Put(parentDir)

	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestGetAll_Empty(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)

	directories, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Empty(t, directories)
}

func TestGetAll_Length(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	const AMOUNT int = 3
	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)

	expected := []model.ParentDirectory{
		{
			Path:        filepath.Join(os.TempDir(), "project"),
			LastUpdated: time.Now(),
			Projects:    []*model.Project{},
		},
		{
			Path:        filepath.Join(os.TempDir(), "project"),
			LastUpdated: time.Now(),
			Projects:    []*model.Project{},
		},
		{
			Path:        filepath.Join(os.TempDir(), "project"),
			LastUpdated: time.Now(),
			Projects:    []*model.Project{},
		},
	}
	repo.Put(expected[0])
	repo.Put(expected[1])
	repo.Put(expected[2])

	projects, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Len(t, projects, AMOUNT)
}
