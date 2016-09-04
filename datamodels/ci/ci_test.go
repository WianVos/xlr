package ci

import (
	"reflect"
	"testing"
)

func TestPruneXLRInternalSettings(t *testing.T) {
	input := getInputCis()
	expected := getExpectedCis()

	input.PruneXLRInternalSettings()

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("ci.PruneXLRInternalSettings returned %+v, expected %+v", input, expected)
	}

}

func getInputCis() Cis {
	return Cis{
		Ci{ID: "Applications/Release6999264",
			CreatedAt:      "2016-08-01T16:26:29.298+0000",
			CreatedBy:      "admin",
			LastModifiedAt: "2016-08-05T14:51:55.858+0000",
			LastModifiedBy: "admin",
			Token:          "8198a254-ce39-4075-9581-a65ec2ab72f1",
			Title:          "test_template",
			Type:           "xlrelease.Release"},
		Ci{
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

func getExpectedCis() Cis {
	return Cis{
		Ci{
			ID:             "Applications/Release6999264",
			CreatedAt:      "",
			CreatedBy:      "",
			LastModifiedAt: "",
			LastModifiedBy: "",
			Token:          "",
			Title:          "test_template",
			Type:           "xlrelease.Release"},
		Ci{
			ID:             "Applications/Release6999266",
			CreatedAt:      "",
			CreatedBy:      "",
			LastModifiedAt: "",
			LastModifiedBy: "",
			Token:          "",
			Title:          "test_template2",
			Type:           "xlrelease.Release"},
	}
}
