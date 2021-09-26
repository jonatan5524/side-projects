package repository_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/stretchr/testify/assert"
)

var (
	testDB *objectbox.ObjectBox = nil
	dbPath string               = ""
)

func TestMain(m *testing.M) {
	testDB, dbPath = config.InitTestDB()

	code := m.Run()

	os.Exit(code)
}

func cleanDB() {
	testDB.Close()
	os.RemoveAll(dbPath)

	testDB, dbPath = config.InitTestDB()
}

func TestPut(t *testing.T) {
	t.Cleanup(cleanDB)

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

func TestPut_Duplicate(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)
	project := model.Project{
		Name:        "project",
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
	}
	projectSecond := model.Project{
		Name:        "project",
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
	}

	id, err := repo.Put(project)
	assert.Nil(t, err)
	assert.NotNil(t, id)

	_, err = repo.Put(projectSecond)
	assert.NotNil(t, err)
}

func TestGetAll_Empty(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)

	projects, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Empty(t, projects)
}

func TestGetAll_Length(t *testing.T) {
	t.Cleanup(cleanDB)

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
			Path:               filepath.Join(os.TempDir(), "project2"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project3",
			Path:               filepath.Join(os.TempDir(), "project3"),
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

func TestDeleteMany_NotFound(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)

	projects := []*model.Project{
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

	err := repo.DeleteMany(projects...)

	assert.IsType(t, err, &core.DBError{})
}

func TestDeleteMany_Found(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)

	projects := []*model.Project{
		{
			Name:               "project",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project2",
			Path:               filepath.Join(os.TempDir(), "project2"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project3",
			Path:               filepath.Join(os.TempDir(), "project3"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
	}
	id, _ := repo.Put(*projects[0])
	projects[0].Id = id
	id, _ = repo.Put(*projects[1])
	projects[1].Id = id
	id, _ = repo.Put(*projects[2])
	projects[2].Id = id

	err := repo.DeleteMany(projects...)

	assert.Nil(t, err)

	projects, _ = repo.GetAll()

	assert.Empty(t, projects)
}
