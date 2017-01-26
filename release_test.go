package sentry

import (
	"testing"
)

func TestReleaseResource(t *testing.T) {
	org, err := client.GetOrganization("sentry")
	if err != nil {
		t.Fatal(err)
	}
	team, err := client.CreateTeam(org, "test team for go project", nil)
	if err != nil {
		t.Fatal(err)
	}
	project, err := client.CreateProject(org, team, "Test python project", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Create a new release", func(t *testing.T) {

		newrelease := NewRelease{
			Version: "abcdefgh123456",
		}

		rel, relerr := client.CreateRelease(org, project, newrelease)
		if relerr != nil {
			t.Error(err)
		}
		if rel.Version != "abcdefgh123456" {
			t.Error("version does not match expected")
		}

		t.Run("Update the release ref", func(t *testing.T) {
			ref := "123456123"
			rel.Ref = &ref
			updateerr := client.UpdateRelease(org, project, rel)
			if updateerr != nil {
				t.Error(updateerr)
			}
			if rel.Ref != &ref {
				t.Error("Ref did not get updated")
			}

			t.Run("Delete the release", func(t *testing.T) {
				delerr := client.DeleteRelease(org, project, rel)
				if delerr != nil {
					t.Error(err)
				}
			})
		})

	})
	delprojerr := client.DeleteProject(org, project)
	if delprojerr != nil {
		t.Fatal(delprojerr)
	}
	if delteam := client.DeleteTeam(org, team); delteam != nil {
		t.Error(delteam)
	}
}
