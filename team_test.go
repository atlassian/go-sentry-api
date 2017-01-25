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

	t.Run("Update the team name", func(t *testing.T) {
		team.Name = "Updated team name for testing"
		err := client.UpdateTeam(org, team)
		if err != nil {
			t.Error(err)
		}
		if team.Name != "Updated team name for testing" {
			t.Error("Failed to update team on server side")
		}
	})

	t.Run("Create new project for team", func(t *testing.T) {
		if proj, err := client.CreateProject(org, team, "Python test project", nil); err != nil {
			t.Error(err)
		} else {
			if proj.Name != "Python test project" {
				t.Error("Project name does not match")
			}
			t.Run("Delete project for org", func(t *testing.T) {
				err := client.DeleteProject(org, proj)
				if err != nil {
					t.Error(err)
				}
			})
		}
	})

	t.Run("Get all projects for this team", func(t *testing.T) {

		newproject, err := client.CreateProject(org, team, "Example project for sentry", nil)
		if err != nil {
			t.Fatal(err)
		}
		projects, err := client.GetTeamProjects(org, team)
		if err != nil {
			t.Error(err)
		}

		first := projects[0]
		if first.Name != newproject.Name {
			t.Error("First project in list not project created")
		}

	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Fatal(err)
	}
}
