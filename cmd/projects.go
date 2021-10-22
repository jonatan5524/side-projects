package cmd

import (
	"strings"

	"github.com/jonatan5524/side-projects-manager/pkg/model"
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/project"
	"github.com/spf13/cobra"
)

var (
	listProjectsCmd = &cobra.Command{
		Use:   "projects",
		Short: "List all the side projects",
		Long:  "List all the side projects under parent directories",
		Run:   ListProjects,
	}
)

const LIST_PROJECTS_FILTERS = "git"

func ListProjects(cmd *cobra.Command, args []string) {
	db := initDB()
	defer db.Close()

	isVerbose, filter := parseListProjectsFlags(cmd)
	service := initProjectUsecase(db)

	printListProjects(service, isVerbose, filter)
}

func parseListProjectsFlags(cmd *cobra.Command) (bool, string) {
	isVerbose, err := cmd.Flags().GetBool(VERBOSE_FLAG)

	if err != nil {
		panic(err)
	}

	filter, err := cmd.Flags().GetString(FILTER_FLAG)

	if err != nil {
		panic(err)
	}

	if !strings.Contains(LIST_PROJECTS_FILTERS, filter) {
		panic("Invalid filter")
	}

	return isVerbose, filter
}

func printListProjects(service usecase.ProjectUsecase, isVerbose bool, filter string) {
	var projects []*model.Project
	var err error

	if filter != "" {
		projects, err = service.GetAllFiltered(filter)
	} else {
		projects, err = service.GetAll()
	}

	if err != nil {
		panic(err)
	}

	if len(projects) == 0 {
		outputHandler.PrintString("you don't have side projects,\nmaybe you have not assign directories\nto assign directory use add-dir command.")
	} else {
		outputHandler.PrintString("projects:")

		if isVerbose {
			printVerboseListProjects(projects)
		} else {
			printNormalListProjects(projects)
		}
	}
}

func printVerboseListProjects(projects []*model.Project) {
	outputHandler.PrintTable(model.ConvertProjectToTablerSlice(projects))
}

func printNormalListProjects(projects []*model.Project) {
	for _, project := range projects {
		outputHandler.PrintString(project.Name)
	}
}

func init() {
	rootCmd.AddCommand(listProjectsCmd)

	listProjectsCmd.Flags().BoolP(VERBOSE_FLAG, "v", false, "Print list projects verbose with all the data on the projects")
	listProjectsCmd.Flags().StringP(FILTER_FLAG, "f", "", "filter list of projects")
}
