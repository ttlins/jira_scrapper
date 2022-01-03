package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/titolins/jira_scrapper/config"
)

const appName = "jira_scrapper"

var (
	rootCmd = &cobra.Command{
		Use:   "jira_scrapper",
		Short: "",
		Long:  ``,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//initConfig()
	config.Init()

	rootCmd.PersistentFlags().BoolP("print", "p", true, "prints the result")
	viper.BindPFlag("print", rootCmd.PersistentFlags().Lookup("print"))

	rootCmd.PersistentFlags().BoolP("skip-cached", "s", true, "auto skips already cached entities")
	viper.BindPFlag("skip-cached", rootCmd.PersistentFlags().Lookup("skip-cached"))

	rootCmd.AddCommand(boardsCmd)
	rootCmd.AddCommand(sprintsCmd)
	rootCmd.AddCommand(issuesCmd)
}
