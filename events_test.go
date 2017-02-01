package sentry

import (
	"fmt"
	"testing"

	"github.com/getsentry/raven-go"
)

func TestEventsResource(t *testing.T) {
	t.Parallel()
	org, err := client.GetOrganization("sentry")
	if err != nil {
		t.Fatal(err)
	}
	team, err := client.CreateTeam(org, "test team for go project", nil)
	if err != nil {
		t.Fatal(err)
	}
	project, err := client.CreateProject(org, team, "Test go project issues", nil)
	if err != nil {
		t.Fatal(err)
	}
	dsnkey, err := client.CreateClientKey(org, project, "testing key")
	if err != nil {
		t.Fatal(err)
	}

	ravenClient, err := raven.New(dsnkey.DSN.Secret)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		ravenClient.CaptureMessageAndWait(fmt.Sprintf("Testing message %d", i), map[string]string{
			"server": fmt.Sprintf("%d-app-node", i),
		}, &raven.Message{
			Message: fmt.Sprintf("Testing some stuff out Brah %d", i),
		})
	}

	issues, _, err := client.GetIssues(org, project)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Read out all events for a issue", func(t *testing.T) {

		for _, issue := range issues {

			latest, err := client.GetLatestEvent(issue)
			if err != nil {
				t.Error(err)
			}

			oldest, err := client.GetOldestEvent(issue)
			if err != nil {
				t.Error(err)
			}

			t.Logf("Issue ID: %s Latest: %v, Oldest: %v \n", *issue.ID, latest, oldest)

		}

	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
