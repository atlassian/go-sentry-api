package sentry

import (
	"testing"
)

const testEmail = "default@email.com"

func createMemberHelper(t *testing.T, org Organization) (Member, func() error) {
	client := newTestClient(t)

	member, err := client.CreateMember(org, testEmail)
	if err != nil {
		t.Fatal(err)
	}

	return member, func() error {
		return client.DeleteMember(org, member)
	}
}

func getMemberHelper(t *testing.T, org Organization, email string) Member {
	client := newTestClient(t)

	member, err := client.GetMemberByEmail(org, testEmail)
	if err != nil {
		t.Error(err)
	}

	return member
}

func TestMemberResource(t *testing.T) {
	t.Parallel()

	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanUpTeam := createTeamHelper(t)
	defer cleanUpTeam()

	member, cleanUpMember := createMemberHelper(t, org)
	defer cleanUpMember()

	t.Run("Fetch member by email returns error if non existing member", func(t *testing.T) {
		_, err := client.GetMemberByEmail(org, "notexisting@email.com")
		if err == nil {
			t.Error("Should have returned error for non existing member")
		}
	})

	t.Run("Add member to team", func(t *testing.T) {
		if err := client.AddExistingMemberToTeam(org, team, member); err != nil {
			t.Error(err)
		}
	})

	t.Run("Make member admin", func(t *testing.T) {
		if err := client.MakeAdmin(org, member); err != nil {
			t.Error(err)
		}

		member := getMemberHelper(t, org, testEmail)
		if member.Role != "admin" {
			t.Error("Member not made admin")
		}
	})
}
