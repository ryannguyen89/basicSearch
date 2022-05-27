package basicSearch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"basicSearch"
)

func TestIndexOrg(t *testing.T) {
	data := []*basicSearch.Organization{{
		ID:            1,
		URL:           "12",
		ExternalID:    "",
		Name:          "",
		DomainNames:   nil,
		CreatedAt:     "",
		Details:       "",
		SharedTickets: false,
		Tags:          nil,
	},
	}
	idx := basicSearch.IndexOrg(data)
	assert.Equal(t, 1, len(idx.Indexes["_id"]))
	assert.Equal(t, 1, len(idx.Indexes["url"]))
	t.Log(idx.Indexes["url"])
}

func TestIndexUser(t *testing.T) {
	data := []*basicSearch.User{
		{
			ID:             123,
			URL:            "http://initech.tokoin.io.com/api/v2/users/123.json",
			ExternalID:     "74341f74-9c79-49d5-9611-87ef9b6eb75f",
			Name:           "Francisca Rasmussen",
			Alias:          "Miss Coffey",
			CreatedAt:      "2016-04-15T05:19:46 -10:00",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-AU",
			Timezone:       "Sri Lanka",
			LastLoginAt:    "2013-08-04T01:03:27 -10:00",
			Email:          "coffeyrasmussen@flotonic.com",
			Phone:          "8335-422-718",
			Signature:      "Don't Worry Be Happy!",
			OrganizationID: 106,
			Tags:           []string{"Springville", "Sutton", "Hartsville/Hartley", "Diaperville"},
			Suspended:      true,
			Role:           "admin",
		},
		{
			ID:             2,
			URL:            "http://initech.tokoin.io.com/api/v2/users/2.json",
			ExternalID:     "c9995ea4-ff72-46e0-ab77-dfe0ae1ef6c2",
			Name:           "Cross Barlow",
			Alias:          "Miss Joni",
			CreatedAt:      "2016-06-23T10:31:39 -10:00",
			Active:         false,
			Verified:       false,
			Shared:         false,
			Locale:         "zh-CN",
			Timezone:       "Armenia",
			LastLoginAt:    "2012-04-12T04:03:28 -10:00",
			Email:          "jonibarlow@flotonic.com",
			Phone:          "9575-552-585",
			Signature:      "Don't Worry Be Happy!",
			OrganizationID: 119,
			Tags:           []string{"Foxworth", "Woodlands", "Herlong", "Henrietta"},
			Suspended:      false,
			Role:           "admin",
		},
	}
	idx := basicSearch.IndexUser(data)

	tests := map[string]struct {
		Field          string
		ExpectedLength int
		Values         []string
	}{
		"ID index": {
			Field:          "_id",
			ExpectedLength: 2,
			Values:         []string{"123", "2"},
		},
		"URL index": {
			Field:          "url",
			ExpectedLength: 2,
			Values: []string{"http://initech.tokoin.io.com/api/v2/users/2.json",
				"http://initech.tokoin.io.com/api/v2/users/123.json"},
		},
		"External ID index": {
			Field:          "external_id",
			ExpectedLength: 2,
			Values:         []string{"c9995ea4-ff72-46e0-ab77-dfe0ae1ef6c2", "74341f74-9c79-49d5-9611-87ef9b6eb75f"},
		},
		"Name index": {
			Field:          "name",
			ExpectedLength: 2,
			Values:         []string{"Cross Barlow", "Francisca Rasmussen"},
		},
		"Alias index": {
			Field:          "alias",
			ExpectedLength: 2,
			Values:         []string{"Miss Coffey", "Miss Joni"},
		},
		"Created index": {
			Field:          "created_at",
			ExpectedLength: 2,
			Values:         []string{"2016-04-15T05:19:46 -10:00", "2016-06-23T10:31:39 -10:00"},
		},
		"Active index": {
			Field:          "active",
			ExpectedLength: 2,
			Values:         []string{"true", "false"},
		},
		"verified index": {
			Field:          "verified",
			ExpectedLength: 2,
			Values:         []string{"true", "false"},
		},
		"shared index": {
			Field:          "shared",
			ExpectedLength: 2,
			Values:         []string{"true", "false"},
		},
		"locale index": {
			Field:          "locale",
			ExpectedLength: 2,
			Values:         []string{"en-AU", "zh-CN"},
		},
		"timezone index": {
			Field:          "timezone",
			ExpectedLength: 2,
			Values:         []string{"Sri Lanka", "Armenia"},
		},
		"last_login_at index": {
			Field:          "last_login_at",
			ExpectedLength: 2,
			Values:         []string{"2013-08-04T01:03:27 -10:00", "2012-04-12T04:03:28 -10:00"},
		},
		"email index": {
			Field:          "email",
			ExpectedLength: 2,
			Values:         []string{"jonibarlow@flotonic.com", "coffeyrasmussen@flotonic.com"},
		},
		"phone index": {
			Field:          "phone",
			ExpectedLength: 2,
			Values:         []string{"8335-422-718", "9575-552-585"},
		},
		"signature index": {
			Field:          "signature",
			ExpectedLength: 1,
			Values:         []string{"Don't Worry Be Happy!"},
		},
		"organization_id index": {
			Field:          "organization_id",
			ExpectedLength: 2,
			Values:         []string{"119", "106"},
		},
		"tags index": {
			Field:          "tags",
			ExpectedLength: 8,
			Values:         []string{"Diaperville", "Herlong"},
		},
		"suspended index": {
			Field:          "suspended",
			ExpectedLength: 2,
			Values:         []string{"true", "false"},
		},
		"role index": {
			Field:          "role",
			ExpectedLength: 1,
			Values:         []string{"admin"},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fieldIdx := idx.Indexes[test.Field]
			require.NotNil(t, fieldIdx)
			require.Equal(t, test.ExpectedLength, len(fieldIdx))
			t.Log(fieldIdx)
			for _, v := range test.Values {
				require.NotNil(t, fieldIdx[v])
			}
		})
	}
}

func TestIndexTicket(t *testing.T) {
	data := []*basicSearch.Ticket{
		{
			ID:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
			URL:            "http://initech.tokoin.io.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
			ExternalID:     "9210cdc9-4bee-485f-a078-35396cd74063",
			CreatedAt:      "2016-04-28T11:19:34 -10:00",
			Type:           "incident",
			Subject:        "A Catastrophe in Korea (North)",
			Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
			Priority:       "low",
			Status:         "pending",
			SubmitterID:    38,
			AssigneeID:     24,
			OrganizationID: 116,
			Tags: []string{
				"Puerto Rico",
				"Idaho",
				"Oklahoma",
				"Louisiana",
			},
			HasIncidents: true,
			DueAt:        "2016-07-31T02:37:50 -10:00",
			Via:          "web",
		},
		{
			ID:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
			URL:            "http://initech.tokoin.io.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
			ExternalID:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
			CreatedAt:      "2016-04-14T08:32:31 -10:00",
			Type:           "incident",
			Subject:        "A Catastrophe in Micronesia",
			Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
			Priority:       "high",
			Status:         "hold",
			SubmitterID:    71,
			AssigneeID:     38,
			OrganizationID: 112,
			Tags: []string{
				"Ohio",
				"Pennsylvania",
				"American Samoa",
				"Northern Mariana Islands",
			},
			HasIncidents: false,
			DueAt:        "2016-08-15T05:37:32 -10:00",
			Via:          "chat",
		},
	}

	idx := basicSearch.IndexTicket(data)

	tests := map[string]struct {
		Field          string
		ExpectedLength int
		Values         []string
	}{
		"_id index": {
			Field:          "_id",
			ExpectedLength: 2,
			Values:         []string{"436bf9b0-1147-4c0a-8439-6f79833bff5b", "1a227508-9f39-427c-8f57-1b72f3fab87c"},
		},
		"url index": {
			Field:          "url",
			ExpectedLength: 2,
			Values:         []string{"http://initech.tokoin.io.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json", "http://initech.tokoin.io.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json"},
		},
		"external_id index": {
			Field:          "external_id",
			ExpectedLength: 2,
			Values:         []string{"9210cdc9-4bee-485f-a078-35396cd74063", "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a"},
		},
		"created_at index": {
			Field:          "created_at",
			ExpectedLength: 2,
			Values:         []string{"2016-04-14T08:32:31 -10:00", "2016-04-28T11:19:34 -10:00"},
		},
		"type index": {
			Field:          "type",
			ExpectedLength: 1,
			Values:         []string{"incident"},
		},
		"subject index": {
			Field:          "subject",
			ExpectedLength: 2,
			Values:         []string{"A Catastrophe in Micronesia", "A Catastrophe in Korea (North)"},
		},
		"description index": {
			Field:          "description",
			ExpectedLength: 2,
			Values:         []string{"Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum."},
		},
		"priority index": {
			Field:          "priority",
			ExpectedLength: 2,
			Values:         []string{"high", "low"},
		},
		"status index": {
			Field:          "status",
			ExpectedLength: 2,
			Values:         []string{"pending", "hold"},
		},
		"submitter_id index": {
			Field:          "submitter_id",
			ExpectedLength: 2,
			Values:         []string{"38", "71"},
		},
		"assignee_id index": {
			Field:          "assignee_id",
			ExpectedLength: 2,
			Values:         []string{"38", "24"},
		},
		"organization_id index": {
			Field:          "organization_id",
			ExpectedLength: 2,
			Values:         []string{"112", "116"},
		},
		"tags index": {
			Field:          "tags",
			ExpectedLength: 8,
			Values:         []string{"Louisiana", "Pennsylvania"},
		},
		"has_incidents index": {
			Field:          "has_incidents",
			ExpectedLength: 2,
			Values:         []string{"true", "false"},
		},
		"due_at index": {
			Field:          "due_at",
			ExpectedLength: 2,
			Values:         []string{"2016-08-15T05:37:32 -10:00", "2016-07-31T02:37:50 -10:00"},
		},
		"via index": {
			Field:          "via",
			ExpectedLength: 2,
			Values:         []string{"web", "chat"},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fieldIdx := idx.Indexes[test.Field]
			require.NotNil(t, fieldIdx)
			require.Equal(t, test.ExpectedLength, len(fieldIdx))
			t.Log(fieldIdx)
			for _, v := range test.Values {
				require.NotNil(t, fieldIdx[v])
			}
		})
	}
}
