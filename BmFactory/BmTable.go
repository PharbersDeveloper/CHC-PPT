package BmFactory

import (
	"github.com/PharbersDeveloper/CHC-PPT/BmDataStorage"
	"github.com/PharbersDeveloper/CHC-PPT/BmHandler"
	"github.com/PharbersDeveloper/CHC-PPT/BmModel"
	"github.com/PharbersDeveloper/CHC-PPT/BmResource"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	//"github.com/PharbersDeveloper/Max-Report/BmMiddleware"
)

type BmTable struct{}

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmCreateSlider":                BmModel.CreateSlider{},
	"BmExcel2Chart":                 BmModel.Excel2Chart{},
	"BmExcel2PPT":                	 BmModel.Excel2PPT{},
	"BmExcelPush":                   BmModel.ExcelPush{},
	"BmRequest":                	 BmModel.Request{},
	"BmExportPPT":                 	 BmModel.ExportPPT{},
	"BmTextSetContent":              BmModel.TextSetContent{},
	"BmPPTInformation":              BmModel.PPTInformation{},
	"BmChcppt":                 BmModel.Chcppt{},
	"BmChcppttemplate":                 BmModel.Chcppttemplate{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmRequestResource":               BmResource.RequestResource{},
	"BmPPTInformationResource":               BmResource.PPTInformationResource{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmCreateSliderStorage":               BmDataStorage.CreateSliderStorage{},
	"BmExcel2ChartStorage":                BmDataStorage.Excel2ChartStorage{},
	"BmExcel2PPTStorage":                BmDataStorage.Excel2PPTStorage{},
	"BmExcelPushStorage":                BmDataStorage.ExcelPushStorage{},
	"BmRequestStorage":                BmDataStorage.RequestStorage{},
	"BmExportPPTStorage":                BmDataStorage.ExportPPTStorage{},
	"BmTextSetContentStorage":                BmDataStorage.TextSetContentStorage{},
	"BmPPTInformationStorage":                BmDataStorage.PPTInformationStorage{},
	"BmChcpptStorage":                BmDataStorage.ChcpptStorage{},
	"BmChcppttemplateStorage":                BmDataStorage.ChcppttemplateStorage{},
}

var BLACKMIRROR_MIDDLEWARE_FACTORY = map[string]interface{}{
	//"BmCheckTokenMiddleware": BmMiddleware.BmCheckTokenMiddleware{},
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	 "BmMongodbDaemon": BmMongodb.BmMongodb{},
	 "BmRedisDaemon":   BmRedis.BmRedis{},
}

var BLACKMIRROR_FUNCTION_FACTORY = map[string]interface{}{
	"BmCHC_PPTHandler":     	   BmHandler.CHC_PPTHandler{},
}


func (t BmTable) GetModelByName(name string) interface{} {
	return BLACKMIRROR_MODEL_FACTORY[name]
}

func (t BmTable) GetResourceByName(name string) interface{} {
	return BLACKMIRROR_RESOURCE_FACTORY[name]
}

func (t BmTable) GetStorageByName(name string) interface{} {
	return BLACKMIRROR_STORAGE_FACTORY[name]
}

func (t BmTable) GetDaemonByName(name string) interface{} {
	return BLACKMIRROR_DAEMON_FACTORY[name]
}

func (t BmTable) GetFunctionByName(name string) interface{} {
	return BLACKMIRROR_FUNCTION_FACTORY[name]
}

func (t BmTable) GetMiddlewareByName(name string) interface{} {
	return BLACKMIRROR_MIDDLEWARE_FACTORY[name]
}