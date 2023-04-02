package service

import (
	"Gin_Chat/models"
	"Gin_Chat/utils"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Getuserlst
// @Summary所用户
// @Schemes
// @Description do ping
// @Router /user/getuserlist [get]
func GetUserList(c *gin.Context) {
	userlist := models.GetUserList()
	c.JSON(200, gin.H{
		"message": userlist,
	})
}

// CreateUser
// @Summary 新增用户
// @Schemes
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Description do ping
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/register [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())
	origin_user := models.FindUserByName(user.Name)
	if origin_user.Name != "" {
		c.JSON(404, gin.H{
			"message": "用户名已注册",
		})
		return
	}
	if password != repassword {
		c.JSON(404, gin.H{
			"message": "两次密码不一致",
		})
		return
	} else {
		utils.MakePassword(password,salt)
		// user.Password = password
		models.CreateUser(user)
		c.JSON(200, gin.H{
			"message": "新增用户成功",
		})
		return
	}
}

// DeleteUser
// @Summary 删除用户
// @Schemes
// @param id query string false "id"
// @Description do ping
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteuser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Schemes
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Description do ping
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/updateuser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")

	//使用govalidator进行校验
	_,err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"essae": "修改参数不配",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"mssage": "修改用户成功",
		})
	}
}
