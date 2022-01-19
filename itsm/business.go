package itsm

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestItsm(uri string, ip string, method MethodStr, params interface{}) {
	reqUrl := ip + uri
	req, err := http.NewRequest(string(method), reqUrl, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(" handle error===>", err)
	}
	fmt.Println("请求返回状态码：：",resp.StatusCode,"请求返回数据：：：\n",string(body))
}