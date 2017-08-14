package sentry

import (
	"testing"
)

func TestReleaseResource(t *testing.T) {
	t.Parallel()
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

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
	t.Run("Fetch all releases for a project", func(t *testing.T) {
		newrelease := NewRelease{
			Version: "abcdefgh123456",
		}

		_, relerr := client.CreateRelease(org, project, newrelease)
		if relerr != nil {
			t.Error(err)
		}

		releases, _, relserr := client.GetReleases(org, project)
		if relserr != nil {
			t.Error(relserr)
		}

		if len(releases) == 0 {
			t.Error("Should be at least one release")
		}

	})
	delprojerr := client.DeleteProject(org, project)
	if delprojerr != nil {
		t.Fatal(delprojerr)
	}
	if delteam := client.DeleteTeam(org, team); delteam != nil {
		t.Error(delteam)
	}
}
