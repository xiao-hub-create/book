package config

import (
	"os"

	"github.com/xiao-hub-create/book/config"
)

func main() {
	//从配置文件中加载配置
	//加载配置
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}
	config.LoadConfigFromYaml(path)

	//访问加载后的配置
	conf := config.Get()

}
