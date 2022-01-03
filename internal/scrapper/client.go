package scrapper

import (
	"github.com/andygrunwald/go-jira"
	"github.com/titolins/jira_scrapper/config"
)

type (
	BoardService interface {
		GetAllBoards(*jira.BoardListOptions) (*jira.BoardsList, *jira.Response, error)
		GetAllSprintsWithOptions(boardID int, opts *jira.GetAllSprintsOptions) (*jira.SprintsList, *jira.Response, error)
	}
	SprintService interface {
		GetIssuesForSprint(sprintID int, options *jira.GetQueryOptions) ([]jira.Issue, *jira.Response, error)
	}
)

type Scrapper struct {
	cfg config.Configuration
	BoardService
	SprintService
}

func New(cfg config.Configuration, bs BoardService, ss SprintService) *Scrapper {
	return &Scrapper{cfg, bs, ss}
}
