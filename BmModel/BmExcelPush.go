package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// ExcelPush is the ExcelPush that a user consumes in order to get fat and happy
type ExcelPush struct {
	ID  string            
	Id_ int        	`jsonapi:"primary,ExcelPush"`

	Name         string  	  `jsonapi:"attr,name"`     
	Cells          []string      `jsonapi:"attr,cells"`     
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ExcelPush) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ExcelPush) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *ExcelPush) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}