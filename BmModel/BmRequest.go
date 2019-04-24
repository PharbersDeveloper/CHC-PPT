package BmModel

import (
	// "github.com/alfredyang1986/blackmirror/bmmodel"
	// "github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
	"errors"
)

type Request struct {
	ID  string         
	Id_ int      `jsonapi:"primary,Requests"`

	Command         string  	  `jsonapi:"attr,command"`
	Jobid   		string        `jsonapi:"attr,jobid"`

	CreateSlider   *CreateSlider `jsonapi:"relation,slider"` 
	CreateSliderID string        `jsonapi:"attr,createslider-id"`

	//Slider  *CreateSlider `jsonapi:"relation,Slider"` 

	ExportPPT   	 *ExportPPT 	 `jsonapi:"relation,exp"` 
	ExportPPTID     string        `jsonapi:"attr,exportppt-id-id"`

	Excel2Chart    *Excel2Chart  `jsonapi:"relation,e2c"` 
	Excel2ChartID  string        `jsonapi:"attr,excel2chart-id"`

	ExcelPush      *ExcelPush    `jsonapi:"relation,push"` 
	ExcelPushID    string       `jsonapi:"attr,excelpush-id" ` 

	Excel2PPT      *Excel2PPT   `jsonapi:"relation,e2p"`  
	Excel2PPTID    string        `jsonapi:"attr,excel2ppt-id"`

	TextSetContent      *TextSetContent    `jsonapi:"relation,text"` 
	TextSetContentID    string       `jsonapi:"attr,textsetcontent-id"` 

}
// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Request) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Request) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Request) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "createsliders",
			Name: "createslider",
		},
		{
			Type: "textsetcontents",
			Name: "textsetcontent",
		},
		{
			Type: "exportppts",
			Name: "exportppt",
		},
		{
			Type: "excel2charts",
			Name: "excel2chart",
		},
		{
			Type: "excelpushs",
			Name: "excelpush",
		},
		{
			Type: "excel2ppts",
			Name: "excel2ppt",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Request) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	if u.CreateSliderID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.CreateSliderID,
			Type: "createsliders",
			Name: "createslider",
		})
	}
	if u.TextSetContentID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.TextSetContentID,
			Type: "textsetcontents",
			Name: "textsetcontent",
		})
	}
	if u.ExportPPTID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ExportPPTID,
			Type: "exportppts",
			Name: "exportppt",
		})
	}
	if u.ExcelPushID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ExcelPushID,
			Type: "excelpushs",
			Name: "excelpush",
		})
	}
	if u.Excel2ChartID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.Excel2ChartID,
			Type: "excel2charts",
			Name: "excel2chart",
		})
	}
	if u.Excel2PPTID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.Excel2PPTID,
			Type: "excel2ppts",
			Name: "excel2ppt",
		})
	}
	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Request) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.CreateSliderID != "" && u.CreateSlider != nil {
		result = append(result, u.CreateSlider)
	}
	if u.ExportPPTID != "" && u.ExportPPT != nil {
		result = append(result, u.ExportPPT)
	}
	if u.TextSetContentID != "" && u.TextSetContent != nil {
		result = append(result, u.TextSetContent)
	}
	if u.Excel2ChartID != "" && u.Excel2Chart != nil {
		result = append(result, u.Excel2Chart)
	}
	if u.ExcelPushID != "" && u.ExcelPush != nil {
		result = append(result, u.ExcelPush)
	}
	if u.Excel2PPTID != "" && u.Excel2PPT != nil {
		result = append(result, u.Excel2PPT)
	}

	return result
}

func (u *Request) SetToOneReferenceID(name, ID string) error {
	if name == "createslider" {
		u.CreateSliderID = ID
		return nil
	}
	if name == "exportppt" {
		u.ExportPPTID = ID
		return nil
	}
	if name == "textsetcontent" {
		u.TextSetContentID = ID
		return nil
	}
	if name == "excel2chart" {
		u.Excel2ChartID = ID
		return nil
	}
	if name == "excelpush" {
		u.ExcelPushID = ID
		return nil
	}
	if name == "excel2ppt" {
		u.Excel2PPTID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}
func (u *Request) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	return rst
}