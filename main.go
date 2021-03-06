package main

import (
	"fmt"
	"github.com/PharbersDeveloper/CHC-PPT/BmFactory"
	"github.com/PharbersDeveloper/CHC-PPT/BmMaxDefine"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"net/http"
)

func main() {
	version := "v2"
	fmt.Println("pod archi begins")

	fac := BmFactory.BmTable{}

	var pod = BmMaxDefine.Pod{ Name: "swl test", Factory:fac }
	pod.RegisterSerFromYAML("resource/def.yaml")

	//本地测试
	//os.Setenv("BM_HOME", ".")
	//os.Setenv("BM_KAFKA_CONF_HOME", "resource/kafkaconfig.json")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig("BM_HOME")
	addr := bmRouter.Host + ":" + bmRouter.Port
	fmt.Println("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)
	pod.RegisterAllFunctions(version, api)
	//pod.RegisterAllMiddleware(api)
	handler := api.Handler().(*httprouter.Router)
	//pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+bmRouter.Port, handler)

	fmt.Println("pod archi ends")
}
