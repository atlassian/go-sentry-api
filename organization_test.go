package sentry

import (
	"os"
	"testing"
)

func TestOrganization(t *testing.T) {
	t.Parallel()

	client := newTestClient(t)
	t.Run("Organization Create", func(t *testing.T) {
		if os.Getenv("TEST_ORGS") == "" {
			t.Skip("Skipping testing of orgs since TEST_ORGS not set")
		}
		org, err := client.CreateOrganization("Test Org via Go Client")
		if err != nil {
			if err.(APIError).StatusCode == 429 {
				t.Skip("Cant create organization due to rate limiting skipping tests suite")
			} else {
				t.Fatal(err)
			}
		}
		t.Run("Organization Update", func(t *testing.T) {
			if org.Name != "Test Org via Go Client" {
				t.Error("New org is not the right slug")
			}

			org.Name = "New Updated Name"

			if err := client.UpdateOrganization(org); err != nil {
				t.Fatal(err)
			}

			if org.Name != "New Updated Name" {
				t.Error("Org didnt have new name after update")
			}

			t.Run("Delete the organization", func(t *testing.T) {
				if err := client.DeleteOrganization(org); err != nil {
					t.Error(err)
				}
			})
		})
	})
	t.Run("Fetch organizations", func(t *testing.T) {
		orgs, link, err := client.GetOrganizations()
		if err != nil {
			t.Fatal(err)
		}
		if len(orgs) <= 0 {
			t.Error("Didnt fetch any orgs")
		}
		if link.Next.Results {
			t.Error("Should only be one instance but got more")
		}
	})
}
