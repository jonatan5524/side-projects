package cmd

import (
	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
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

func ListProjects(cmd *cobra.Command, args []string) {
	db, err := config.InitDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProjectObjectBoxRepository(db)
	service := usecase.NewProjectService(repository)

	isVerbose, err := cmd.Flags().GetBool(VERBOSE_FLAG)

	if err != nil {
		panic(err)
	}

	printListProjects(service, isVerbose)
}

func printListProjects(service usecase.ProjectUsecase, isVerbose bool) {
	projects, err := service.GetAll()

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
}
