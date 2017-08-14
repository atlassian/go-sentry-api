package sentry

import (
	"testing"
)

func TestIssueResource(t *testing.T) {
	t.Parallel()

	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

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
		if len(issues) < 0 {
			t.Error("No issues found for this project")
		}
		if link.Previous.Results {
			t.Error("Should be no new results")
		}

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
				t.Error("Should be at least more than 1 event")
			}
		})

		t.Run("Modify first issue found in project", func(t *testing.T) {
			firstIssue := issues[0]

			resolved := Resolved
			firstIssue.Status = &resolved

			if err := client.UpdateIssue(firstIssue); err != nil {
				t.Error(err)
			}

			if *firstIssue.Status != Resolved {
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
