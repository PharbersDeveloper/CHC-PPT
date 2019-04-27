package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Citylist is the Citylist that a user consumes in order to get fat and happy
type Citylist struct {
	ID  string            `json:"-" bson:"-"`
	Id_ bson.ObjectId         	`json:"-" bson:"_id"`
	Name     	string  	  		`json:"name" bson:"name"` 
	Mark   	   string   	`json:"mark" bson:"mark"` 
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Citylist) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Citylist) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Citylist) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}