package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Marketlist is the Marketlist that a user consumes in order to get fat and happy
type Marketlist struct {
	ID  string            `json:"-" bson:"-"`
	Id_ bson.ObjectId         	`json:"-" bson:"_id"`
	Name     	string  	  		`json:"name" bson:"name"` 
	Mark   	   string   	`json:"mark" bson:"mark"` 
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Marketlist) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Marketlist) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Marketlist) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}