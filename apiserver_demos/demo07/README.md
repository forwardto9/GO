实战：用户业务逻辑处理（业务处理）

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo07/ apiserver/
cd apiserver
less conf/config.yaml
go mod init apiserver
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver 
```

添加用户

```shell
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"uwei","password":"uwei123"}'

// {"code":0,"message":"OK","data":{"username":"uwei"}

curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"testor","password":"test123"}'
// {"code":0,"message":"OK","data":{"username":"testor"}}
```

查询前20个用户

```shell
curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"offset": 0, "limit": 20}'
```



删除id=4的用户

```shell
curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/4

{"code":0,"message":"OK","data":null}
```



更新id=2的用户信息

```shell
curl -XPUT -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2 -d'{"username":"kong001","password":"kongmodify"}'


{"code":0,"message":"OK","data":null}
```

