package repository_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()

	repo := repository.NewProjectObjectBoxRepository(testDB)
	project := model.Project{
		Name:        "project",
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
	}

	id, err := repo.Put(project)

	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestGetAll_Empty(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	repo := repository.NewProjectObjectBoxRepository(testDB)

	projects, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Empty(t, projects)
}

func TestGetAll_Length(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	const AMOUNT int = 3
	repo := repository.NewProjectObjectBoxRepository(testDB)

	expected := []model.Project{
		{
			Name:               "project",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project2",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project3",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
	}
	repo.Put(expected[0])
	repo.Put(expected[1])
	repo.Put(expected[2])

	projects, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Len(t, projects, AMOUNT)
}
