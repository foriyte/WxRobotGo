package main

import (
	"flag"
	"github.com/foriyte/WxRobotGo/server"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	configFile = flag.String("config", "", "配置文件路径")
)

func main() {
	flag.Parse()
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigFile(*configFile)
	if err := config.ReadInConfig(); err != nil {
		log.Println("读取配置文件失败:", err)
	}
	dbConf := config.Sub("mysql")
	server.InitWxToken(config)
	server.InitTuLingApi(config)
	server.InitDb(dbConf)
	log.Println("start server ...")
	http.HandleFunc("/", server.MessagePush)
	err := http.ListenAndServe(config.GetString("port"), nil)
	//err := s.ListenAndServe()
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("success!")

}
