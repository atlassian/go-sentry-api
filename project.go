package sentry

import (
	"fmt"
	"time"
)

// Project is your project in sentry
type Project struct {
	Status             string     `json:"status,omitempty"`
	Slug               *string    `json:"slug,omitempty"`
	DefaultEnvironment *string    `json:"defaultEnvironment,omitempty"`
	Color              string     `json:"color,omitempty"`
	IsPublic           bool       `json:"isPublic,omitempty"`
	DateCreated        time.Time  `json:"dateCreated,omitempty"`
	CallSign           string     `json:"callSign,omitempty"`
	FirstEvent         *time.Time `json:"firstEvent,omitempty"`
	IsBookmarked       bool       `json:"isBookmarked,omitempty"`
	CallSignReviewed   bool       `json:"callSignReviewed,omitempty"`
	ID                 string     `json:"id,omitempty"`
	Name               string     `json:"name"`
	Platforms          *[]string  `json:"platforms,omitempty"`
}

// CreateProject will create a new project in your org and team
func (c *Client) CreateProject(o Organization, t Team, name string, slug *string) (Project, error) {
	var proj Project
	projreq := &Project{
		Name: name,
		Slug: slug,
	}

	err := c.do("POST", fmt.Sprintf("teams/%s/%s/projects", *o.Slug, *t.Slug), &proj, projreq)
	return proj, err
}

// DeleteProject will take your org, team, and proj and delete it from sentry.
func (c *Client) DeleteProject(o Organization, p Project) error {
	return c.do("DELETE", fmt.Sprintf("projects/%s/%s", *o.Slug, *p.Slug), nil, nil)
}
