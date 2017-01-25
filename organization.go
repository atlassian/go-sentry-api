package sentry

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (

	// OrgEndpointName is set to roganizations
	OrgEndpointName = "organizations"
	// StatReceived is set to received for sending to /stats/ endpoints
	StatReceived StatQuery = "received"
	// StatRejected is set to rejected for sending to /stats/ endpoints
	StatRejected StatQuery = "rejected"
	// StatBlacklisted is set to blacklisted for sending to /stats/ endpoints
	StatBlacklisted StatQuery = "blacklisted"
)

// StatQuery is semple type for sending to /stats/ endpoints
type StatQuery string

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

type orgStatRequest struct {
	Stat       StatQuery `json:"stat"`
	Since      int64     `json:"since"`
	Until      int64     `json:"until"`
	Resolution string    `json:"resolution,omitempty"`
}

func (o *orgStatRequest) ToQueryString() string {
	query := url.Values{}
	query.Add("stat", string(o.Stat))
	query.Add("since", strconv.FormatInt(o.Since, 10))
	query.Add("until", strconv.FormatInt(o.Until, 10))
	return query.Encode()
}

// OrganizationStat is used for tetting a time in seconds and the metric in a float
type OrganizationStat [2]float64

// GetOrganizationStats fetches stats from the org. Needs a Organization, a StatQuery, a timestamp in seconds since epoch and a optional resolution
func (c *Client) GetOrganizationStats(org Organization, stat StatQuery, since, until int64, resolution *string) ([]OrganizationStat, error) {
	var orgstats []OrganizationStat

	optionalResolution := ""
	if resolution != nil {
		optionalResolution = *resolution
	}
	orgstatrequest := &orgStatRequest{
		Stat:       stat,
		Since:      since,
		Until:      until,
		Resolution: optionalResolution,
	}

	err := c.do("GET", fmt.Sprintf("%s/%s/stats", OrgEndpointName, *org.Slug), &orgstats, orgstatrequest)
	return orgstats, err
}

// GetOrganization takes a org slug and returns back the org
func (c *Client) GetOrganization(orgslug string) (Organization, error) {
	var org Organization

	err := c.do("GET", fmt.Sprintf("%s/%s", OrgEndpointName, orgslug), &org, nil)
	return org, err
}

// GetOrganizations will return back every organization in the sentry instance
func (c *Client) GetOrganizations() ([]Organization, error) {
	orgs := make([]Organization, 0)
	err := c.do("GET", OrgEndpointName, &orgs, nil)
	return orgs, err
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
