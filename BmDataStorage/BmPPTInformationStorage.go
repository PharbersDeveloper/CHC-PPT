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

// PPTInformationStorage stores all PPTInformationes
type PPTInformationStorage struct {
	db *BmMongodb.BmMongodb
}

func (s PPTInformationStorage) NewPPTInformationStorage(args []BmDaemons.BmDaemon) *PPTInformationStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &PPTInformationStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s PPTInformationStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.PPTInformation {
	in := BmModel.PPTInformation{}
	var out []BmModel.PPTInformation
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.PPTInformation
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.PPTInformation)
	}
}

// GetOne model
func (s PPTInformationStorage) GetOne(id string) (BmModel.PPTInformation, error) {
	in := BmModel.PPTInformation{ID: id}
	model := BmModel.PPTInformation{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("PPTInformation for id %s not found", id)
	return BmModel.PPTInformation{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *PPTInformationStorage) Insert(c BmModel.PPTInformation) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *PPTInformationStorage) Delete(id string) error {
	in := BmModel.PPTInformation{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("PPTInformation with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *PPTInformationStorage) Update(c BmModel.PPTInformation) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("PPTInformation with id does not exist")
	}

	return nil
}

func (s *PPTInformationStorage) Count(req api2go.Request, c BmModel.PPTInformation) int {
	r, _ := s.db.Count(req, &c)
	return r
}
