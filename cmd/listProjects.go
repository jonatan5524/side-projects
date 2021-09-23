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
		Use:   "list-projects",
		Short: "list all the side projects",
		Long:  "list all the side projects under parent directories",
		Run:   ListProjects,
	}
)

const (
	VERBOSE_FLAG = "verbose"
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

	outputHandler.PrintString("projects:")

	if isVerbose {
		printVerboseList(projects)
	} else {
		printNormalList(projects)
	}
}

func printVerboseList(projects []*model.Project) {
	outputHandler.PrintTable(model.ConvertProjectToTablerSlice(projects))
}

func printNormalList(projects []*model.Project) {
	for _, project := range projects {
		outputHandler.PrintString(project.Name)
	}
}

func init() {
	rootCmd.AddCommand(listProjectsCmd)
	listProjectsCmd.Flags().BoolP(VERBOSE_FLAG, "v", false, "Print list projects verbose with all the data on the projects")
}
