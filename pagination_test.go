package sentry

import (
	"testing"
)

func TestLinkPagination(t *testing.T) {
	examplelink := `<https://sentry.io/api/0/projects/the-interstellar-jurisdiction/pump-station/releases/?&cursor=100:-1:1>; rel="previous"; results="true"; cursor="100:-1:1", <https://sentry.io/api/0/projects/the-interstellar-jurisdiction/pump-station/releases/?&cursor=100:1:0>; rel="next"; results="true"; cursor="100:1:0`

	link := NewLink(examplelink)

	if link.Next.URL != "https://sentry.io/api/0/projects/the-interstellar-jurisdiction/pump-station/releases/?&cursor=100:1:0" {
		t.Errorf("Link next isnt correct: %s", link.Next.URL)
	}

	if link.Previous.URL != "https://sentry.io/api/0/projects/the-interstellar-jurisdiction/pump-station/releases/?&cursor=100:-1:1" {
		t.Errorf("Link previous isnt correct: %s", link.Previous.URL)
	}

	if !link.Next.Results {
		t.Error("Results should be set to true for next")
	}

	if !link.Previous.Results {
		t.Error("Results should be set to true for previous")
	}

}
