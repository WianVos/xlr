package xlr

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/wianvos/xlr/datamodels/ci"
	"github.com/wianvos/xlr/datamodels/template"
)

var mockTemplateListResponse = `[
    {
        "id": "Applications/Release6999264",
        "type": "xlrelease.Release",
        "$token": "8198a254-ce39-4075-9581-a65ec2ab72f1",
        "$createdBy": "admin",
        "$createdAt": "2016-08-01T16:26:29.298+0000",
        "$lastModifiedBy": "admin",
        "$lastModifiedAt": "2016-08-05T14:51:55.858+0000",
        "title": "test_template",
        "scheduledStartDate": "2016-08-01T09:00:00Z",
        "flagStatus": "OK",
        "maxConcurrentReleases": 100,
        "releaseTriggers": [],
        "teams": [],
        "roleViewers": [],
        "attachments": [],
        "phases": [],
        "queryableStartDate": "2016-08-01T09:00:00Z",
        "realFlagStatus": "OK",
        "status": "TEMPLATE",
        "tags": [],
        "variables": [],
        "calendarPublished": false,
        "tutorial": false,
        "abortOnFailure": false,
        "allowConcurrentReleasesFromTrigger": false,
        "runningTriggeredReleasesCount": 0,
        "createdFromTrigger": false,
        "extensions": []
    },{
		"id": "Applications/Release6999266",
		"type": "xlrelease.Release",
		"$token": "8198a254-ce39-4075-9581-a65ec2ab72f1",
		"$createdBy": "admin",
		"$createdAt": "2016-08-01T16:26:29.298+0000",
		"$lastModifiedBy": "admin",
		"$lastModifiedAt": "2016-08-05T14:51:55.858+0000",
		"title": "test_template2",
		"scheduledStartDate": "2016-08-01T09:00:00Z",
		"flagStatus": "OK",
		"maxConcurrentReleases": 100,
		"releaseTriggers": [],
		"teams": [],
		"roleViewers": [],
		"attachments": [],
		"phases": [],
		"queryableStartDate": "2016-08-01T09:00:00Z",
		"realFlagStatus": "OK",
		"status": "TEMPLATE",
		"tags": [],
		"variables": [],
		"calendarPublished": false,
		"tutorial": false,
		"abortOnFailure": false,
		"allowConcurrentReleasesFromTrigger": false,
		"runningTriggeredReleasesCount": 0,
		"createdFromTrigger": false,
		"extensions": []
	}
]`

func TestTemplatesList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mockTemplateListResponse)
	})

	acct, err := client.Templates.List()
	if err != nil {
		t.Errorf("Templates.List returned error: %v", err)
	}

	expected := getCis()

	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Template.List returned %+v, expected %+v", acct, expected)
	}
}

func TestTemplatesShow(t *testing.T) {
	setup()
	defer teardown()

	// setup the api route we're going to need for this test
	mux.HandleFunc("/api/v1/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mockTemplateListResponse)
	})

	cases := []struct {
		searchID         string
		byTitle          bool
		expectedTemplate template.Template
		expectedErr      error
	}{
		{

			searchID:         "Release6999264",
			byTitle:          false,
			expectedTemplate: getFullTemplate(),
			expectedErr:      nil,
		}, {
			searchID:         "Applications/Release6999264",
			byTitle:          false,
			expectedTemplate: getFullTemplate(),
			expectedErr:      nil,
		}, {
			searchID:         "test_template",
			byTitle:          true,
			expectedTemplate: getFullTemplate(),
			expectedErr:      nil,
		}, {
			searchID:         "bogus",
			byTitle:          false,
			expectedTemplate: template.Template{},
			expectedErr:      errors.New("unable to find template with id: Applications/bogus"),
		},
	}

	for _, c := range cases {
		acct, err := client.Templates.Show(c.searchID, c.byTitle)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}

		if c.expectedTemplate.ID != acct.ID {
			t.Errorf("Expected %v but got %v", c.expectedTemplate, acct)
		}
	}

}

func TestTemplateCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mockTemplateListResponse)
	})

	mux.HandleFunc("/api/v1/templates/import", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

	})

	cases := []struct {
		template         template.Template
		overWrite        bool
		expectedTemplate template.Template
		expectedErr      error
	}{
		{
			template:         getFullTemplate(),
			overWrite:        false,
			expectedTemplate: template.Template{},
			expectedErr:      errors.New("Template with the same title already exists:test_template"),
		}, {
			template:         getFullTemplate(),
			overWrite:        true,
			expectedTemplate: getFullTemplate(),
			expectedErr:      nil,
		},
	}

	for _, c := range cases {
		acct, err := client.Templates.CreateTemplate(c.template, c.overWrite)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}

		if c.expectedTemplate.ID != acct.ID {
			t.Errorf("Expected %v but got %v", c.expectedTemplate, acct)
		}
	}

}

// private helper functions ..
// usually to keep the code nice, clean and readable

func getFullTemplate() template.Template {
	return template.Template{
		Ci: ci.Ci{
			ID:             "Applications/Release6999264",
			Type:           "xlrelease.Release",
			Token:          "8198a254-ce39-4075-9581-a65ec2ab72f1",
			CreatedBy:      "admin",
			CreatedAt:      "2016-08-01T16:26:29.298+0000",
			LastModifiedBy: "admin",
			LastModifiedAt: "2016-08-05T14:51:55.858+0000",
			Title:          "test_template",
		},
		ScheduledStartDate:                 "2016-08-01T09:00:00Z",
		FlagStatus:                         "OK",
		MaxConcurrentReleases:              100,
		QueryableStartDate:                 "2016-08-01T09:00:00Z",
		RealFlagStatus:                     "OK",
		Status:                             "TEMPLATE",
		CalendarPublished:                  false,
		Tutorial:                           false,
		AbortOnFailure:                     false,
		AllowConcurrentReleasesFromTrigger: false,
		RunningTriggeredReleasesCount:      0,
		CreatedFromTrigger:                 false,
	}
}

func getCis() []ci.Ci {
	return []ci.Ci{
		ci.Ci{ID: "Applications/Release6999264",
			CreatedAt:      "2016-08-01T16:26:29.298+0000",
			CreatedBy:      "admin",
			LastModifiedAt: "2016-08-05T14:51:55.858+0000",
			LastModifiedBy: "admin",
			Token:          "8198a254-ce39-4075-9581-a65ec2ab72f1",
			Title:          "test_template",
			Type:           "xlrelease.Release"},
		ci.Ci{
			ID:             "Applications/Release6999266",
			CreatedAt:      "2016-08-01T16:26:29.298+0000",
			CreatedBy:      "admin",
			LastModifiedAt: "2016-08-05T14:51:55.858+0000",
			LastModifiedBy: "admin",
			Token:          "8198a254-ce39-4075-9581-a65ec2ab72f1",
			Title:          "test_template2",
			Type:           "xlrelease.Release"},
	}
}
