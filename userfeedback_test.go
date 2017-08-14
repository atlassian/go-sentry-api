package sentry

import (
	"testing"
)

func TestUserFeedbackResource(t *testing.T) {
	t.Parallel()

	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	t.Run("Submit user feedback without a issue", func(t *testing.T) {
		issues, _, _ := client.GetIssues(org, project, nil, nil, nil)

		if len(issues) <= 0 {
			t.Fatal("No issues found.")
		}

		issue := issues[0]

		events, _, _ := client.GetIssueEvents(issue)

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
