package sentry

import (
	"strconv"
	"strings"
)

// Page is a link and if it has results or not
type Page struct {
	URL     string
	Results bool
}

// Link represents a link object as per: https://docs.sentry.io/api/pagination/
type Link struct {
	Previous Page
	Next     Page
}

// NewLink creates a new link object via the link header string
func NewLink(linkheader string) *Link {
	link := &Link{}
	links := strings.SplitN(linkheader, ",", 2)
	for _, page := range links {
		data := strings.SplitN(page, ";", 4)

		pagelink := strings.TrimLeft(strings.TrimSpace(data[0]), "<")
		pagelink = strings.TrimRight(pagelink, ">")

		pagetype := strings.Trim(strings.Split(data[1], "=")[1], `"`)
		results, err := strconv.ParseBool(strings.Trim(strings.Split(strings.TrimSpace(data[2]), "=")[1], `"`))
		if err != nil {
			results = false
		}

		if pagetype == "previous" {
			link.Previous.URL = pagelink
			link.Previous.Results = results
		} else {
			link.Next.URL = pagelink
			link.Next.Results = results
		}
	}

	return link
}

// GetPage will fetch a page via the Link object and decode it from out.
// Should be used like `client.GetPage(link.Previous, make([]Organization, 0))`
func (c *Client) GetPage(p Page, out interface{}) (*Link, error) {
	return c.rawWithPagination("GET", strings.TrimPrefix(p.URL, c.Endpoint), out, nil)
}
