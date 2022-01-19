package main

import (
	"fmt"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"zc-poc/cmdb"
)

func HmacSha1(key []byte,data string)  string{
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte(nil)))
}


func main() {

	cmdb.RequestCmdb("/cmdb/toolkit/tools/APP",cmdb.EasyopsOpenApiHost,cmdb.MethodStrGet, "")
	fmt.Println("CMDB、ITSM、YOUWEI、CONTAINER")
}


// 获取（/cmdb/object/BUSINESS/instance/_search ） 系统 -> 应用 -> 项目 ---> 流水
func business_search()  {
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
	cmdb.RequestCmdb("/cmdb/object/BUSINESS/instance/_search",cmdb.EasyopsOpenApiHost,cmdb.MethodStrPOST, businessParams)
}