package sentry

import (
	"fmt"
	"net/http"
	"time"
)

// Member is a sentry member
type Member struct {
	Role        string    `json:"role,omitempty"`
	Name        string    `json:"name"`
	RoleName    string    `json:"roleName"`
	User        *User     `json:"user"`
	DateCreated time.Time `json:"dateCreated"`
	ID          *string   `json:"id,omitempty"`
	Teams       []string  `json:"teams"`
	Pending     *bool     `json:"pending"`
	Email       string    `json:"email,omitempty"`
}

// InviteMember invites a member to join a organization
func (c *Client) InviteMember(o Organization, email, role string, teams []string) (Member, error) {
	var member Member
	memberreq := &Member{
		Email: email,
		Role:  role,
		Teams: teams,
	}

	err := c.do(http.MethodPost, fmt.Sprintf("organizations/%s/members", *o.Slug), &member, memberreq)
	return member, err
}

// RemoveMember removes a member from a organization
func (c *Client) RemoveMember(o Organization, m Member) error {
	return c.do(http.MethodDelete, fmt.Sprintf("organizations/%s/members/%s", *o.Slug, *m.ID), nil, nil)
}

// UpdateMember updates role of a member
func (c *Client) UpdateMember(o Organization, m Member) error {
	return c.do(http.MethodPut, fmt.Sprintf("organizations/%s/members/%s", *o.Slug, *m.ID), &m, &m)
}
