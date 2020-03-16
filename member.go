package sentry

import (
	"fmt"
	"net/url"
	"time"
)

// Member is a sentry member
type Member struct {
	Email       *string    `json:"email,omitempty"`
	Expired     *bool      `json:"expired,omitempty"`
	Name        *string    `json:"name,omitempty"`
	IsPending   *bool      `json:"isPending,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	Role        *string    `json:"role,omitempty"`
	ID          *string    `json:"id,omitempty"`
	RoleName    *string    `json:"roleName,omitempty"`
}

type memberQuery struct {
	query string `json:"query,omitempty"`
}

func (o *memberQuery) ToQueryString() string {
	query := url.Values{}
	query.Add("query", string(o.query))
	return query.Encode()
}

// GetMember takes a user email and returns back the user
func (c *Client) GetMemberByEmail(o Organization, memberEmail string) (Member, error) {
	var members []Member

	err := c.doWithQuery("GET", fmt.Sprintf("organizations/%s/members", *o.Slug), &members, nil, &memberQuery{memberEmail})
	if err != nil {
		return Member{}, fmt.Errorf("failed to get member: %w", err)
	}

	if len(members) == 0 {
		return Member{}, fmt.Errorf("no member with that email found")
	}

	return members[0], err
}

// AddMemberToTeam takes a member and adds them to a team
func (c *Client) AddMemberToTeam(o Organization, t Team, m Member) error {
	return c.do("POST", fmt.Sprintf("organizations/%s/members/%s/teams/%s", *o.Slug, *m.ID, *t.Slug), nil, nil)
}

// MakeAdmin takes a member and makes them admin
func (c *Client) MakeAdmin(o Organization, a Member) error {
	adminRole := "admin"
	a.Role = &adminRole

	return c.do("PUT", fmt.Sprintf("organizations/%s/members/%s", *o.Slug, *a.ID), nil, &a)
}
