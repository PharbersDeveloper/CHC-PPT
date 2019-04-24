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

// ExcelPushStorage stores all ExcelPushes
type ExcelPushStorage struct {
	db *BmMongodb.BmMongodb
}

func (s ExcelPushStorage) NewExcelPushStorage(args []BmDaemons.BmDaemon) *ExcelPushStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &ExcelPushStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s ExcelPushStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.ExcelPush {
	in := BmModel.ExcelPush{}
	var out []BmModel.ExcelPush
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.ExcelPush
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.ExcelPush)
	}
}

// GetOne model
func (s ExcelPushStorage) GetOne(id string) (BmModel.ExcelPush, error) {
	in := BmModel.ExcelPush{ID: id}
	model := BmModel.ExcelPush{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("ExcelPush for id %s not found", id)
	return BmModel.ExcelPush{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *ExcelPushStorage) Insert(c BmModel.ExcelPush) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *ExcelPushStorage) Delete(id string) error {
	in := BmModel.ExcelPush{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ExcelPush with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *ExcelPushStorage) Update(c BmModel.ExcelPush) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ExcelPush with id does not exist")
	}

	return nil
}

func (s *ExcelPushStorage) Count(req api2go.Request, c BmModel.ExcelPush) int {
	r, _ := s.db.Count(req, &c)
	return r
}
