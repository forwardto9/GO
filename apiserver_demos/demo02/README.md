实战：配置文件读取

```shell
rm -rf $GOPATH/src/apiserver 
cp -a demo02/ $GOPATH/src/apiserver
cd $GOPATH/src/apiserver 
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver
```

测试热更新配置

```go
// 在main.go中添加以下代码
for {
        fmt.Println(viper.GetString("runmode"))
        time.Sleep(4*time.Second)
    }
```