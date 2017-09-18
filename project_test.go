package sentry

import (
	"testing"
)

func createProjectHelper(t *testing.T, team Team) (Project, func() error) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {

		t.Fatal(err)
	}

	project, err := client.CreateProject(org, team, generateIdentifier("project"), nil)
	if err != nil {
		t.Fatal(err)
	}
	return project, func() error {
		return client.DeleteProject(org, project)
	}
}

func TestProjectResource(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	t.Run("Fetch project from endpoint", func(t *testing.T) {
		endpointproject, err := client.GetProject(org, *project.Slug)
		if err != nil {
			t.Error(err)
		}
		if endpointproject.Team.Name != team.Name {
			t.Error("Project fetch didnt have the right team name")
		}
	})

	t.Run("Fetch all projects", func(t *testing.T) {
		projects, err := client.GetProjects()
		if err != nil {
			t.Error(err)
		}
		if len(projects) <= 0 {
			t.Error("Should have at least on project but got 0")
		}
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
