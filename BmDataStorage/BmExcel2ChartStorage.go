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

// Excel2ChartStorage stores all Excel2Chartes
type Excel2ChartStorage struct {
	db *BmMongodb.BmMongodb
}

func (s Excel2ChartStorage) NewExcel2ChartStorage(args []BmDaemons.BmDaemon) *Excel2ChartStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &Excel2ChartStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s Excel2ChartStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.Excel2Chart {
	in := BmModel.Excel2Chart{}
	var out []BmModel.Excel2Chart
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.Excel2Chart
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Excel2Chart)
	}
}

// GetOne model
func (s Excel2ChartStorage) GetOne(id string) (BmModel.Excel2Chart, error) {
	in := BmModel.Excel2Chart{ID: id}
	model := BmModel.Excel2Chart{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("Excel2Chart for id %s not found", id)
	return BmModel.Excel2Chart{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *Excel2ChartStorage) Insert(c BmModel.Excel2Chart) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *Excel2ChartStorage) Delete(id string) error {
	in := BmModel.Excel2Chart{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Excel2Chart with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *Excel2ChartStorage) Update(c BmModel.Excel2Chart) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Excel2Chart with id does not exist")
	}

	return nil
}

func (s *Excel2ChartStorage) Count(req api2go.Request, c BmModel.Excel2Chart) int {
	r, _ := s.db.Count(req, &c)
	return r
}
