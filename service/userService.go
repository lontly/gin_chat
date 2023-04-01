package service

import (
	"Gin_Chat/models"
	"github.com/gin-gonic/gin"
)

// Getuserlist
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags 获取用户列表
// @Success 200 {string} Helloworld
// @Router /getuserlist [get]
func GetUserList(c *gin.Context) {
	userlist := models.GetUserList()
	c.JSON(200, gin.H{
		"message": userlist,
	})
}

// Getuserlist
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags 获取用户列表
// @Success 200 {string} Helloworld
// @Router /getuserlist [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(404, gin.H{
			"message": "两次密码不一致",
		})
	}else{
		user.Password = password
		models.CreateUser(user)
	}

}
