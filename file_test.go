package sentry

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReleaseFileResource(t *testing.T) {
	t.Parallel()

	client := newTestClient(t)

	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	newrelease := NewRelease{
		Version: "1.0.0",
	}
	release, err := client.CreateRelease(org, project, newrelease)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Create a new release file", func(t *testing.T) {

		var data = `
			Hello World!
		`

		filename := fmt.Sprintf("%s-example.txt", project.ID)

		file, err := client.UploadReleaseFile(org, project, release,
			filename,
			bytes.NewBuffer([]byte(data)),
			"Content-Type:text/plain; encoding=utf-8")

		if err != nil {
			t.Error("Failed to save file to sentry", err)
		}

		if file.Name != filename {
			t.Error("File did not save as example.txt")
		}
		t.Run("Fetch the release files for this release", func(t *testing.T) {
			files, err := client.GetReleaseFiles(org, project, release)
			if err != nil {
				t.Error(err)
			}
			if len(files) <= 0 {
				t.Error("Should have at least one file")
			}
		})
		t.Run("Get a previously created release file for this release", func(t *testing.T) {
			singlefile, err := client.GetReleaseFile(org, project, release, file.ID)
			if err != nil {
				t.Error(err)
			}

			if singlefile.ID != file.ID {
				t.Error("Should have gotten a file with the same id but didnt")
			}

		})
		t.Run("Update name of release file", func(t *testing.T) {

			file.Name = "Something/else/related"
			err := client.UpdateReleaseFile(org, project, release, file)
			if err != nil {
				t.Error(err)
			}
			if file.Name != "Something/else/related" {
				t.Error("File update did not change name")
			}

			t.Run("Delete file uploaded to release", func(t *testing.T) {
				err := client.DeleteReleaseFile(org, project, release, file)
				if err != nil {
					t.Error(err)
				}
			})
		})
	})
}
