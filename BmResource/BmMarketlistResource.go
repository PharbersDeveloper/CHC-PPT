package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/CHC-PPT/BmDataStorage"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type MarketlistResource struct {
	MarketlistStorage *BmDataStorage.MarketlistStorage
}

func (c MarketlistResource) NewMarketlistResource(args []BmDataStorage.BmStorage) MarketlistResource {
	var cs *BmDataStorage.MarketlistStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "MarketlistStorage" {
			cs = arg.(*BmDataStorage.MarketlistStorage)
		}	
	}
	return MarketlistResource{MarketlistStorage: cs}
}

// FindAll Marketlists
func (c MarketlistResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.MarketlistStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c MarketlistResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.MarketlistStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c MarketlistResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Marketlist)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.MarketlistStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c MarketlistResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.MarketlistStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c MarketlistResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Marketlist)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.MarketlistStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
