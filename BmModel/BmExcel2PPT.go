package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Excel2PPT is the Excel2PPT that a user consumes in order to get fat and happy
type Excel2PPT struct {
	ID  string            
	Id_ int        	`jsonapi:"primary,Excel2PPT"`

	Name         string  	  `jsonapi:"attr,name"`     
	Slider         int  	  `jsonapi:"attr,slider"`     
	Pos          []int       `jsonapi:"attr,pos"`      
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Excel2PPT) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Excel2PPT) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Excel2PPT) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}