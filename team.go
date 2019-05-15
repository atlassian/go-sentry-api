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

// GetTeam takes a team slug and returns back the team
func (c *Client) GetTeam(o Organization, teamSlug string) (Team, error) {
	var team Team
	err := c.do("GET", fmt.Sprintf("teams/%s/%s", *o.Slug, teamSlug), &team, nil)
	return team, err
}

// UpdateTeam will update a team on the server side
func (c *Client) UpdateTeam(o Organization, t Team) error {
	return c.do("PUT", fmt.Sprintf("teams/%s/%s", *o.Slug, *t.Slug), &t, &t)
}

// DeleteTeam deletes a team from a organization
func (c *Client) DeleteTeam(o Organization, t Team) error {
	return c.do(http.MethodDelete, fmt.Sprintf("teams/%s/%s", *o.Slug, *t.Slug), nil, nil)
}

// GetTeamProjects fetchs all projects for a Team
func (c *Client) GetTeamProjects(o Organization, t Team) ([]Project, error) {
	projects := make([]Project, 0)
	err := c.do("GET", fmt.Sprintf("teams/%s/%s/projects", *o.Slug, *t.Slug), &projects, nil)
	return projects, err
}

// GetTeamMembers fetchs all members of a Team
func (c *Client) GetTeamMembers(o Organization, t Team) ([]Member, error) {
	members := make([]Member, 0)
	err := c.do("GET", fmt.Sprintf("teams/%s/%s/members", *o.Slug, *t.Slug), &members, nil)
	return members, err
}
