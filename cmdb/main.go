package main

import (
	"fmt"
	"os"
	"log"

	cmdb "zc-poc/cmdb/demo"

	"zc-poc/go-simplejson"
	"reflect"
	"encoding/json"
	"time"
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

	// getPipelineDetail("测试应用", "5d5ff3fcaffb4")

	// 执行构建
	runId := execute_pipelines()

	// 构建日志
	buildLog(runId)

	return

	search_app_project()

	return
	//get_pipeline()

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
	log.SetPrefix("\n\n[/cmdb/object/BUSINESS/instance/_search]返回结果::")
	log.Println(string(res))
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
	log.SetPrefix("\n\n[/cmdb/object/APP/instance/_search]返回结果::")
	log.Println(string(json_str))

	res, _ := simplejson.NewJson([]byte(json_str))
	projects, _ := res.Get("data").Get("list").Array()
	// 项目S
	for i := range projects {
		projectId, _ := res.Get("data").Get("list").GetIndex(i).Get("instanceId").String()
		projectName, _ := res.Get("data").Get("list").GetIndex(i).Get("name").String()
		fmt.Println("应用名称：", projectName, "应用ID：", projectId)
		// 流水线S  #应用所拥有的流水线_代码项目
		pipelines, _ := res.Get("data").Get("list").GetIndex(i).Get("PIPELINE_PROJECT").Array()
		// fmt.Print(pipelines)
		for _, pipeline := range pipelines {
			if data, ok := pipeline.(map[string]interface{}); ok {
				fmt.Println("流水线名称：", data["name"], "流水线ID：", data["instanceId"], "创建人", data["creator"], "创建时间", data["ctime"])
				//  pipeline["_PIPELINE"] 流水线，变量variables
				fmt.Print("流水线变量：", data["_PIPELINE"])

			}
		}
	}

	return ""
}

// 临时接口，开发时候（甲方配置流水线详情接口）
func getPipelineDetail(projectName, pipelineId string) {
	businessParams := make(map[string]interface{})
	businessParams["sort"] = map[string]interface{}{
		"name": 1,
	}
	businessParams["fields"] = map[string]interface{}{
		"*":                          true,
		"PIPELINE_PROJECT":           true,
		"PIPELINE_PROJECT._PIPELINE": true,
	}
	businessParams["query"] = map[string]interface{}{
		"name": map[string]interface{}{
			"$in": []string{projectName},
		},
	}
	// page_size 最大不能超过 3千 ，默认20个
	businessParams["page_size"] = 3000
	businessParams["page"] = 1
	json_str := cmdb.RequestCmdb("/cmdb/object/APP/instance/_search", cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, businessParams)
	log.SetPrefix("\n\n[/cmdb/object/APP/instance/_search]返回结果::")
	log.Println(string(json_str))

	res, _ := simplejson.NewJson([]byte(json_str))
	projects, _ := res.Get("data").Get("list").Array()
	// 应用信息
	for i := range projects {
		projectId, _ := res.Get("data").Get("list").GetIndex(i).Get("instanceId").String()
		projectName, _ := res.Get("data").Get("list").GetIndex(i).Get("name").String()
		fmt.Println("应用名称：", projectName, "应用ID：", projectId)
		// 流水线项目S  #应用所拥有的流水线_代码项目
		pipelineProjects, _ := res.Get("data").Get("list").GetIndex(i).Get("PIPELINE_PROJECT").Array()
		// fmt.Print(pipelines)
		for _, pipelineProject := range pipelineProjects {
			if data, ok := pipelineProject.(map[string]interface{}); ok {
				fmt.Println("流水线项目名称：", data["name"], "流水线项目ID：", data["instanceId"], "创建人", data["creator"], "创建时间", data["ctime"])
				//  pipeline["_PIPELINE"] 流水线，变量variables

				// 流水线项目下的 -> 流水线
				if pipelines, ok := data["_PIPELINE"].([]interface{}); ok {
					for _, v := range pipelines {
						if pipeline, ok := v.(map[string]interface{}); ok {
							fmt.Println("流水线可读名称：", pipeline["alias_name"], "流水线名称：", pipeline["name"], "流水线ID：", pipeline["instanceId"], "创建人", pipeline["creator"], "创建时间", pipeline["ctime"])

							fmt.Println("类型", reflect.TypeOf(pipeline["variables"]))
							bytes, _ := json.Marshal(pipeline["variables"])
							fmt.Print("流水线变量：", string(bytes))

						}
					}
				}

			}
		}
	}

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
	project_id := "5a3f2e1cdb2ff"  //项目ID
	pipeline_id := "5a40382f43d16" //
	uri := fmt.Sprintf("/pipeline/api/pipeline/v1/projects/%s/pipelines/%s", project_id, pipeline_id)
	res := cmdb.RequestCmdb(uri, cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, "")
	log.SetPrefix("[get_pipeline]")
	log.Println(res)
}

// （触发）执行流水线
func execute_pipelines() string {
	project_id := "596404f7fd420"  // 流水线项目ID
	pipeline_id := "5d5ff3fcaffb4" // 流水线ID
	uri := fmt.Sprintf("/pipeline/api/pipeline/v1/projects/%s/pipelines/%s/execute", project_id, pipeline_id)
	businessParams := make(map[string]interface{})

	businessParams["branch"] = "master"
	// businessParams["tag"] = "aaaaaaaa"

	// 流水线变量 [ variables ]  {"name":"qqqq","value":"aaaa"}
	variable := make(map[string]interface{})
	variable["name"] = "VM"
	variable["value"] = nil

	variables := make([]map[string]interface{}, 0)
	variables = append(variables, variable)

	businessParams["inputs"] = variables

	requestCmdb := cmdb.RequestCmdb(uri, cmdb.EasyopsOpenApiHost, cmdb.MethodStrPOST, businessParams)
	log.SetPrefix("\n\n[/pipeline/api/pipeline/v1/projects/%s/pipelines/%s/execute]返回结果::")
	log.Println(string(requestCmdb))
	// 触发流水线信息  {"code":0,"codeExplain":"","error":"","data":{"id":"61ea651c9bcfbd3c1e902e15"}}
	// 触发流水线信息  {"code":0,"codeExplain":"","error":"","data":{"id":"61ea65f79bcfbd3c1e902e17"}}
	fmt.Println("触发流水线信息 ", string(requestCmdb))

	res, _ := simplejson.NewJson([]byte(requestCmdb))
	runId, _ := res.Get("data").Get("id").String()
	fmt.Print("运行的id：", runId)
	return runId
}

// # 查看流水线构建日志
func buildLog(runId string) {
	for {
		time.Sleep(time.Second * 2)
		uri := fmt.Sprintf("/pipeline/api/pipeline/v1/builds/%s", runId)
		requestCmdb := cmdb.RequestCmdb(uri, cmdb.EasyopsOpenApiHost, cmdb.MethodStrGet, "")

		log.SetPrefix("\n\n[/pipeline/api/pipeline/v1/builds/]返回结果::")
		log.Println(string(requestCmdb))
		fmt.Print("构建日志", string(requestCmdb))
		res, _ := simplejson.NewJson([]byte(requestCmdb))

		// 执行状态 [pending \ succeeded 执行完成 \ failed 执行失败]
		runSate, _ := res.Get("data").Get("status").Get("state").String()

		var isBreak bool
		if runSate == "succeeded" {
			// 执行完成
			isBreak = true
		}
		if runSate == "failed" {
			// 流水线执行失败
			isBreak = true
		}
		stages, _ := res.Get("data").Get("stages").Array()
		for _, v := range stages {
			if logInfo, ok := v.(map[string]interface{}); ok {
				fmt.Println("步骤名称 : ", logInfo["stage_name"])
				if steps, ok := logInfo["steps"].([]interface{}); ok {
					for _, stepMap := range steps {
						if step, ok := stepMap.(map[string]interface{}); ok {
							uri := fmt.Sprintf("/pipeline/api/pipeline/v1/progress_log/%s", step["log_id"])
							requestCmdb := cmdb.RequestCmdb(uri, cmdb.EasyopsOpenApiHost, cmdb.MethodStrGet, "")
							log.SetPrefix("\n\n[/pipeline/api/pipeline/v1/progress_log]返回结果::")
							log.Println(string(requestCmdb))

							fmt.Print("日志信息请求信息：", string(requestCmdb))
							res, _ := simplejson.NewJson([]byte(requestCmdb))
							log, _ := res.Get("data").Get("logs").String()
							fmt.Println("日志信息 : ", log)
						}
					}
				}
			}
		}

		if isBreak {
			break
		}

	}
}

// 流水线执行记录
