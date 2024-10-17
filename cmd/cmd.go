package cmd

import (
	"fmt"
	"gin-vue/conf"
	"gin-vue/global"
	"gin-vue/router"
)

func Start() {
	// init config
	conf.InitConfig()
	// init logger
	global.Logger = conf.InitLogger()
	// init router
	router.InitRounter()
}

func Clean() {
	fmt.Println("=========Clean===============")
}
