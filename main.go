package main

import (
	"fmt"
	"ig570_go/conf"
	"ig570_go/router"
)

func main() {
	conf.DefaultInit()

	fmt.Println(conf.Env)

	r := router.RoutersInit()
	fmt.Println("开始运行")
	_ = r.Run(":8081")
}
