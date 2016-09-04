package template

import "github.com/wianvos/xlr/datamodels/ci"

type Tasks []Task

//Task  the datamodel for tasks
type Task struct {
	ci.Ci
	Attachments               []interface{} `json:"attachments,omitempty"`
	Comments                  []Comment     `json:"comments,omitempty"`
	Container                 string        `json:"container,omitempty"`
	Tasks                     Tasks         `json:"tasks,omitempty"`
	Description               string        `json:"description,omitempty"`
	FailuresCount             int           `json:"failuresCount,omitempty"`
	FlagStatus                string        `json:"flagStatus,omitempty"`
	HasBeenDelayed            bool          `json:"hasBeenDelayed,omitempty"`
	HasBeenFlagged            bool          `json:"hasBeenFlagged,omitempty"`
	OverdueNotified           bool          `json:"overdueNotified,omitempty"`
	Owner                     string        `json:"owner,omitempty"`
	Status                    string        `json:"status,omitempty"`
	WaitForScheduledStartDate bool          `json:"waitForScheduledStartDate,omitempty"`
}

func (t Task) GetCiInfo() ci.Ci {
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

func (ts *Tasks) PruneXLRInternalSettings() {

	var newTasks Tasks

	rts := *ts

	for _, t := range rts {
		t.Ci.PruneXLRInternalSettings()
		newTasks = append(newTasks, t)
	}

	*ts = newTasks
}

func (t *Task) PruneXLRInternalSettings() {
	rt := *t
	rt.Ci.PruneXLRInternalSettings()
	rt.Tasks.PruneXLRInternalSettings()
	*t = rt
}
