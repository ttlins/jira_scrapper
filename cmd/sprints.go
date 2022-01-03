package cmd

import (
	"github.com/spf13/cobra"
	"github.com/titolins/jira_scrapper/config"
	"github.com/titolins/jira_scrapper/internal/service"
)

var sprintsCmd = &cobra.Command{
	Use:   "sprints",
	Short: "lists projects sprints",
	Long:  `iterates through all projects defined in the configuration file, and fetches the closed sprints for each of them.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := buildScrapper()
		runProjects(func(p config.Project) error {
			return service.NewSprints(s, p.Key, p.BoardID).Call()
		})
	},
}
