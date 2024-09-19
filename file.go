package sentry

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// File is used to create a new file for a release
type File struct {
	SHA1        string            `json:"sha1,omitempty"`
	Name        string            `json:"name,omitempty"`
	DateCreated time.Time         `json:"dateCreated,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	ID          string            `json:"id,omitempty"`
	Size        int               `json:"size,omitempty"`
}

// UploadReleaseFile will upload a file to release
func (c *Client) UploadReleaseFile(o Organization, p Project, r Release,
	name string, buffer io.Reader, header string) (File, error) {
	var file File

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("header", header)
	writer.WriteField("name", name)

	filewriter, createerr := writer.CreateFormFile("file", name)
	if createerr != nil {
		return file, createerr
	}

	data, readerr := ioutil.ReadAll(buffer)
	if readerr != nil {
		return file, readerr
	}

	filewriter.Write(data)
	closeerr := writer.Close()
	if closeerr != nil {
		return file, closeerr
	}

	endpoint := fmt.Sprintf("%sprojects/%s/%s/releases/%s/files/", c.Endpoint, *o.Slug, *p.Slug, r.Version)

	req, err := http.NewRequest("POST", endpoint, body)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Close = true

	if err != nil {
		return file, err
	}

	decoderr := c.send(req, &file)
	return file, decoderr
}

// DeleteReleaseFile will remove the file from a sentry release
func (c *Client) DeleteReleaseFile(o Organization, p Project, r Release, f File) error {
	return c.do("DELETE", fmt.Sprintf("projects/%s/%s/releases/%s/files/%s", *o.Slug, *p.Slug, r.Version, f.ID),
		nil, nil)
}

// UpdateReleaseFile will update just the name of the release file
func (c *Client) UpdateReleaseFile(o Organization, p Project, r Release, f File) error {
	return c.do("PUT", fmt.Sprintf("projects/%s/%s/releases/%s/files/%s", *o.Slug, *p.Slug, r.Version, f.ID), &f, &f)
}

// GetReleaseFiles will fetch all files in a release
func (c *Client) GetReleaseFiles(o Organization, p Project, r Release) ([]File, error) {
	var files []File
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/releases/%s/files", *o.Slug, *p.Slug, r.Version), &files, nil)
	return files, err
}

// GetReleaseFile will get the release file
func (c *Client) GetReleaseFile(o Organization, p Project, r Release, id string) (File, error) {
	var file File
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/releases/%s/files/%s", *o.Slug, *p.Slug, r.Version, id), &file, nil)
	return file, err
}
