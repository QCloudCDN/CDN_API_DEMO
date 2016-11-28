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
	params["Action"] = "RefreshCdnUrl"
	params["urls.0"] = "http://hello.world.att.oa.com/test1.php"

	/*use qcloudcdn_api.Signature to obtain signature and params with correct signature**/
	signature, request_params := cdnapi.Signature(SecretKey, params, Method, Requesturl)
	fmt.Println("signature : ", signature)

	/*use qcloudcdn_api.SendRequest to send request**/
	response := cdnapi.SendRequest(Requesturl, request_params, Method)
	fmt.Println(response)
}
