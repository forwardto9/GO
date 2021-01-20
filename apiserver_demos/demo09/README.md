实战：API身份验证

重点变化文件： auth.go, router.go, login.go, code.go

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo09/ ./apiserver
// 查看配置文件
less apiserver/conf/config.yaml 
cd apiserver
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver
```

API 登录，申请Token

```shell
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/login -d'{"username":"admin","password":"admin"}'

//
{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MTExMDYyODEsImlkIjowLCJuYmYiOjE2MTExMDYyODEsInVzZXJuYW1lIjoiYWRtaW4ifQ.Qd9I_dzOAMWotLy-72ocrdBLVWWHmeG_11GVxLpj1Z0"}}
```

不带Token的创建用户的API

```shell
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"user1","password":"user1234"}'

//
{"code":20103,"message":"The token was invalid.","data":null}
```

带上Token的创建用户的API

```shell
curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MTExMDYyODEsImlkIjowLCJuYmYiOjE2MTExMDYyODEsInVzZXJuYW1lIjoiYWRtaW4ifQ.Qd9I_dzOAMWotLy-72ocrdBLVWWHmeG_11GVxLpj1Z0" -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"user2","password":"user12342"}

//
{"code":0,"message":"OK","data":{"username":"user2"}}
```







