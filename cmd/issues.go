package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/titolins/jira_scrapper/config"
	"github.com/titolins/jira_scrapper/internal/cache"
	"github.com/titolins/jira_scrapper/internal/scrapper"
	"github.com/titolins/jira_scrapper/internal/service"
)

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "lists sprints issues",
	Long:  `iterates through all projects defined in the configuration file, tries to fetch the sprints cache for each of them and then fetches the issues for each of the sprints in the cache. Requires the sprints command to be run first`,
	Run: func(cmd *cobra.Command, args []string) {
		s := buildScrapper()
		runProjects(func(p config.Project) error {
			for _, sp := range loadCachedSprints(p).Sprints {
				if err := service.NewIssues(s, sp, p.Key).Call(); err != nil {
					log.Printf("failed to fetch issues for sprint %d: %v", sp.ID, err)
				}
			}
			return nil
		})
	},
}

func loadCachedSprints(p config.Project) (d scrapper.GetSprintsResult) {
	f := fmt.Sprintf("%s_%d", p.Key, p.BoardID)
	c := cache.New("sprints", f)
	if _, exists := c.Exists(); !exists {
		log.Printf("cache for project %q board %d doesn't exist\n", p.Key, p.BoardID)
	}
	if err := c.Load(&d); err != nil {
		log.Printf("failed to load cache for project %q board %d: %v", p.Key, p.BoardID, err)
	}

	return
}
