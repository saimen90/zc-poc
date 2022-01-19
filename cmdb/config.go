package cmdb

// API 秘钥管理
// access_key
var AccessKey = "ffca5ae5093820c032289a7b"

// scretc_key
var ScrectKey = "736e6b6e556c616251764d6b71564f75797453486c67676a66434a6d524a6249"

// 请求地址
var EasyopsOpenApiHost = "http://192.168.124.124"

type MethodStr string

const (
	MethodStrGet    MethodStr = "GET"
	MethodStrPOST   MethodStr = "POST"
	MethodStrPUT    MethodStr = "PUT"
	MethodStrDELETE MethodStr = "DELETE"
)
