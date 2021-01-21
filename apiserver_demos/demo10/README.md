进阶：用HTTPS加密API请求

```shell
cd $GOPATH/src/
ls apiserver/conf
cd apiserver
go mod init apiserver
go build
go get -u github.com/shirou/gopsutil/disk
go build
./apiserver
```

不带CA信息的请求

```shell
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://127.0.0.1:8081/v1/user

// 报错信息：
curl: (60) SSL certificate problem: self signed certificate
More details here: https://curl.haxx.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.

```

带CA信息的请求

```shell
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://127.0.0.1:8081/v1/user --cacert conf/server.crt --cert conf/server.crt --key conf/server.key

// 查询结果：
{"code":0,"message":"OK","data":{"totalCount":5,"userList":[{...}]}}
```



// 注意token的有效性，如果无效，可以通过以下接口获取一个有效的token

````shell
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/login -d'{"username":"admin","password":"admin"}'
````

备注：

在 apiserver 中添加 HTTPS 功能，步骤很简单，大概分为以下三步。

生成私钥文件（server.key）和自签发的数字证书（server.crt）：

```shell
$ openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
```

