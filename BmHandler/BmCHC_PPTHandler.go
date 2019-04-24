package BmHandler

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"reflect"
	"os"
	//"bytes"
	//"strings"
	//"strconv"
	//"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/julienschmidt/httprouter"
)
type CHC_PPTHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}
type PersonInfo struct {
    Name    string
    age     int32
    Sex     bool
    Hobbies []string
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
	Request := BmModel.Request{}
	
	// response := map[string]interface{}{
	// 	"status": "",
	// 	"result": nil,
	// 	"error":  nil,
	// }
	data2 := []byte("\n")
	data,_ := jsonapi.ToJsonByteArr(Request)
	filePtr, err := os.Create("CHC_PPT.json")
    if err != nil {
        fmt.Println("Create file failed", err.Error())
        return 0
    }
	defer filePtr.Close()
	filePtr.Write(data)
	filePtr.Write(data2)
	return 0
}
/*
func (h CHC_PPTHandler) CHC_PPT(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	data2 := []byte("\n")
	creatPpt := h.CreatPPT()
	creatPage := h.creatPage()
	creatTxtRect := h.creatTxtRect()
	insertTable := h.insertTable()
	makeExcel := h.makeExcel()
	insertPpt := h.insertPpt()
	
    // 创建文件
    filePtr, err := os.Create("CHC_PPT.json")
    if err != nil {
        fmt.Println("Create file failed", err.Error())
        return 0
    }
    defer filePtr.Close()

	//creatPpt
	data, err := json.MarshalIndent(creatPpt, "", "  ")
    if err != nil {
	   fmt.Println("CreatPPT failed", err.Error())
	} 
	filePtr.Write(data)
	filePtr.Write(data2)

	//creatPage
	data, err = json.MarshalIndent(creatPage, "", "  ")
    if err != nil {
	   fmt.Println("CreatPPT failed", err.Error())
	} 
	filePtr.Write(data)
	filePtr.Write(data2)
	
	//creatTxtRect
	data, err = json.MarshalIndent(creatTxtRect, "", "  ")
    if err != nil {
	   fmt.Println("CreatPPT failed", err.Error())
	} 
	filePtr.Write(data)
	filePtr.Write(data2)

	//insertTable
	data, err = json.MarshalIndent(insertTable, "", "  ")
    if err != nil {
	   fmt.Println("CreatPPT failed", err.Error())
	} 
	filePtr.Write(data)
	filePtr.Write(data2)

	//makeExcel
	data, err = json.MarshalIndent(makeExcel, "", "  ")
    if err != nil {
	   fmt.Println("CreatPPT failed", err.Error())
	} 
	filePtr.Write(data)
	filePtr.Write(data2)

	//insertPpt
	data, err = json.MarshalIndent(insertPpt, "", "  ")
    if err != nil {
	   fmt.Println("CreatPPT failed", err.Error())
	} 
	filePtr.Write(data)
	filePtr.Write(data2)

	return 0
}
*/
func (h CHC_PPTHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h CHC_PPTHandler) GetHandlerMethod() string {
	return h.Method
}

func (h CHC_PPTHandler) CreatPPT() interface{} {
	attributes := map[string]interface{}{
		"jobid" : "2b76e502-67c5-4523-92c5-3342a15dae31",
		"command" : "GenPPT",
	}
	data := map[string]interface{}{
		"attributes" : attributes,
		"id" : "f8f0539b-a562-4d51-b67a-0d8a421911f3",
		"type" : "PhRequest",
	}
	response := map[string]interface{}{
		"data" : data,
	}
	return response
}

func (h CHC_PPTHandler) creatPage() interface{} {
	incattributes := map[string]interface{}{
		"sliderType" : "title",
		"slider" : 0,
		"title" : "口服降糖药市场CHC数据分析报告\n2018Q3YTD",
	}
	included := map[string]interface{}{
		"type" : "PhCreateSlider",
		"id" : "tmp_08d8363d-5abb-45c5-9a16-5905fe09fa25",
		"attributes" : incattributes,
	}
	sliderdata := map[string]interface{}{
		"type" : "PhCreateSlider",
		"id" : "tmp_08d8363d-5abb-45c5-9a16-5905fe09fa25" ,
	}
	slider := map[string]interface{}{
		"data" : sliderdata,
	}
	relationships := map[string]interface{}{
		"slider" : slider,
	}
	attributes := map[string]interface{}{
			"jobid" : "2b76e502-67c5-4523-92c5-3342a15dae31",
			"command" : "CreateSlider",
	}
	data := map[string]interface{}{
		"type" : "PhRequest",
		"id" : "644e4897-a4e0-4468-9d2f-8419428b08dd",
		"attributes" : attributes,
		"relationships" : relationships,
	}
	response := map[string]interface{}{
		"data" : data,
		"included" : included,
	}
	return response
}
func (h CHC_PPTHandler) creatTxtRect() interface{} {
	var pos = []float32{892086, 2658273, 7550359,690647}
	incattributes := map[string]interface{}{
		"content" : "#{##[#口服降糖药市场CHC数据分析报告\n2018Q3YTD#C#20Blue#]##P#center#}#",
		"pos" : pos,
		"slider" : 0,
		"shapeType" : "Rectangle",
	}
	included := map[string]interface{}{
		"type" : "PhTextPPT",
		"id" : "tmp_9db785a7-462d-42a5-8b92-5762162bf73d",
		"attributes" : incattributes,
	}
	textdata := map[string]interface{}{
		"type" : "PhTextPPT",
		"id" : "tmp_9db785a7-462d-42a5-8b92-5762162bf73d",
	}
	text := map[string]interface{}{
		"data" : textdata,
	}
	relationships := map[string]interface{}{
		"text" : text,
	}

	attributes := map[string]interface{}{
		"jobid" : "2b76e502-67c5-4523-92c5-3342a15dae31",
		"command" : "PushText",
	}
	data := map[string]interface{}{
		"type" : "PhRequest",
		"id" : "f2f4f241-017a-428a-9bdd-e92ed777c974",
		"attributes" : attributes,
		"relationships" : relationships,
	}
	response := map[string]interface{}{
		"data" : data,
		"included" : included,
	}
	return response
}

func (h CHC_PPTHandler) insertTable() interface{} {
	var pos = []float32{892086, 2658273, 7550359,690647}
	incattributes := map[string]interface{}{
		"name" : "1",
		"pos" : pos,
		"slider" : 1,
		"chartType" : "Stacked",
		"css" : "",
	}
	included := map[string]interface{}{
	"type" : "PhExcel2Chart",
    "id" : "bb611104-3a2c-49fa-bd12-5d1fe1a6c730",
    "attributes" : incattributes,
	}

	e2cdata := map[string]interface{}{
		"type" : "PhExcel2Chart",
		"id" : "bb611104-3a2c-49fa-bd12-5d1fe1a6c730",
	}
	e2c := map[string]interface{}{
		"data" : e2cdata,
	}
	relationships := map[string]interface{}{
		"e2c" : e2c,
	}
	attributes := map[string]interface{}{
		"jobid" : "2b76e502-67c5-4523-92c5-3342a15dae31",
        "command" : "Excel2Chart",
	}
	data := map[string]interface{}{
	"type" : "PhRequest",
    "id" : "d27e778d-973c-4826-8ee7-573c9b1f649c",
	"attributes" : attributes,
	"relationships" : relationships,
	}
	response := map[string]interface{}{
		"data" : data,
		"included" : included,
	}
	return response
}

func (h CHC_PPTHandler) makeExcel() interface{} {
	var cells = []string{ "#c#C6#s#col_common3*row_5#t#Number#v#",
	"#c#E10#s#col_common1*row_5#t#Number#v#",
	"#c#F7#s#col_common3*row_5#t#Number#v#",
	"#c#B5#s#col_common1*row_5#t#Number#v#",
	"#c#C3#s#col_common3*row_7#t#Number#v#",
	"#c#C19#s#col_common3*row_5#t#Number#v#",
	"#c#C8#s#col_common3*row_5#t#Number#v#",
	"#c#G15#s#col_common3*row_5#t#Number#v#",
	"#c#C7#s#col_common3*row_5#t#Number#v#"}
	incattributes := map[string]interface{}{
		"name" : "0d2d1e3f-299b-4fe1-8d17-4820a37bb03b",
		"cells" : cells,
	}
	included := map[string]interface{}{
		"type" : "PhExcelPush",
		"id" : "4c6892db-74f7-43b0-bd25-b02d0e042fa4",
		"attributes" : incattributes,
	}
	pushdata := map[string]interface{}{
		"type" : "PhExcelPush",
		"id" : "4c6892db-74f7-43b0-bd25-b02d0e042fa4",
	}
	push := map[string]interface{}{
		"data" : pushdata,
	}
	relationships := map[string]interface{}{
		"push" : push,
	}
	attributes := map[string]interface{}{
		"jobid" : "71d87f12-53a0-404b-939a-d06518b77ea9",
		"command" : "ExcelPush",
	}
	data := map[string]interface{}{
	"type" : "PhRequest",
    "id" : "26fabbb3-e78b-4152-8e2a-08b2bdf36366",
    "attributes" : attributes,
    "relationships" : relationships,
	}
	response := map[string]interface{}{
		"data" : data,
		"included" : included,
	}
	return response
}

func (h CHC_PPTHandler) insertPpt() interface{} {
	var pos = []float32{643884,1352517,7705035,2438848}
	incattributes := map[string]interface{}{
		"name" : "0d2d1e3f-299b-4fe1-8d17-4820a37bb03b",
		"pos" : pos,
		"slider" : 0,
	}
	included := map[string]interface{}{
	"type" : "PhExcel2PPT",
    "id" : "2b0ece78-2688-4bd7-9d93-b430cfb62622",
    "attributes" : incattributes,
	}

	e2pdata := map[string]interface{}{
		"type" : "PhExcel2PPT",
		"id" : "2b0ece78-2688-4bd7-9d93-b430cfb62622",
	}
	e2p := map[string]interface{}{
		"data" : e2pdata,
	}
	relationships := map[string]interface{}{
		"e2p" : e2p,
	}
	attributes := map[string]interface{}{
		"jobid" : "71d87f12-53a0-404b-939a-d06518b77ea9",
        "command" : "Excel2PPT",
	}
	data := map[string]interface{}{
	"type" : "PhRequest",
    "id" : "0310a1ae-a225-4971-b2a6-52260887cc18",
	"attributes" : attributes,
	"relationships" : relationships,
	}
	response := map[string]interface{}{
		"data" : data,
		"included" : included,
	}
	return response
}