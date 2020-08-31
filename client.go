// Package sentry provides a client to access https://sentry.io/api and sentry
// instances apis.
package sentry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultHost is the default host that is used
	DefaultHost = "https://sentry.io"
	// DefaultEndpoint is the entry point for the api
	DefaultEndpoint = "/api/0/"
	// DefaultTimeout is the default timeout and is set to 60 seconds
	DefaultTimeout = time.Duration(60) * time.Second
)

// Client is used to talk to a sentry endpoint.
// Needs a auth token.
// If no endpoint this defaults to https://sentry.io/api/0/
type Client struct {
	AuthToken  string
	Endpoint   string
	HTTPClient *http.Client
}

// NewClient takes a auth token a optional endpoint and optional timeout and
// will return back a client and error
func NewClient(authtoken string, endpoint *string, timeout *int) (*Client, error) {
	var (
		clientEndpoint string
		clientTimeout  time.Duration
	)

	if endpoint == nil {
		clientEndpoint = fmt.Sprintf("%s%s", DefaultHost, DefaultEndpoint)
	} else {
		if *endpoint == "" {
			return nil, fmt.Errorf("Endpoint can not be a empty string")
		}
		clientEndpoint = *endpoint
	}

	if timeout == nil {
		clientTimeout = DefaultTimeout
	} else {
		clientTimeout = time.Duration(*timeout) * time.Second
	}

	return &Client{
		AuthToken: authtoken,
		Endpoint:  clientEndpoint,
		HTTPClient: &http.Client{
			Timeout: clientTimeout,
		},
	}, nil
}

func (c *Client) hasError(response *http.Response) error {

	if response.StatusCode > 299 || response.StatusCode < 200 {
		apierror := APIError{
			StatusCode: response.StatusCode,
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, &apierror); err != nil {
			apierror.Detail = string(body)
		}

		return error(apierror)
	}
	return nil
}

func (c *Client) decodeOrError(response *http.Response, out interface{}) error {

	if err := c.hasError(response); err != nil {
		return err
	}

	defer response.Body.Close()

	if out != nil {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(body, &out); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) encodeOrError(in interface{}) (io.Reader, error) {
	bytedata, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bytedata), nil
}

func (c *Client) newRequest(method, endpoint string, in interface{}) (*http.Request, error) {

	var bodyreader io.Reader

	if in != nil {
		newbodyreader, err := c.encodeOrError(in)
		if err != nil {
			return nil, err
		}
		bodyreader = newbodyreader
	}

	finalEndpoint := c.Endpoint+endpoint
	if !strings.HasSuffix(endpoint, "/")  {
		finalEndpoint = finalEndpoint + "/"
	}

	req, err := http.NewRequest(method, finalEndpoint, bodyreader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Add("Accept", "application/json")
	req.Close = true

	return req, nil
}

func (c *Client) rawRequest(method, endpoint string, in interface{}) (*http.Request, error) {
	var bodyreader io.Reader

	if in != nil {
		newbodyreader, err := c.encodeOrError(in)
		if err != nil {
			return nil, err
		}
		bodyreader = newbodyreader
	}

	req, err := http.NewRequest(method, c.Endpoint+endpoint, bodyreader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Close = true

	return req, nil

}

func (c *Client) doWithQuery(method string, endpoint string, out interface{}, in interface{}, query QueryArgs) error {
	request, err := c.newRequest(method, endpoint, in)
	if err != nil {
		return err
	}
	request.URL.RawQuery = query.ToQueryString()
	return c.send(request, out)
}
func (c *Client) doWithPaginationQuery(method, endpoint string, out, in interface{}, query QueryArgs) (*Link, error) {
	request, err := c.newRequest(method, endpoint, in)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = query.ToQueryString()
	return c.sendGetLink(request, out)
}

func (c *Client) do(method string, endpoint string, out interface{}, in interface{}) error {
	request, err := c.newRequest(method, endpoint, in)
	if err != nil {
		return err
	}
	return c.send(request, out)
}

func (c *Client) doWithPagination(method, endpoint string, out, in interface{}) (*Link, error) {
	request, err := c.newRequest(method, endpoint, in)
	if err != nil {
		return nil, err
	}
	return c.sendGetLink(request, out)
}

//rawWithPagination is used when we need to get a raw URL vs a url we combine and comb with newrequest
func (c *Client) rawWithPagination(method, endpoint string, out, in interface{}) (*Link, error) {
	request, err := c.rawRequest(method, endpoint, in)
	if err != nil {
		return nil, err
	}
	return c.sendGetLink(request, out)
}

func (c *Client) fetchLink(r *http.Response) *Link {
	link := &Link{}
	if r.Header.Get("Link") != "" {
		link = NewLink(r.Header.Get("Link"))
	}

	return link
}

func (c *Client) sendGetLink(req *http.Request, out interface{}) (*Link, error) {
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	link := c.fetchLink(response)
	decodeerr := c.decodeOrError(response, out)
	return link, decodeerr
}

func (c *Client) send(req *http.Request, out interface{}) error {
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	return c.decodeOrError(response, out)
}
