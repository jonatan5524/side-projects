package cmd

import (
	"fmt"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/parentDirectory"
	"github.com/jonatan5524/side-projects-manager/pkg/util"
	"github.com/spf13/cobra"
)

var addDirCmd = &cobra.Command{
	Use:   "add-dir",
	Short: "Adding directory of side projects",
	Long:  `Adding directory to list of directories that contains side projects`,
	Run:   addDir,
}

type ParentDirectoryConstructor func(string, model.DirectoryGetter) (model.ParentDirectory, error)

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

func addParentDirectoryToDB(service usecase.ParentDirectoryUsecase, path string, parentDirectoryConstructor ParentDirectoryConstructor) {
	parentDirectory, err := parentDirectoryConstructor(path, util.GetDirectory)

	if err != nil {
		panic(err)
	}

	_, err = service.Put(parentDirectory)

	if err != nil {
		panic(err)
	}

	fmt.Println("Directory added!")
}

func init() {
	rootCmd.AddCommand(addDirCmd)
}
