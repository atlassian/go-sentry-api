package sentry

import (
	"fmt"
	"path"
	"time"
)

const (
	MembersEndpointName = "members"
	OrgEndpointName     = "organizations"
	UsersEndpointName   = "users"
)

// Quota is your quote for a project limit and max rate
type Quota struct {
	ProjectLimit int `json:"projectLimit"`
	MaxRate      int `json:"maxRate"`
}

// Organization is your sentry organization.
type Organization struct {
	PendingAccessRequest *int       `json:"pendingAccessRequests,omitempty"`
	Slug                 *string    `json:"slug,omitempty"`
	Name                 string     `json:"name"`
	Quota                *Quota     `json:"quota,omitempty"`
	DateCreated          *time.Time `json:"dateCreated,omitempty"`
	Teams                *[]Team    `json:"teams,omitempty"`
	ID                   *string    `json:"id,omitempty"`
	IsEarlyAdopter       *bool      `json:"isEarlyAdopter,omitempty"`
	Features             *[]string  `json:"features,omitempty"`
}

// GetOrganization takes a org slug and returns back the org
func (c *Client) GetOrganization(orgslug string) (Organization, error) {
	var org Organization

	err := c.do("GET", fmt.Sprintf("%s/%s", OrgEndpointName, orgslug), &org, nil)
	return org, err
}

// GetOrganizationMember returns the member associated with orgslug and memberID
func (c *Client) GetOrganizationMember(orgslug string, memberID string) (Member, error) {
	var member Member
	err := c.do("GET", path.Join(OrgEndpointName, orgslug, MembersEndpointName, memberID), &member, nil)
	return member, err
}

// GetOrganizations will return back every organization in the sentry instance
func (c *Client) GetOrganizations() ([]Organization, *Link, error) {
	orgs := make([]Organization, 0)
	link, err := c.doWithPagination("GET", OrgEndpointName, &orgs, nil)
	return orgs, link, err
}

// ListOrganizationUsers returns users in the organization identified by orgslug
func (c *Client) ListOrganizationUsers(orgslug string) ([]User, error) {
	users := make([]User, 0)
	err := c.do("GET", path.Join(OrgEndpointName, orgslug, UsersEndpointName), &users, nil)
	return users, err
}

// CreateOrganization creates a organization with a name
func (c *Client) CreateOrganization(orgname string) (Organization, error) {
	var org Organization
	orgreq := &Organization{
		Name: orgname,
	}
	err := c.do("POST", OrgEndpointName, &org, orgreq)
	return org, err
}

// UpdateOrganization takes a organization and updates it on the server side
func (c *Client) UpdateOrganization(org Organization) error {
	return c.do("PUT", fmt.Sprintf("%s/%s", OrgEndpointName, *org.Slug), &org, &org)
}

// DeleteOrganization will delete the Org. There is not way to revert this if you do.
func (c *Client) DeleteOrganization(org Organization) error {
	return c.do("DELETE", fmt.Sprintf("%s/%s", OrgEndpointName, *org.Slug), nil, nil)
}

// GetOrganizationTeams will fetch all teams for this org
func (c *Client) GetOrganizationTeams(o Organization) ([]Team, error) {
	teams := make([]Team, 0)
	err := c.do("GET", fmt.Sprintf("organizations/%s/teams", *o.Slug), &teams, nil)
	return teams, err
}
