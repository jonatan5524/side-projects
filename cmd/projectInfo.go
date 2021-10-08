package cmd

import (
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/project"
	"github.com/spf13/cobra"
)

var projectInfoCmd = &cobra.Command{
	Use:   "project-info [path]",
	Short: "Represent a full info of registered db project",
	Long:  "Represent a full info of registered db project",
	Run:   InfoProjectCMD,
}

func InfoProjectCMD(cmd *cobra.Command, args []string) {
	if args[0] == "" {
		panic("path not provided")
	}

	db := initDB()
	defer db.Close()

	service := initProjectUsecase(db)

	projectInfo(service, args[0])
}

func projectInfo(service usecase.ProjectUsecase, path string) {
	project, err := service.Get(path)

	if err != nil {
		panic(err)
	}

	outputHandler.PrintString(project.FullInfo())
}

func init() {
	rootCmd.AddCommand(projectInfoCmd)
}
