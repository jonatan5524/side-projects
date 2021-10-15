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

func TestDeleteByPath_NotFound(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)

	path := filepath.Join(os.TempDir(), "project")

	err := repo.DeleteByPath(path)

	assert.IsType(t, err, &core.DBError{})
}

func TestDeleteByPath_Found(t *testing.T) {
	t.Cleanup(cleanDB)
	defer testDB.Close()
	const AMOUNT int = 3
	repo := repository.NewProjectObjectBoxRepository(testDB)

	project := model.Project{
		Name:        "project",
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
	}

	repo.Put(project)

	err := repo.DeleteByPath(project.Path)

	assert.Nil(t, err)
}

func TestGet_Found(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)
	project := model.Project{
		Name:        "project",
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
	}

	id, _ := repo.Put(project)
	project.Id = id

	retProject, err := repo.Get(project.Path)

	assert.Nil(t, err)
	assert.Equal(t, project.Id, retProject.Id)
}

func TestGet_NotFound(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)
	project := model.Project{
		Name:        "project",
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
	}

	retProject, err := repo.Get(project.Path)

	assert.Equal(t, model.NilProject, retProject)
	assert.IsType(t, err, &core.DBError{})
}

func TestGetRecent(t *testing.T) {
	t.Cleanup(cleanDB)

	repo := repository.NewProjectObjectBoxRepository(testDB)

	const AMOUNT = 2

	projects := []model.Project{
		{
			Name:               "project",
			Path:               filepath.Join(os.TempDir(), "project"),
			LastUpdated:        time.Now(),
			HaveVersionControl: false,
		},
		{
			Name:               "project2",
			Path:               filepath.Join(os.TempDir(), "project2"),
			LastUpdated:        time.Now().AddDate(0, -1, 0),
			HaveVersionControl: false,
		},
		{
			Name:               "project3",
			Path:               filepath.Join(os.TempDir(), "project3"),
			LastUpdated:        time.Now().AddDate(0, 0, -1),
			HaveVersionControl: false,
		},
	}
	repo.Put(projects[0])
	repo.Put(projects[1])
	repo.Put(projects[2])

	expected := []model.Project{
		projects[0],
		projects[2],
	}

	actual, err := repo.GetRecent(AMOUNT)

	assert.Nil(t, err)
	assert.Len(t, actual, AMOUNT)
	assert.Equal(t, expected[0].Name, actual[0].Name)
	assert.Equal(t, expected[1].Name, actual[1].Name)
}
