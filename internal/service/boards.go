package service

import (
	"github.com/titolins/jira_scrapper/internal/cache"
	"github.com/titolins/jira_scrapper/internal/handler"
	"github.com/titolins/jira_scrapper/internal/scrapper"
)

type Boards struct {
	h *handler.Handler
}

func NewBoards(
	s *scrapper.Scrapper,
	projectKey string,
) *Boards {
	h := handler.New(
		cache.New("boards", projectKey),
		func() (interface{}, error) {
			return s.GetBoards(projectKey)
		},
	)
	return &Boards{h}
}

func (b *Boards) Call() (err error) {
	var res scrapper.GetBoardsResult

	if err = b.h.Handle(&res); err != nil {
		return
	}

	return
}
