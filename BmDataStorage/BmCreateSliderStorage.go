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

// CreateSliderStorage stores all PhCreateSlideres
type CreateSliderStorage struct {
	db *BmMongodb.BmMongodb
}

func (s CreateSliderStorage) NewCreateSliderStorage(args []BmDaemons.BmDaemon) *CreateSliderStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &CreateSliderStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s CreateSliderStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.CreateSlider {
	in := BmModel.CreateSlider{}
	var out []BmModel.CreateSlider
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.CreateSlider
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.PhCreateSlider)
	}
}

// GetOne model
func (s CreateSliderStorage) GetOne(id string) (BmModel.CreateSlider, error) {
	in := BmModel.CreateSlider{ID: id}
	model := BmModel.CreateSlider{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("CreateSlider for id %s not found", id)
	return BmModel.CreateSlider{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *CreateSliderStorage) Insert(c BmModel.CreateSlider) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *CreateSliderStorage) Delete(id string) error {
	in := BmModel.CreateSlider{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("CreateSlider with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *CreateSliderStorage) Update(c BmModel.CreateSlider) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("CreateSlider with id does not exist")
	}

	return nil
}

func (s *CreateSliderStorage) Count(req api2go.Request, c BmModel.CreateSlider) int {
	r, _ := s.db.Count(req, &c)
	return r
}
