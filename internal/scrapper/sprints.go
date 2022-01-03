package scrapper

import (
	"bytes"
	"fmt"

	"github.com/andygrunwald/go-jira"
)

const (
	maxResults        = 50
	closedSprintState = "closed"
)

type GetSprintsResult struct {
	ProjectKey string        `json:"project_key"`
	BoardID    int           `json:"board_id"`
	Sprints    []jira.Sprint `json:"sprints,omitempty"`
}

func (r *GetSprintsResult) String() string {
	if len(r.Sprints) == 0 {
		return fmt.Sprintf("No sprints for project %s, board %d\n", r.ProjectKey, r.BoardID)
	}

	var buf bytes.Buffer
	buf.WriteString(
		fmt.Sprintf(
			"Project: %s\n, Board: %d\nSprints:\n",
			r.ProjectKey,
			r.BoardID,
		),
	)
	for i, s := range r.Sprints {
		buf.WriteString(printSprint(i, s))
	}

	return buf.String()
}

func (s *Scrapper) GetSprints(
	projectKey string,
	boardID int,
) (*GetSprintsResult, error) {
	var (
		startAt int
		ss      []jira.Sprint
	)

	for {
		res, _, err := s.GetAllSprintsWithOptions(boardID, getSprintsOpts(startAt))
		if err != nil {
			return nil, err
		} else if len(res.Values) == 0 {
			break
		}

		startAt += maxResults
		ss = append(ss, res.Values...)
	}

	return &GetSprintsResult{
		ProjectKey: projectKey,
		BoardID:    boardID,
		Sprints:    ss,
	}, nil
}

func getSprintsOpts(startAt int) *jira.GetAllSprintsOptions {
	return &jira.GetAllSprintsOptions{
		State: closedSprintState,
		SearchOptions: jira.SearchOptions{
			StartAt:    startAt,
			MaxResults: maxResults,
		},
	}
}

func printSprint(i int, s jira.Sprint) string {
	return fmt.Sprintf(
		"\t%d:\n\t\tID: %d\n\t\tName: %s\n\t\tState: %s\n",
		i,
		s.ID,
		s.Name,
		s.State,
	)
}
