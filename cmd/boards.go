package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/titolins/jira_scrapper/internal/service"
)

var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "lists the boards by project key",
	Long:  `iterates all projects in the configuration and returns all boards for each of them. Can be used to help in the configuration (finding board_id for your projects)`,
	Run: func(cmd *cobra.Command, args []string) {
		s := buildScrapper()
		for _, p := range getProjects() {
			if err := service.NewBoards(s, p.Key).Call(); err != nil {
				log.Printf("failed to fetch boards for project %q: %v", p.Key, err)
			}
		}
	},
}
