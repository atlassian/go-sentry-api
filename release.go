package sentry

import (
	"fmt"
	"time"
)

// Release is your release for a orgs teams project
type Release struct {
	DateReleased *time.Time `json:"dateReleased,omitempty"`
	URL          *string    `json:"url,omitempty"`
	Ref          *string    `json:"ref,omitempty"`
	Owner        *string    `json:"owner,omitempty"`
	DateCreated  *time.Time `json:"dateCreated,omitempty"`
	LastEvent    *time.Time `json:"lastEvent,omitempty"`
	Version      string     `json:"version"`
	FirstEvent   *time.Time `json:"firstEvent,omitempty"`
	ShortVersion string     `json:"shortVersion"`
	DateStarted  *time.Time `json:"dateStarted,omitempty"`
	NewGroups    int        `json:"newGroups,omitempty"`
}

// GetRelease will fetch a release from your org and project this does need a version string
func (c *Client) GetRelease(o Organization, p Project, version string) (Release, error) {
	var rel Release
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/releases/%s", *o.Slug, *p.Slug, version), &rel, nil)
	return rel, err
}
