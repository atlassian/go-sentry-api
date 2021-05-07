package sentry

import (
	"fmt"
	"time"
)

type Author struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"name,omitempty"`
}

type Repository struct {
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name"`
	Status      *string    `json:"status,omitempty"`
	Provider    *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"provider,omitempty"`
	URL           *string `json:"url,omitempty"`
	IntegrationId *string `json:"integrationId,omitempty"`
	ExternalSlug  *string `json:"externalSlug,omitempty"`
}

type Commit struct {
	PatchSet *[]struct {
		Path string `json:"path"`
		Type string `json:"type"`
	} `json:"patch_set,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Author     *Author     `json:"author,omitempty"`
	Timestamp  *time.Time  `json:"timestamp,omitempty"`
	Message    *string     `json:"message,omitempty"`
	ID         string      `json:"id"`
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
	link, err := c.doWithPagination("GET", fmt.Sprintf("organizations/%s/repos/%s/commits/", *o.Slug, r.ID), &commits, nil)
	return commits, link, err
}
