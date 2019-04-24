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

// ExportPPTStorage stores all ExportPPTes
type ExportPPTStorage struct {
	db *BmMongodb.BmMongodb
}

func (s ExportPPTStorage) NewExportPPTStorage(args []BmDaemons.BmDaemon) *ExportPPTStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &ExportPPTStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s ExportPPTStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.ExportPPT {
	in := BmModel.ExportPPT{}
	var out []BmModel.ExportPPT
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.ExportPPT
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.ExportPPT)
	}
}

// GetOne model
func (s ExportPPTStorage) GetOne(id string) (BmModel.ExportPPT, error) {
	in := BmModel.ExportPPT{ID: id}
	model := BmModel.ExportPPT{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("ExportPPT for id %s not found", id)
	return BmModel.ExportPPT{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *ExportPPTStorage) Insert(c BmModel.ExportPPT) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *ExportPPTStorage) Delete(id string) error {
	in := BmModel.ExportPPT{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ExportPPT with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *ExportPPTStorage) Update(c BmModel.ExportPPT) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ExportPPT with id does not exist")
	}

	return nil
}

func (s *ExportPPTStorage) Count(req api2go.Request, c BmModel.ExportPPT) int {
	r, _ := s.db.Count(req, &c)
	return r
}
