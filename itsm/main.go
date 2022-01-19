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
	requestItsm := itsm.RequestItsm("/api/misc/get_change_list", itsm.ItsmOpenApiHost, itsm.MethodStrGet, nil)
	log.SetPrefix("[requestItsm结果]")
	log.Println(requestItsm)
}
