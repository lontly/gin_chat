package models

import (
	"Gin_Chat/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserBasic继承gorm.Model
type UserBasic struct {
	gorm.Model
	Identity     string
	Name         string    //用户名称
	Password     string    //用户密码
	Phone        string    `valid:"matches(^1[3-9]{1}\\d{9}$)"`//手机号 可以校验
	Email        string    `valid:"email"`//邮箱   可以校验
	ClientIp     string    //客户端Ip
	ClientPort   string    //客户端端口
	LoginTime    time.Time //登录时间
	Salt 		 string    //随机数
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

//通过姓名搜索
func FindUserByName(Name string)UserBasic{
	user := UserBasic{}
	utils.DB.Where("name=?",Name).First(&user)
	return user
}
//通过手机号搜索
func FindUserByPhone(phone string)UserBasic{
	user := UserBasic{}
	utils.DB.Where("phone=?",phone).First(&user)
	return user
}
//通过邮箱搜索
func FindUserByEmail(email string)UserBasic{
	user := UserBasic{}
	utils.DB.Where("email=?",email).First(&user)
	return user
}

//添加用户
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}


//删除用户
func DeleteUser(user UserBasic) *gorm.DB{
	return utils.DB.Delete(&user)
}

//修改用户
func UpdateUser(user UserBasic) *gorm.DB{
	return utils.DB.Model(&user).Updates(user)
}