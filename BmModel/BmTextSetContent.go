package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// TextSetContent is the TextSetContent that a user consumes in order to get fat and happy
type TextSetContent struct {
	ID  string            
	Id_ int        	`jsonapi:"primary,TextSetContent"`

	Slider         int  	   `jsonapi:"attr,slider"`   
	Pos          []int        `jsonapi:"attr,pos"`    
	Content        string  	  `jsonapi:"attr,content"`    
	ShapeType     string   	   `jsonapi:"attr,shapeType"`   
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c TextSetContent) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *TextSetContent) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *TextSetContent) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}