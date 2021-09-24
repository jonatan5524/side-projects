package model

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	core "github.com/jonatan5524/side-projects-manager/pkg/core/io"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen clean
//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type ParentDirectory struct {
	Id          uint64
	Path        string
	LastUpdated time.Time `objectbox:"date"`
	Projects    []*Project
}

var NilParentDirectory = ParentDirectory{}

type ParentDirectoryConstructor func(string, util.DirectoryGetter) (ParentDirectory, error)

func NewParentDirectory(path string, directoryGetter util.DirectoryGetter) (ParentDirectory, error) {
	directoryInfo, err := directoryGetter(path)

	if err != nil {
		return NilParentDirectory, err
	}

	return ParentDirectory{
		Path:        path,
		LastUpdated: directoryInfo.ModTime(),
		Projects:    []*Project{},
	}, nil
}

func (parentDir *ParentDirectory) LoadProjects() error {
	projectsInfo, err := util.ListDirectory(parentDir.Path, util.FilterByDirectories)

	if err != nil {
		return err
	}

	projects := []*Project{}

	for _, fileInfo := range projectsInfo {
		project, err := NewProject(filepath.Join(parentDir.Path, fileInfo.Name()), util.GetDirectory)

		if err != nil {
			return err
		}

		projects = append(projects, &project)
	}

	parentDir.Projects = projects

	return nil
}

func (parentDir ParentDirectory) String() string {
	return fmt.Sprintf(`ParentDirectory {
	Id: %d
	Path: %s
	LastUpdated: %v
	Projects: %v
}`,
		parentDir.Id,
		parentDir.Path,
		parentDir.LastUpdated,
		parentDir.Projects,
	)
}

func (ParentDirectory) TableHeader() table.Row {
	return table.Row{
		"Path",
		"Last Updated",
		"Number of projects",
	}
}

func (dir ParentDirectory) TableData() table.Row {
	return table.Row{
		dir.Path,
		dir.LastUpdated,
		len(dir.Projects),
	}
}

func ConvertParentDirectoryToTablerSlice(dirs []*ParentDirectory) []core.Tabler {
	tablers := []core.Tabler{}

	for _, dir := range dirs {
		tablers = append(tablers, dir)
	}

	return tablers
}
