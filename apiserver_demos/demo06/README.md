实战：读取和返回HTTP请求

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo06/ apiserver/
cp ~/Workspace/GitHub/forwardto9/GO/apiserver_demos/demo04/conf/config.yaml apiserver/conf/
less apiserver/conf/config.yaml
cd apiserver
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver
```

```shell
// 注意这里的curl后的地址地址有desc字段，需要对地址整体双引号。
curl  -H "Content-Type: application/json" -d'{"username":"admin","password":"admin"}' -XPOST "http://127.0.0.1:8080/v1/user/admin2?desc=test"
```

```json
{
  "code":0,
  "message":"OK",
  "data":
  		{
        "username":"admin"
      }
}
```

