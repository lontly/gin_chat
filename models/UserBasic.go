package models

import (
	"Gin_Chat/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// UserBasic继承gorm.Model
type UserBasic struct {
	gorm.Model
	Identity     string
	Name         string    //用户名称
	Password     string    //用户密码
	Phone        string    //手机号
	Email        string    //邮箱
	ClientIp     string    //客户端Ip
	ClientPort   string    //客户端端口
	LoginTime    time.Time //登录时间
	LoginOutTime time.Time `gorm:"login_out_time" json:"login_out_time"` //下线时间
	Heartbeat    time.Time //心跳时间
	IsLoginout   bool      //是否下线
	DeviceInfo   string    //设备相关信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []UserBasic {
	data := make([]UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
