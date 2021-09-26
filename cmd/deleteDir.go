package cmd

import (
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/parentDirectory"
	"github.com/spf13/cobra"
)

var deleteDirCmd = &cobra.Command{
	Use:   "delete-dir",
	Short: "Deleting an assign directory with his path",
	Long:  "Deleting an assign directory with his path",
	Run:   DeleteDirectoryCMD,
}

func DeleteDirectoryCMD(cmd *cobra.Command, args []string) {
	if args[0] == "" {
		panic("path not added")
	}

	db := initDB()
	defer db.Close()

	service := initParentDirectoryUsecase(db)

	deleteDirectory(service, args[0])
}

func deleteDirectory(service usecase.ParentDirectoryUsecase, path string) {
	err := service.DeleteByPath(path)

	if err != nil {
		panic(err)
	}

	outputHandler.PrintString("Directory deleted from db")
}

func init() {
	rootCmd.AddCommand(deleteDirCmd)
}
