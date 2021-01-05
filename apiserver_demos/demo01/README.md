实战：启动一个最简单的RESTful API服务器

```shell
// 将demo01目录下的全部文件拷贝到指定目录
cp -a demo01/ $GOPATH/src/apiserver
```

```shell
cd $GOPATH/src/apiserver
// 建立mod管理模块
go mod init apiserver
// 解决包依赖
go mod vendor
// 构建，如果此步提示依赖的包有问题，可能是包依赖问题，可以尝试使用 go get -u <packagename>, 再执行 go mod vendor 解决依赖
go build
```

```shell
./apiserver 
```

```shell
// 开启另一个终端，进行接口测试
curl -XGET http://127.0.0.1:8080/sd/health
curl -XGET http://127.0.0.1:8080/sd/disk
curl -XGET http://127.0.0.1:8080/sd/cpu
curl -XGET http://127.0.0.1:8080/sd/ram
```

