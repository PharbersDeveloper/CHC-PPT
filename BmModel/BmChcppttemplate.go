package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Chcppttemplate is the Chcppttemplate that a user consumes in order to get fat and happy
type Chcppttemplate struct {
	ID  string            
	Id_ bson.ObjectId         	`jsonapi:"primary,Chcppttemplate"  bson:"_id"`
	Slider_Index  string        `jsonapi:"attr,silder_index"  bson:"silder_index"` 
	Slider_Type  string			`jsonapi:"attr,slider_type" bson:"slider_type"` 
	Shapes    []interface{}     `jsonapi:"attr,shapes" bson:"shapes"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Chcppttemplate) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Chcppttemplate) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Chcppttemplate) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}