package sentry

import (
	"testing"
)

func TestProjectResource(t *testing.T) {
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

	t.Run("Fetch project from endpoint", func(t *testing.T) {
		endpointproject, err := client.GetProject(org, *project.Slug)
		if err != nil {
			t.Error(err)
		}
		if endpointproject.Team.Name != "test team for go project" {
			t.Error("Project fetch didnt have the right team name")
		}
		if endpointproject.Organization.Name != "Sentry" {
			t.Error("Projects organization is not sentry")
		}
	})

}
