package sentry

import (
	"testing"
)

func TestBulkResourceModifyDelete(t *testing.T) {

	client := newTestClient(t)

	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	project, cleanupproj := createProjectHelper(t, team)

	defer cleanupproj()
	defer cleanup()

	createMessagesHelper(t, client, org, project, 10)

	t.Run("Fetch all messages for project", func(t *testing.T) {
		issues, link, err := client.GetIssues(org, project, nil, nil, nil)
		if err != nil {
			t.Error(err)
		}

		ids := make([]string, 0)
		for _, issue := range issues {
			ids = append(ids, *issue.ID)
		}

		for link.Next.Results {
			for _, issue := range issues {
				ids = append(ids, *issue.ID)
			}
			link, err = client.GetPage(link.Next, &issues)
			if err != nil {
				t.Error(err)
			}
		}

		t.Run("Modify all messages to be resolved", func(t *testing.T) {
			resolved := Resolved
			resp, err := client.BulkMutateIssues(org, project, IssueBulkRequest{
				Status: &resolved,
			}, &ids, nil)

			if err != nil {
				t.Skip(err)
			}
			if resp.Status != nil {
				if *resp.Status != Resolved {
					t.Error("Should have made this resolved")
				}
			}

		})

		t.Run("Delete all of the messages", func(t *testing.T) {
			err := client.BulkDeleteIssues(org, project, ids)
			if err != nil {
				t.Error(err)
			}
		})
	})

}
