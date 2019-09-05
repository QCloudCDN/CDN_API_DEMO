## qcloud cdn openapi go版本sdk

## 使用方法
- 在GOPATH下的src 下添加openapi的sdk qcloudcdn_api
- qcloudcdn_api.Signature 生成签名 qcloudcdn_api.SendRequest 发送请求

## 代码示例
```
package main

import (
	"fmt"
	cdnapi "qcloudcdn_api"
)

func main() {
	/**get SecretKey & SecretId from https://console.qcloud.com/capi**/
	var Requesturl string = "cdn.api.qcloud.com/v2/index.php"
	var SecretKey string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	var Method string = "POST"

	/**params to signature**/
	params := make(map[string]interface{})
	params["SecretId"] = "ooooooooooooooooooooooooooooooo"
	params["Action"] = "DescribeCdnHosts"

	/*use qcloudcdn_api.Signature to obtain signature and params with correct signature**/
	request, _ := cdnapi.Signature(SecretKey, params, Method, Requesturl)
	fmt.Println("request : ", request)

	/*use qcloudcdn_api.SendRequest to send request**/
	response := cdnapi.SendRequest(Requesturl, request, Method)
	fmt.Println(response)
}    
```
