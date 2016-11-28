##腾讯云CDN API 概览
## 消耗及统计量查询

| API                                      | 功能说明                                     |
| ---------------------------------------- | ---------------------------------------- |
| [DescribeCdnHostInfo](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2CDN%E6%B6%88%E8%80%97%E7%BB%9F%E8%AE%A1) | 查询流量、带宽、请求数、IP访问数、命中率等总统计信息              |
| [DescribeCdnHostDetailedInfo](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2CDN%E6%B6%88%E8%80%97%E6%98%8E%E7%BB%86) | 查询流量、带宽、请求数、IP访问数、命中率等统计明细信息，1-3日时间粒度为5分钟，3-7日明细粒度为1小时，7-30日明细粒度为1天 |
| [GetCdnStatusCode](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2%E8%BF%94%E5%9B%9E%E7%A0%81%E7%BB%9F%E8%AE%A1) | 查询返回码 200、206、304、404、416、500 统计明细       |
| [GetCdnStatTop](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2%E6%B6%88%E8%80%97%E6%8E%92%E5%90%8D) | 查询省份、运营商、URL的流量/带宽排名情况，TOP100            |



## 域名查询

| API                                      | 功能说明                             |
| ---------------------------------------- | -------------------------------- |
| [DescribeCdnHosts](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2%E5%9F%9F%E5%90%8D%E4%BF%A1%E6%81%AF) | 查询所有域名详细信息，包括配置信息，支持分页查询         |
| [GetHostInfoByHost](https://www.qcloud.com/doc/api/231/%E6%A0%B9%E6%8D%AE%E5%9F%9F%E5%90%8D%E6%9F%A5%E8%AF%A2%E5%9F%9F%E5%90%8D%E4%BF%A1%E6%81%AF) | 根据域名查询域名详细信息，包括配置信息，支持多个域名查询     |
| [GetHostInfoById](https://www.qcloud.com/doc/api/231/%E6%A0%B9%E6%8D%AE%E5%9F%9F%E5%90%8DID%E6%9F%A5%E8%AF%A2%E5%9F%9F%E5%90%8D%E4%BF%A1%E6%81%AF) | 根据域名ID查询域名详细信息，包括配置信息，支持多个域名ID查询 |



## 域名管理

| API                                      | 功能说明                     |
| ---------------------------------------- | ------------------------ |
| [AddCdnHost](https://www.qcloud.com/doc/api/231/%E6%96%B0%E5%A2%9E%E5%8A%A0%E9%80%9F%E5%9F%9F%E5%90%8D)| 接入域名至腾讯云CDN              |
| [OnlineHost](https://www.qcloud.com/doc/api/231/%E4%B8%8A%E7%BA%BFCDN%E5%9F%9F%E5%90%8D) | 上线指定CDN域名                |
| [OfflineHost](https://www.qcloud.com/doc/api/231/%E4%B8%8B%E7%BA%BFCDN%E5%9F%9F%E5%90%8D) | 下线指定CDN域名                |
| [DeleteCdnHost](https://www.qcloud.com/doc/api/231/%E5%88%A0%E9%99%A4%E5%8A%A0%E9%80%9F%E5%9F%9F%E5%90%8D) | 删除指定CDN域名                |
| [UpdateCdnHost](https://www.qcloud.com/doc/api/231/%E8%AE%BE%E7%BD%AE%E6%BA%90%E7%AB%99%E4%BF%A1%E6%81%AF) | 修改域名源站设置                 |
| [UpdateCdnProject](https://www.qcloud.com/doc/api/231/%E8%AE%BE%E7%BD%AE%E5%9F%9F%E5%90%8D%E6%89%80%E5%B1%9E%E9%A1%B9%E7%9B%AE) | 修改域名所属项目                 |
| [UpdateCache](https://www.qcloud.com/doc/api/231/%E8%AE%BE%E7%BD%AE%E7%BC%93%E5%AD%98%E8%A7%84%E5%88%99) | 修改域名对应的缓存规则配置            |
| [UpdateCdnConfig](https://www.qcloud.com/doc/api/231/%E8%AE%BE%E7%BD%AE%E5%9F%9F%E5%90%8D%E9%85%8D%E7%BD%AE) | 对指定域名的缓存、防盗链、回源等各项信息进行设置 |



## 域名刷新

| API                                      | 功能说明                |
| ---------------------------------------- | ------------------- |
| [GetCdnRefreshLog](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2%E5%88%B7%E6%96%B0%E7%BA%AA%E5%BD%95) | 查询指定时间区间内，刷新日志、刷新次数 |
| [RefreshCdnUrl](https://www.qcloud.com/doc/api/231/%E5%88%B7%E6%96%B0URL) | 刷新URL               |
| [RefreshCdnDir](https://www.qcloud.com/doc/api/231/%E5%88%B7%E6%96%B0%E7%9B%AE%E5%BD%95) | 刷新目录                |

 

## 日志查询

| API                                      | 功能说明     |
| ---------------------------------------- | -------- |
| [GenerateLogList](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2%E6%97%A5%E5%BF%97%E4%B8%8B%E8%BD%BD%E9%93%BE%E6%8E%A5) | 查询日志下载链接 |



## 辅助工具

| API                                      | 功能说明         |
| ---------------------------------------- | ------------ |
| [GetCdnMiddleSourceList](https://www.qcloud.com/doc/api/231/%E6%9F%A5%E8%AF%A2CDN%E4%B8%AD%E9%97%B4%E6%BA%90) | 查询CDN中间源IP列表 |

