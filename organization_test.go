package sentry

import (
	"log"
	"os"
	"testing"
)

func TestOrganizationGet(t *testing.T) {

	authtoken := os.Getenv("SENTRY_AUTH_TOKEN")

	client := NewClient(authtoken, nil, nil)
	org, err := client.GetOrganization("atlassian-2y")
	if err != nil {
		t.Fatal(err)
	}
	if org.Name != "Atlassian" {
		t.Error("Name is not atlassian")
	}
	log.Printf("%v", org)
}

func TestOrganizationsList(t *testing.T) {
	authtoken := os.Getenv("SENTRY_AUTH_TOKEN")

	client := NewClient(authtoken, nil, nil)
	orgs, err := client.GetOrganizations()
	if err != nil {
		t.Fatal(err)
	}
	if len(orgs) <= 0 {
		t.Error("Didnt fetch any orgs")
	}
}

func TestOrganizationCreateUpdateDelete(t *testing.T) {
	authtoken := os.Getenv("SENTRY_AUTH_TOKEN")

	client := NewClient(authtoken, nil, nil)
	org, err := client.CreateOrganization("Test Org via Go Client")
	if err != nil {
		t.Fatal(err)
	}

	if org.Name != "Test Org via Go Client" {
		t.Error("New org is not the right slug: %v", org)
	}

	org.Name = "New Updated Name"

	if err := client.UpdateOrganization(org); err != nil {
		t.Fatal(err)
	}

	if err := client.DeleteOrganization(org); err != nil {
		t.Fatal(err)
	}
}
