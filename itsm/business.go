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
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" handle error===>", err)
	}
	fmt.Println("\nITSM::请求返回数据：：：", string(body))
}