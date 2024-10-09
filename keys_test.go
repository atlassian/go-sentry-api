package sentry

import (
	"testing"
)

const (
	testClientKeyName = "Test client key"
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
		key, err := client.CreateClientKey(org, project, testClientKeyName)
		if err != nil {
			t.Error(err)
		}
		if key.Label != testClientKeyName {
			t.Error("Key does not have correct label")
		}
		if key.Label != testClientKeyName {
			t.Error("Key does not have correct name")
		}
		if key.RateLimit != nil {
			t.Error("freshly created keys should not have rate limiting")
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
		t.Run("Update rate limit for client key", func(t *testing.T) {
			resKey, err := client.SetClientKeyRateLimit(org, project, key, 1000, 60)
			if err != nil {
				t.Error(err)
			}
			if resKey.RateLimit == nil {
				t.Error("missing rate limit in updated key")
			}
			if resKey.RateLimit.Count != 1000 || resKey.RateLimit.Window != 60 {
				t.Error("failed to apply rate limit for key")
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
		t.Run("Get client key", func(t *testing.T) {
			key, err := client.GetClientKey(org, project, key.ID)
			if err != nil {
				t.Error(err)
			}

			if key.ID != key.ID {
				t.Error("Failed to fetch key")
			}

			if key.Name != testClientKeyName {
				t.Error("Key does not have correct name")
			}
		})
	})

}
