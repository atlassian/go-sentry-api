package sentry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	DefaultEndpoint = "https://sentry.io/api/0/"
	DefaultTimeout  = time.Duration(60) * time.Second
)

type Client struct {
	AuthToken  string
	Endpoint   string
	HttpClient *http.Client
}

func NewClient(authtoken string, endpoint *string, timeout *int) *Client {
	var (
		clientEndpoint string
		clientTimeout  time.Duration
	)

	if endpoint == nil {
		clientEndpoint = DefaultEndpoint
	} else {
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
		HttpClient: &http.Client{
			Timeout: clientTimeout,
		},
	}
}

func CreateURI(resource string) string {
	return ""
}

func (c *Client) do(method string, endpoint string, out interface{}, in interface{}) error {

	var (
		bodyreader io.Reader
		request    *http.Request
		response   *http.Response
		err        error
	)

	log.Printf("Sending %s request to endpoint %s%s", method, c.Endpoint, endpoint)

	if in != nil {
		bytedata, err := json.Marshal(in)
		if err != nil {
			return err
		}
		bodyreader = bytes.NewReader(bytedata)
	}

	request, err = http.NewRequest(method, c.Endpoint+endpoint+"/", bodyreader)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))

	response, err = c.HttpClient.Do(request)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode > 299 || response.StatusCode < 200 {
		apierror := SentryApiError{}

		if err := json.Unmarshal(body, &apierror); err != nil {
			return err
		}

		apierror.StatusCode = response.StatusCode
		return error(apierror)
	} else {
		if out != nil {
			if err := json.Unmarshal(body, &out); err != nil {
				return err
			}
		}
	}

	return nil
}
