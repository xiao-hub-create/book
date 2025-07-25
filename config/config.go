package config

import "encoding/json"

// 定义程序配置
type Config struct {
	App   *App   `json:"app"`
	MySQL *MySQL `json:"mysql"`
}

func (c *Config) String() string {
	v, _ := json.Marshal(c)
	return string(v)
}

type App struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`
}
