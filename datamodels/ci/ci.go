package ci

import (
	"reflect"
	"strings"
)

//Cis represents A collection fo ci's
type Cis []Ci

//Ci is the base type for most fo the xlr data models
type Ci struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"$createdAt,omitempty"`
	CreatedBy      string `json:"$createdBy,omitempty"`
	LastModifiedAt string `json:"$lastModifiedAt,omitempty"`
	LastModifiedBy string `json:"$lastModifiedBy,omitempty"`
	Token          string `json:"$token,omitempty"`
	Title          string `json:"title,omitempty"`
	Type           string `json:"type,omitempty"`
}

//PruneXLRInternalSettings replaces all internal settings for xlr in a template ot a nil string value
func (c *Cis) PruneXLRInternalSettings() {

	var newCis Cis

	cis := *c

	for _, ci := range cis {
		ci.PruneXLRInternalSettings()
		newCis = append(newCis, ci)
	}

	*c = newCis
}

//PruneXLRInternalSettings replaces all internal settings for xlr in a template ot a nil string value
//  to accomplish this we're using reflection(reflection is bad but very convenient)
//  we range the struct and search for json tag values that contain a dollar sign ($)
//  is we find on we update in place and overwrite the existing value with a empty string
//  we do assume that we are dealing with a string .. this might break if we are going to do time value crap
func (c *Ci) PruneXLRInternalSettings() {
	ci := *c
	value := reflect.ValueOf(ci)            // turns struct type to the reflection interface
	for i := 0; i < value.NumField(); i++ { // iterates through every struct type field
		tag := value.Type().Field(i).Tag // returns the tag string
		if strings.Contains(tag.Get("json"), "$") {

			reflect.ValueOf(&ci).Elem().Field(i).SetString("")

		}

	}

	*c = ci
}
