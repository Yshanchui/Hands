package cmd

import (
	"Hands/conf"
	"Hands/global"
	"Hands/router"
	"Hands/utils"
	"fmt"
)

func Start() {
	var initErr error
	//初始化系统配置文件
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()
	//初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	//初始化redis连接
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	//初始化过程中是否存在错误
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	//初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("--------------clean--------------")
}
