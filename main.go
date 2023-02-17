package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/taoruicheng/tour/cmd"
)

func main() {
	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件异常:%v", err)
	}
	log.Printf("server.port: %v", viper.GetInt("server.port"))

	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err = cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

}
