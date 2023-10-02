package routes

import (
	"log"
	"net/http"
	"rookieCode/config"
	"rookieCode/dao"
	"rookieCode/utils"
	"time"
)

func InitGlobalVariable() {
	//初始化viper
	utils.InitViper()
	//初始化数据库
	dao.DB = utils.InitMySQLDB()
	// dao.DB = utils.InitSQLiteDB("gorm.db")
	// 初始化 Logger
	utils.InitLogger()
	//初始化redis
	utils.InitRedis()
	// 初始化 Casbin
	utils.InitCasbin(dao.DB)
}

// 后台服务
func AdminServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      AdminRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
