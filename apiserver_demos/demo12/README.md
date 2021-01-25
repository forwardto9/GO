进阶：给API命令增加版本功能

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo12/ apiserver/
cd apiserver
make mod
make
```



```shell
./apiserver -v
```

