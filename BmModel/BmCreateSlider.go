package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// CreateSlider is the CreateSlider that a user consumes in order to get fat and happy
type CreateSlider struct {
	ID  string            
	Id_ int         	`jsonapi:"primary,CreateSlider"`
	SliderType         string  	  `jsonapi:"attr,slidertype"` 
	Title   	   string        `jsonapi:"attr,title"` 
	Slider         int  		 `jsonapi:"attr,slider"` 
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c CreateSlider) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *CreateSlider) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *CreateSlider) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}