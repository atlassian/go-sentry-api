package sentry

import (
	"fmt"
	"net/http"
	"time"
)

// Team is a sentry team
type Team struct {
	Slug        *string    `json:"slug,omitempty"`
	Name        string     `json:"name"`
	HasAccess   *bool      `json:"hasAccess,omitempty"`
	IsPending   *bool      `json:"isPending,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	IsMember    *bool      `json:"isMember,omitempty"`
	ID          *string    `json:"id,omitempty"`
	Projects    *[]Project `json:"projects,omitempty"`
}

// CreateTeam creates a team with a organization object and requires a name and a optional slug
func (c *Client) CreateTeam(o Organization, name string, slug *string) (Team, error) {
	var team Team
	teamreq := &Team{
		Name: name,
		Slug: slug,
	}

	err := c.do(http.MethodPost, fmt.Sprintf("organizations/%s/teams", *o.Slug), &team, teamreq)
	return team, err
}

// DeleteTeam deletes a team from a organization
func (c *Client) DeleteTeam(o Organization, t Team) error {
	return c.do(http.MethodDelete, fmt.Sprintf("teams/%s/%s", *o.Slug, *t.Slug), nil, nil)
}
