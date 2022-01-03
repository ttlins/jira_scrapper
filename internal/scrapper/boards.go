package scrapper

import (
	"bytes"
	"fmt"

	"github.com/andygrunwald/go-jira"
)

type GetBoardsResult struct {
	ProjectKey string
	Boards     []jira.Board
}

func (r *GetBoardsResult) String() string {
	if len(r.Boards) == 0 {
		return fmt.Sprintf("No boards for project %s\n", r.ProjectKey)
	}

	var buf bytes.Buffer
	buf.WriteString(
		fmt.Sprintf(
			"Project: %s\nBoards:\n",
			r.ProjectKey,
		),
	)
	for i, b := range r.Boards {
		buf.WriteString(printBoard(i, b))
	}

	return buf.String()
}

func (s *Scrapper) GetBoards(projectKey string) (*GetBoardsResult, error) {
	boards, _, err := s.GetAllBoards(&jira.BoardListOptions{
		ProjectKeyOrID: projectKey,
	})
	if err != nil {
		return nil, err
	}

	return &GetBoardsResult{
		ProjectKey: projectKey,
		Boards:     boards.Values,
	}, nil
}

func printBoard(i int, b jira.Board) string {
	return fmt.Sprintf(
		"\t%d:\n\t\tID: %d\n\t\tName: %s\n",
		i,
		b.ID,
		b.Name,
	)
}
