package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/titolins/jira_scrapper/internal/service"
)

var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := buildScrapper()
		for _, p := range getProjects() {
			if err := service.NewBoards(s, p.Key).Call(); err != nil {
				log.Printf("failed to fetch boards for project %q: %v", p.Key, err)
			}
		}
	},
}
