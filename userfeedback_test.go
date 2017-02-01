package sentry

import (
	"testing"
)

func TestUserFeedbackResource(t *testing.T) {
	t.Parallel()
	org, err := client.GetOrganization("sentry")
	if err != nil {
		t.Fatal(err)
	}
	team, err := client.CreateTeam(org, "test team for go project", nil)
	if err != nil {
		t.Fatal(err)
	}
	project, err := client.CreateProject(org, team, "Test python project", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Submit user feedback without a issue", func(t *testing.T) {
		issues, _, _ := client.GetIssues(org, project)
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

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}
}
