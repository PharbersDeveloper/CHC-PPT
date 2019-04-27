package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/manyminds/api2go"
)

// PptinformationStorage stores all Pptinformationes
type PptinformationStorage struct {
	db *BmMongodb.BmMongodb
}

func (s PptinformationStorage) NewPptinformationStorage(args []BmDaemons.BmDaemon) *PptinformationStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &PptinformationStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s PptinformationStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Pptinformation {
	in := BmModel.Pptinformation{}
	var out []BmModel.Pptinformation
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Pptinformation
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Pptinformation)
	}
}

// GetOne model
func (s PptinformationStorage) GetOne(id string) (BmModel.Pptinformation, error) {
	in := BmModel.Pptinformation{ID: id}
	model := BmModel.Pptinformation{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Pptinformation for id %s not found", id)
	return BmModel.Pptinformation{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *PptinformationStorage) Insert(c BmModel.Pptinformation) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *PptinformationStorage) Delete(id string) error {
	in := BmModel.Pptinformation{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Pptinformation with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *PptinformationStorage) Update(c BmModel.Pptinformation) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Pptinformation with id does not exist")
	}

	return nil
}

func (s *PptinformationStorage) Count(req api2go.Request, c BmModel.Pptinformation) int {
	r, _ := s.db.Count(req, &c)
	return r
}
