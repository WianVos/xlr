package template

import "github.com/wianvos/xlr/datamodels/ci"

type Phases []Phase

//Phase holds the datamodel for a Phase construct
type Phase struct {
	ci.Ci
	FlagStatus string `json:"flagStatus,omitempty"`
	Release    string `json:"release,omitempty"`
	Status     string `json:"status,omitempty"`
	Tasks      Tasks  `json:"tasks,omitempty"`
}

func (p Phase) GetCiInfo() ci.Ci {
	return ci.Ci{
		ID:             p.ID,
		CreatedAt:      p.CreatedAt,
		CreatedBy:      p.CreatedBy,
		LastModifiedAt: p.LastModifiedAt,
		LastModifiedBy: p.LastModifiedBy,
		Token:          p.Token,
		Title:          p.Title,
		Type:           p.Type,
	}
}

func (ps *Phases) PruneXLRInternalSettings() {

	var newPhases Phases

	rps := *ps

	for _, p := range rps {
		p.PruneXLRInternalSettings()
		newPhases = append(newPhases, p)
	}

	*ps = newPhases
}

func (p *Phase) PruneXLRInternalSettings() {
	rp := *p
	rp.Ci.PruneXLRInternalSettings()

	rp.Tasks.PruneXLRInternalSettings()

	*p = rp
}
