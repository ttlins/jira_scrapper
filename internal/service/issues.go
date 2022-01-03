package service

import (
	"strconv"

	"github.com/andygrunwald/go-jira"
	"github.com/titolins/jira_scrapper/internal/cache"
	"github.com/titolins/jira_scrapper/internal/handler"
	"github.com/titolins/jira_scrapper/internal/scrapper"
)

type Issues struct {
	h *handler.Handler
}

func NewIssues(
	s *scrapper.Scrapper,
	sp jira.Sprint,
	projectKey string,
) *Issues {
	h := handler.New(
		cache.New("issues", strconv.Itoa(sp.ID)),
		func() (interface{}, error) {
			return s.GetIssues(sp, projectKey)
		},
	)
	return &Issues{h}
}

func (i *Issues) Call() (err error) {
	var res scrapper.GetIssuesResult

	if err = i.h.Handle(&res); err != nil {
		return
	}

	return
}
