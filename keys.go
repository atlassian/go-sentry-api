package sentry

import (
	"fmt"
	"time"
)

// DSN is the actual connection strings used to send errors
type DSN struct {
	Secret string `json:"secret,omitempty"`
	CSP    string `json:"csp,omitempty"`
	Public string `json:"public,omitempty"`
}

// Key is a DSN that sentry has made
type Key struct {
	Label       string        `json:"label,omitempty"`
	DSN         DSN           `json:"dsn,omitempty"`
	Secret      string        `json:"secret,omitempty"`
	ID          string        `json:"id,omitempty"`
	DateCreated time.Time     `json:"dateCreated,omitempty"`
	Public      string        `json:"public,omitempty"`
	RateLimit   *KeyRateLimit `json:"rateLimit,omitempty"`
}

// KeyRateLimit is the rate limit for a DSN
type KeyRateLimit struct {
	Count  int `json:"count"`
	Window int `json:"window"`
}

type nameReq struct {
	Name string `json:"name"`
}

// CreateClientKey creates a new client key for a project and org
func (c *Client) CreateClientKey(o Organization, p Project, name string) (Key, error) {
	var key Key
	req := &nameReq{
		Name: name,
	}
	err := c.do("POST", fmt.Sprintf("projects/%s/%s/keys", *o.Slug, *p.Slug), &key, &req)
	return key, err
}

// DeleteClientKey deletes a client key for a project and org
func (c *Client) DeleteClientKey(o Organization, p Project, k Key) error {
	return c.do("DELETE", fmt.Sprintf("projects/%s/%s/keys/%s", *o.Slug, *p.Slug, k.ID), nil, nil)
}

// UpdateClientKey updates the name only of a key
func (c *Client) UpdateClientKey(o Organization, p Project, k Key, name string) (Key, error) {
	var key Key
	req := &nameReq{
		Name: name,
	}
	err := c.do("PUT", fmt.Sprintf("projects/%s/%s/keys/%s", *o.Slug, *p.Slug, k.ID), &key, &req)
	return key, err
}

// GetClientKeys fetches all client keys of the given project
func (c *Client) GetClientKeys(o Organization, p Project) ([]Key, error) {
	var keys []Key
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/keys", *o.Slug, *p.Slug), &keys, nil)
	return keys, err
}

// SetClientKeyRateLimit updates the rate limit only of a key. window is in seconds.
func (c *Client) SetClientKeyRateLimit(o Organization, p Project, k Key, count, window int) (Key, error) {
	var key Key
	req := &Key{RateLimit: &KeyRateLimit{Count: count, Window: window}}
	err := c.do("PUT", fmt.Sprintf("projects/%s/%s/keys/%s", *o.Slug, *p.Slug, k.ID), &key, &req)
	return key, err
}
