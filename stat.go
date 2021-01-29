package sentry

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	// StatReceived is set to received for sending to /stats/ endpoints
	StatReceived StatQuery = "received"
	// StatRejected is set to rejected for sending to /stats/ endpoints
	StatRejected StatQuery = "rejected"
	// StatBlacklisted is set to blacklisted for sending to /stats/ endpoints
	StatBlacklisted StatQuery = "blacklisted"
)

type statRequest struct {
	Stat       StatQuery `json:"stat"`
	Since      int64     `json:"since"`
	Until      int64     `json:"until"`
	Resolution *string   `json:"resolution,omitempty"`
}

func (o *statRequest) ToQueryString() string {
	query := url.Values{}
	query.Add("stat", string(o.Stat))
	query.Add("since", strconv.FormatInt(o.Since, 10))
	query.Add("until", strconv.FormatInt(o.Until, 10))

	if o.Resolution != nil {
		query.Add("resolution", string(*o.Resolution))
	}

	return query.Encode()
}

// Stat is used for tetting a time in seconds and the metric in a float
type Stat [2]float64

// StatQuery is sample type for sending to /stats/ endpoints
type StatQuery string

// GetOrganizationStats fetches stats from the org. Needs a Organization, a StatQuery, a timestamp in seconds since epoch and a optional resolution
func (c *Client) GetOrganizationStats(org Organization, stat StatQuery, since, until int64, resolution *string) ([]Stat, error) {
	var orgstats []Stat

	orgstatrequest := &statRequest{
		Stat:       stat,
		Since:      since,
		Until:      until,
		Resolution: resolution,
	}

	err := c.doWithQuery("GET", fmt.Sprintf("%s/%s/stats", OrgEndpointName, *org.Slug), &orgstats, nil, orgstatrequest)
	return orgstats, err
}

// GetTeamStats will fetch all stats for a specific team. Similar to GetOrganizationStats
func (c *Client) GetTeamStats(o Organization, t Team, stat StatQuery, since, until int64, resolution *string) ([]Stat, error) {
	var teamstats []Stat
	teamstatrequest := &statRequest{
		Stat:       stat,
		Since:      since,
		Until:      until,
		Resolution: resolution,
	}

	err := c.doWithQuery("GET", fmt.Sprintf("teams/%s/%s/stats", *o.Slug, *t.Slug), &teamstats, nil, teamstatrequest)
	return teamstats, err
}

// GetProjectStats will fetch all stats for a specific project. Similar to GetOrganizationStats
func (c *Client) GetProjectStats(o Organization, p Project, stat StatQuery, since, until int64, resolution *string) ([]Stat, error) {
	var stats []Stat
	statrequest := &statRequest{
		Stat:       stat,
		Since:      since,
		Until:      until,
		Resolution: resolution,
	}

	err := c.doWithQuery("GET", fmt.Sprintf("projects/%s/%s/stats", *o.Slug, *p.Slug), &stats, nil, statrequest)
	return stats, err
}
