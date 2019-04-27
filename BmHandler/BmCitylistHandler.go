package BmHandler

import (
	"encoding/json"
	"net/http"
	"reflect"
	//"github.com/google/jsonapi"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/julienschmidt/httprouter"
)

type CitylistHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h CitylistHandler) NewCitylistHandler(args ...interface{}) CitylistHandler {
	var m *BmMongodb.BmMongodb
	var hm string
	var md string
	var ag []string
	//sts := args[0].([]BmDaemons.BmDaemon)
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}
	return CitylistHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}
type City struct {
	Cityname     string    `json:"name"`
	Citymark   string	   `json:"mark"`
}
//TODO: load files
func (h CitylistHandler) Citylist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	citys := []City{{"北京","bj"},{"上海","sh"}}
		//_ = jsonapi.MarshalPayload(w,citys)
	enc := json.NewEncoder(w)
	enc.Encode(citys)
	return 0
}

func (h CitylistHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h CitylistHandler) GetHandlerMethod() string {
	return h.Method
}
