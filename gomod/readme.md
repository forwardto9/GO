go module使用说明

为了解除对GOPATH的依赖编译运行，使用go module，可以让go项目在任意位置都可以编译运行

1. 创建需要编译的go源码文件(main.go)
2. 按照main.go源文件中对github上的包的依赖导入语法完成远程包的导入
3. 按照main.go源文件中对本地包的依赖导入路径规则(绝对路径，且在同一个项目文件夹下)完成本地包的导入
4. 初始化module，使用命令`go mod init yourprojectname`
   - go.sum，记录所依赖的远程包的版本信息，module信息
   - go.mod，包含module声明，go版本信息、远程依赖库的版本信息

> 如果需要使用GOPATH进行编译，则需要在项目文件夹下将go.sum, go.mod文件删除，然后将项目文件夹放到GOPATH目录下的src目录即可