实战：初始化Mysql数据库并建立连接

```shell
rm -rf $GOPATH/src/apiserver
cp -a ~/Workspace/GitHub/apiserver_demos/demo04/ ./apiserver
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver
```

```ruby
Authentication plugin 'caching_sha2_password' cannot be loaded: dlopen(/usr/local/mysql/lib/plugin/caching_sha2_password.so, 2): image not found
```

- 是MySQL兼容问题，需要修改数据库的认证方式
- MySQL8.0版本默认的认证方式是caching_sha2_password
- MySQL5.7版本则为mysql_native_password。

[解决办法](// https://www.jianshu.com/p/9a645c473676)

本文使用

```sql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'yourpassword';
```