package template

import "github.com/wianvos/xlr/datamodels/ci"

//Variables is a collection of variables
type Variables []Variable

//Variable represents a variable structure
type Variable struct {
	ci.Ci
	Key                string `json:"key,omitempty"`
	RequiresValue      bool   `json:"requiresValue,omitempty"`
	ShowOnReleaseStart bool   `json:"showOnReleaseStart,omitempty"`
	Label              string `json:"label,omitempty"`
	Description        string `json:"description,omitempty"`
	Value              string `json:"value,omitempty"`
}

func (t Variable) GetCiInfo() ci.Ci {
	return ci.Ci{
		ID:             t.ID,
		CreatedAt:      t.CreatedAt,
		CreatedBy:      t.CreatedBy,
		LastModifiedAt: t.LastModifiedAt,
		LastModifiedBy: t.LastModifiedBy,
		Token:          t.Token,
		Title:          t.Title,
		Type:           t.Type,
	}
}

func (ts *Variables) PruneXLRInternalSettings() {

	var newVariables Variables

	rts := *ts

	for _, t := range rts {
		t.Ci.PruneXLRInternalSettings()
		newVariables = append(newVariables, t)
	}

	*ts = newVariables
}

func (t *Variable) PruneXLRInternalSettings() {
	rt := *t
	rt.Ci.PruneXLRInternalSettings()
	*t = rt
}
