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

// ChcppttemplateStorage stores all Chcppttemplatees
type ChcppttemplateStorage struct {
	db *BmMongodb.BmMongodb
}

func (s ChcppttemplateStorage) NewChcppttemplateStorage(args []BmDaemons.BmDaemon) *ChcppttemplateStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &ChcppttemplateStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s ChcppttemplateStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Chcppttemplate {
	in := BmModel.Chcppttemplate{}
	var out []BmModel.Chcppttemplate
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Chcppttemplate
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Chcppttemplate)
	}
}

// GetOne model
func (s ChcppttemplateStorage) GetOne(id string) (BmModel.Chcppttemplate, error) {
	in := BmModel.Chcppttemplate{ID: id}
	model := BmModel.Chcppttemplate{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Chcppttemplate for id %s not found", id)
	return BmModel.Chcppttemplate{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *ChcppttemplateStorage) Insert(c BmModel.Chcppttemplate) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *ChcppttemplateStorage) Delete(id string) error {
	in := BmModel.Chcppttemplate{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Chcppttemplate with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *ChcppttemplateStorage) Update(c BmModel.Chcppttemplate) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Chcppttemplate with id does not exist")
	}

	return nil
}

func (s *ChcppttemplateStorage) Count(req api2go.Request, c BmModel.Chcppttemplate) int {
	r, _ := s.db.Count(req, &c)
	return r
}
