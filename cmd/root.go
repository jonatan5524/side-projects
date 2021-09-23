package cmd

import (
	"fmt"
	"os"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/io"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var outputHandler = core.NewOutputStdout()

var rootCmd = &cobra.Command{
	Use:   "side-projects",
	Short: "side projects manager",
	Long: `Side Projects Manager manage all your side projects in one place,

you can use it to list all your projects from different directories
get details on each project in one place on terminal or on graphical user interface.

Use "side-projects <command> --help" for more information about a given command`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.side-projects.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".side-projects")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
