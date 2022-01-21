package main

import (
	"os"
	"log"

	itsm "zc-poc/itsm/demo"
)

// ITSM 系统
func init() {
	file := "./" + "itsm_response" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[ITSM]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {

	search_where()

	// search_all()
}

func search_all() {
	requestItsm := itsm.RequestItsm("/api/misc/get_change_list", itsm.ItsmOpenApiHost, itsm.MethodStrGet, nil)
	log.SetPrefix("[requestItsm结果]")
	log.Println(requestItsm)
}

func search_where() {

	businessParams := make(map[string]string)

	businessParams["page"] = "1"
	businessParams["page_size"] = "10"

	businessParams["chg_app_id"] = "5a7a1f39e9801"    // 应用
	businessParams["chg_system_id"] = "5a7a1fe531fcd" // 系统

	requestItsm := itsm.RequestItsm("/api/misc/get_change_list", itsm.ItsmOpenApiHost, itsm.MethodStrGet, businessParams)
	log.SetPrefix("[requestItsm 带条件查询 结果]")
	log.Println(requestItsm)
}
