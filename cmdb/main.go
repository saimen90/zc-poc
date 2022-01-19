package main

import (
	"fmt"
	"os"
	"log"

	cmdb "zc-poc/cmdb/demo"

	"zc-poc/go-simplejson"
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

	search_app_project()
	return
	get_pipeline()

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
	return ""
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
	app_name := "测试应用"

	businessParams["query"] = map[string]interface{}{
		"name": map[string]interface{}{
			"$in": []string{app_name},
		},
	}
	// page_size 最大不能超过 3千 ，默认20个
	businessParams["page_size"] = 3000
	businessParams["page"] = 1
	json_str := cmdb.RequestCmdb("/cmdb/object/APP/instance/_search", cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, businessParams)


	res, _ := simplejson.NewJson([]byte(json_str))
	projects, _ := res.Get("data").Get("list").Array()
	for i := range projects {
		projectId, _ := res.Get("data").Get("list").GetIndex(i).Get("instanceId").String()
		projectName, _ := res.Get("data").Get("list").GetIndex(i).Get("name").String()
		fmt.Println("应用名称：",projectName,"应用ID：",projectId)

		pipelines, _ := res.Get("data").Get("list").GetIndex(i).Get("PIPELINE_PROJECT").Array()


		// 流水线名称 ， 流水线id
		for _, pipeline := range pipelines {
			if data, ok := pipeline.(map[string]interface{}); ok {
				fmt.Println("流水线名称：",data["name"],"流水线ID：",data["instanceId"],"创建人",data["creator"],"创建时间",data["ctime"])

				// 需要升级  pipeline["_PIPELINE"]
			}
		}

	}

	return ""
}

// 测试-调试使用
func test_app() string {
	res := cmdb.RequestCmdb("/cmdb/toolkit/tools/APP", cmdb.EasyopsOpenApiHost, cmdb.MethodStrGet, "")
	log.SetPrefix("[test_app]")
	log.Println(string(res))
	return string(res)
}

// 获取流水线信息
func get_pipeline() {
	project_id := "5a3f2e1cdb2ff" //项目ID
	pipeline_id := "5a40382f43d16" //
	uri := fmt.Sprintf("/pipeline/api/pipeline/v1/projects/%s/pipelines/%s", project_id, pipeline_id)
	res := cmdb.RequestCmdb(uri, cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, "")
	log.SetPrefix("[get_pipeline]")
	log.Println(res)
}

// （触发）执行流水线
func execute_pipelines() string {
	project_id := "5a3f2e1cdb2ff"
	pipeline_id := "5a40382f43d16"
	uri := fmt.Sprintf("/pipeline/api/pipeline/v1/projects/%s/pipelines/%s/execute", project_id, pipeline_id)
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
	 cmdb.RequestCmdb(uri, cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, businessParams)

	return ""
}

//  终端控制输出

// 流水线执行记录
