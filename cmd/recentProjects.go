package cmd

import (
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/project"
	"github.com/spf13/cobra"
)

var recentProjectsCmd = &cobra.Command{
	Use:   "recent",
	Short: "List recent side projects edited",
	Long:  "List recent side projects edited",
	Run:   ListProjectsRecent,
}

func ListProjectsRecent(cmd *cobra.Command, args []string) {
	db := initDB()
	defer db.Close()

	service := initProjectUsecase(db)

	printListProjectsRecent(service)
}

func printListProjectsRecent(service usecase.ProjectUsecase) {
	projects, err := service.GetRecent()

	if err != nil {
		panic(err)
	}

	if len(projects) == 0 {
		outputHandler.PrintString("you don't have side projects,\nmaybe you have not assign directories\nto assign directory use add-dir command.")
	} else {
		outputHandler.PrintString("projects:")

		printVerboseListProjects(projects)
	}
}

func init() {
	rootCmd.AddCommand(recentProjectsCmd)
}
