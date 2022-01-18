package cmdb

// API 秘钥管理

// 请求地址
var EasyopsOpenApiHost = "http://192.168.124.124"

type MethodStr string

const (
	MethodStrGet    MethodStr = "GET"
	MethodStrPOST   MethodStr = "POST"
	MethodStrPUT    MethodStr = "PUT"
	MethodStrDELETE MethodStr = "DELETE"
)
