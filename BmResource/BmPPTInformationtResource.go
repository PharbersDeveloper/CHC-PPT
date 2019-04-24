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
	
)

type PPTInformationResource struct {
	RequestStorage 		  *BmDataStorage.RequestStorage
	PPTInformationStorage *BmDataStorage.PPTInformationStorage
	ChcpptStorage         *BmDataStorage.ChcpptStorage
	ChcppttemplateStorage         *BmDataStorage.ChcppttemplateStorage
}

func (c PPTInformationResource) NewPPTInformationResource(args []BmDataStorage.BmStorage) PPTInformationResource {
	var cs *BmDataStorage.RequestStorage
	var ss *BmDataStorage.PPTInformationStorage
	var ts *BmDataStorage.ChcpptStorage
	var ps *BmDataStorage.ChcppttemplateStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "RequestStorage" {
			cs = arg.(*BmDataStorage.RequestStorage)
		}else if tp.Name() == "PPTInformationStorage" {
			ss = arg.(*BmDataStorage.PPTInformationStorage)
		}else if tp.Name() == "ChcpptStorage" {
			ts = arg.(*BmDataStorage.ChcpptStorage)
		} else if tp.Name() == "ChcppttemplateStorage" {
			ps = arg.(*BmDataStorage.ChcppttemplateStorage)
		}
	}
	return PPTInformationResource{RequestStorage: cs,PPTInformationStorage:ss,ChcpptStorage: ts,ChcppttemplateStorage: ps}
}

// FindAll Requests
func (c PPTInformationResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	results := c.PPTInformationStorage.GetAll(r,-1,-1)
	if len(results) <= 0{
		return &Response{Res: results}, nil
	}
	// fmt.Println(results[0])
	// return &Response{Res: results}, nil
	var url string
	uuid, _ := uuid.GenerateUUID()
	url = c.GenPPT(uuid)
	time.Sleep(2000)
	result := results[0]
	var iscreat int
	var temp string
	if result.Url == ""{
		result.Url = url
		for i,data:=range result.Data{
			iscreat=0
			fmt.Println(i)
			if i==8{
				fmt.Println(i)
				url  = c.CreateSlider(uuid,"end","end",i)
				break
			}
			dataMap,_:= data.(bson.M)
			tempint:=dataMap["temp"]
			if tempint!=nil{
				temp=tempint.(string)
			}
			
			tmp,_ := c.ChcppttemplateStorage.GetOne(temp)
			Shapes:=tmp.Shapes
		
			contentsint,_:= dataMap["contents"]
			contentints,_:=contentsint.([]interface{})
			for j,contentint := range contentints{
				var txt string
				var name string
				var css string
				var shapeType string
				//var chartType string
				var table string
				var chart string
				var pos []int
				var cells []string
				Shapeint,_:=Shapes[j].(interface{})
				Shape,_:= Shapeint.(bson.M)
				posint:=Shape["pos"]
				if posint!=nil{
					posints,_:=posint.([]interface{})
					for _,posint:=range posints{
						posint,_:=posint.(int)
						tt:=int32(posint)
						float:=float32(tt)
						float=float/0.000278
						pos=append(pos,int(float))
					}
				}
				fmt.Println(pos)
				shapeTypeint:=Shape["shapeType"]
				if shapeTypeint!=nil{
					shapeType=shapeTypeint.(string)
				}
				formatint:=Shape["format"]
				var formatstr string
				if formatint!=nil{
					formatstr=formatint.(string)
				}
			
				cellsint:=Shape["cells"]
				if cellsint!=nil{
					cellsints,_:=cellsint.([]interface{})
					for _,cellint:=range cellsints{
						cell,_:=cellint.(string)
						cells=append(cells,cell)
					}
				}	
				nameint:=Shape["name"]
				if nameint!=nil{
					name="tmp"
				}
				cssint:=Shape["css"]
				if cssint!=nil{
					css=cssint.(string)
				}
				// chartTypeint:=Shape["chartType"]
				// if chartTypeint!=nil{
				// 	chartType=chartTypeint.(string)
				// }
				fmt.Println(shapeType)
				content,_:= contentint.(bson.M)
				txtint:=content["txt"]
				if txtint!=nil{
					txttmp:=txtint.(string)
					txt=strings.Replace(formatstr,"txt",txttmp, 1)
				}

				tableint:=content["table"]
				if tableint!=nil{
					table=tableint.(string)
				}
				
				chartint:=content["chart"]
				if chartint!=nil{
					chart=chartint.(string)
				}
				if iscreat==0{
					url  = c.CreateSlider(uuid,tmp.Slider_Type,txt,i)
					time.Sleep(2000)
					iscreat=1
				}
				if temp!=""&&txt!=""{
					c.PushText(uuid,txt,pos,i,shapeType)
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
						fmt.Println(coordinate,value)
					}		
					name=table+temp
					url  = c.ExcelPush(uuid ,name,cells)
					time.Sleep(2000)
					url  = c.Excel2PPT(uuid,name,pos,i)
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
						fmt.Println(coordinate,value)
					}		
					name=chart+temp	
					url  = c.ExcelPush(uuid ,name,cells)
					time.Sleep(3000)
					url  = c.Excel2Chart(uuid,name,pos,i,shapeType,css)	
					time.Sleep(3000)
					
				}
			}
		}
		url  = c.PushPPT(uuid)
		_= c.PPTInformationStorage.Update(*result)
	}
	return &Response{Res: results}, nil
	// 	jobid:="12345678"
	// 	var url string
	// 	url = c.GenPPT(jobid,"123123")
	// 	time.Sleep(2000)
	// 	url  = c.CreateSlider(jobid , "123123", "title","口服降糖药市场CHC数据分析报告\n2018Q3YTD",0)
	// 	time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药市场CHC数据分析报告\n2018Q3YTD#C#20Blue#]##P#center#}#",[]int{892086,2658273,7550359,690647},0,"Rectangle")
	// 	time.Sleep(2000)
	// 	url  = c.CreateSlider(jobid , "123123", "default","市场分析",1)
	// 	time.Sleep(2000)
	// 	url  = c.ExcelPush(jobid , "123123", "1",[]string{ 
	// 	"#c#A1#s#col_common1*row_5#t#String#v#", 
	// 	"#c#A2#s#col_common1*row_6#t#String#v#其他",
	// 	"#c#A3#s#col_common1*row_7#t#String#v#GLP-1激动剂",
	// 	"#c#A4#s#col_common1*row_8#t#String#v#DPP4抑制剂",
	// 	"#c#A5#s#col_common1*row_9#t#String#v#格列酮类",
	// 	"#c#A6#s#col_common1*row_10#t#String#v#格列奈类",
	// 	"#c#A7#s#col_common1*row_11#t#String#v#磺脲类",
	// 	"#c#A8#s#col_common1*row_12#t#String#v#双胍类",
	// 	"#c#A9#s#col_common1*row_13#t#String#v#糖苷类",
	// 	"#c#B1#s#col_common1*row_5#t#String#v#2017Q3YTD",
	// 	"#c#B2#s#col_common1*row_6#t#Number#v#8952147.219",
	// 	"#c#B3#s#col_common1*row_7#t#Number#v#57955.23183",
	// 	"#c#B4#s#col_common1*row_8#t#Number#v#231121.6192",
	// 	"#c#B5#s#col_common1*row_9#t#Number#v#9941337.422",
	// 	"#c#B6#s#col_common1*row_10#t#Number#v#17155958.46",
	// 	"#c#B7#s#col_common1*row_11#t#Number#v#71104098.57",
	// 	"#c#B8#s#col_common1*row_12#t#Number#v#66350640.52",
	// 	"#c#B9#s#col_common1*row_13#t#Number#v#270920150.6",
	// 	"#c#C1#s#col_common1*row_5#t#String#v#2018Q3YTD",
	// 	"#c#C2#s#col_common1*row_6#t#Number#v#13217709.42",
	// 	"#c#C3#s#col_common1*row_7#t#Number#v#208363.6539",
	// 	"#c#C4#s#col_common1*row_8#t#Number#v#7380419.939",
	// 	"#c#C5#s#col_common1*row_9#t#Number#v#12691516.72",
	// 	"#c#C6#s#col_common1*row_10#t#Number#v#22029627.82",
	// 	"#c#C7#s#col_common1*row_11#t#Number#v#94492996.64",
	// 	"#c#C8#s#col_common1*row_12#t#Number#v#122650916.3",
	// 	"#c#C9#s#col_common1*row_13#t#Number#v#381883258.3",
	// 	"#c#D1#s#col_common1*row_5#t#String#v#",
	// 	"#c#D2#s#col_common1*row_6#t#Number#v#2.0%",
	// 	"#c#D3#s#col_common1*row_7#t#Number#v#0.0%",
	// 	"#c#D4#s#col_common1*row_8#t#Number#v#0.1%",
	// 	"#c#D5#s#col_common1*row_9#t#Number#v#2.2%",
	// 	"#c#D6#s#col_common1*row_10#t#Number#v#3.9%",
	// 	"#c#D7#s#col_common1*row_11#t#Number#v#16.0%",
	// 	"#c#D8#s#col_common1*row_12#t#Number#v#14.9%",
	// 	"#c#D9#s#col_common1*row_13#t#Number#v#60.9%",
	// 	"#c#E1#s#col_common1*row_5#t#String#v#",
	// 	"#c#E2#s#col_common1*row_6#t#Number#v#2.0%",
	// 	"#c#E3#s#col_common1*row_7#t#Number#v#0.0%",
	// 	"#c#E4#s#col_common1*row_8#t#Number#v#1.1%",
	// 	"#c#E5#s#col_common1*row_9#t#Number#v#1.9%",
	// 	"#c#E6#s#col_common1*row_10#t#Number#v#3.4%",
	// 	"#c#E7#s#col_common1*row_11#t#Number#v#14.4%",
	// 	"#c#E8#s#col_common1*row_12#t#Number#v#18.7%",
	// 	"#c#E9#s#col_common1*row_13#t#Number#v#58.3%"})

	// 	url  = c.Excel2Chart(jobid , "123123", "1",[]int{892086,2658273,7550359,690647},1,"Stacked","")
	// 	time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物市场规模在北京市CHC的2018Q3YTD年以47.2%的增长速度达到6.55亿人民币#C#18Black#]##P#left#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{   899280,1723021,3528776,525179},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushText(jobid , "123123", "#{##[#口服降糖药物CHC市场\n（百万人民币)#C#14Black#]##P#center#}#",
	// 						[]int{892086,2658273,7550359,690647},1,"Rectangle")
	// 						time.Sleep(2000)
	// 	url  = c.PushPPT(jobid , "123123")		
	// 	fmt.Println(url)
	
		// result := results[0]
		// if result.Url == ""{
		// 	for date:=range result.Date{

		// 	}
		// a := []string {"GenPPT","CreateSlider"} 
		// if 1==1 {
		// 	for _,com := range a{
		// 		switch com{
		// 		case "GenPPT":
		// 			url := c.GenPPT(jobid , 123123 )
		// 			fmt.Println(url)
		// 		case "CreateSlider":
		// 			url := c.CreateSlider(jobid , 123123 )
		// 			fmt.Println(url)
		// 		case "PushText":
		// 		case "Excel2Chart":
		// 		case "Excel2PPT":
		// 		case "ExportPPT":
		// 		case "TextSetContent":
		// 		case "PushPPT":
		// 		}
		// 	}
		// 	//result.Url = url
		// }
	// return &Response{Res: results}, nil
}

// FindOne choc
func (c PPTInformationResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.PPTInformationStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c PPTInformationResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.PPTInformation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.PPTInformationStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c PPTInformationResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.PPTInformationStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c PPTInformationResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.PPTInformation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.PPTInformationStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
func (c PPTInformationResource) GenPPT(jobid string) string {
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
func (c PPTInformationResource) CreateSlider(jobid string ,sliderType string , title string,slider int) string {
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

func (c PPTInformationResource) PushText(jobid string ,content string , pos []int,slider int,shapeType string) string {
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
func (c PPTInformationResource) ExcelPush(jobid string ,name string , cells []string) string {
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
func (c PPTInformationResource) Excel2Chart(jobid string ,name string , pos []int,slider int, chartType string,css string) string {
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
func (c PPTInformationResource) Excel2PPT(jobid string ,name string , pos []int,slider int) string {
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
func (c PPTInformationResource) PushPPT(jobid string ) string {
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