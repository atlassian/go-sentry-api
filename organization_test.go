package sentry

import (
	"testing"
	"time"
)

func TestOrganization(t *testing.T) {
	org, err := client.CreateOrganization("Test Org via Go Client")
	if err != nil {
		if err.(SentryApiError).StatusCode == 429 {
			t.Skip("Cant create organization due to rate limiting skipping tests suite")
		} else {
			t.Fatal(err)
		}
	}
	t.Run("Organization Get", func(t *testing.T) {
		org, err := client.GetOrganization(*org.Slug)
		if err != nil {
			t.Fatal(err)
		}
		if org.Name != "Test Org via Go Client" {
			t.Error("Name is not atlassian")
		}
		t.Run("Organization Update", func(t *testing.T) {
			if org.Name != "Test Org via Go Client" {
				t.Error("New org is not the right slug: %v", org)
			}

			org.Name = "New Updated Name"

			if err := client.UpdateOrganization(org); err != nil {
				t.Fatal(err)
			}

			if org.Name != "New Updated Name" {
				t.Error("Org didnt have new name after update")
			}
		})
	})
	t.Run("Fetch organizations", func(t *testing.T) {
		orgs, err := client.GetOrganizations()
		if err != nil {
			t.Fatal(err)
		}
		if len(orgs) <= 0 {
			t.Error("Didnt fetch any orgs")
		}
	})
	if err := client.DeleteOrganization(org); err != nil {
		t.Fatal(err)
	}
}

func TestOrganizationStat(t *testing.T) {
	now := time.Now().Unix()
	hourlater := time.Duration(1) * time.Hour
	later := now - int64(hourlater.Seconds())

	org, err := client.GetOrganization("sentry")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.GetOrganizationStats(org, StatReceived, later, now, nil)
	if err != nil {
		t.Error(err)
	}

}
