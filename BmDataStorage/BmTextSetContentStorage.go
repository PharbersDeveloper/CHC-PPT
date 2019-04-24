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

// TextSetContentStorage stores all TextSetContentes
type TextSetContentStorage struct {
	db *BmMongodb.BmMongodb
}

func (s TextSetContentStorage) NewTextSetContentStorage(args []BmDaemons.BmDaemon) *TextSetContentStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &TextSetContentStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s TextSetContentStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.TextSetContent {
	in := BmModel.TextSetContent{}
	var out []BmModel.TextSetContent
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.TextSetContent
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.TextSetContent)
	}
}

// GetOne model
func (s TextSetContentStorage) GetOne(id string) (BmModel.TextSetContent, error) {
	in := BmModel.TextSetContent{ID: id}
	model := BmModel.TextSetContent{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("TextSetContent for id %s not found", id)
	return BmModel.TextSetContent{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *TextSetContentStorage) Insert(c BmModel.TextSetContent) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *TextSetContentStorage) Delete(id string) error {
	in := BmModel.TextSetContent{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("TextSetContent with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *TextSetContentStorage) Update(c BmModel.TextSetContent) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("TextSetContent with id does not exist")
	}

	return nil
}

func (s *TextSetContentStorage) Count(req api2go.Request, c BmModel.TextSetContent) int {
	r, _ := s.db.Count(req, &c)
	return r
}
