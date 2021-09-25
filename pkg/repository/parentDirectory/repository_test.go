package repository_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
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

func TestDelete_NotFound(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)

	dir := model.ParentDirectory{
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}

	err := repo.Delete(dir)

	assert.IsType(t, err, &core.DBError{})
}

func TestDelete_Found(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)

	dir := model.ParentDirectory{
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}

	id, _ := repo.Put(dir)
	dir.Id = id

	err := repo.Delete(dir)

	assert.Nil(t, err)
}

func TestDeleteByPath_NotFound(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)

	path := filepath.Join(os.TempDir(), "project")

	err := repo.DeleteByPath(path)

	assert.IsType(t, err, &core.DBError{})
}

func TestDeleteByPath_Found(t *testing.T) {
	testDB := config.InitTestDB(t)
	defer testDB.Close()
	const AMOUNT int = 3
	repo := repository.NewParentDirectoryObjectBoxRepository(testDB)

	dir := model.ParentDirectory{
		Path:        filepath.Join(os.TempDir(), "project"),
		LastUpdated: time.Now(),
		Projects:    []*model.Project{},
	}

	repo.Put(dir)

	err := repo.DeleteByPath(dir.Path)

	assert.Nil(t, err)
}
