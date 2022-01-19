package demo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"crypto/md5"
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
func RequestCmdb(uri string, ip string, method MethodStr, params interface{}) []byte {
	// 请求地址
	reqUrl := ip + uri
	// 时间戳
	requestTime := time.Now().Unix()
	var bytesData []byte
	if params != nil && params != "" {
		bytesData, _ = json.Marshal(params)
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
	fmt.Println("签名signature ==> ", signature)
	keys := make(map[string]string)
	keys["accesskey"] = AccessKey
	keys["signature"] = signature
	keys["expires"] = strconv.FormatInt(requestTime, 10)
	var urlParams string
	if method == MethodStrGet || method == MethodStrDELETE {
		if data, ok := params.(map[string]string); ok {
			urlParams = HttpBuildQuery(data) + "&" + HttpBuildQuery(keys)
		} else {
			urlParams = HttpBuildQuery(keys)
			//urlParams = HttpBuildQuery(keys)
			//urlParams = "accesskey="+AccessKey+"&signature="+signature+"&expires="+keys["expires"]
		}
	} else {
		urlParams = HttpBuildQuery(keys)
	}
	if urlParams != "" {
		if find := strings.Contains(reqUrl, "?"); find {
			reqUrl = reqUrl + "&" + urlParams
		} else {
			reqUrl = reqUrl + "?" + urlParams
		}
	}
	/*	fmt.Println("\n reqUrl ==>",reqUrl,bytesData)
		fmt.Println("new_str1",bytes.NewBuffer(bytesData))
		fmt.Println("new_str2",bytes.NewReader(bytesData))
		fmt.Println("new_str3",string(bytesData))
	*/
	req, err := http.NewRequest(string(method), reqUrl, bytes.NewBuffer(bytesData))
	req.Host = "openapi.easyops-only.com"
	req.Header.Add("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(" handle error===>", err)
	}
	fmt.Println("请求返回状态码：：", resp.StatusCode, "请求返回数据：：：", string(body))
	return body
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
			fmt.Println("\nurlParams =>>>>>", urlParams)
		}
	}
	if sign.Method == MethodStrPOST || sign.Method == MethodStrPUT {
		if data, ok := sign.Data.(string); ok {
			fmt.Print("sign.Data===>", data)
			bodyContent = Md5([]byte(data))
			//bodyContent = md5V2(data)
		} else {
			/*indent, _ := json.MarshalIndent(sign.Data, "", "")
			replace := strings.Replace(string(indent), "\n", "", -1)
			replace = strings.Replace(replace, ",\"", ", \"", -1)
			bodyContent = Md5([]byte(replace))
*/
			//fmt.Print("replace=>",replace)
			marshal, _ := json.Marshal(sign.Data)
			bodyContent = Md5String(string(marshal))

		}
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
	fmt.Println("\n strSign ==>", strSign)
	return HmacSha1([]byte(sign.SecretKey), strSign)
}

func md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func HttpBuildQuery(params map[string]string) (paramStr string) {
	paramsArr := make([]string, 0, len(params))
	for k, v := range params {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	}
	paramStr = strings.Join(paramsArr, "&")
	return paramStr
}
