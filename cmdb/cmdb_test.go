package cmdb

import (
	"fmt"
	"testing"
)

// 获取产品和模块数据
func TestGetProductAndModule(t *testing.T) {
	fmt.Println("获取 CMDB ，系统及应用信息")
	/*params1 := ` {
        "sort": {"name": 1},
        "fields": {"name": True, "_sub_system": True, "instanceId": True, "_sub_system._businesses_APP": True,"_businesses_APP":True, "_sub_system._sub_system._businesses_APP": True},
        "query": {"$and": [{"_parent_system": {"$exists": False}}]},
        "page_size": 3000,
        "page": aaaa
    }`*/
	params := make(map[string]string)

	params["aaa"] = "test"
	params["bb"] = "3"

	// RequestCmdb("/cmdb/object/BUSINESS/instance/_search",EasyopsOpenApiHost,MethodStrPOST,params)
	RequestCmdb("/cmdb/object/BUSINESS/instance/_search",EasyopsOpenApiHost,MethodStrPOST,params)

}