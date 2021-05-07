package sentry

import (
	"fmt"
	"time"
)

// Release is your release for a orgs teams project
type Release struct {
	DateCreated  *time.Time              `json:"dateCreated,omitempty"`
	DateReleased *time.Time              `json:"dateReleased,omitempty"`
	FirstEvent   *time.Time              `json:"firstEvent,omitempty"`
	LastEvent    *time.Time              `json:"lastEvent,omitempty"`
	NewGroups    *int                    `json:"newGroups,omitempty"`
	Projects     []Project               `json:"projects,omitempty"`
	Ref          *string                 `json:"ref,omitempty"`
	ShortVersion string                  `json:"shortVersion"`
	URL          *string                 `json:"url,omitempty"`
	Version      string                  `json:"version"`
	Authors      *[]Author               `json:"authors,omitempty"`
	CommitCount  *int                    `json:"commitCount,omitempty"`
	LastCommit   *Commit                 `json:"lastCommit,omitempty"`
	Data         *map[string]interface{} `json:"data,omitempty"`
	DeployCount  *int                    `json:"deployCount,omitempty"`
	LastDeploy   *Deploy                 `json:"lastDeploy,omitempty"`
}

// NewRelease is used to create a new release
type NewRelease struct {
	// Optional commit ref.
	Ref *string `json:"ref,omitempty"`

	// Optional URL to point to the online source code
	URL *string `json:"url,omitempty"`

	// Required for creating the release
	Version string `json:"version"`

	// Optional to set when it started
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// Optional to set when it was released to the public
	DateReleased *time.Time `json:"dateReleased,omitempty"`
}

type Deploy struct {
	ID           string     `json:"id"`
	Name         *string    `json:"name,omitempty"`
	URL          *string    `json:"url,omitempty"`
	DateStarted  *time.Time `json:"dateStarted,omitempty"`
	DateFinished *time.Time `json:"dateFinished,omitempty"`
	Environment  string     `json:"environment"`
}

type NewDeploy struct {
	// Required for creating the deploy
	Environment string `json:"environment"`

	// Optional parameters for creating the deploy
	Name         *string    `json:"name,omitempty"`
	URL          *string    `json:"url,omitempty"`
	DateStarted  *time.Time `json:"dateStarted,omitempty"`
	DateFinished *time.Time `json:"dateReleased,omitempty"`
}

// GetRelease will fetch a release from your org and project this does need a version string
func (c *Client) GetRelease(o Organization, p Project, version string) (Release, error) {
	var rel Release
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/releases/%s", *o.Slug, *p.Slug, version), &rel, nil)
	return rel, err
}

// GetReleases will fetch all releases from your org and project
func (c *Client) GetReleases(o Organization, p Project) ([]Release, *Link, error) {
	var rel []Release
	link, err := c.doWithPagination("GET", fmt.Sprintf("projects/%s/%s/releases", *o.Slug, *p.Slug), &rel, nil)
	return rel, link, err
}

//CreateRelease will create a new release for a project in a org
func (c *Client) CreateRelease(o Organization, p Project, r NewRelease) (Release, error) {
	var rel Release
	err := c.do("POST", fmt.Sprintf("projects/%s/%s/releases", *o.Slug, *p.Slug), &rel, &r)
	return rel, err
}

//UpdateRelease will update ref, url, started, released for a release.
//Version should not change.
func (c *Client) UpdateRelease(o Organization, p Project, r Release) error {
	return c.do("PUT", fmt.Sprintf("projects/%s/%s/releases/%s", *o.Slug, *p.Slug, r.Version), &r, &r)
}

//DeleteRelease will delete the release from your project
func (c *Client) DeleteRelease(o Organization, p Project, r Release) error {
	return c.do("DELETE", fmt.Sprintf("projects/%s/%s/releases/%s", *o.Slug, *p.Slug, r.Version), nil, nil)
}

// CreateDeploy will create a deploy in your org for a given version
func (c *Client) CreateDeploy(o Organization, r Release, d NewDeploy) (Deploy, error) {
	var dep Deploy
	err := c.do("POST", fmt.Sprintf("organizations/%s/releases/%s/deploys", *o.Slug, r.Version), &dep, &d)
	return dep, err
}

// ListDeploys will list the deploys for a given release.
func (c *Client) ListDeploys(o Organization, r Release) ([]Deploy, *Link, error) {
	var dep []Deploy
	link, err := c.doWithPagination("GET", fmt.Sprintf("organizations/%s/releases/%s/deploys", *o.Slug, r.Version), &dep, nil)
	return dep, link, err
}
