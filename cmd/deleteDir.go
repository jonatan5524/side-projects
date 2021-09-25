package cmd

import (
	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
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

	db, err := config.InitDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewParentDirectoryObjectBoxRepository(db)
	service := usecase.NewParentDirectoryService(repository)

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
