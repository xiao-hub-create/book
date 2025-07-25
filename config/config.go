package config

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func (m *MySQL) DB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FShanghai&allowNativePasswords=true", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败:%v", err))
	}
	if m.Debug {
		db = db.Debug()
	}
	return db
}

// func (m *MySQL) String() string {
// 	v, _ := json.Marshal(m)
// 	return string(v)
// }
