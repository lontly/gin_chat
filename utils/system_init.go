package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper" //用于读取配置文件
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	//自定义日志，打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout,"\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL的阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	//初始化mysql  gorm连接mysql,viper.GetString()用于获取配置文件中的mysql配置
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println("mysql inited")
}


