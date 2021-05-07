package sentry

import (
	"fmt"
	"time"
)

type Repository struct {
	DateCreated		*time.Time `json:"dateCreated,omitempty"`
	ID              *string    `json:"id,omitempty"`
	Name            string     `json:"name"`
}

type Commit struct {
	DateCreated		*time.Time `json:"dateCreated,omitempty"`
	ID              *string    `json:"id,omitempty"`
	Message         *string    `json:"message,omitempty"`
}

// GetRepositories returns a list of version control repositories for a given organization.
func (c *Client) GetRepositories(o Organization) ([]Repository, *Link, error) {
	var repos []Repository
	link, err := c.doWithPagination("GET", fmt.Sprintf("organizations/%s/repos", *o.Slug), &repos, nil)
	return repos, link, err
}

// GetRepositoryCommits returns a list of commits for a given repository.
func (c *Client) GetRepositoryCommits(o Organization, r Repository) ([]Commit, *Link, error) {
	var commits []Commit
	link, err := c.doWithPagination("GET", fmt.Sprintf("organizations/%s/repos/%s/commits/", *o.Slug, *r.ID), &commits, nil)
	return commits, link, err
}