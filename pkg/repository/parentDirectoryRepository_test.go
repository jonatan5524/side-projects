package repository

import (
	"testing"
	"time"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestPut_success(t *testing.T) {
	testDB := config.InitTestDB()
	defer testDB.Close()

	repo := NewParentDirectoryRepositoryObjectBox(testDB)
	parentDir := &model.ParentDirectory{ Path: "temp", LastUpdated: time.Now(), Projects: []*model.Project{}}

	id, err := repo.Put(parentDir)

	assert.Nil(t, err)
	assert.NotNil(t, id)
}
