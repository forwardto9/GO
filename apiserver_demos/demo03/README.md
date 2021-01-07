实战：记录和管理API日志

```shell
cd $GOPATH/src 
cp -a ~/Workspace/GitHub/apiserver_demos/demo03/ apiserver/
cd apiserver 
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver 
```

测试日志自动打包能力

```go
在main.go中增加以下代码：
for {
    log.Info("zip log");
}
```

