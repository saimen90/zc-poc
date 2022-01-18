package cmdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 签名构建
type Signature struct {
	AccessKey   string    // access_key
	SecretKey   string    // srect_key
	RequestTime int64     // 时间戳
	Method      MethodStr // 请求方式 post 、 get
	Uri         string
	Data        interface{}
	ContentType string
}

// 获取签名（加密）
func RequestCmdb(uri string, ip string, method MethodStr, params interface{}) {
	// 请求地址
	reqUrl := ip + uri
	// 时间戳
	requestTime := time.Now().Unix()
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 创建签名
	signature := createSignature(Signature{
		AccessKey:   AccessKey,
		SecretKey:   ScrectKey,
		RequestTime: requestTime,
		Method:      method,
		Uri:         uri,
		Data:        params,
		ContentType: "application/json",
	})
	fmt.Println("\n签名signature ==> ", signature)
	keys := make(map[string]string)
	keys["accesskey"] = AccessKey
	keys["signature"] = signature
	keys["expires"] = strconv.FormatInt(requestTime, 10)
	var urlParams string
	if method == MethodStrGet || method == MethodStrDELETE {
		m := params.(map[string]string)
		urlParams = HttpBuildQuery(m) + "&" + HttpBuildQuery(keys)
	} else {
		urlParams = HttpBuildQuery(keys)
	}
	if find := strings.Contains(reqUrl, "?"); find {
		reqUrl = reqUrl + "&" + urlParams
	} else {
		reqUrl = reqUrl + "?" + urlParams
	}
	fmt.Println("\n reqUrl ==>",reqUrl)
	req, err := http.NewRequest(string(method), reqUrl, bytes.NewReader(bytesData))
	req.Header.Set("Host", "openapi.easyops-only.com")
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	return
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" handle error===>", err)
	}
	fmt.Println("请求返回数据：：：", string(body))
}

// 生产signature信息
func createSignature(sign Signature) string {
	urlParams := ""
	bodyContent := ""
	if sign.Method == MethodStrGet || sign.Method == MethodStrDELETE {
		var keys []string

		fmt.Print(sign.Data)
		if data, ok := sign.Data.(map[string]string); ok {
			for k := range data {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			var urlParamsArray []string
			for _, k := range keys {
				urlParamsArray = append(urlParamsArray, k+data[k])
			}
			urlParams = strings.Join(urlParamsArray, "")
			fmt.Println("\nurlParams =>>>>>",urlParams)
		}
	}
	if sign.Method == MethodStrPOST || sign.Method == MethodStrPUT {
		data, _ := json.Marshal(sign.Data)
		bodyContent = Md5(data)
	}
	signData := []string{
		string(sign.Method),
		sign.Uri,
		urlParams,
		sign.ContentType,
		bodyContent,
		strconv.FormatInt(sign.RequestTime, 10),
		sign.AccessKey,
	}
	strSign := strings.Join(signData, "\n")
	return HmacSha1([]byte(sign.SecretKey), strSign)
}


func HttpBuildQuery(params map[string]string) (paramStr string) {
	paramsArr := make([]string, 0, len(params))
	for k, v := range params {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	}
	paramStr = strings.Join(paramsArr, "&")
	return paramStr
}