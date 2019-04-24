package BmModel

import (
	bson "gopkg.in/mgo.v2/bson"
	//"fmt"
)

// Excel2Chart is the Excel2Chart that a user consumes in order to get fat and happy
type Excel2Chart struct {
	ID  string            
	Id_ int       		 `jsonapi:"primary,Excel2Chart"`

	Css          string  	  `jsonapi:"attr,css"`   
	Name         string  	  `jsonapi:"attr,name"`   
	Slider         int  	 `jsonapi:"attr,slider"`    
	ChartType    string   	 `jsonapi:"attr,charttype"`     
	Pos          []int      `jsonapi:"attr,pos"`      
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Excel2Chart) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Excel2Chart) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Excel2Chart) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}