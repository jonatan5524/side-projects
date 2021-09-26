package cmd

import (
	"github.com/jonatan5524/side-projects-manager/pkg/model"
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
	db := initDB()
	defer db.Close()

	isVerbose := parseListDirectoriesFlags(cmd)
	service := initParentDirectoryUsecase(db)

	printListDirectories(service, isVerbose)
}

func parseListDirectoriesFlags(cmd *cobra.Command) bool {
	isVerbose, err := cmd.Flags().GetBool(VERBOSE_FLAG)

	if err != nil {
		panic(err)
	}

	return isVerbose
}

func printListDirectories(service usecase.ParentDirectoryUsecase, isVerbose bool) {
	directories, err := service.GetAll()

	if err != nil {
		panic(err)
	}

	if len(directories) == 0 {
		outputHandler.PrintString("you don't have directories assign\nto assign directory use add-dir command")
	} else {
		outputHandler.PrintString("directories:")

		if isVerbose {
			printVerboseListDirectories(directories)
		} else {
			printNormalListDirectories(directories)
		}
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
