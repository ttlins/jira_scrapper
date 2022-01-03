package service

import (
	"fmt"

	"github.com/titolins/jira_scrapper/internal/cache"
	"github.com/titolins/jira_scrapper/internal/handler"
	"github.com/titolins/jira_scrapper/internal/scrapper"
)

type Sprints struct {
	h *handler.Handler
}

func NewSprints(
	s *scrapper.Scrapper,
	projectKey string,
	boardID int,
) *Sprints {
	f := fmt.Sprintf("%s_%d", projectKey, boardID)
	h := handler.New(
		cache.New("sprints", f),
		func() (interface{}, error) {
			return s.GetSprints(projectKey, boardID)
		},
	)
	return &Sprints{h}
}

func (s *Sprints) Call() (err error) {
	var res scrapper.GetSprintsResult

	if err = s.h.Handle(&res); err != nil {
		return
	}

	return
}
