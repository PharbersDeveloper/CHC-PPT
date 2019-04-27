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

// MarketlistStorage stores all Marketlistes
type MarketlistStorage struct {
	db *BmMongodb.BmMongodb
}

func (s MarketlistStorage) NewMarketlistStorage(args []BmDaemons.BmDaemon) *MarketlistStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &MarketlistStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s MarketlistStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Marketlist {
	in := BmModel.Marketlist{}
	var out []BmModel.Marketlist
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Marketlist
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Marketlist)
	}
}

// GetOne model
func (s MarketlistStorage) GetOne(id string) (BmModel.Marketlist, error) {
	in := BmModel.Marketlist{ID: id}
	model := BmModel.Marketlist{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Marketlist for id %s not found", id)
	return BmModel.Marketlist{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *MarketlistStorage) Insert(c BmModel.Marketlist) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *MarketlistStorage) Delete(id string) error {
	in := BmModel.Marketlist{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Marketlist with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *MarketlistStorage) Update(c BmModel.Marketlist) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Marketlist with id does not exist")
	}

	return nil
}

func (s *MarketlistStorage) Count(req api2go.Request, c BmModel.Marketlist) int {
	r, _ := s.db.Count(req, &c)
	return r
}
