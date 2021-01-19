实战：HTTP调用添加自定义处理逻辑

新增两个功能：

1. 在请求和返回的 Header 中插入 `X-Request-Id`（`X-Request-Id` 值为 32 位的 UUID，用于唯一标识一次 HTTP 请求）
2. 日志记录每一个收到的请求

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo08/ apiserver
cp ~/Workspace/GitHub/forwardto9/GO/apiserver_demos/demo07/conf/config.yaml apiserver/conf/
cd apiserver
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver
```



```shell
curl -v -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user
```

```
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /v1/user HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Type: application/json
> 
< HTTP/1.1 200 OK
< Access-Control-Allow-Origin: *
< Cache-Control: no-cache, no-store, max-age=0, must-revalidate, value
< Content-Type: application/json; charset=utf-8
< Expires: Thu, 01 Jan 1970 00:00:00 GMT
< Last-Modified: Tue, 19 Jan 2021 01:48:27 GMT
< X-Content-Type-Options: nosniff
< X-Frame-Options: DENY
< X-Request-Id: 3205aec7-87ad-43a6-9bdf-a7c694f695b2 // 返回请求头中添加了新字段
< X-Xss-Protection: 1; mode=block
< Date: Tue, 19 Jan 2021 01:48:27 GMT
< Content-Length: 662
< 
* Connection #0 to host 127.0.0.1 left intact
{"code":0,"message":"OK","data":{"totalCount":3,"userList":[{"id":3,"username":"uwei","sayHello":"Hello 46Y0C2fMR","password":"$2a$10$KwrYJqyTJXu.ouaXASXh5.wTP4CH7HUc2cszHB6INpZjtOiXOIqbO","createdAt":"2021-01-18 09:36:35","updatedAt":"2021-01-18 09:36:35"},{"id":2,"username":"kong001","sayHello":"Hello VeLAC2BMgz","password":"$2a$10$ne7n5cYDfpNYTZYoBr4KaOAUPlf/ca40nEl1jUi8mhIv6lkBnLJuO","createdAt":"2021-01-18 09:36:01","updatedAt":"2021-01-18 09:48:00"},{"id":0,"username":"admin","sayHello":"Hello V6LAjhfMgm","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}* Closing connection 0
```



日志记录

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /v1/user                  --> apiserver/handler/user.Create (7 handlers)
[GIN-debug] DELETE /v1/user/:id              --> apiserver/handler/user.Delete (7 handlers)
[GIN-debug] PUT    /v1/user/:id              --> apiserver/handler/user.Update (7 handlers)
[GIN-debug] GET    /v1/user                  --> apiserver/handler/user.List (7 handlers)
[GIN-debug] GET    /v1/user/:username        --> apiserver/handler/user.Get (7 handlers)
[GIN-debug] GET    /sd/health                --> apiserver/handler/sd.HealthCheck (7 handlers)
[GIN-debug] GET    /sd/disk                  --> apiserver/handler/sd.DiskCheck (7 handlers)
[GIN-debug] GET    /sd/cpu                   --> apiserver/handler/sd.CPUCheck (7 handlers)
[GIN-debug] GET    /sd/ram                   --> apiserver/handler/sd.RAMCheck (7 handlers)
{"level":"INFO","timestamp":"2021-01-19 09:46:59.699","file":"apiserver/main.go:59","msg":"Start to listening the incoming requests on http address: :8080"}
{"level":"INFO","timestamp":"2021-01-19 09:46:59.700","file":"apiserver/main.go:73","msg":"Waiting for the router, retry in 1 second."}
{"level":"INFO","timestamp":"2021-01-19 09:47:00.702","file":"apiserver/main.go:56","msg":"The router has been deployed successfully."}
{"level":"INFO","timestamp":"2021-01-19 09:48:27.494","file":"middleware/logging.go:84","msg":"2.944ms       | 127.0.0.1    | GET /v1/user | {code: 0, message: OK}"} // 记录了请求
```

