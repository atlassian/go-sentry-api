package sentry

import (
	"testing"
)

func TestTeamResource(t *testing.T) {
	org, err := client.GetOrganization("sentry")
	if err != nil {
		t.Fatal(err)
	}
	team, err := client.CreateTeam(org, "Test team for Go Client", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Verify team creation", func(t *testing.T) {
		if team.Name != "Test team for Go Client" {
			t.Error("Team name is not correct")
		}
	})

	t.Run("Create new project for team", func(t *testing.T) {
		if proj, err := client.CreateProject(org, team, "Python test project", nil); err != nil {
			t.Error(err)
		} else {
			if proj.Name != "Python test project" {
				t.Error("Project name does not match")
			}
		}
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Fatal(err)
	}
}
