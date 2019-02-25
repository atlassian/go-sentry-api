package sentry

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/atlassian/go-sentry-api/datatype"
)

// Tag is used for a event
type Tag struct {
	Value *string `json:"value,omitempty"`
	Key   *string `json:"key,omitempty"`
}

//User is the user that was affected
type User struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// Entry is the entry for the message/stacktrace/etc...
type Entry struct {
	Type string          `json:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

//GetInterface will convert the entry into a go interface
func (e *Entry) GetInterface() (string, interface{}, error) {
	var destination interface{}

	switch e.Type {
	case "message":
		destination = new(datatype.Message)
	case "stacktrace":
		destination = new(datatype.Stacktrace)
	case "exception":
		destination = new(datatype.Exception)
	case "request":
		destination = new(datatype.Request)
	case "template":
		destination = new(datatype.Template)
	case "user":
		destination = new(datatype.User)
	case "query":
		destination = new(datatype.Query)
	case "breadcrumbs":
		destination = new(datatype.Breadcrumb)
	}

	err := json.Unmarshal(e.Data, &destination)
	return e.Type, destination, err
}

// Event is the event that was created on the app and sentry reported on
type Event struct {
	EventID         string                  `json:"eventID,omitempty"`
	UserReport      *string                 `json:"userReport,omitempty"`
	NextEventID     *string                 `json:"nextEventID,omitempty"`
	PreviousEventID *string                 `json:"previousEventID,omitempty"`
	Message         *string                 `json:"message,omitempty"`
	ID              *string                 `json:"id,omitempty"`
	Size            *int                    `json:"size,omitempty"`
	Platform        *string                 `json:"platform,omitempty"`
	Type            *string                 `json:"type,omitempty"`
	Metadata        *map[string]string      `json:"metadata,omitempty"`
	Tags            *[]Tag                  `json:"tags,omitempty"`
	DateCreated     *time.Time              `json:"dateCreated,omitempty"`
	DateReceived    *time.Time              `json:"dateReceived,omitempty"`
	User            *User                   `json:"user,omitempty"`
	Entries         *[]Entry                `json:"entries,omitempty"`
	Packages        *map[string]string      `json:"packages,omitempty"`
	SDK             *map[string]interface{} `json:"sdk,omitempty"`
	Contexts        *map[string]interface{} `json:"contexts,omitempty"`
	Context         *map[string]interface{} `json:"context,omitempty"`
	Release         *Release                `json:"release,omitempty"`
	GroupID         *string                 `json:"groupID,omitempty"`
}

// GetProjectEvent will fetch a event on a project
func (c *Client) GetProjectEvent(o Organization, p Project, eventID string) (Event, error) {
	var event Event
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/events/%s", *o.Slug, *p.Slug, eventID), &event, nil)
	return event, err
}

//GetLatestEvent will fetch the latest event for a issue
func (c *Client) GetLatestEvent(i Issue) (Event, error) {
	var event Event
	err := c.do("GET", fmt.Sprintf("issues/%s/events/latest", *i.ID), &event, nil)
	return event, err
}

//GetOldestEvent will fetch the latest event for a issue
func (c *Client) GetOldestEvent(i Issue) (Event, error) {
	var event Event
	err := c.do("GET", fmt.Sprintf("issues/%s/events/oldest", *i.ID), &event, nil)
	return event, err
}
