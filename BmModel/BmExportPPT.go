package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// ExportPPT is the ExportPPT that a user consumes in order to get fat and happy
type ExportPPT struct {
	ID  string            
	Id_ int        	`jsonapi:"primary,ExportPPT"`

	Path         string  	   `jsonapi:"attr,path"`   

}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ExportPPT) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ExportPPT) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *ExportPPT) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}