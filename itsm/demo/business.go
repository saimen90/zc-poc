package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func RequestItsm(uri string, ip string, method MethodStr, params interface{}) string {
	reqUrl := ip + uri
	var urlParams string
	if method == MethodStrGet || method == MethodStrDELETE {
		if data, ok := params.(map[string]string); ok {
			urlParams = HttpBuildQuery(data)
		}
	}
	if urlParams != "" {
		if find := strings.Contains(reqUrl, "?"); find {
			reqUrl = reqUrl + "&" + urlParams
		} else {
			reqUrl = reqUrl + "?" + urlParams
		}
	}
	req, err := http.NewRequest(string(method), reqUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(" handle error===>", err)
	}
	fmt.Println("请求返回状态码：：", resp.StatusCode, "请求返回数据：：：\n", string(body))
	return string(body)
}

func HttpBuildQuery(params map[string]string) (paramStr string) {
	paramsArr := make([]string, 0, len(params))
	for k, v := range params {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	}
	paramStr = strings.Join(paramsArr, "&")
	return paramStr
}
