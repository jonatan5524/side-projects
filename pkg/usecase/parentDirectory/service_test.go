package usecase_test

import (
	"os"
	"testing"
	"time"

	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/jonatan5524/side-projects-manager/pkg/repository/mocks"
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/parentDirectory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPut(t *testing.T) {
	repoMock := new(mocks.ParentDirectoryRepository)
	service := usecase.NewParentDirectoryService(repoMock)
	parentDir := model.ParentDirectory{Path: os.TempDir(), LastUpdated: time.Now(), Projects: []*model.Project{}}

	const ID_RET uint64 = 1
	var ERR_RET error = nil
	repoMock.On("Put", mock.Anything).Return(ID_RET, ERR_RET)

	newParentDir, err := service.Put(parentDir)

	assert.Nil(t, err)
	assert.Equal(t, newParentDir.Id, ID_RET)
}
