package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// PPTInformation is the PPTInformation that a user consumes in order to get fat and happy
type PPTInformation struct {
	ID  string           
	Id_ bson.ObjectId         	`jsonapi:"primary,PPTInformation" bson:"_id"`

	Data        []interface{}  `json:"-" bson:"data"`
	Url         string  	   `jsonapi:"url" bson:"url"`
	City		string  	   `jsonapi:"city" bson:"city"`
	Market		string  	   `jsonapi:"market" bson:"market"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c PPTInformation) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *PPTInformation) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *PPTInformation) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}