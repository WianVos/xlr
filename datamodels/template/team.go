package template

import "github.com/wianvos/xlr/datamodels/ci"

//Teams is a slice with teams
type Teams []Team

//Team holds datamodel for team descriptions
type Team struct {
	ci.Ci
	Members     []string      `json:"members,omitempty"`
	Permissions []string      `json:"permissions,omitempty"`
	Roles       []interface{} `json:"roles"`
	TeamName    string        `json:"teamName,omitempty"`
}

func (t Team) GetCiInfo() ci.Ci {
	return ci.Ci{
		ID:             t.ID,
		CreatedAt:      t.CreatedAt,
		CreatedBy:      t.CreatedBy,
		LastModifiedAt: t.LastModifiedAt,
		LastModifiedBy: t.LastModifiedBy,
		Token:          t.Token,
		Type:           t.Type,
	}
}

func (ts *Teams) PruneXLRInternalSettings() {

	var newTeams Teams

	rts := *ts

	for _, t := range rts {
		t.PruneXLRInternalSettings()
		newTeams = append(newTeams, t)
	}

	*ts = newTeams
}

func (t *Team) PruneXLRInternalSettings() {
	t.Ci.PruneXLRInternalSettings()
	rt := *t
	rt.Ci.PruneXLRInternalSettings()
	*t = rt
}
