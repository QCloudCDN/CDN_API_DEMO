## CDN openapi java demo
- 函数入口位于cdnapi_demo.java，指定接口名，公共参数（secretKey,secretId）以及接口参数。
- 签名方法和请求方法位于./Utilities/Sign.java 和 ./Utilities/Request.java。
- 可通过Request.generateUrl方法校验签名及请求url。
