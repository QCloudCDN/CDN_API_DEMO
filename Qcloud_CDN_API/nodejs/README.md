## qcloud cdn openapi nodejs版本sdk

## 安装

    npm install qcloud-cdn-node-sdk --save

## API

参见[API文档](https://github.com/QCloudCDN/CDN_API_SDK/blob/master/README.md)


## 错误码

参见[错误码文档](https://www.qcloud.com/document/product/228/5078)

## 使用

### 准备工作

qcloud账号的secret_id和secret_key可以从[https://console.qcloud.com/capi](https://console.qcloud.com/capi) 获取

### 初始化SDK配置

```js
const qcloudSDK = require('qcloud-cdn-node-sdk');

qcloudSDK.config({
    secretId: 'qcloud账号的secretId',
    secretKey: 'qcloud账号的参数表示secretKey'
})
```

### 调用具体的CDN方法

```js
/************Action对应的名字************/

//API文档见 https://github.com/QCloudCDN/CDN_API_SDK/blob/master/README.md

// DescribleCdnHosts
// GetHostInfoByHost
// GetHostInfoById 
// RefreshCdnUrl
// RefreshCdnDir
// UpdateCache 
// UpdateCdnProject
// UpdateCdnHost
// UpdateCdnConfig
// OfflineHost
// AddCdnHost
// OnlineHost
// DeleteCdnHost
// GenerateLogList
// GetCdnRefreshLog
// GetCdnStatTop
// GetCdnStatusCode
// DescribeCdnHostDetailedInfo
// DescribeCdnHostInfo
/************************/

// action对应的参数无需传递公共请求参数
qcloudSDK.request('Action的名字', action对应的参数对象, callback)

```

## 示例

```js
const qcloudSDK = require('qcloud-cdn-node-sdk');

qcloudSDK.config({
    secretId: 'AAAAAAAAAAAAAAAAA',
    secretKey: 'BBBBBBBBBBBBBBBBBB'
})

qcloudSDK.request('DescribeCdnHostInfo', {
    'startDate': '2016-12-01',
    'endDate': '2016-12-01',
    'statType': 'bandwidth',
    'projects.0': '123',
    'hosts.0': 'www.test.com', 
    'hosts.1': 'www.test2.com' 
}, (res) => {
    // res为json格式
    // do something
})
```

## 反馈

欢迎提issue

## LICENSE

MIT
