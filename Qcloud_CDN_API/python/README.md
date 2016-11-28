## python demo 工具使用方法
其中 -u参数表示 SECRET_ID -p 参数表示SECRET_KEY,可从https://console.qcloud.com/capi 获取，其余参数参考https://www.qcloud.com/doc/api/231/API%E6%A6%82%E8%A7%88

## DescribleCdnHosts
python QcloudCdnTools_V2.py DescribeCdnHosts -u xxxxx -p xxxxxxx

## GetHostInfoByHost
python QcloudCdnTools_V2.py GetHostInfoByHost -u xxxxx -p xxxxxxx --hosts xxxxxxxtang.oa.com --hosts www.xxxxxxx.217.oa.com 

## GetHostInfoById 
python QcloudCdnTools_V2.py GetHostInfoById -u xxxxx -p xxxxxxx --ids 184736 --ids 206348

## RefreshCdnUrl
python QcloudCdnTools_V2.py RefreshCdnUrl -u xxxxx -p xxxxxxx --urls http://xxxxxxxtang.sp.oa.com/test.php --urls http://xxxxxxxtang.sp.oa.com/test1.php 
 
## RefreshCdnDir
python QcloudCdnTools_V2.py RefreshCdnDir -u xxxxx -p xxxxxxx --dirs http://xxxxxxxtang.sp.oa.com/test/ --dirs http://xxxxxxxtang.sp.oa.com/a/ --dirs http://xxxxxxxtang.sp.oa.com/b/ 

## UpdateCache 
python QcloudCdnTools_V2.py UpdateCache -u xxxxx -p xxxxxxx  --hostId 206092 --cache [[0,\"all\",1023448]]

## UpdateCdnProject
python QcloudCdnTools_V2.py UpdateCdnProject -u xxxxx -p xxxxxxx  --hostId 206092 --projectId 0

## UpdateCdnHost
python QcloudCdnTools_V2.py UpdateCdnHost -u xxxxx -p xxxxxxx  --hostId 206092 --host xxxxxxxtang.vip2.oa.com --origin 1.1.1.2

## UpdateCdnConfig
python QcloudCdnTools_V2.py UpdateCdnConfig -u xxxxx -p xxxxxxx  --hostId 206084 --projectId 0 --cacheMode custom --cache  [[0,\"all\",1023448]] --refer [1,[\"qq.baidu.com\",\"v.qq.com\"]] --fwdHost qq.com.cn --fullUrl off --debug

## OfflineHost
python QcloudCdnTools_V2.py OfflineHost -u xxxxx -p xxxxxxx  --hostId 206092

## AddCdnHost
python QcloudCdnTools_V2.py AddCdnHost -u xxxxx -p xxxxxxx  --host evincai.oa.com --projectId 0 --hostType cname --origin 1.1.1.1

## OnlineHost
python QcloudCdnTools_V2.py OnlineHost -u xxxxx -p xxxxxxx  --hostId 206092


## DeleteCdnHost
python QcloudCdnTools_V2.py DeleteCdnHost -u xxxxx -p xxxxxxx  --hostId 81094

##GenerateLogList
python QcloudCdnTools_V2.py GenerateLogList -u xxxxx -p xxxxxxx  --hostId 206092

##GetCdnRefreshLog
python QcloudCdnTools_V2.py GetCdnRefreshLog -u xxxxxxxxxxxx -p xxxxxxxxxxxx  --startDate 2016-04-25 --endDate 2016-04-26

##GetCdnStatTop
python QcloudCdnTools_V2.py GetCdnStatTop -u xxxxxxxxxxxx -p xxxxxxxxxxxx  --startDate 2016-05-08 --endDate 2016-05-09 --statType bandwidth --projects 0 --hosts ping.cdn.qcloud.com --hosts ts6.cache.qcloudcdn.com

##GetCdnStatusCode
python QcloudCdnTools_V2.py GetCdnStatusCode -u xxxxxxxxxxxx -p xxxxxxxxxxxx  --startDate 2016-05-08 --endDate 2016-05-09 --projects 0 --hosts ping.cdn.qcloud.com --hosts ts6.cache.qcloudcdn.com

##DescribeCdnHostDetailedInfo
python QcloudCdnTools_V2.py DescribeCdnHostDetailedInfo -u xxxxxxxxxxxx -p xxxxxxxxxxxx  --startDate 2016-05-08 --endDate 2016-05-09 --projects 0 --hosts ping.cdn.qcloud.com --hosts ts6.cache.qcloudcdn.com  --statType bandwidth

##DescribeCdnHostInfo
python QcloudCdnTools_V2.py DescribeCdnHostInfo -u xxxxxxxxxxxx -p xxxxxxxxxxxx  --startDate 2016-05-08 --endDate 2016-05-09  --projects 0  --statType bandwidth  --debug
