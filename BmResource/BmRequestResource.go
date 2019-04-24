package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/CHC-PPT/BmDataStorage"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	// "github.com/alfredyang1986/blackmirror/jsonapi"
	// "os"
	// "fmt"
	// "strings"
)

type RequestResource struct {
	RequestStorage 		*BmDataStorage.RequestStorage
	CreateSliderStorage *BmDataStorage.CreateSliderStorage
	Excel2ChartStorage  *BmDataStorage.Excel2ChartStorage
	ExcelPushStorage    *BmDataStorage.ExcelPushStorage
	ExportPPTStorage 	*BmDataStorage.ExportPPTStorage
	Excel2PPTStorage    *BmDataStorage.Excel2PPTStorage
	TextSetContentStorage    *BmDataStorage.TextSetContentStorage
}

func (c RequestResource) NewRequestResource(args []BmDataStorage.BmStorage) RequestResource {
	var cs *BmDataStorage.RequestStorage
	var ss *BmDataStorage.CreateSliderStorage
	var ds *BmDataStorage.Excel2ChartStorage
	var ys *BmDataStorage.ExcelPushStorage
	var rs *BmDataStorage.ExportPPTStorage
	var ts *BmDataStorage.TextSetContentStorage
	var us *BmDataStorage.Excel2PPTStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "RequestStorage" {
			cs = arg.(*BmDataStorage.RequestStorage)
		}else if tp.Name() == "CreateSliderStorage" {
			ss = arg.(*BmDataStorage.CreateSliderStorage)
		} else if tp.Name() == "Excel2ChartStorage" {
			ds = arg.(*BmDataStorage.Excel2ChartStorage)
		} else if tp.Name() == "ExcelPushStorage" {
			ys = arg.(*BmDataStorage.ExcelPushStorage)
		} else if tp.Name() == "ExportPPTStorage" {
			rs = arg.(*BmDataStorage.ExportPPTStorage)
		} else if tp.Name() == "TextSetContentStorage" {
			ts = arg.(*BmDataStorage.TextSetContentStorage)
		} else if tp.Name() == "Excel2PPTStorage" {
			us = arg.(*BmDataStorage.Excel2PPTStorage)
		}	
	}
	return RequestResource{RequestStorage: cs,CreateSliderStorage:ss,Excel2ChartStorage:ds,ExcelPushStorage:ys,ExportPPTStorage:rs,Excel2PPTStorage:us,TextSetContentStorage:ts}
}

// FindAll Requests
func (c RequestResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	result := c.RequestStorage.GetAll(r,-1,-1)
	// if len(result) <= 0{
	// 	filePtr, err := os.Open("person_info.json")
	// 	if err != nil {
	// 		//fmt.Println("Open file failed [Err:%s]", err.Error())
	// 		return nil,nil
	// 	}
	// 	defer filePtr.Close()
	// 	var a BmModel.Request
	// 	jsonapi.ToJsonAPI(a,filePtr)

	// scheme := "http://"
	// //version := strings.Split(r.URL.Path, "/")[1]

	// resource := fmt.Sprint("192.168.100.195:9999/api/ppt")
	// mergeURL := strings.Join([]string{scheme, resource}, "")

	// // 转发
	// client := &http.Client{}
	// req, _ := http.NewRequest("POST", mergeURL, nil)

	// for k, v := range r.Header {
	// 	req.Header.Add(k, v[0])
	// }
	// response, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Fuck Error")
	// }
	// result, err := ioutil.ReadAll(response.Body)
	// data := map[string]string {}
	// json.Unmarshal(result, &data)
	// }
	return &Response{Res: result}, nil
}

// FindOne choc
func (c RequestResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.RequestStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c RequestResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Request)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.RequestStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c RequestResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.RequestStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c RequestResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Request)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.RequestStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
