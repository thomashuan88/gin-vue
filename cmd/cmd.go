package cmd

import (
	"fmt"
	"gin-vue/conf"
	"gin-vue/router"
)

func Start() {
	conf.InitConfig()
	router.InitRounter()
}

func Clean() {
	fmt.Println("=========Clean===============")
}
