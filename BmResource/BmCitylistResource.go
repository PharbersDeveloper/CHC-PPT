package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/CHC-PPT/BmDataStorage"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type CitylistResource struct {
	CitylistStorage *BmDataStorage.CitylistStorage
}

func (c CitylistResource) NewCitylistResource(args []BmDataStorage.BmStorage) CitylistResource {
	var cs *BmDataStorage.CitylistStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "CitylistStorage" {
			cs = arg.(*BmDataStorage.CitylistStorage)
		}	
	}
	return CitylistResource{CitylistStorage: cs}
}

// FindAll Citylists
func (c CitylistResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.CitylistStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c CitylistResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.CitylistStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c CitylistResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Citylist)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.CitylistStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c CitylistResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.CitylistStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c CitylistResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Citylist)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.CitylistStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
