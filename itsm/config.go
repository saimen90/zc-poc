package itsm



// 请求地址
var ItsmOpenApiHost = "http://192.168.125.162:8001"

type MethodStr string

const (
	MethodStrGet    MethodStr = "GET"
	MethodStrPOST   MethodStr = "POST"
	MethodStrPUT    MethodStr = "PUT"
	MethodStrDELETE MethodStr = "DELETE"
)
