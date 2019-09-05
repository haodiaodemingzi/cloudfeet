# cloudfeet api 接口

>SS配置,域名采集,域名鉴定 

## 拉取SS服务器配置

URI地址: `/api/v1/proxy`
请求方法: `GET`
请求参数: 无

响应JSON:

```json
{
    "code": 200,
    "msg": "ok",
    "data": {
        "id": 1,
        "server":"ss.csdc.io",
        "port": 7007,
        "method": "chacha20",
        "password":""
    }
}
```



## 上传采集到的域名

URI地址:  `/api/v1/pac/domains`
请求方法: `POST`
请求json:

```
{
	"source": "mac app",
	"domains": "www.baidu.com,www.google.com,www.facebook.com"
}
```

响应JSON:
失败:

```json
{
    "code":100,
    "msg":"失败原因"
}
```

成功:

```json
{
    "code":200,
    "msg":""
}
```

##  拉取尚未判定结果的域名列表

URI地址: `/api/v1/pac/domains
请求方法: `GET
请求参数: 无                                                                                                                                                                  

最大拉取1000个域名

响应JSON:
失败:

```json
{
    "code":100,
    "msg":"no ok"
    "data":null
}
```

成功:

```json
{
    "code":200,
    "msg":"ok"
    "data": [
      "www.baidu.com",
      "www.google.com"
    ],
	"limit": 1000
}
```

## 拉取被屏蔽的域名

URI地址: `/api/v1/pac/domains?status=1

请求参数: 无

请求方法：GET

响应JSON:
失败:

```json
{
    "code":100,
    "msg":"no ok"
    "data":null
}
```

成功:

```json
{
    "code":200,
    "msg":"ok"
    "data": [
      "www.baidu.com",
      "www.google.com"
    ],
	"limit": 1000
}
```



## 提交判定后的域名结果

URI地址: `/api/v1/checker/domains
请求方法: ` PUT
请求JSON:

```json
{
    "source": "app",
    "domains": {
        "google.com": 2,
        "qq.com":1,
        "wabac.com": 1
    }
}
```


响应JSON:
成功/失败:

```json
{
    "code": 200/100,
    "msg":""
}
```
