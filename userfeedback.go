package sentry

import (
	"fmt"
	"time"
)

// UserFeedback is the feedback on a given project and issue
type UserFeedback struct {
	EventID     *string    `json:"event_id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Comments    *string    `json:"comments,omitempty"`
	Email       *string    `json:"email,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	Issue       *Issue     `json:"issue,omitempty"`
	ID          *string    `json:"id,omitempty"`
}

// NewUserFeedback will generate a new UserFeedback and type correctly
func NewUserFeedback(name, comments, email, eventID string) UserFeedback {
	return UserFeedback{
		Name:     &name,
		Comments: &comments,
		Email:    &email,
		EventID:  &eventID,
	}
}

// SubmitUserFeedback is used when you want to submit feedback to a organizations project
func (c *Client) SubmitUserFeedback(o Organization, p Project, u *UserFeedback) error {
	return c.do("POST", fmt.Sprintf("projects/%s/%s/user-feedback", *o.Slug, *p.Slug), &u, &u)
}

// GetProjectUserFeedback is used to fetch all feedback given for a certain project
func (c *Client) GetProjectUserFeedback(o Organization, p Project) ([]UserFeedback, *Link, error) {
	var feedback []UserFeedback
	link, err := c.doWithPagination("GET", fmt.Sprintf("projects/%s/%s/user-feedback", *o.Slug, *p.Slug), &feedback, nil)
	return feedback, link, err
}
