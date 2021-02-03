package main

import (
	"fmt"
	"ig570_go/conf"
	"ig570_go/router"
)

func main() {
	conf.DefaultInit()

	fmt.Println(conf.Env)
	fmt.Println(conf.Cnf.Url570)

	r := router.RoutersInit()
	fmt.Println("开始运行")
	_ = r.Run(":8081")
}
