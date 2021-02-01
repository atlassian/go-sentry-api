package sentry

import (
	"fmt"
	"net/url"
	"time"
)

// Member is a sentry member
type Member struct {
	Email       string     `json:"email,omitempty"`
	Expired     *bool      `json:"expired,omitempty"`
	Name        *string    `json:"name,omitempty"`
	IsPending   *bool      `json:"isPending,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	Role        string     `json:"role,omitempty"`
	ID          *string    `json:"id,omitempty"`
	RoleName    *string    `json:"roleName,omitempty"`
	Teams       []string   `json:"teams,omitempty"`
}

type memberQuery struct {
	query string `json:"query,omitempty"`
}

func (o *memberQuery) ToQueryString() string {
	query := url.Values{}
	query.Add("query", string(o.query))
	return query.Encode()
}

// CreateMember takes an email and creates a new member
func (c *Client) CreateMember(o Organization, email string) (Member, error) {
	var member Member
	memberRole := "member"
	memberreq := Member{
		Email: email,
		Role:  memberRole,
	}

	err := c.do("POST", fmt.Sprintf("organizations/%s/members", *o.Slug), &member, &memberreq)
	return member, err
}

// GetMemberByEmail takes a user email and returns back the user
func (c *Client) GetMemberByEmail(o Organization, memberEmail string) (Member, error) {
	var members []Member

	err := c.doWithQuery("GET", fmt.Sprintf("organizations/%s/members", *o.Slug), &members, nil, &memberQuery{memberEmail})
	if err != nil {
		return Member{}, fmt.Errorf("failed to get member with error: %s", err.Error())
	}

	if len(members) == 0 {
		return Member{}, fmt.Errorf("no member with that email found")
	}

	return members[0], err
}

// AddExistingMemberToTeam takes a member and adds them to a team
func (c *Client) AddExistingMemberToTeam(o Organization, t Team, m Member) error {
	return c.do("POST", fmt.Sprintf("organizations/%s/members/%s/teams/%s", *o.Slug, *m.ID, *t.Slug), nil, nil)
}

// DeleteMember takes a member and deletes from the org
func (c *Client) DeleteMember(o Organization, m Member) error {
	return c.do("DELETE", fmt.Sprintf("organizations/%s/members/%s", *o.Slug, *m.ID), nil, nil)
}

// MakeAdmin takes a member and makes them admin
func (c *Client) MakeAdmin(o Organization, a Member) error {
	a.Role = "admin"

	return c.do("PUT", fmt.Sprintf("organizations/%s/members/%s", *o.Slug, *a.ID), nil, &a)
}
