package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	//读取配置文件
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("app inited")
	fmt.Println("config mysql:", viper.Get("mysql"))
}

func InitMysql() {
	//初始化mysql  gorm连接mysql,viper.GetString()用于获取配置文件中的mysql配置
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	fmt.Println("mysql inited")
}
