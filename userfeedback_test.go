package sentry

import (
	"testing"
)

func TestUserFeedbackResource(t *testing.T) {

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

	t.Run("Submit user feedback with a issue", func(t *testing.T) {
		issues, _, _ := client.GetIssues(org, project, nil, nil, nil)

		if len(issues) == 0 {
			t.Fatal("No issues found.")
		}

		issue := issues[0]

		events, _, _ := client.GetIssueEvents(issue)
		if len(events) == 0 {
			t.Fatal("no events found")
		}

		feedback := NewUserFeedback("Colin Wood", "This is a great feature", "cwood@testing.com", events[0].EventID)

		err := client.SubmitUserFeedback(org, project, &feedback)
		if err != nil {
			t.Error(err)
		}
		if feedback.DateCreated == nil {
			t.Error("Date created didnt get updated when posted to the backend")
		}

		t.Run("Fetch all feedback for this project", func(t *testing.T) {
			feedbacks, link, err := client.GetProjectUserFeedback(org, project)
			if err != nil {
				t.Error(err)
			} else {

				if len(feedbacks) == 0 {
					t.Error("Should at least be one feedback for this project")
				}
				if link.Next.Results {
					t.Error("Should only be one page of results")
				}
			}
		})
	})

}
