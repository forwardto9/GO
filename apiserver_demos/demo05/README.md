实战：自定义业务错误信息

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo05/ ./apiserver
cd apiserver
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver 
```



测试API

```shell
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin"}'
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"password":"admin"}'
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin","password":"admin"}'
```