package scrapper

import (
	"bytes"
	"fmt"

	"github.com/andygrunwald/go-jira"
)

type GetIssuesResult struct {
	ProjectKey string       `json:"project_key"`
	BoardID    int          `json:"board_id"`
	SprintID   int          `json:"sprint_id"`
	Issues     []jira.Issue `json:"issues,omitempty"`
}

func (r *GetIssuesResult) String() string {
	if len(r.Issues) == 0 {
		return fmt.Sprintf("No issues for board %d, sprint %d\n", r.BoardID, r.SprintID)
	}

	var buf bytes.Buffer
	buf.WriteString(
		fmt.Sprintf(
			"Board: %d\nSprint:%d\nIssues:\n",
			r.BoardID,
			r.SprintID,
		),
	)
	for ix, i := range r.Issues {
		buf.WriteString(printIssue(ix, i))
	}

	return buf.String()
}

func (s *Scrapper) GetIssues(
	sp jira.Sprint,
	projectKey string,
) (*GetIssuesResult, error) {
	is, _, err := s.GetIssuesForSprint(sp.ID, getQueryOpts())
	if err != nil {
		return nil, err
	}
	return &GetIssuesResult{
		ProjectKey: projectKey,
		BoardID:    sp.OriginBoardID,
		SprintID:   sp.ID,
		Issues:     is,
	}, nil
}

func getQueryOpts() *jira.GetQueryOptions {
	return &jira.GetQueryOptions{
		Expand: "changelog",
	}
}

func printIssue(ix int, i jira.Issue) string {
	return fmt.Sprintf(
		"\t%d:\n\t\tID: %s\n\t\tKey: %s\n\t\tChangelog: %v\n\t\tTransitions:%v\n",
		ix,
		i.ID,
		i.Key,
		i.Changelog,
		i.Transitions,
	)
}
