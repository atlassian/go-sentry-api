package sentry

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	// Resolved helps mark a issue or others as resolved
	Resolved Status = "resolved"
	// Unresolved helps mark a issue or others as unresolved
	Unresolved Status = "unresolved"
	// Ignored helps mark a issue or others as ignored
	Ignored Status = "ignored"
)

// Hash is returned via the issue_id/hashes
type Hash struct {
	ID string `json:"id,omitempty"`
}

// Status is used to make consts for statuses
type Status string

// IssueStats is the stats of a issue
type IssueStats struct {
	TwentyFourHour *[]Stat `json:"24h,omitempty"`
	FourteenDays   *[]Stat `json:"14d,omitempty"`
	ThirtyDays     *[]Stat `json:"30d,omitempty"`
}

// IssueTagValue represents a tags value
type IssueTagValue struct {
	Count     *int64     `json:"count,omitempty"`
	FirstSeen *time.Time `json:"firstSeen,omitempty"`
	ID        *string    `json:"iD,omitempty"`
	Key       *string    `json:"key,omitempty"`
	LastSeen  *time.Time `json:"lastSeen,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Value     *string    `json:"value,omitempty"`
}

// IssueTag is a tag used in a sentry issue
type IssueTag struct {
	UniqueValues int             `json:"uniqueValues,omitempty"`
	ID           string          `json:"id,omitempty"`
	Key          string          `json:"key,omitempty"`
	Name         string          `json:"name,omitempty"`
	TopValues    []IssueTagValue `json:"topValues,omitempty"`
}

// Avatar is used for a users avatar
type Avatar struct {
	AvatarType *string `json:"avatarType,omitempty"`
	AvatarUUID *string `json:"avatarUuid,omitempty"`
}

// InternalUser is a user on sentry and not a external customer
type InternalUser struct {
	AvatarURL  *string    `json:"avatarUrl,omitempty"`
	DateJoined *time.Time `json:"dateJoined,omitempty"`
	Email      *string    `json:"email,omitempty"`
	Has2FA     *bool      `json:"has2fa,omitempty"`
	ID         *string    `json:"iD,omitempty"`
	IsActive   *bool      `json:"isActive,omitempty"`
	IsManaged  *bool      `json:"isManaged,omitempty"`
	LastLogin  *time.Time `json:"lastLogin,omitempty"`
	Name       *string    `json:"name,omitempty"`
	Username   *string    `json:"username,omitempty"`
}

// Activity is what current activity has happend on a issue
type Activity struct {
	Data        *map[string]interface{} `json:"data,omitempty"`
	DateCreated *time.Time              `json:"dateCreated,omitempty"`
	ID          *string                 `json:"id,omitempty"`
	Type        *string                 `json:"type,omitempty"`
	User        *InternalUser           `json:"user,omitempty"`
}

// Issue returns a issue found in sentry
type Issue struct {
	Annotations         *[]string               `json:"annotations,omitempty"`
	AssignedTo          *InternalUser           `json:"assignedTo,omitempty"`
	Activity            *[]Activity             `json:"activity,omitempty"`
	Count               *string                 `json:"count,omitempty"`
	Culprit             *string                 `json:"culprit,omitempty"`
	FirstSeen           *time.Time              `json:"firstSeen,omitempty"`
	HasSeen             *bool                   `json:"hasSeen,omitempty"`
	ID                  *string                 `json:"id,omitempty"`
	IsBookmarked        *bool                   `json:"isBookmarked,omitempty"`
	IsPublic            *bool                   `json:"isPublic,omitempty"`
	IsSubscribed        *bool                   `json:"isSubscribed,omitempty"`
	LastSeen            *time.Time              `json:"lastSeen,omitempty"`
	Level               *string                 `json:"level,omitempty"`
	Logger              *string                 `json:"logger,omitempty"`
	Metadata            *map[string]interface{} `json:"metadata,omitempty"`
	NumComments         *int                    `json:"numComments,omitempty"`
	Permalink           *string                 `json:"permalink,omitempty"`
	Project             *Project                `json:"project,omitempty"`
	ShareID             *string                 `json:"shareId,omitempty"`
	ShortID             *string                 `json:"shortId,omitempty"`
	Stats               *IssueStats             `json:"stats,omitempty"`
	Status              *Status                 `json:"status,omitempty"`
	StatusDetails       *map[string]interface{} `json:"statusDetails,omitempty"`
	SubscriptionDetails *map[string]string      `json:"subscriptionDetails,omitempty"`
	Tags                *[]IssueTag             `json:"tags,omitempty"`
	Title               *string                 `json:"title,omitempty"`
	Type                *string                 `json:"type,omitempty"`
	UserCount           *int                    `json:"userCount,omitempty"`
	UserReportCount     *int                    `json:"userReportCount,omitempty"`
}

type issueQuery struct {
	StatsPeriod   *string
	ShortIDLookup *bool
	Query         *string
}

func (i *issueQuery) ToQueryString() string {
	query := url.Values{}
	if i.StatsPeriod != nil {
		query.Add("statsPeriod", *i.StatsPeriod)
	}
	if i.ShortIDLookup != nil {
		query.Add("shortIdLookup", strconv.FormatBool(*i.ShortIDLookup))
	}
	if i.Query != nil {
		query.Add("query", *i.Query)
	}

	return query.Encode()
}

// GetIssues will fetch all issues for organization and project
func (c *Client) GetIssues(o Organization, p Project, StatsPeriod *string, ShortIDLookup *bool, query *string) ([]Issue, *Link, error) {
	var issues []Issue

	issueFilter := &issueQuery{
		StatsPeriod:   StatsPeriod,
		ShortIDLookup: ShortIDLookup,
		Query:         query,
	}

	link, err := c.doWithPaginationQuery(
		"GET", fmt.Sprintf("projects/%s/%s/issues", *o.Slug, *p.Slug), &issues, nil, issueFilter)
	return issues, link, err
}

// GetIssue will fetch a issue by its ID as a string
func (c *Client) GetIssue(id string) (Issue, error) {
	var issue Issue
	err := c.do("GET", fmt.Sprintf("issues/%s", id), &issue, nil)
	return issue, err
}

// GetIssueHashes will fetch all hashes for a issue
func (c *Client) GetIssueHashes(i Issue) ([]Hash, *Link, error) {
	var hashes []Hash
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/hashes", *i.ID), &hashes, nil)
	return hashes, link, err
}

// GetIssueTags will fetch all tags for a issue
func (c *Client) GetIssueTags(i Issue) ([]IssueTag, *Link, error) {
	var tags []IssueTag
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/tags", *i.ID), &tags, nil)
	return tags, link, err
}

// GetIssueTag will fetch a tag used in a issue. Eg; environment, release, server
func (c *Client) GetIssueTag(i Issue, tagname string) (IssueTag, error) {
	var tag IssueTag
	err := c.do("GET", fmt.Sprintf("issues/%s/tags/%s", *i.ID, tagname), &tag, nil)
	return tag, err
}

// GetIssueTagValues will fetch all values for a issues tag
func (c *Client) GetIssueTagValues(i Issue, tag IssueTag) ([]IssueTagValue, *Link, error) {
	var values []IssueTagValue
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/tags/%s/values", *i.ID, tag.Key), &values, nil)
	return values, link, err
}

// GetIssueEvents will fetch all events for a issue
func (c *Client) GetIssueEvents(i Issue) ([]Event, *Link, error) {
	var events []Event
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/events", *i.ID), &events, nil)
	return events, link, err
}

// UpdateIssue will update status, assign to, hasseen, isbookmarked and issubscribed
func (c *Client) UpdateIssue(i Issue) error {
	return c.do("PUT", fmt.Sprintf("issues/%s", *i.ID), &i, &i)
}

// DeleteIssue will delete an issue
func (c *Client) DeleteIssue(i Issue) error {
	return c.do("DELETE", fmt.Sprintf("issues/%s", *i.ID), nil, nil)
}
