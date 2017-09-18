package sentry

import (
	"testing"
)

func TestKeysResource(t *testing.T) {
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

	t.Run("Create a new key for project", func(t *testing.T) {
		key, err := client.CreateClientKey(org, project, "Test client key")
		if err != nil {
			t.Error(err)
		}
		if key.Label != "Test client key" {
			t.Error("Key does not have correct label")
		}
		t.Run("List client keys", func(t *testing.T) {
			keys, err := client.GetClientKeys(org, project)
			if err != nil {
				t.Error(err)
			}
			if len(keys) != 2 {
				t.Errorf("Expected 2 keys, got %d", len(keys))
			}
		})
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

}
