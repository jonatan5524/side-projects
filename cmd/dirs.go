package cmd

import (
	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
	usecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/parentDirectory"
	"github.com/spf13/cobra"
)

var listDirectoriesCmd = &cobra.Command{
	Use:   "dirs",
	Short: "List all the directories assign",
	Long:  "List all the directories assign",
	Run:   ListDirectories,
}

func ListDirectories(cmd *cobra.Command, args []string) {
	db, err := config.InitDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewParentDirectoryObjectBoxRepository(db)
	service := usecase.NewParentDirectoryService(repository)

	isVerbose, err := cmd.Flags().GetBool(VERBOSE_FLAG)

	if err != nil {
		panic(err)
	}

	printListDirectories(service, isVerbose)
}

func printListDirectories(service usecase.ParentDirectoryUsecase, isVerbose bool) {
	directories, err := service.GetAll()

	if err != nil {
		panic(err)
	}

	outputHandler.PrintString("directories:")

	if isVerbose {
		printVerboseListDirectories(directories)
	} else {
		printNormalListDirectories(directories)
	}
}

func printVerboseListDirectories(directories []*model.ParentDirectory) {
	outputHandler.PrintTable(model.ConvertParentDirectoryToTablerSlice(directories))
}

func printNormalListDirectories(directories []*model.ParentDirectory) {
	for _, directory := range directories {
		outputHandler.PrintString(directory.Path)
	}
}

func init() {
	rootCmd.AddCommand(listDirectoriesCmd)
	listDirectoriesCmd.Flags().BoolP(VERBOSE_FLAG, "v", false, "Print list directories verbose with all the data on the directories")
}
