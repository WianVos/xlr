package template

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/wianvos/xlr/datamodels/ci"
)

const (
	idBase = "Applications/"
)

type CiObject interface {
	GetCiInfo() ci.Ci
}

//Templates is a collection of templates
type Templates []Template

//Template is a structure to hold an xlr template
type Template struct {
	ci.Ci
	AbortOnFailure                     bool            `json:"abortOnFailure,omitempty"`
	AllowConcurrentReleasesFromTrigger bool            `json:"allowConcurrentReleasesFromTrigger,omitempty"`
	Attachments                        []Attachment    `json:"attachments,omitempty"`
	CalendarPublished                  bool            `json:"calendarPublished,omitempty"`
	CreatedFromTrigger                 bool            `json:"createdFromTrigger,omitempty"`
	Extensions                         []interface{}   `json:"extensions,omitempty"`
	FlagStatus                         string          `json:"flagStatus,omitempty"`
	MaxConcurrentReleases              int             `json:"maxConcurrentReleases,omitempty"`
	MemberViewers                      []string        `json:"memberViewers,omitempty"`
	OriginTemplateID                   string          `json:"originTemplateId,omitempty"`
	Phases                             Phases          `json:"phases,omitempty"`
	QueryableStartDate                 string          `json:"queryableStartDate,omitempty"`
	RealFlagStatus                     string          `json:"realFlagStatus,omitempty"`
	ReleaseTriggers                    ReleaseTriggers `json:"releaseTriggers,omitempty"`
	RoleViewers                        []interface{}   `json:"roleViewers,omitempty"`
	ScheduledStartDate                 string          `json:"scheduledStartDate,number,omitempty"`
	Status                             string          `json:"status,omitempty"`
	Tags                               []string        `json:"tags,omitempty"`
	Teams                              Teams           `json:"teams,omitempty"`
	Tutorial                           bool            `json:"tutorial,omitempty"`
	Variables                          Variables       `json:"variables,omitempty"`
	RunningTriggeredReleasesCount      int             `json:"runningTriggeredReleasesCount,omitempty"`
}

//GetCiInfo returns a struct holding the ci settings for a template
func (t Template) GetCiInfo() ci.Ci {
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

func (t *Template) PruneXLRInternalSettings() {

	t.Ci.PruneXLRInternalSettings()
	t.Phases.PruneXLRInternalSettings()
	t.Teams.PruneXLRInternalSettings()
	t.ReleaseTriggers.PruneXLRInternalSettings()
	t.Variables.PruneXLRInternalSettings()
}

//Comment represents a comment ci
type Comment struct {
	Author string `json:"author,omitempty"`
	Date   string `json:"date,omitempty"`
	Text   string `json:"text,omitempty"`
}

//

type Attachment struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Release     string `json:"release,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	FileURI     string `json:"fileUri,omitempty"`
}

//Variables a collection of variable definitions

//FilterReleases pulls all templates from the collection that aren't real templates
func (t *Templates) FilterOutReleases() {

	//GetReleases prunes all the non-released templates from the Templates collectio
	td := *t
	var nt Templates

	for _, tx := range td {
		if tx.IsRelease() == false {
			nt = append(nt, tx)
		}
	}

	*t = nt
}

//IsRelease returns true if the release is only a template
//returns false if it is a release
func (t Template) IsRelease() bool {

	if t.Status == "TEMPLATE" {
		return false
	}
	return true
}

//GetCisInfo from the templates
func (t Templates) GetCiInfo() []ci.Ci {
	var cis []ci.Ci
	for _, tl := range t {
		cis = append(cis, tl.GetCiInfo())
	}

	return cis
}

type randomizer func(int, int) int

//PruneForUpload prunes an inputed template and clears any xlr internal fieds that should not have a value
func (t *Template) PruneForUpload(fn randomizer) {
	randomNr := fn(100000, 999999)
	defaultStatus := "PLANNED"

	td := *t

	templateID := idBase + strings.Replace(td.Title, " ", "_", -1) + strconv.Itoa(randomNr)
	td.ID = templateID
	td.CreatedAt = ""
	td.CreatedBy = ""
	td.LastModifiedAt = ""
	td.LastModifiedBy = ""
	td.Token = ""

	for pr, p := range td.Phases {
		randomNr++
		phaseID := templateID + "/Phase" + strconv.Itoa(randomNr)
		td.Phases[pr].ID = phaseID
		td.Phases[pr].CreatedAt = ""
		td.Phases[pr].CreatedBy = ""
		td.Phases[pr].LastModifiedAt = ""
		td.Phases[pr].LastModifiedBy = ""
		td.Phases[pr].Token = ""
		td.Phases[pr].Status = defaultStatus

		for tr, ta := range p.Tasks {
			randomNr++
			taskID := phaseID + "/Task" + strconv.Itoa(randomNr)
			td.Phases[pr].Tasks[tr].ID = taskID
			td.Phases[pr].Tasks[tr].CreatedAt = ""
			td.Phases[pr].Tasks[tr].CreatedBy = ""
			td.Phases[pr].Tasks[tr].LastModifiedAt = ""
			td.Phases[pr].Tasks[tr].LastModifiedBy = ""
			td.Phases[pr].Tasks[tr].Token = ""
			td.Phases[pr].Tasks[tr].Status = defaultStatus
			td.Phases[pr].Tasks[tr].Container = phaseID

			for st := range ta.Tasks {
				randomNr++
				sTaskID := taskID + "/Task" + strconv.Itoa(randomNr)
				td.Phases[pr].Tasks[tr].Tasks[st].ID = sTaskID
				td.Phases[pr].Tasks[tr].Tasks[st].CreatedAt = ""
				td.Phases[pr].Tasks[tr].Tasks[st].CreatedBy = ""
				td.Phases[pr].Tasks[tr].Tasks[st].LastModifiedAt = ""
				td.Phases[pr].Tasks[tr].Tasks[st].LastModifiedBy = ""
				td.Phases[pr].Tasks[tr].Tasks[st].Token = ""
				td.Phases[pr].Tasks[tr].Tasks[st].Status = defaultStatus
				td.Phases[pr].Tasks[tr].Tasks[st].Container = taskID

			}
		}
	}

	for tr := range td.Teams {
		randomNr++
		teamID := templateID + "/Team" + strconv.Itoa(randomNr)
		td.Teams[tr].ID = teamID
		td.Teams[tr].CreatedAt = ""
		td.Teams[tr].CreatedBy = ""
		td.Teams[tr].LastModifiedAt = ""
		td.Teams[tr].LastModifiedBy = ""
		td.Teams[tr].Token = ""
	}
	for vr := range td.Variables {
		td.Variables[vr].ID = idBase
		td.Variables[vr].CreatedAt = ""
		td.Variables[vr].CreatedBy = ""
		td.Variables[vr].LastModifiedAt = ""
		td.Variables[vr].LastModifiedBy = ""
		td.Variables[vr].Token = ""
	}

	for trig := range td.ReleaseTriggers {
		triggerID := templateID + "/Trigger" + strconv.Itoa(randomNr)
		td.ReleaseTriggers[trig].ID = triggerID
		td.ReleaseTriggers[trig].CreatedAt = ""
		td.ReleaseTriggers[trig].CreatedBy = ""
		td.ReleaseTriggers[trig].LastModifiedAt = ""
		td.ReleaseTriggers[trig].Token = ""
		td.ReleaseTriggers[trig].TriggerState = ""
		td.ReleaseTriggers[trig].Template = templateID
		td.ReleaseTriggers[trig].ExecutionID = ""

	}
	*t = td

}

func RandomIDNr(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
