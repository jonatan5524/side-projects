package cmd

import (
	"fmt"

	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	"github.com/jonatan5524/side-projects-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var addDirCmd = &cobra.Command{
	Use:   "add-dir",
	Short: "Adding directory of side projects",
	Long: `Adding directory to list of directories that contains side projects`,
	Run: AddDir,
}

// TODO: making service and testing
func AddDir(cmd *cobra.Command, args []string) {
	db, err := config.InitDB()
	fmt.Println("Initalizing db")

	if err != nil {
		panic(fmt.Sprintf("error while initalizing db: %v", err))
	}

	repo := repository.NewParentDirectoryRepositoryObjectBox(db)
	fmt.Printf("saving new parent directory: %s\n", args[0])
	parentDir, err := model.NewParentDirectory(args[0])

	if err != nil {
		panic(fmt.Sprintf("error while adding new dir: %v\n", err))
	}

	if _, err := repo.Put(parentDir); err != nil {
		panic(fmt.Sprintf("error while adding new dir: %v", err))
	} 

	fmt.Println("new directory saved!")
}

func init() {
	rootCmd.AddCommand(addDirCmd)

}
