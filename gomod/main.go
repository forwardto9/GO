package main

import (
	"fmt"

	"gomod/clib" // 依赖导入本地包

	testgopackageongithub "github.com/forwardto9/GO" // 依赖导入远程包
)

func main() {
	fmt.Println("Hello go mod")
	demoString := testgopackageongithub.DemoFunction()
	fmt.Println(demoString)
	clib.DemoMethod()

}
