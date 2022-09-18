package sentry

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetAlertRules(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	rules, _, err := client.GetAlertRules(org, project)
	require.NoError(t, err, "unable to get alert rules")

	require.Greater(t, len(rules), 0, "no alert rules defined")
	require.Greater(t, len(rules[0].Conditions), 0, "no conditions for rule")
	require.Greater(t, len(rules[0].Actions), 0, "no actions for rule")
}

func TestClient_AddAlertRule(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	rule, err := client.AddAlertRule(org, project, AlertRule{
		Name:        "Test alert rule",
		ActionMatch: AlertRuleMatchAll,
		Frequency:   30,
		Conditions: []AlertRuleCondition{
			{ID: "sentry.rules.conditions.regression_event.RegressionEventCondition"},
		},
		Actions: []AlertRuleAction{
			{ID: "sentry.rules.actions.notify_event_service.NotifyEventServiceAction", Service: "mail"},
		},
	})
	require.NoError(t, err, "unable to create rule")
	require.NotEqual(t, 0, rule.ID, "missing rule ID")
	require.Equal(t, 1, len(rule.Conditions), "wrong condition count")
	require.Equal(t, 1, len(rule.Actions), "wrong action count")

	rules, _, err := client.GetAlertRules(org, project)
	require.NoError(t, err, "unable to get alert rules")
	require.Equal(t, len(rules), 2, "rule wasn't created")
	for _, r := range rules {
		if r.ID == rule.ID {
			require.Equal(t, rule, r, "rules do not match")
			break
		}
	}
}

func TestClient_UpdateAlertRule(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	rules, _, err := client.GetAlertRules(org, project)
	require.NoError(t, err, "unable to get alert rules")

	require.Greater(t, len(rules), 0, "no alert rules defined")

	rule := rules[0]
	require.Equal(t, 1, len(rule.Conditions), "unexpected condition count for rule")

	rule.Conditions = append(
		rule.Conditions,
		AlertRuleCondition{ID: "sentry.rules.conditions.regression_event.RegressionEventCondition"},
	)

	rule, err = client.UpdateAlertRule(org, project, rule)
	require.NoError(t, err, "unable to get update rule")
	require.Equal(t, 2, len(rule.Conditions), "conditions weren't updated")
}

func TestClient_DeleteAlertRule(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	rules, _, err := client.GetAlertRules(org, project)
	require.NoError(t, err, "unable to get alert rules")

	require.Greater(t, len(rules), 0, "no alert rules defined")

	err = client.DeleteAlertRule(org, project, rules[0])
	require.NoError(t, err, "unable to get delete rule")

	rules, _, err = client.GetAlertRules(org, project)
	require.NoError(t, err, "unable to get alert rules")
	require.Equal(t, 0, len(rules), "rule wasn't deleted")

}
