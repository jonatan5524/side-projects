package usecase_test

import (
	"testing"

	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/jonatan5524/side-projects-manager/pkg/repository/mocks"
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/project"
)

func TestGetAllFiltered_Git(t *testing.T) {
	repoMock := new(mocks.ProjectRepository)
	service := usecase.NewProjectService(repoMock)

	PROJECTS := []*model.Project{}
	var ERR_RET error = nil
	repoMock.On("GetAllFilteredGit").Return(PROJECTS, ERR_RET)

	service.GetAllFiltered("git")

	repoMock.AssertCalled(t, "GetAllFilteredGit")
}

func TestGetAllFiltered_InvalidFilter(t *testing.T) {
	repoMock := new(mocks.ProjectRepository)
	service := usecase.NewProjectService(repoMock)

	PROJECTS := []*model.Project{}
	var ERR_RET error = nil
	repoMock.On("GetAllFilteredGit").Return(PROJECTS, ERR_RET)

	service.GetAllFiltered("")

	repoMock.AssertNotCalled(t, "GetAllFilteredGit")
}
