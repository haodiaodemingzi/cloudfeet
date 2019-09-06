# cloudfeet api 接口


>SS配置,域名采集,域名鉴定 

## 获取API访问token
URI地址: `/api/v1/auth/token`
请求方法: `POST`
请求参数: json

```json
// 测试阶段随便填
{
    "username": "ooxxx",
    "password": "ooxx123" 
}
```
响应JSON:

```json
{
    "code": 200,
    "msg": "ok",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbWVzIiwicGFzc3dvcmQiOiJqYW1lczEyMyIsImV4cCI6MTU2Nzc3MDk2NSwiaXNzIjoiZ2luLWJsb2cifQ.VB1PVKTcwQ9V43SOt3BuVQCiDGhNj036G3k4_mJrWMo"
    }
}
```

## 拉取SS服务器配置

URI地址: `/api/v1/proxy`
请求方法: `GET`
请求参数: 
header:
    token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbWVzIiwicGFzc3dvcmQiOiJqYW1lczEyMyIsImV4cCI6MTU2Nzc3MDk2NSwiaXNzIjoiZ2luLWJsb2cifQ.VB1PVKTcwQ9V43SOt3BuVQCiDGhNj036G3k4_mJrWMo

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
header:
    token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbWVzIiwicGFzc3dvcmQiOiJqYW1lczEyMyIsImV4cCI6MTU2Nzc3MDk2NSwiaXNzIjoiZ2luLWJsb2cifQ.VB1PVKTcwQ9V43SOt3BuVQCiDGhNj036G3k4_mJrWMo

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
请求参数: 无                                                                                                                               header:
    token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbWVzIiwicGFzc3dvcmQiOiJqYW1lczEyMyIsImV4cCI6MTU2Nzc3MDk2NSwiaXNzIjoiZ2luLWJsb2cifQ.VB1PVKTcwQ9V43SOt3BuVQCiDGhNj036G3k4_mJrWMo
                                   

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
header:
    token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbWVzIiwicGFzc3dvcmQiOiJqYW1lczEyMyIsImV4cCI6MTU2Nzc3MDk2NSwiaXNzIjoiZ2luLWJsb2cifQ.VB1PVKTcwQ9V43SOt3BuVQCiDGhNj036G3k4_mJrWMo

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
请求方法: `PUT`
header:
    token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbWVzIiwicGFzc3dvcmQiOiJqYW1lczEyMyIsImV4cCI6MTU2Nzc3MDk2NSwiaXNzIjoiZ2luLWJsb2cifQ.VB1PVKTcwQ9V43SOt3BuVQCiDGhNj036G3k4_mJrWMo

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
