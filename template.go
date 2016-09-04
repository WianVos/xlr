package xlr

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/wianvos/xlr/datamodels/ci"
	"github.com/wianvos/xlr/datamodels/template"
)

const (
	templateBasePath = "api/v1/templates"
	templateIDPrefix = "Applications/"
)

//TemplateService is interface
type TemplateService interface {
	List() ([]ci.Ci, error)
	Show(id string, title bool) (template.Template, error)
	CreateTemplate(t template.Template, o bool) (template.Template, error)

	getTemplateList() (template.Templates, error)
}

//TemplateServiceOp holds the communication service for templates
type TemplateServiceOp struct {
	client *Client
}

var _ TemplateService = &TemplateServiceOp{}

//List will return a list of templates
//
func (s *TemplateServiceOp) List() ([]ci.Ci, error) {

	//retrieve templates
	templates, err := s.getTemplateList()
	if err != nil {
		return nil, err
	}

	// this is a list .. and by default we only return the ci info on the release to the requestor

	return templates.GetCiInfo(), nil

}

//Show returns on single template
// t bool: indicate that the search need to be performed by title
func (s *TemplateServiceOp) Show(id string, title bool) (template.Template, error) {

	//if we are searching by ID then we have to make sure the id is up to par
	if !title {
		if !strings.Contains(id, templateIDPrefix) {
			id = templateIDPrefix + id
		}
	}

	// get the template list
	templates, err := s.getTemplateList()
	if err != nil {
		return template.Template{}, err
	}

	// loop over the list with templates
	for _, t := range templates {
		if !title {
			if t.ID == id {
				return t, nil
			}
		} else {
			if t.Title == id {
				return t, nil
			}
		}
	}

	// apparently we have not been able to find the requested template so lets complain about that
	newError := errors.New("unable to find template with id: " + id)

	return template.Template{}, newError

}

//CreateTemplate in xl-release
// to do this we need to take the following Applications
// remove any ID settings
// marshall the struct into json
// upload it to xlr
func (s *TemplateServiceOp) CreateTemplate(t template.Template, o bool) (template.Template, error) {

	exists, _ := s.templateExists(t.Title)

	if exists {
		err := errors.New("Template with the same title already exists:" + t.Title)
		return template.Template{}, err
	}

	s.client.NewRequest(templateBasePath, "POST", nil)

	var tn template.Templates
	var ta template.Templates

	t.PruneForUpload(template.RandomIDNr)

	ta = append(ta, t)

	data, err := json.Marshal(ta)
	if err != nil {
		return tn[0], err
	}
	req, err := s.client.NewRequest(templateBasePath+"/import", "POST", data)
	if err != nil {
		return tn[0], err
	}

	resp, err := s.client.Do(req, &tn)
	if err != nil {
		return tn[0], err
	}

	defer resp.Body.Close()

	return tn[0], nil
}

//getTemplateList retrieves a full list of templates from the xlr server
func (s *TemplateServiceOp) getTemplateList() (template.Templates, error) {
	var templates template.Templates

	req, err := s.client.NewRequest(templateBasePath, "GET", nil)

	resp, err := s.client.Do(req, &templates)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// filter out all the releases, we're only interrested in templates at this point
	templates.FilterOutReleases()

	return templates, nil

}

func (s *TemplateServiceOp) templateExists(id string) (bool, error) {

	templates, err := s.getTemplateList()
	if err != nil {
		return false, err
	}

	for _, t := range templates {

		if t.Title == id {
			return true, nil
		}
		if t.ID == id {
			return true, nil
		}
	}

	return false, nil
}
