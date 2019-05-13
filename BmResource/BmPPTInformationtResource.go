package BmResource

import (
	//"bytes"
	"errors"
	"github.com/PharbersDeveloper/CHC-PPT/BmDataStorage"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"github.com/google/jsonapi"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	bson "gopkg.in/mgo.v2/bson"
	"time"
	"regexp"
	"github.com/hashicorp/go-uuid"
	"github.com/alfredyang1986/blackmirror/bmkafka"
)
var url string
type PptinformationResource struct {
	RequestStorage 		  *BmDataStorage.RequestStorage
	PptinformationStorage *BmDataStorage.PptinformationStorage
	ChcpptStorage         *BmDataStorage.ChcpptStorage
	ChcppttemplateStorage         *BmDataStorage.ChcppttemplateStorage
}

func (c PptinformationResource) NewPptinformationResource(args []BmDataStorage.BmStorage) PptinformationResource {
	var cs *BmDataStorage.RequestStorage
	var ss *BmDataStorage.PptinformationStorage
	var ts *BmDataStorage.ChcpptStorage
	var ps *BmDataStorage.ChcppttemplateStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "RequestStorage" {
			cs = arg.(*BmDataStorage.RequestStorage)
		}else if tp.Name() == "PptinformationStorage" {
			ss = arg.(*BmDataStorage.PptinformationStorage)
		}else if tp.Name() == "ChcpptStorage" {
			ts = arg.(*BmDataStorage.ChcpptStorage)
		} else if tp.Name() == "ChcppttemplateStorage" {
			ps = arg.(*BmDataStorage.ChcppttemplateStorage)
		}
	}
	return PptinformationResource{RequestStorage: cs,PptinformationStorage:ss,ChcpptStorage: ts,ChcppttemplateStorage: ps}
}

// FindAll Requests
func (c PptinformationResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	results := c.PptinformationStorage.GetAll(r,-1,-1)
	if len(results) <= 0{
		return &Response{Res: results}, nil
	}
	if results[0].Url == ""{
		c.GetUrl(results[0])
		_= c.PptinformationStorage.Update(*results[0])
		//time.Sleep(40*time.Second)
	}
	return &Response{Res: results}, nil
}

func (c PptinformationResource) GetUrl( result *BmModel.Pptinformation){
	var iscreat int
	var temp string
	var url string
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println(err)
		return 
	}
	url = c.GenPPT(uuid)
	if url == ""{
		painc("err")
	}else{
		url=""
	}
	time.Sleep(2000)
	for i,data:=range result.Data{
		var contentints []interface{}
		iscreat=0	
		dataMap,_:= data.(bson.M)
		c.GetDatamap(dataMap,&temp,&contentints)
		tmp,err := c.ChcppttemplateStorage.GetOne(temp)
		if err != nil {
			fmt.Println(err)
			return 
		}
		Shapes:=tmp.Shapes
		for j,contentint := range contentints{
			var txt string
			var name string
			var css string
			var shapeType string
			var table string
			var chart string
			var pos []int
			var cells []string
			var formatstr string
			if len(Shapes)>0{
				Shapeint,_:=Shapes[j].(interface{})
				Shape,_:= Shapeint.(bson.M)
				c.GetShape(Shape,&pos,&shapeType,&formatstr,&cells,&name,&css)
			}
			
			content,_:= contentint.(bson.M)
			c.GetContent(content,formatstr,&txt,&table,&chart)
			if txt=="end"{
				url  = c.CreateSlider(uuid,"end","end",i)
				result.Url = url
				break
			}
			if iscreat==0{
				url  = c.CreateSlider(uuid,tmp.Slider_Type,txt,i)
				if url == ""{
					painc("err")
				}else{
					url=""
				}
				time.Sleep(2000)
				iscreat=1
			}
			if temp!=""&&txt!=""{
				url=c.PushText(uuid,txt,pos,i,shapeType)
				if url == ""{
					painc("err")
				}else{
					url=""
				}
				time.Sleep(2000)
			}else if temp!=""&&table!=""{
				chcppts,_:=c.ChcpptStorage.GetOne(table)
				for _,Cellint:=range chcppts.Cells{
					var coordinate string
					var value string
					cell,_:=Cellint.(bson.M)

					coordinateint:=cell["coordinate"]
					if coordinateint!=nil{
						coordinate=coordinateint.(string)
					}
					valueint:=cell["value"]
					if valueint!=nil{
						value=valueint.(string)
					}
					for t,celldata:=range cells{
						reg := regexp.MustCompile("#c#[^#]+")
						sitearr := reg.Find([]byte(celldata))
						site:=string(sitearr)[3:]
						if site==coordinate{
							cells[t]=cells[t]+value
						}
					}
				}		
				name=table+temp
				url  = c.ExcelPush(uuid ,name,cells)
				if url == ""{
					painc("err")
				}else{
					url=""
				}
				time.Sleep(2000)
				url  = c.Excel2PPT(uuid,name,pos,i)
				if url == ""{
					painc("err")
				}else{
					url=""
				}
				time.Sleep(2000)			
			}else if temp!=""&&chart!=""{
				chcppts,_:=c.ChcpptStorage.GetOne(chart)
				for _,Cellint:=range chcppts.Cells{
					var coordinate string
					var value string
					cell,_:=Cellint.(bson.M)

					coordinateint:=cell["coordinate"]
					if coordinateint!=nil{
						coordinate=coordinateint.(string)
					}
					valueint:=cell["value"]
					if valueint!=nil{
						value=valueint.(string)
					}
					for t,celldata:=range cells{
						reg := regexp.MustCompile("#c#[^#]+")
						sitearr := reg.Find([]byte(celldata))
						site:=string(sitearr)[3:]
						if site==coordinate{
							cells[t]=cells[t]+value
						}
					}
				}		
				name=chart+temp	
				url  = c.ExcelPush(uuid ,name,cells)
				if url == ""{
					painc("err")
				}else{
					url=""
				}
				time.Sleep(2000)
				url  = c.Excel2Chart(uuid,name,pos,i,shapeType,css)	
				if url == ""{
					painc("err")
				}else{
					url=""
				}
				time.Sleep(2000)	
			}
		}
	}
	url  = c.PushPPT(uuid)
}

func (c PptinformationResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.PptinformationStorage.GetOne(ID)
	return &Response{Res: res}, err
}

func (c PptinformationResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Pptinformation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.PptinformationStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

func (c PptinformationResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.PptinformationStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

func (c PptinformationResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.Pptinformation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.PptinformationStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}

func (c PptinformationResource) GenPPT(jobid string) string {
	var arr BmModel.Request
	arr.Command="GenPPT"
	arr.Jobid=jobid
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	//filePtr,_=os.Open("person_info.json")
	strbyt, err := ioutil.ReadFile("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	str:=string(strbyt)
	fmt.Println(str)

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}
	topic := "test"
	bkc.Produce(&topic, str)

	topics := []string{"test"}
	bkc.SubscribeTopics(topics, c.subscribeFunc)

	// request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	// request.Header.Set("Content-Type", "application/json")
	// response, _:= client.Do(request)
	// result, _ := ioutil.ReadAll(response.Body)	
	//os.Remove("person_info.json")
	//url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}
func (c PptinformationResource)subscribeFunc(a interface{}) {
	url=a.(string)
}
func (c PptinformationResource) CreateSlider(jobid string ,sliderType string , title string,slider int) string {
	var arr BmModel.Request
	var cs BmModel.CreateSlider
	cs.Slider=slider
	cs.SliderType=sliderType
	cs.Title=title
	arr.CreateSlider=&cs
	//arr.Slider=&cs

	arr.Command="CreateSlider"
	arr.Jobid=jobid
	
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	//requestBody := bytes.NewBuffer(nil)
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	filePtr, _= os.Open("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	request.Header.Set("Content-Type", "application/json")
	response, _:= client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)	
	os.Remove("person_info.json")
	url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}

func (c PptinformationResource) PushText(jobid string ,content string , pos []int,slider int,shapeType string) string {
	var arr BmModel.Request
	var cs BmModel.TextSetContent
	cs.Slider=slider
	cs.Content=content
	cs.Pos=pos

	cs.ShapeType=shapeType
	
	arr.TextSetContent=&cs
	arr.Command="PushText"
	arr.Jobid=jobid
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	filePtr, _= os.Open("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	request.Header.Set("Content-Type", "application/json")
	response, _:= client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)	
	os.Remove("person_info.json")
	url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}

func (c PptinformationResource) ExcelPush(jobid string ,name string , cells []string) string {
	var arr BmModel.Request
	var cs BmModel.ExcelPush
	cs.Name=name
	cs.Cells=cells
	arr.ExcelPush=&cs
	arr.Command="ExcelPush"
	arr.Jobid=jobid
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	filePtr, _= os.Open("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	request.Header.Set("Content-Type", "application/json")
	response, _:= client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)	
	os.Remove("person_info.json")
	url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}

func (c PptinformationResource) Excel2Chart(jobid string ,name string , pos []int,slider int, chartType string,css string) string {
	var arr BmModel.Request
	var cs BmModel.Excel2Chart
	cs.Name=name
	cs.Pos=pos
	cs.Css=css
	cs.ChartType = chartType
	cs.Slider=slider
	arr.Excel2Chart=&cs
	arr.Command="Excel2Chart"
	arr.Jobid=jobid
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	filePtr, _= os.Open("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	request.Header.Set("Content-Type", "application/json")
	response, _:= client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)	
	os.Remove("person_info.json")
	url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}

func (c PptinformationResource) Excel2PPT(jobid string ,name string , pos []int,slider int) string {
	var arr BmModel.Request
	var cs BmModel.Excel2PPT
	cs.Name=name
	cs.Pos=pos
	cs.Slider=slider
	arr.Excel2PPT=&cs
	arr.Command="Excel2PPT"
	arr.Jobid=jobid
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	filePtr, _= os.Open("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	request.Header.Set("Content-Type", "application/json")
	response, _:= client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)	
	os.Remove("person_info.json")
	url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}

func (c PptinformationResource) PushPPT(jobid string ) string {
	var arr BmModel.Request
	arr.Command="PushPPT"
	arr.Jobid=jobid
	client := http.Client{}
	filePtr, err := os.Create("person_info.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	err = jsonapi.MarshalPayload(filePtr,&arr)
	filePtr.Close()

	filePtr, _= os.Open("person_info.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	request, err := http.NewRequest("POST", "http://192.168.100.195:9999/api/ppt", filePtr)
	request.Header.Set("Content-Type", "application/json")
	response, _:= client.Do(request)
	result, _ := ioutil.ReadAll(response.Body)	
	os.Remove("person_info.json")
	url := string(result)
	s := strings.Split(url, ":")
	url = s[4][1:len(s[4])-4]
	return url
}

func (c PptinformationResource) GetShape(Shape bson.M,pos *[]int,shapeType *string,formatstr *string,cells *[]string,name *string,css *string) {
	posint:=Shape["pos"]
	if posint!=nil{
		posints,_:=posint.([]interface{})
		for _,posint:=range posints{
			posint,_:=posint.(int)
			tt:=int32(posint)
			float:=float32(tt)
			float=float/0.000278
			*pos=append(*pos,int(float))
		}
	}

	shapeTypeint:=Shape["shapeType"]
	if shapeTypeint!=nil{
		*shapeType=shapeTypeint.(string)
	}

	formatint:=Shape["format"]
	if formatint!=nil{
		*formatstr=formatint.(string)
	}

	cellsint:=Shape["cells"]
	if cellsint!=nil{
		cellsints,_:=cellsint.([]interface{})
		for _,cellint:=range cellsints{
			cell,_:=cellint.(string)
			*cells=append(*cells,cell)
		}
	}

	nameint:=Shape["name"]
	if nameint!=nil{
		*name="tmp"
	}

	cssint:=Shape["css"]
	if cssint!=nil{
		*css=cssint.(string)
	}
}

func (c PptinformationResource) GetContent(content bson.M,formatstr string,txt *string,table *string,chart*string){
	txtsint:=content["txts"]
	if txtsint!=nil{
		txtints := txtsint.([]interface{})
		for _,txtint:=range txtints{
			if txtint!=nil{
				txttmp:=txtint.(bson.M)
				for txtx,txtcent := range txttmp{
					txtstr:=txtcent.(string)
					if txtstr=="end"{
						*txt=txtstr
						return
					}
					formatstr=strings.Replace(formatstr,txtx,txtstr,1)
				}
				
			}
		}
	}
	*txt=formatstr
	tableint:=content["table"]
	if tableint!=nil{
		*table=tableint.(string)
	}

	chartint:=content["chart"]
	if chartint!=nil{
		*chart=chartint.(string)
	}
}

func (c PptinformationResource) GetDatamap(dataMap bson.M,temp*string,contentints*[]interface{}){
	tempint:=dataMap["temp"]
	if tempint!=nil{
		*temp=tempint.(string)
	}	
	contentsint,_:= dataMap["contents"]
	*contentints,_=contentsint.([]interface{})
}


