package sentry

import (
	"time"
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
	Type *string `json:"type,omitempty"`
}

// Event is the event that was created on the app and sentry reported on
type Event struct {
	EventID         string                    `json:"eventID,omitempty"`
	UserReport      *string                   `json:"userReport,omitempty"`
	NextEventID     *string                   `json:"nextEventID,omitempty"`
	PreviousEventID *string                   `json:"previousEventID,omitempty"`
	Message         *string                   `json:"message,omitempty"`
	ID              *string                   `json:"id,omitempty"`
	Size            *int                      `json:"size,omitempty"`
	Platform        *string                   `json:"platform,omitempty"`
	Type            *string                   `json:"type,omitempty"`
	Metadata        *map[string]string        `json:"metadata,omitempty"`
	Tags            *[]Tag                    `json:"tags,omitempty"`
	DateCreated     *time.Time                `json:"dateCreated,omitempty"`
	DateReceived    *time.Time                `json:"dateReceived,omitempty"`
	User            *User                     `json:"user,omitempty"`
	Entries         *[]map[string]interface{} `json:"entries,omitempty"`
	Packages        *map[string]string        `json:"packages,omitempty"`
	SDK             *string                   `json:"sdk,omitempty"`
	Contexts        *map[string]string        `json:"contexts,omitempty"`
	Context         *map[string]interface{}   `json:"context,omitempty"`
	Release         *Release                  `json:"release,omitempty"`
	GroupID         *string                   `json:"groupID,omitempty"`
}
