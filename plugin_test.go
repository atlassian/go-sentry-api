package sentry

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const testPlugin = "webhooks"

func pluginEnabled(t *testing.T, client *Client, org Organization, project Project, pluginID string) bool {
	project, err := client.GetProject(org, *project.Slug)
	require.NoError(t, err, "failed to get project details")

	require.NotNil(t, project.Plugins, "project details missing plugins section")
	require.Greater(t, len(*project.Plugins), 0, "no plugins installed")

	for _, p := range *project.Plugins {
		if p.ID == pluginID {
			return p.Enabled
		}
	}
	require.Fail(t, "plugin not found")

	return false
}

func TestClient_EnablePlugin(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	require.False(t, pluginEnabled(t, client, org, project, testPlugin), "plugin already enabled")

	err = client.EnablePlugin(org, project, testPlugin)
	require.NoError(t, err, "unable to enable plugin")

	require.True(t, pluginEnabled(t, client, org, project, testPlugin), "plugin wasn't enabled")

	err = client.EnablePlugin(org, project, testPlugin)
	require.NoError(t, err, "unable to enable plugin")

	require.True(t, pluginEnabled(t, client, org, project, testPlugin), "plugin wasn't enabled")

	err = client.DisablePlugin(org, project, testPlugin)
	require.NoError(t, err, "unable to disable plugin")

	require.False(t, pluginEnabled(t, client, org, project, testPlugin), "plugin wasn't disabled")
}

func TestClient_GetPlugin(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	plugin, err := client.GetPlugin(org, project, testPlugin)
	require.NoError(t, err, "unable to get plugin")

	require.Greater(t, len(plugin.Config), 0, "plugin config is missing")
}

func TestClient_SetPluginConfig(t *testing.T) {
	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	plugin, err := client.SetPluginConfig(org, project, testPlugin, map[string]interface{}{
		"urls": "https://test.com/",
	})
	require.NoError(t, err, "unable to get plugin")
	require.NotEmpty(t, plugin.Config, "unable to get plugin config")
	require.Equal(t, "https://test.com/", plugin.Config[0].Value)
}
