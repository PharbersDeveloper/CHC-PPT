package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Chcppt is the Chcppt that a user consumes in order to get fat and happy
type Chcppt struct {
	ID  string            
	Id_ bson.ObjectId         	`jsonapi:"primary,Chcppt" bson:"_id"`
	TableIndex     string  	  `jsonapi:"attr,tableindex" bson:"tableIndex"` 
	Cells   	   []interface{}   `jsonapi:"attr,cells" bson:"cells"` 
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Chcppt) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Chcppt) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Chcppt) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}