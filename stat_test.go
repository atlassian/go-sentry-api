package sentry

import (
	"testing"
	"time"
)

func TestStat(t *testing.T) {
	t.Parallel()

	now := time.Now().Unix()
	hourlater := time.Duration(1) * time.Hour
	later := now - int64(hourlater.Seconds())

	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Run Getting Organization Stats", func(t *testing.T) {
		if err != nil {
			t.Fatal(err)
		}
		stats, err := client.GetOrganizationStats(org, StatReceived, later, now, nil)
		if err != nil {
			t.Error(err)
		}

		if len(stats) <= 0 {
			t.Error("No stats were returned")
		}
	})

	t.Run("Run Getting Team Stats", func(t *testing.T) {
		team, cleanup := createTeamHelper(t)
		defer cleanup()

		_, err = client.GetTeamStats(org, team, StatReceived, later, now, nil)
		if err != nil {
			t.Error(err)
		}

		err = client.DeleteTeam(org, team)
		if err != nil {
			t.Error("Failed to delete test team")
		}
	})

}
