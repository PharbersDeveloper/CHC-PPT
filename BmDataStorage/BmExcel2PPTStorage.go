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

// Excel2PPTStorage stores all Excel2PPTes
type Excel2PPTStorage struct {
	db *BmMongodb.BmMongodb
}

func (s Excel2PPTStorage) NewExcel2PPTStorage(args []BmDaemons.BmDaemon) *Excel2PPTStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &Excel2PPTStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s Excel2PPTStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Excel2PPT {
	in := BmModel.Excel2PPT{}
	var out []BmModel.Excel2PPT
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Excel2PPT
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Excel2PPT)
	}
}

// GetOne model
func (s Excel2PPTStorage) GetOne(id string) (BmModel.Excel2PPT, error) {
	in := BmModel.Excel2PPT{ID: id}
	model := BmModel.Excel2PPT{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Excel2PPT for id %s not found", id)
	return BmModel.Excel2PPT{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *Excel2PPTStorage) Insert(c BmModel.Excel2PPT) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *Excel2PPTStorage) Delete(id string) error {
	in := BmModel.Excel2PPT{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Excel2PPT with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *Excel2PPTStorage) Update(c BmModel.Excel2PPT) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Excel2PPT with id does not exist")
	}

	return nil
}

func (s *Excel2PPTStorage) Count(req api2go.Request, c BmModel.Excel2PPT) int {
	r, _ := s.db.Count(req, &c)
	return r
}
