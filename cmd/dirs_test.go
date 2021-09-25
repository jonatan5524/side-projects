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

func TestPrintListDirectories_NotVerbose(t *testing.T) {
	directories := []*model.ParentDirectory{
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
	mockService := new(usecaseMocks.ParentDirectoryUsecase)
	mockService.On("GetAll").Return(directories, nil)
	mockOutputHandler := new(coreMocks.OutputHandler)
	mockOutputHandler.On("PrintString", mock.Anything).Return()
	defer func() { outputHandler = core.NewOutputStdout() }()

	outputHandler = mockOutputHandler

	printListDirectories(mockService, false)

	mockOutputHandler.AssertNotCalled(t, "PrintTable")
}

func TestPrintListDirectories_Verbose(t *testing.T) {
	directories := []*model.ParentDirectory{
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
	mockService := new(usecaseMocks.ParentDirectoryUsecase)
	mockService.On("GetAll").Return(directories, nil)
	mockOutputHandler := new(coreMocks.OutputHandler)
	mockOutputHandler.On("PrintString", mock.Anything).Return()
	mockOutputHandler.On("PrintTable", mock.Anything).Return()
	defer func() { outputHandler = core.NewOutputStdout() }()

	outputHandler = mockOutputHandler

	printListDirectories(mockService, true)

	mockOutputHandler.AssertNumberOfCalls(t, "PrintTable", 1)
}

func TestPrintNormalListDirectories(t *testing.T) {
	directories := []*model.ParentDirectory{
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
	mockOutputHandler := new(coreMocks.OutputHandler)
	mockOutputHandler.On("PrintString", mock.Anything).Return()
	defer func() { outputHandler = core.NewOutputStdout() }()

	outputHandler = mockOutputHandler

	printNormalListDirectories(directories)

	mockOutputHandler.AssertNumberOfCalls(t, "PrintString", len(directories))
}
