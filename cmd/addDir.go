package cmd

import (
	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/parentDirectory"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
	"github.com/spf13/cobra"
)

var (
	addDirCmd = &cobra.Command{
		Use:   "add-dir",
		Short: "Adding directory of side projects",
		Long:  `Adding directory to list of directories that contains side projects`,
		Run:   addDir,
	}
)

func addDir(cmd *cobra.Command, args []string) {
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

	addParentDirectoryToDB(service, args[0], model.NewParentDirectory)
}

func addParentDirectoryToDB(service usecase.ParentDirectoryUsecase, path string, parentDirectoryConstructor model.ParentDirectoryConstructor) {
	parentDirectory, err := parentDirectoryConstructor(path, util.GetDirectory)

	if err != nil {
		panic(err)
	}

	err = parentDirectory.LoadProjects()

	if err != nil {
		panic(err)
	}

	_, err = service.Put(parentDirectory)

	if err != nil {
		panic(err)
	}

	outputHandler.PrintString("Directory added!")
}

func init() {
	rootCmd.AddCommand(addDirCmd)
}
