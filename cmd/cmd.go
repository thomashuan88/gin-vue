package cmd

import (
	"fmt"
	"gin-vue/conf"
	"gin-vue/global"
	"gin-vue/router"
	"gin-vue/utils"
)

func Start() {
	var initErr error
	// init config
	conf.InitConfig()
	// init logger
	global.Logger = conf.InitLogger()
	// init db
	db, err := conf.InitDB()
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.DB = db

	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// init router
	router.InitRounter()
}

func Clean() {
	fmt.Println("=========Clean===============")
}
