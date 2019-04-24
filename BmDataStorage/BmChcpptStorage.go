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

// ChcpptStorage stores all Chcpptes
type ChcpptStorage struct {
	db *BmMongodb.BmMongodb
}

func (s ChcpptStorage) NewChcpptStorage(args []BmDaemons.BmDaemon) *ChcpptStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &ChcpptStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s ChcpptStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Chcppt {
	in := BmModel.Chcppt{}
	var out []BmModel.Chcppt
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Chcppt
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Chcppt)
	}
}

// GetOne model
func (s ChcpptStorage) GetOne(id string) (BmModel.Chcppt, error) {
	in := BmModel.Chcppt{ID: id}
	model := BmModel.Chcppt{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Chcppt for id %s not found", id)
	return BmModel.Chcppt{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *ChcpptStorage) Insert(c BmModel.Chcppt) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *ChcpptStorage) Delete(id string) error {
	in := BmModel.Chcppt{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Chcppt with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *ChcpptStorage) Update(c BmModel.Chcppt) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Chcppt with id does not exist")
	}

	return nil
}

func (s *ChcpptStorage) Count(req api2go.Request, c BmModel.Chcppt) int {
	r, _ := s.db.Count(req, &c)
	return r
}
