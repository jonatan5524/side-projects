package model

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	core "github.com/jonatan5524/side-projects-manager/pkg/core/io"
	gitUtil "github.com/jonatan5524/side-projects-manager/pkg/util/git"
	ioUtil "github.com/jonatan5524/side-projects-manager/pkg/util/io"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen clean
//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Project struct {
	Id                 uint64
	Name               string
	Path               string    `objectbox"unique"`
	LastUpdated        time.Time `objectbox:"date"`
	HaveVersionControl bool
}

var NilProject = Project{}

func NewProject(path string, directoryGetter ioUtil.DirectoryGetter) (Project, error) {
	projectInfo, err := directoryGetter(path)

	if err != nil {
		return NilProject, err
	}

	haveVersionControl, err := gitUtil.IsInVersionControl(path)

	if err != nil {
		return NilProject, err
	}

	return Project{
		Name:               projectInfo.Name(),
		Path:               path,
		LastUpdated:        projectInfo.ModTime(),
		HaveVersionControl: haveVersionControl,
	}, nil
}

func (project Project) String() string {
	return fmt.Sprintf(`
	Project {
		Id: %d
		Name: %s
		Path: %s
		LastUpdated: %v
		HaveVersionControl: %t
	}`,
		project.Id,
		project.Name,
		project.Path,
		project.LastUpdated,
		project.HaveVersionControl,
	)
}

func (Project) TableHeader() table.Row {
	return table.Row{
		"Name",
		"Path",
		"Last Updated",
		"Have Version Control",
	}
}

func (project Project) TableData() table.Row {
	return table.Row{
		project.Name,
		project.Path,
		project.LastUpdated,
		project.HaveVersionControl,
	}
}

func ConvertProjectToTablerSlice(projects []*Project) []core.Tabler {
	tablers := []core.Tabler{}

	for _, project := range projects {
		tablers = append(tablers, project)
	}

	return tablers
}
