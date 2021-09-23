package repository_test

import (
	"os"
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
