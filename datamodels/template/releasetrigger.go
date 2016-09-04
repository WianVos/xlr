package template

import "github.com/wianvos/xlr/datamodels/ci"

type ReleaseTriggers []ReleaseTrigger

//Trigger is the data structure to store triggers
type ReleaseTrigger struct {
	ci.Ci
	TriggerState string        `json:"triggerState"`
	InitialFire  bool          `json:"initialFire"`
	Title        string        `json:"title"`
	ReleaseTitle string        `json:"releaseTitle"`
	PollType     string        `json:"pollType"`
	Periodicity  string        `json:"periodicity"`
	Enabled      bool          `json:"enabled"`
	ExecutionID  string        `json:"executionId"`
	Variables    []Variables   `json:"variables"`
	Template     string        `json:"template"`
	Tags         []interface{} `json:"tags"`
	URL          string        `json:"url"`
	JSONPath     string        `json:"jsonPath"`
}

func (t ReleaseTrigger) GetCiInfo() ci.Ci {
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

func (ts *ReleaseTriggers) PruneXLRInternalSettings() {

	var newTriggers ReleaseTriggers

	rts := *ts

	for _, t := range rts {
		t.Ci.PruneXLRInternalSettings()
		newTriggers = append(newTriggers, t)
	}

	*ts = newTriggers
}

func (t *ReleaseTrigger) PruneXLRInternalSettings() {
	rt := *t
	rt.Ci.PruneXLRInternalSettings()
	*t = rt
}
