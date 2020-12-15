package sentry

import (
	"fmt"
	"testing"

	raven "github.com/getsentry/sentry-go"
)

func createMessagesHelper(t *testing.T, client *Client, org Organization, project Project, numOfMessages int) {

	dsnkey, err := client.CreateClientKey(org, project, "testing key")
	if err != nil {
		t.Fatal(err)
	}

	scope := raven.NewScope()
	scope.SetExtra("server", "app-node-01")

	ravenClient, err := raven.NewClient(raven.ClientOptions{
		Dsn:       dsnkey.DSN.Secret,
		Transport: raven.NewHTTPSyncTransport(),
	})

	hub := raven.NewHub(ravenClient, scope)

	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= numOfMessages; i++ {
		hub.CaptureMessage(fmt.Sprintf("failed to execute on id %d", i))
	}
}

func TestIssueResource(t *testing.T) {

	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	createMessagesHelper(t, client, org, project, 10)

	t.Run("Get all issues with a query of resolved", func(t *testing.T) {
		query := "is:resolved"
		issues, _, err := client.GetIssues(org, project, nil, nil, &query)
		if err != nil {
			t.Error(err)
		}
		if len(issues) >= 1 {
			t.Error("Should have not had any issues marked as resolved")
		}
	})

	t.Run("Get all issues for a project", func(t *testing.T) {
		issues, link, err := client.GetIssues(org, project, nil, nil, nil)
		if err != nil {
			t.Error(err)
		}
		if len(issues) <= 0 {
			t.Fatal("No issues found for this project")
		}
		if link.Previous.Results {
			t.Error("Should be no new results")
		}

		t.Run("Get issues with statsPeriod of 14 days", func(t *testing.T) {
			period := "14d"
			issues, _, err := client.GetIssues(org, project, &period, nil, nil)
			if err != nil {
				t.Error(err)
			}
			if len(issues) <= 0 {
				t.Fatal("No issues found for this project")
			}
			for _, issue := range issues {
				if issue.Stats.FourteenDays == nil {
					t.Fatal("We should be able to get 14 days of stats for this issue but didn't.")
				}
			}
		})

		t.Run("Get hashes for issue", func(t *testing.T) {
			hashes, link, err := client.GetIssueHashes(issues[0])
			if err != nil {
				t.Error(err)
			}
			if len(hashes) == 0 {
				t.Error("Should be at least one hash in list")
			}
			if link.Next.Results {
				t.Error("Should not be any other results but there is some")
			}

		})

		t.Run("Get the issue only", func(t *testing.T) {
			issue, err := client.GetIssue(*issues[0].ID)
			if err != nil {
				t.Errorf("Failed to get issue: %s", err)
			}
			if *issue.ID != *issues[0].ID {
				t.Error("Somehow not the same ID? How is this possible")
			}
		})

		t.Run("Get events for this issue", func(t *testing.T) {
			events, _, err := client.GetIssueEvents(issues[0])
			if err != nil {
				t.Error(err)
			}
			if len(events) == 0 {
				t.Errorf("Should be at least more than 1 event %v", events)
			}
		})

		t.Run("Modify first issue found in project", func(t *testing.T) {
			firstIssue := issues[0]

			resolved := Resolved
			firstIssue.Status = &resolved
			firstIssue.StatusDetails = &map[string]interface{}{
				"inNextRelease": true,
			}

			if err := client.UpdateIssue(firstIssue); err != nil {
				t.Error(err)
			}

			if *firstIssue.Status != Resolved {
				t.Error("Status did not get updated")
			}

			details, ok := (*firstIssue.StatusDetails)["inNextRelease"].(bool)
			if !ok {
				t.Error("Status did not get updated")
			}

			if !details {
				t.Error("Status did not get updated")
			}

			t.Run("Delete the first issue in this project", func(t *testing.T) {
				err := client.DeleteIssue(firstIssue)
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
