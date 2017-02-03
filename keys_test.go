package sentry

import (
	"testing"
)

func TestKeysResource(t *testing.T) {
	t.Parallel()
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}
	team, err := client.CreateTeam(org, "test team for keys", nil)
	if err != nil {
		t.Fatal(err)
	}
	project, err := client.CreateProject(org, team, "Test python project keys", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Create a new key for project", func(t *testing.T) {
		key, err := client.CreateClientKey(org, project, "Test client key")
		if err != nil {
			t.Error(err)
		}
		if key.Label != "Test client key" {
			t.Error("Key does not have correct label")
		}
		t.Run("Update name of client key", func(t *testing.T) {

			key, err = client.UpdateClientKey(org, project, key, "This is a new name")
			if err != nil {
				t.Error(err)
			}

			if key.Label != "This is a new name" {
				t.Error("Failed to update to a new name")
			}

			t.Run("Deleting the client key", func(t *testing.T) {
				err := client.DeleteClientKey(org, project, key)
				if err != nil {
					t.Error(err)
				}
			})
		})
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
