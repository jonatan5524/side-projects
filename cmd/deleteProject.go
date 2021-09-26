package cmd

import (
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/project"
	"github.com/spf13/cobra"
)

var deleteProjectCmd = &cobra.Command{
	Use:   "delete-project [path of project]",
	Short: "Deleting an assign project with his path (not deleting from file system)",
	Long:  "Deleting an assign project with his path (not deleting from file system)",
	Run:   DeleteProjectCMD,
}

func DeleteProjectCMD(cmd *cobra.Command, args []string) {
	if args[0] == "" {
		panic("path not provided")
	}

	db := initDB()
	defer db.Close()

	service := initProjectUsecase(db)

	deleteProject(service, args[0])
}

func deleteProject(service usecase.ProjectUsecase, path string) {
	err := service.DeleteByPath(path)

	if err != nil {
		panic(err)
	}

	outputHandler.PrintString("Directory deleted from db")
}

func init() {
	rootCmd.AddCommand(deleteProjectCmd)
}
