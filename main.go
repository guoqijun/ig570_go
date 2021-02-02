package main

import (
	"fmt"
	"ig570_go/conf"
)

func main() {
	conf.DefaultInit()
	fmt.Println(conf.Env)
	fmt.Println(conf.Cnf.Url570)
	fmt.Println("i love you")
}
