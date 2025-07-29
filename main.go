package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xiao-hub-create/book/api"
	"github.com/xiao-hub-create/book/config"
)

func main() {
	//从配置文件中加载配置
	//加载配置
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}
	if err := config.LoadConfigFromYaml(path); err != nil {
		fmt.Printf("加载配置错误：%s", err)
		os.Exit(1)
	}

	//访问加载后的配置
	conf := config.Get()

	r := gin.Default()

	api.NewBookHandler().Registry(r)

	if err := r.Run(conf.App.Address()); err != nil {
		log.Println(err)
	}

}
