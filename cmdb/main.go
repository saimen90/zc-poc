package main

import (
	"fmt"
	"os"
	"log"

	cmdb "zc-poc/cmdb/demo"
)

// CMDB（优维系统）
func init() {
	file := "./" + "cmdb_response" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[CMDB]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {

	business_search := business_search()

	fmt.Print("business_search结果：", business_search)

	search_app_project := search_app_project()
	fmt.Print("search_app_project结果：", search_app_project)

	test_app()

	fmt.Println("CMDB、ITSM、YOUWEI、CONTAINER")
}

// 获取（/cmdb/object/BUSINESS/instance/_search ） 系统 -> 应用 -> 项目 ---> 流水
func business_search() string {
	businessParams := make(map[string]interface{})
	businessParams["sort"] = map[string]interface{}{
		"name": 1,
	}
	businessParams["fields"] = map[string]interface{}{
		"name":                                       true,
		"_businesses_APP":                            true,
		"instanceId":                                 true,
		"_businesses_APP.memo":                       true,
		"developer":                                  true,
		"_businesses_APP.owner":                      true,
		"_businesses_APP.tester":                     true,
		"_businesses_APP.developer":                  true,
		"_businesses_APP.USER_WB":                    true,
		"_businesses_APP.PIPELINE_PROJECT":           true,
		"_businesses_APP.PIPELINE_PROJECT._PIPELINE": true,
	}
	// page_size 最大不能超过 3千 ，默认20个
	businessParams["page_size"] = 3000
	businessParams["page"] = 1
	res := cmdb.RequestCmdb("/cmdb/object/BUSINESS/instance/_search", cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, businessParams)
	log.SetPrefix("[business_search]")
	log.Println(res)
	return res
}

// 查看所有应用下的流水线（/cmdb/object/BUSINESS/instance/_search ） 系统 -> 应用 -> 项目 ---> 流水
func search_app_project() string {
	businessParams := make(map[string]interface{})
	businessParams["sort"] = map[string]interface{}{
		"name": 1,
	}
	businessParams["fields"] = map[string]interface{}{
		"*":                          true,
		"PIPELINE_PROJECT":           true,
		"PIPELINE_PROJECT._PIPELINE": true,
	}
	app_name := "测试系统"
	businessParams["query"] = map[string]interface{}{
		"name": map[string]interface{}{
			"$in": []string{app_name},
		},
	}
	// page_size 最大不能超过 3千 ，默认20个
	businessParams["page_size"] = 3000
	businessParams["page"] = 1
	res := cmdb.RequestCmdb("/cmdb/object/APP/instance/_search", cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, businessParams)
	log.SetPrefix("[search_app_project]")
	log.Println(res)
	return res
}

// 测试-调试使用
func test_app() string {
	res := cmdb.RequestCmdb("/cmdb/toolkit/tools/APP", cmdb.EasyopsOpenApiHost, cmdb.MethodStrGet, "")
	log.SetPrefix("[test_app]")
	log.Println(res)
	return res
}
