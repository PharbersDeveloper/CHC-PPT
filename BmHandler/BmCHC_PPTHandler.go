package BmHandler

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"reflect"
	//"strconv"
	//"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	//"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/julienschmidt/httprouter"
)
type CHC_PPTHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h CHC_PPTHandler) NewBmCHC_PPTHandler(args ...interface{}) CHC_PPTHandler {
	var m *BmMongodb.BmMongodb
	var hm string
	var md string
	var ag []string
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
	return CHC_PPTHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h CHC_PPTHandler) CHC_PPT(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	jso := jsonapiobj.JsResult{}
	response := map[string]interface{}{
		"status": "",
		"sum": nil,
		"same":  nil,
		"ring":  nil,
		"error":  nil,
	}

	response["status"] = "ok"
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h CHC_PPTHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h CHC_PPTHandler) GetHandlerMethod() string {
	return h.Method
}


