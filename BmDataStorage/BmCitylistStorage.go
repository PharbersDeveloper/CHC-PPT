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

// CitylistStorage stores all Citylistes
type CitylistStorage struct {
	db *BmMongodb.BmMongodb
}

func (s CitylistStorage) NewCitylistStorage(args []BmDaemons.BmDaemon) *CitylistStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &CitylistStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s CitylistStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Citylist {
	in := BmModel.Citylist{}
	var out []BmModel.Citylist
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Citylist
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Citylist)
	}
}

// GetOne model
func (s CitylistStorage) GetOne(id string) (BmModel.Citylist, error) {
	in := BmModel.Citylist{ID: id}
	model := BmModel.Citylist{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Citylist for id %s not found", id)
	return BmModel.Citylist{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *CitylistStorage) Insert(c BmModel.Citylist) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *CitylistStorage) Delete(id string) error {
	in := BmModel.Citylist{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Citylist with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *CitylistStorage) Update(c BmModel.Citylist) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Citylist with id does not exist")
	}

	return nil
}

func (s *CitylistStorage) Count(req api2go.Request, c BmModel.Citylist) int {
	r, _ := s.db.Count(req, &c)
	return r
}
