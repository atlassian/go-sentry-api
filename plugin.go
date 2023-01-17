package sentry

import (
	"fmt"
	"net/http"
)

func (c *Client) touchPlugin(o Organization, p Project, pluginID string, enabled bool) error {
	method := http.MethodDelete
	if enabled {
		method = http.MethodPost
	}

	return c.do(method, fmt.Sprintf("projects/%s/%s/plugins/%s/", *o.Slug, *p.Slug, pluginID), nil, nil)
}

func (c *Client) EnablePlugin(o Organization, p Project, pluginID string) error {
	return c.touchPlugin(o, p, pluginID, true)
}

func (c *Client) DisablePlugin(o Organization, p Project, pluginID string) error {
	return c.touchPlugin(o, p, pluginID, false)
}

func (c *Client) GetPlugin(o Organization, p Project, pluginID string) (plugin Plugin, err error) {
	err = c.do(http.MethodGet, fmt.Sprintf("projects/%s/%s/plugins/%s/", *o.Slug, *p.Slug, pluginID), &plugin, nil)

	return
}

func (c *Client) SetPluginConfig(o Organization, p Project, pluginID string, config map[string]interface{}) (plugin Plugin, err error) {
	err = c.do(http.MethodPut, fmt.Sprintf("projects/%s/%s/plugins/%s/", *o.Slug, *p.Slug, pluginID), &plugin, config)

	return
}
