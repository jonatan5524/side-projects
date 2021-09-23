package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/io"
	coreMocks "github.com/jonatan5524/side-projects-manager/pkg/core/io/mocks"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	usecaseMocks "github.com/jonatan5524/side-projects-manager/pkg/usecase/mocks"
	"github.com/stretchr/testify/mock"
)

func TestPrintListProjects_NotVerbose(t *testing.T) {
	mockService := new(usecaseMocks.ProjectUsecase)
	mockService.On("GetAll").Return([]*model.Project{}, nil)
	mockOutputHandler := new(coreMocks.OutputHandler)
	mockOutputHandler.On("PrintString", mock.Anything).Return()
	defer func() { outputHandler = core.NewOutputStdout() }()

	outputHandler = mockOutputHandler

	printListProjects(mockService, false)

	mockOutputHandler.AssertNotCalled(t, "PrintTable")
}

func TestPrintListProjects_Verbose(t *testing.T) {
	mockService := new(usecaseMocks.ProjectUsecase)
	mockService.On("GetAll").Return([]*model.Project{}, nil)
	mockOutputHandler := new(coreMocks.OutputHandler)
	mockOutputHandler.On("PrintString", mock.Anything).Return()
	mockOutputHandler.On("PrintTable", []core.Tabler{}).Return()
	defer func() { outputHandler = core.NewOutputStdout() }()

	outputHandler = mockOutputHandler

	printListProjects(mockService, true)

	mockOutputHandler.AssertCalled(t, "PrintTable", []core.Tabler{})
}

func TestPrintNormalList(t *testing.T) {
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
	mockOutputHandler := new(coreMocks.OutputHandler)
	mockOutputHandler.On("PrintString", mock.Anything).Return()
	defer func() { outputHandler = core.NewOutputStdout() }()

	outputHandler = mockOutputHandler

	printNormalList(projects)

	mockOutputHandler.AssertNumberOfCalls(t, "PrintString", len(projects))
}
