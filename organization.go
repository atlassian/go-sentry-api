package sentry

import (
	"fmt"
	"time"
)

const (
	OrgEndpointName = "organizations"
)

type Quota struct {
	ProjectLimit int `json:"projectLimit"`
	MaxRate      int `json:"maxRate"`
}

type Project struct {
	Status             string     `json:"status"`
	Slug               string     `json:"slug"`
	DefaultEnvironment *string    `json:"defaultEnvironment"`
	Color              string     `json:"color"`
	IsPublic           bool       `json:"isPublic"`
	DateCreated        time.Time  `json:"dateCreated"`
	CallSign           string     `json:"callSign"`
	FirstEvent         *time.Time `json:"firstEvent"`
	IsBookmarked       bool       `json:"isBookmarked"`
	CallSignReviewed   bool       `json:"callSignReviewed"`
	Id                 string     `json:"id"`
	Name               string     `json:"name"`
	Platforms          *[]string  `json:"platforms"`
}

type Team struct {
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	HasAccess   bool      `json:"hasAccess"`
	IsPending   bool      `json:"isPending"`
	DateCreated time.Time `json:"dateCreated"`
	IsMember    bool      `json:"isMember"`
	Id          string    `json:"id"`
	Projects    []Project `json:"projects,omitempty"`
}

type Organization struct {
	PendingAccessRequest *int       `json:"pendingAccessRequests,omitempty"`
	Slug                 *string    `json:"slug,omitempty"`
	Name                 string     `json:"name"`
	Quota                *Quota     `json:"quota,omitempty"`
	DateCreated          *time.Time `json:"dateCreated,omitempty"`
	Teams                *[]Team    `json:"teams,omitempty"`
	Id                   *string    `json:"id,omitempty"`
	IsEarlyAdopter       *bool      `json:"isEarlyAdopter,omitempty"`
	Features             *[]string  `json:"features,omitempty"`
}

func (c *Client) GetOrganization(orgslug string) (Organization, error) {
	var org Organization

	if err := c.do("GET", fmt.Sprintf("%s/%s", OrgEndpointName, orgslug), &org, nil); err != nil {
		return org, err
	}
	return org, nil
}

func (c *Client) GetOrganizations() ([]Organization, error) {
	orgs := make([]Organization, 0)
	if err := c.do("GET", OrgEndpointName, &orgs, nil); err != nil {
		return orgs, err
	}
	return orgs, nil
}

func (c *Client) CreateOrganization(orgname string) (Organization, error) {
	var org Organization
	orgreq := &Organization{
		Name: orgname,
	}
	if err := c.do("POST", OrgEndpointName, &org, orgreq); err != nil {
		return org, err
	}
	return org, nil
}

func (c *Client) UpdateOrganization(org Organization) error {
	if err := c.do("PUT", OrgEndpointName+*org.Slug, &org, &org); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteOrganization(org Organization) error {
	if err := c.do("DELETE", OrgEndpointName+*org.Slug, nil, nil); err != nil {
		return err
	}
	return nil
}
