进阶：基于Nginx的API部署方案

## 安装Nginx

```shell
brew install nginx
nginx -v
brew services start nginx
```

## 配置Nginx

```json
# 记录 Nginx 进程 id
pid        logs/nginx.pid; 
# 配置 http
http {
    include       mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';
		# 记录接口访问
    access_log  logs/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    #keepalive_timeout  0;
    keepalive_timeout  65;

    #gzip  on;

    server {
  			# 配置监听端口号
        listen       80;
  			# 配置 host
        server_name  apiserver.com;

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            proxy_set_header Host $http_host;
            proxy_set_header X-Forwarded-Host $http_host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  					#proxy_pass：反向代理的路径（1.本机；2.端口要和 API 服务端口一致：8080）
            proxy_pass  http://127.0.0.1:8080/;
            client_max_body_size 5m;
        }
		}
}

```

macOS上，Nginx配置需要指定路径

```shell
nginx -c /usr/local/etc/nginx/nginx.conf
#如果遇到提示文件不存在，则创建提示中的路径文件
nginx: [error] open() "/usr/local/Cellar/nginx/1.19.4/logs/nginx.pid" failed (2: No such file or directory)
nginx: [emerg] open() "/usr/local/Cellar/nginx/1.19.4/logs/access.log" failed (2: No such file or directory)

cd /usr/local/Cellar/nginx/1.19.4/
mkdir logs
touch nginx.pid access.log
```



## 配置本机 host 映射

```
在 /etc/hosts 中添加一行：127.0.0.1 apiserver.com
```



## 启动服务

```shell
cd $GOPATH/src
rm -rf apiserver/*
cp -a ~/Workspace/GitHub/apiserver_demos/demo14/ apiserver
cd apiserver
less Makefile
make mod
make
less admin.sh
./admin.sh -h
./admin.sh status
./admin.sh start
```



## API 访问

```shell
curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" http://apiserver.com/v1/user

# {"code":10001,"message":"sql: database is closed","data":null}%  
# 此时可以查看 access.log 文件的记录
```



## 配置 Nginx 负载均衡

```json
#在nginx.conf文件的http标识字段内新增以下：
upstream apiserver.com {
     server 127.0.0.1:8080;
     server 127.0.0.1:8082;
}

#修改以下字段:
proxy_pass http://apiserver.com/;
```

注意：

负载均衡的演示需要多个后端服务，为此我们在同一个服务器上启动多个 apiserver，配置不同的端口（8080、8082），并采用 Nginx 默认的轮询转发策略（轮询：每个请求按时间顺序逐一分配到不同的后端服务器）。