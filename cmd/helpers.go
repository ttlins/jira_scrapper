package cmd

import (
	"log"

	jira "github.com/andygrunwald/go-jira"
	"github.com/titolins/jira_scrapper/config"
	"github.com/titolins/jira_scrapper/internal/scrapper"
)

func getProjects() []config.Project {
	cfg := config.Config()
	if len(cfg.Jira.Projects) == 0 {
		log.Fatal("No 'projects' defined in config. Please add the projects you'd like to query for")
	}
	return cfg.Jira.Projects
}

func getClient() *jira.Client {
	cfg := config.Config()
	tp := jira.PATAuthTransport{
		Token: cfg.Jira.Token,
	}
	client, err := jira.NewClient(tp.Client(), cfg.Jira.Host)
	if err != nil {
		log.Fatalf("failed to create jira client: %v", err)
	}
	return client
}

func buildScrapper() *scrapper.Scrapper {
	c := getClient()
	return scrapper.New(
		config.Config(),
		c.Board,
		c.Sprint,
	)
}

func runProjects(run func(p config.Project) error) {
	for _, p := range getProjects() {
		if p.BoardID == 0 {
			log.Printf("board_id for project %q not defined. skipping project", p.Key)
			continue
		}
		if err := run(p); err != nil {
			log.Printf("error running command for project %q board %d: %v", p.Key, p.BoardID, err)
		}
	}
}
