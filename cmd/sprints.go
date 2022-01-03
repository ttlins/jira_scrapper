package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/titolins/jira_scrapper/internal/service"
)

var sprintsCmd = &cobra.Command{
	Use:   "sprints",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := buildScrapper()
		for _, p := range getProjects() {
			ss := service.NewSprints(s, p.Key, p.BoardID)
			if err := ss.Call(); err != nil {
				log.Fatal(err)
			}
		}
	},
}
