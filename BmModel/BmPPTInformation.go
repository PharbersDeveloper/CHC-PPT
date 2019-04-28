package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Pptinformation is the Pptinformation that a user consumes in order to get fat and happy
type Pptinformation struct {
	ID  string           
	Id_ bson.ObjectId         	`json:"pptinformation" bson:"_id"`
	Data        []interface{}  `json:"-" bson:"data"`
	Url         string  	   `json:"url" bson:"url"`
	City		string  	   `json:"city" bson:"city"`
	Market		string  	   `json:"market" bson:"market"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Pptinformation) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Pptinformation) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Pptinformation) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "market":
			rst[k] = v[0]
		case "city":
			rst[k] = v[0]
		}
	}
	return rst
}