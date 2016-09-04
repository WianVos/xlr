package template

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/wianvos/xlr/datamodels/ci"
)

func TestPruneXLRInternalSettings(t *testing.T) {
	input := getLongTemplate()
	expected := getLongTemplatePruned()

	input.PruneXLRInternalSettings()
	if !reflect.DeepEqual(input, expected) {
		//t.Errorf("template.PruneXLRInternalSettings returned %+#v, expected %+#v", string(ppInput), string(ppExpected))
		jInput, _ := json.MarshalIndent(input, " ", "")
		jExpected, _ := json.MarshalIndent(expected, " ", "")

		t.Errorf("template.PruneXLRInternalSettings returned %s, expected %s", string(jInput), string(jExpected))
	}

}

// this mockRandomizer renders a very specific result :-)
func mockRandomizer(min int, max int) int {
	ret := max - min

	return ret
}

func TestPruneForUpload(t *testing.T) {
	input := getLongTemplate()
	expected := getLongTemplateReset()

	// passing in a mock function here
	input.PruneForUpload(mockRandomizer)

	if !reflect.DeepEqual(input, expected) {
		fmt.Printf("%#v", pretty.Compare(input, expected))
		t.Errorf("template.PruneXLRInternalSettings returned %s, expected %s", prettySimple(input), prettySimple(expected))
	}

}

func prettySimple(i interface{}) string {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

//helper methods

func getFullTemplate() Template {
	return Template{
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

func getPrunedTemplate() Template {
	return Template{
		Ci: ci.Ci{ID: "Applications/Release6999264",
			Type:           "xlrelease.Release",
			Token:          "",
			CreatedBy:      "",
			CreatedAt:      "",
			LastModifiedBy: "",
			LastModifiedAt: "",
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

func getLongTemplate() Template {
	var t Template
	j := `{
    "id": "Applications/Release6290848",
    "type": "xlrelease.Release",
    "$token": "612b03af-c8c2-4f37-8b27-8de45f6132b1",
    "$createdBy": "admin",
    "$createdAt": "2016-08-12T16:31:09.118+0000",
    "$lastModifiedBy": "admin",
    "$lastModifiedAt": "2016-08-13T03:46:39.264+0000",
    "title": "test",
    "scheduledStartDate": "2016-08-12T13:00:00Z",
    "flagStatus": "OK",
    "maxConcurrentReleases": 100,
    "releaseTriggers": [
        {
            "id": "Applications/Release6290848/Trigger4365752",
            "type": "rel.jsonTrigger",
            "$token": "8c137799-5224-4d28-9553-23a176daf3ed",
            "$createdBy": "admin",
            "$createdAt": "2016-08-12T16:32:11.852+0000",
            "$lastModifiedAt": "2016-08-13T03:27:36.748+0000",
            "triggerState": "test_build_11",
            "initialFire": false,
            "title": "test trigger",
            "releaseTitle": "test ${returnValue}",
            "pollType": "REPEAT",
            "periodicity": "30",
            "enabled": true,
            "executionId": "d6eef2cf-bf02-459a-8e70-59aed9428e0c",
            "variables": [
                {
                    "id": "Applications/Release6290848/Trigger4365752/Variable9853959",
                    "type": "xlrelease.StringVariable",
                    "$token": "b47da1da-6160-43b5-9bd6-85cf6cd9119a",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-12T16:32:11.858+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-12T16:32:11.858+0000",
                    "key": "testrelease",
                    "requiresValue": true,
                    "showOnReleaseStart": true,
                    "value": "${outputValue}"
                }
            ],
            "template": "Applications/Release6290848",
            "tags": [],
            "url": "http://192.168.99.100:8080/triggertest1",
            "jsonPath": "properties/last_good_build",
            "returnValue": "test_build_11"
        }
    ],
    "teams": [
        {
            "id": "Applications/Release6290848/Team2032544",
            "type": "xlrelease.Team",
            "$token": "0532cfd1-ec99-45d4-ac71-3bafed8c03e2",
            "$createdBy": "admin",
            "$createdAt": "2016-08-12T16:31:09.190+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-12T16:31:09.190+0000",
            "teamName": "Template Owner",
            "members": [
                "admin"
            ],
            "roles": [],
            "permissions": [
                "template#create_release",
                "template#view",
                "template#edit",
                "template#edit_security"
            ]
        },
        {
            "id": "Applications/Release6290848/Team7909521",
            "type": "xlrelease.Team",
            "$token": "4b4e12f3-bc54-4b9b-88df-7b1525a86bf2",
            "$createdBy": "admin",
            "$createdAt": "2016-08-12T16:31:09.238+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-12T16:31:09.238+0000",
            "teamName": "Release Admin",
            "members": [],
            "roles": [],
            "permissions": [
                "template#view",
                "release#view",
                "release#edit",
                "release#edit_security",
                "release#start",
                "release#abort",
                "release#edit_task",
                "release#reassign_task"
            ]
        }
    ],
    "memberViewers": [
        "admin"
    ],
    "roleViewers": [],
    "attachments": [],
    "phases": [
        {
            "id": "Applications/Release6290848/Phase4047557",
            "type": "xlrelease.Phase",
            "$token": "f091d3e2-c538-421b-b2d5-75a1910a0f8a",
            "$createdBy": "admin",
            "$createdAt": "2016-08-12T16:31:09.165+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-12T16:31:09.165+0000",
            "title": "New Phase",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase4047557/Task7204988",
                    "type": "xlrelease.NotificationTask",
                    "$token": "aa2a7225-ae15-4f6c-8bfc-e34f11fd65b3",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-12T16:31:25.335+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-12T16:31:25.335+0000",
                    "title": "test ${testrelease}",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase4047557",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "addresses": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED"
        },
        {
            "id": "Applications/Release6290848/Phase9310923",
            "type": "xlrelease.Phase",
            "$token": "8224a58a-e1f3-4b5d-bc7d-1179fec1c537",
            "$createdBy": "admin",
            "$createdAt": "2016-08-13T03:46:08.006+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-13T03:46:15.430+0000",
            "title": "tseting123",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase9310923/Task9228130",
                    "type": "xlrelease.GateTask",
                    "$token": "b98ec529-6aa5-4c4e-90a4-7a2798c25c84",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-13T03:46:22.650+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-13T03:46:22.650+0000",
                    "title": "test",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase9310923",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "conditions": [],
                    "dependencies": []
                },
                {
                    "id": "Applications/Release6290848/Phase9310923/Task9449663",
                    "type": "xlrelease.UserInputTask",
                    "$token": "987b373a-6ead-4bf7-bf4d-fb6367b44669",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-13T03:46:28.737+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-13T03:46:28.737+0000",
                    "title": "wait",
                    "description": "Please enter the required information below.",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase9310923",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "variables": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED",
            "color": "#009CDB"
        },
        {
            "id": "Applications/Release6290848/Phase7970307",
            "type": "xlrelease.Phase",
            "$token": "ec335870-801d-4a5c-925c-7e4b121ea170",
            "$createdBy": "admin",
            "$createdAt": "2016-08-13T03:46:08.006+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-13T03:46:32.947+0000",
            "title": "tseting123",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase7970307/Task9228130",
                    "type": "xlrelease.GateTask",
                    "$token": "3fa64cb8-ffdc-42e4-af31-e40f0022dff3",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-13T03:46:22.650+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-13T03:46:32.947+0000",
                    "title": "test",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase7970307",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "conditions": [],
                    "dependencies": []
                },
                {
                    "id": "Applications/Release6290848/Phase7970307/Task9449663",
                    "type": "xlrelease.UserInputTask",
                    "$token": "7ca206ac-7c09-4631-a6a2-a13c8a281239",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-13T03:46:28.737+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-13T03:46:32.948+0000",
                    "title": "wait",
                    "description": "Please enter the required information below.",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase7970307",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "variables": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED",
            "color": "#009CDB"
        },
        {
            "id": "Applications/Release6290848/Phase1529966",
            "type": "xlrelease.Phase",
            "$token": "465761c5-39e4-4e0a-a7df-208bf816be26",
            "$createdBy": "admin",
            "$createdAt": "2016-08-12T16:31:09.165+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-13T03:46:35.222+0000",
            "title": "New Phase",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase1529966/Task7204988",
                    "type": "xlrelease.NotificationTask",
                    "$token": "35345a11-43d9-4e4a-9eb5-8db086de2bd8",
                    "$createdBy": "admin",
                    "$createdAt": "2016-08-12T16:31:25.335+0000",
                    "$lastModifiedBy": "admin",
                    "$lastModifiedAt": "2016-08-13T03:46:35.222+0000",
                    "title": "test ${testrelease}",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase1529966",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "addresses": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED"
        }
    ],
    "queryableStartDate": "2016-08-12T13:00:00Z",
    "realFlagStatus": "OK",
    "status": "TEMPLATE",
    "tags": [],
    "variables": [
        {
            "id": "Applications/Release6290848/Variable9853959",
            "type": "xlrelease.StringVariable",
            "$token": "68a06315-7b3e-4965-9f31-17377022a78b",
            "$createdBy": "admin",
            "$createdAt": "2016-08-12T16:31:25.373+0000",
            "$lastModifiedBy": "admin",
            "$lastModifiedAt": "2016-08-12T16:31:25.373+0000",
            "key": "testrelease",
            "requiresValue": true,
            "showOnReleaseStart": true
        }
    ],
    "calendarPublished": false,
    "tutorial": false,
    "abortOnFailure": false,
    "allowConcurrentReleasesFromTrigger": true,
    "runningTriggeredReleasesCount": 109,
    "createdFromTrigger": false,
    "extensions": []
}`
	json.Unmarshal([]byte(j), &t)
	return t

}
func getLongTemplatePruned() Template {
	var t Template
	j := `{
    "id": "Applications/Release6290848",
    "type": "xlrelease.Release",
"$token": "",
    "$createdBy": "",
    "$createdAt": "",
    "$lastModifiedBy": "",
    "$lastModifiedAt": "",
    "title": "test",
    "scheduledStartDate": "2016-08-12T13:00:00Z",
    "flagStatus": "OK",
    "maxConcurrentReleases": 100,
    "releaseTriggers": [
        {
            "id": "Applications/Release6290848/Trigger4365752",
            "type": "rel.jsonTrigger",
            "$token": "",
            "$createdBy": "",
"$createdAt": "",
"$lastModifiedAt": "",
            "triggerState": "test_build_11",
            "initialFire": false,
            "title": "test trigger",
            "releaseTitle": "test ${returnValue}",
            "pollType": "REPEAT",
            "periodicity": "30",
            "enabled": true,
            "executionId": "d6eef2cf-bf02-459a-8e70-59aed9428e0c",
            "variables": [
                {
                    "id": "Applications/Release6290848/Trigger4365752/Variable9853959",
                    "type": "xlrelease.StringVariable",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "key": "testrelease",
                    "requiresValue": true,
                    "showOnReleaseStart": true,
                    "value": "${outputValue}"
                }
            ],
            "template": "Applications/Release6290848",
            "tags": [],
            "url": "http://192.168.99.100:8080/triggertest1",
            "jsonPath": "properties/last_good_build",
            "returnValue": "test_build_11"
        }
    ],
    "teams": [
        {
            "id": "Applications/Release6290848/Team2032544",
            "type": "xlrelease.Team",
            "$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "teamName": "Template Owner",
            "members": [
                "admin"
            ],
            "roles": [],
            "permissions": [
                "template#create_release",
                "template#view",
                "template#edit",
                "template#edit_security"
            ]
        },
        {
            "id": "Applications/Release6290848/Team7909521",
            "type": "xlrelease.Team",
"$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "teamName": "Release Admin",
            "members": [],
            "roles": [],
            "permissions": [
                "template#view",
                "release#view",
                "release#edit",
                "release#edit_security",
                "release#start",
                "release#abort",
                "release#edit_task",
                "release#reassign_task"
            ]
        }
    ],
    "memberViewers": [
        "admin"
    ],
    "roleViewers": [],
    "attachments": [],
    "phases": [
        {
            "id": "Applications/Release6290848/Phase4047557",
            "type": "xlrelease.Phase",
"$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "title": "New Phase",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase4047557/Task7204988",
                    "type": "xlrelease.NotificationTask",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "title": "test ${testrelease}",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase4047557",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "addresses": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED"
        },
        {
            "id": "Applications/Release6290848/Phase9310923",
            "type": "xlrelease.Phase",
"$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "title": "tseting123",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase9310923/Task9228130",
                    "type": "xlrelease.GateTask",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "title": "test",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase9310923",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "conditions": [],
                    "dependencies": []
                },
                {
                    "id": "Applications/Release6290848/Phase9310923/Task9449663",
                    "type": "xlrelease.UserInputTask",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "title": "wait",
                    "description": "Please enter the required information below.",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase9310923",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "variables": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED",
            "color": "#009CDB"
        },
        {
            "id": "Applications/Release6290848/Phase7970307",
            "type": "xlrelease.Phase",
"$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "title": "tseting123",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase7970307/Task9228130",
                    "type": "xlrelease.GateTask",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "title": "test",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase7970307",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "conditions": [],
                    "dependencies": []
                },
                {
                    "id": "Applications/Release6290848/Phase7970307/Task9449663",
                    "type": "xlrelease.UserInputTask",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "title": "wait",
                    "description": "Please enter the required information below.",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase7970307",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "variables": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED",
            "color": "#009CDB"
        },
        {
            "id": "Applications/Release6290848/Phase1529966",
            "type": "xlrelease.Phase",
"$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "title": "New Phase",
            "flagStatus": "OK",
            "tasks": [
                {
                    "id": "Applications/Release6290848/Phase1529966/Task7204988",
                    "type": "xlrelease.NotificationTask",
"$token": "",
                    "$createdBy": "",
"$createdAt": "",
                    "$lastModifiedBy": "",
"$lastModifiedAt": "",
                    "title": "test ${testrelease}",
                    "flagStatus": "OK",
                    "comments": [],
                    "container": "Applications/Release6290848/Phase1529966",
                    "attachments": [],
                    "status": "PLANNED",
                    "overdueNotified": false,
                    "waitForScheduledStartDate": true,
                    "hasBeenFlagged": false,
                    "hasBeenDelayed": false,
                    "failuresCount": 0,
                    "variableMapping": {},
                    "addresses": []
                }
            ],
            "release": "Applications/Release6290848",
            "status": "PLANNED"
        }
    ],
    "queryableStartDate": "2016-08-12T13:00:00Z",
    "realFlagStatus": "OK",
    "status": "TEMPLATE",
    "tags": [],
    "variables": [
        {
            "id": "Applications/Release6290848/Variable9853959",
            "type": "xlrelease.StringVariable",
"$token": "",
            "$createdBy": "",
"$createdAt": "",
            "$lastModifiedBy": "",
"$lastModifiedAt": "",
            "key": "testrelease",
            "requiresValue": true,
            "showOnReleaseStart": true
        }
    ],
    "calendarPublished": false,
    "tutorial": false,
    "abortOnFailure": false,
    "allowConcurrentReleasesFromTrigger": true,
    "runningTriggeredReleasesCount": 109,
    "createdFromTrigger": false,
    "extensions": []
}`
	json.Unmarshal([]byte(j), &t)
	return t

}

func getLongTemplateReset() Template {
	var tempTemplate Template
	out := `{
  "id": "Applications/test899999",
  "title": "test",
  "type": "xlrelease.Release",
  "allowConcurrentReleasesFromTrigger": true,
  "flagStatus": "OK",
  "maxConcurrentReleases": 100,
  "memberViewers": [
    "admin"
  ],
  "phases": [
    {
      "id": "Applications/test899999/Phase900000",
      "title": "New Phase",
      "type": "xlrelease.Phase",
      "flagStatus": "OK",
      "release": "Applications/Release6290848",
      "status": "PLANNED",
      "tasks": [
        {
          "id": "Applications/test899999/Phase900000/Task900001",
          "title": "test ${testrelease}",
          "type": "xlrelease.NotificationTask",
          "container": "Applications/test899999/Phase900000",
          "flagStatus": "OK",
          "status": "PLANNED",
          "waitForScheduledStartDate": true
        }
      ]
    },
    {
      "id": "Applications/test899999/Phase900002",
      "title": "tseting123",
      "type": "xlrelease.Phase",
      "flagStatus": "OK",
      "release": "Applications/Release6290848",
      "status": "PLANNED",
      "tasks": [
        {
          "id": "Applications/test899999/Phase900002/Task900003",
          "title": "test",
          "type": "xlrelease.GateTask",
          "container": "Applications/test899999/Phase900002",
          "flagStatus": "OK",
          "status": "PLANNED",
          "waitForScheduledStartDate": true
        },
        {
          "id": "Applications/test899999/Phase900002/Task900004",
          "title": "wait",
          "type": "xlrelease.UserInputTask",
          "container": "Applications/test899999/Phase900002",
          "description": "Please enter the required information below.",
          "flagStatus": "OK",
          "status": "PLANNED",
          "waitForScheduledStartDate": true
        }
      ]
    },
    {
      "id": "Applications/test899999/Phase900005",
      "title": "tseting123",
      "type": "xlrelease.Phase",
      "flagStatus": "OK",
      "release": "Applications/Release6290848",
      "status": "PLANNED",
      "tasks": [
        {
          "id": "Applications/test899999/Phase900005/Task900006",
          "title": "test",
          "type": "xlrelease.GateTask",
          "container": "Applications/test899999/Phase900005",
          "flagStatus": "OK",
          "status": "PLANNED",
          "waitForScheduledStartDate": true
        },
        {
          "id": "Applications/test899999/Phase900005/Task900007",
          "title": "wait",
          "type": "xlrelease.UserInputTask",
          "container": "Applications/test899999/Phase900005",
          "description": "Please enter the required information below.",
          "flagStatus": "OK",
          "status": "PLANNED",
          "waitForScheduledStartDate": true
        }
      ]
    },
    {
      "id": "Applications/test899999/Phase900008",
      "title": "New Phase",
      "type": "xlrelease.Phase",
      "flagStatus": "OK",
      "release": "Applications/Release6290848",
      "status": "PLANNED",
      "tasks": [
        {
          "id": "Applications/test899999/Phase900008/Task900009",
          "title": "test ${testrelease}",
          "type": "xlrelease.NotificationTask",
          "container": "Applications/test899999/Phase900008",
          "flagStatus": "OK",
          "status": "PLANNED",
          "waitForScheduledStartDate": true
        }
      ]
    }
  ],
  "queryableStartDate": "2016-08-12T13:00:00Z",
  "realFlagStatus": "OK",
  "releaseTriggers": [
    {
      "id": "Applications/Release6290848/Trigger4365752",
      "$createdAt": "2016-08-12T16:32:11.852+0000",
      "$createdBy": "admin",
      "$lastModifiedAt": "2016-08-13T03:27:36.748+0000",
      "$token": "8c137799-5224-4d28-9553-23a176daf3ed",
      "type": "rel.jsonTrigger",
      "triggerState": "test_build_11",
      "initialFire": false,
      "title": "test trigger",
      "releaseTitle": "test ${returnValue}",
      "pollType": "REPEAT",
      "periodicity": "30",
      "enabled": true,
      "executionId": "d6eef2cf-bf02-459a-8e70-59aed9428e0c",
      "variables": [
        null
      ],
      "template": "Applications/Release6290848",
      "tags": [],
      "url": "http://192.168.99.100:8080/triggertest1",
      "jsonPath": "properties/last_good_build"
    }
  ],
  "scheduledStartDate": "2016-08-12T13:00:00Z",
  "status": "TEMPLATE",
  "teams": [
    {
      "id": "Applications/test899999/Team900010",
      "type": "xlrelease.Team",
      "members": [
        "admin"
      ],
      "permissions": [
        "template#create_release",
        "template#view",
        "template#edit",
        "template#edit_security"
      ],
      "roles": [],
      "teamName": "Template Owner"
    },
    {
      "id": "Applications/test899999/Team900011",
      "type": "xlrelease.Team",
      "permissions": [
        "template#view",
        "release#view",
        "release#edit",
        "release#edit_security",
        "release#start",
        "release#abort",
        "release#edit_task",
        "release#reassign_task"
      ],
      "roles": [],
      "teamName": "Release Admin"
    }
  ],
  "variables": [
    {
      "id": "Applications/",
      "type": "xlrelease.StringVariable",
      "key": "testrelease",
      "requiresValue": true,
      "showOnReleaseStart": true
    }
  ],
  "runningTriggeredReleasesCount": 109,
	"createdFromTrigger": false,
	"extensions": []
}`
	json.Unmarshal([]byte(out), &tempTemplate)

	return tempTemplate
}
