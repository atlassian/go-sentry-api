package sentry

import (
	"fmt"
	"net/url"
)

// IssueBulkRequest is what should be used when bulk mutate issue
type IssueBulkRequest struct {
	Status         *Status `json:"status,omitempty"`
	IgnoreDuration *int    `json:"ignoreDuration,omitempty"`
	IsPublic       *bool   `json:"isPublic,omitempty"`
	Merge          *bool   `json:"merge,omitempty"`
	HasSeen        *bool   `json:"hasSeen,omitempty"`
	IsBookmarked   *bool   `json:"isBookmarked,omitempty"`
}

// IssueBulkResponse is what is returned when your mutation is done
type IssueBulkResponse struct {
	Status        *Status            `json:"status,omitempty"`
	IsPublic      *bool              `json:"isPublic,omitempty"`
	StatusDetails *map[string]string `json:"statusDetails,omitempty"`
}

// issueMutateArgs
type issueMutateArgs struct {
	ID     *[]string
	Status *Status
}

func (i *issueMutateArgs) ToQueryString() string {
	query := url.Values{}
	if i.ID != nil {
		for _, id := range *i.ID {
			query.Add("id", id)
		}
	}
	if i.Status != nil {
		query.Add("status", string(*i.Status))
	}
	return query.Encode()
}

// BulkMutateIssues takes a list of ids and optional status to filter through
func (c *Client) BulkMutateIssues(o Organization, p Project, req IssueBulkRequest, issues *[]string, status *Status) (IssueBulkResponse, error) {
	var issueBulkResponse IssueBulkResponse

	mutatequery := &issueMutateArgs{
		ID:     issues,
		Status: status,
	}

	err := c.doWithQuery("PUT", fmt.Sprintf("projects/%s/%s/issues", *o.Slug, *p.Slug),
		&issueBulkResponse, req, mutatequery)

	return issueBulkResponse, err
}

// BulkDeleteIssues takes a list of IDs and will delete them
func (c *Client) BulkDeleteIssues(o Organization, p Project, issues []string) error {
	mutateQuery := &issueMutateArgs{
		ID: &issues,
	}
	return c.doWithQuery("DELETE", fmt.Sprintf("projects/%s/%s/issues", *o.Slug, *p.Slug),
		nil, nil, mutateQuery)
}
