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

// RequestStorage stores all Requestes
type RequestStorage struct {
	db *BmMongodb.BmMongodb
}

func (s RequestStorage) NewRequestStorage(args []BmDaemons.BmDaemon) *RequestStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &RequestStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s RequestStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Request {
	in := BmModel.Request{}
	var out []BmModel.Request
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Request
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Request)
	}
}

// GetOne model
func (s RequestStorage) GetOne(id string) (BmModel.Request, error) {
	in := BmModel.Request{ID: id}
	model := BmModel.Request{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Request for id %s not found", id)
	return BmModel.Request{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *RequestStorage) Insert(c BmModel.Request) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *RequestStorage) Delete(id string) error {
	in := BmModel.Request{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Request with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *RequestStorage) Update(c BmModel.Request) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Request with id does not exist")
	}

	return nil
}

func (s *RequestStorage) Count(req api2go.Request, c BmModel.Request) int {
	r, _ := s.db.Count(req, &c)
	return r
}
